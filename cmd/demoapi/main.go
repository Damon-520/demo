package main

import (
	"demoapi/internal/conf"
	"demoapi/internal/core/logger"
	"demoapi/internal/elasticsearch"
	"demoapi/internal/pkg/loggerx"
	"demoapi/internal/router"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "demo-api"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs/local/api/", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		log.Fatal("c.Load Err: ", err)
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		log.Fatal("c.Scan Err: ", err)
		panic(err)
	}

	// 初始化 Elasticsearch 连接
	esURL := "https://localhost:9200" // 根据实际情况修改
	esClient, err := elasticsearch.InitES(esURL)
	if err != nil {
		log.Fatal("Elasticsearch connection failed: ", err)
	}

	// 创建索引
	if err := elasticsearch.CreateIndex("test_index"); err != nil {
		log.Fatal("Failed to create index: ", err)
	}
	// 配置Zipkin追踪
	zipkinURL := "http://localhost:9411/api/v2/spans" // Zipkin服务器地址
	reporter := zipkinhttp.NewReporter(zipkinURL)
	defer reporter.Close()

	// 创建Zipkin Tracer
	endpoint, _ := zipkin.NewEndpoint(Name, "localhost:8081")
	tracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint)) // 使用 reporter
	if err != nil {
		log.Fatal("Unable to create Zipkin tracer: ", err)
	}

	// _, _ = utils.CreatePathDir(bc.Log.Path)

	logger_ := log.With(
		logger.NewLogger(logger.Config{
			Path:         bc.Log.Path,
			Level:        bc.Log.Level,
			Rotationtime: bc.Log.Rotationtime.AsDuration(),
			Maxage:       bc.Log.Maxage.AsDuration(),
			OpenStat:     true,
		}),
		"x_time", log.DefaultTimestamp,
		"x_caller", log.DefaultCaller,
		"x_server_name", Name,
		"x_server_version", Version,
		"x_server_ip", id,
		"x_source", logger.Caller(3),
		"x_trace_id", loggerx.TraceID("x-trace-id"),
		"x_rpc_id", loggerx.IncrRpcId("x-rpc-id"),
		"x_zipkin_trace", tracer,
	)
	r := gin.Default()
	router.RegisterRouter(r, bc.Config, nil, nil, tracer, esClient)

	app, cleanup, err := wireApp(bc.Server, bc.Config, bc.Data, logger_)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
