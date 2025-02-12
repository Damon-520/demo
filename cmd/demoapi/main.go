package main

import (
	"demoapi/internal/conf"
	"demoapi/internal/core/logger"
	"demoapi/internal/elasticsearch"
	"demoapi/internal/pkg/loggerx"
	"demoapi/internal/router"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
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

func initNacos() config_client.IConfigClient {
	// ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			Port:        8848,
			ContextPath: "/nacos",
			Scheme:      "http",
		},
	}

	// ClientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "public", // If you need to specify the namespace, fill in the ID of the namespace here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "nacos/log",
		CacheDir:            "nacos/cache",
		LogLevel:            "debug",
	}

	// Create config client
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		log.Fatalf("create config client failed: %v", err)
	}
	fmt.Println("Nacos client initialized successfully!")
	return configClient
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	// 初始化Nacos客户端
	configClient := initNacos()

	// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "example-data-id",
		Group:  "example-group",
	})

	if err != nil {
		log.Fatalf("get config failed: %v", err)
	}

	fmt.Printf("config content: %s", content)

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

	// 如果索引不存在则创建索引
	if err := elasticsearch.CreateIndex("test_index"); err != nil {
		log.Fatal("创建索引失败: ", err)
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
	// todo
	//可以通过设置环境变量或在代码中调用 gin.SetMode(gin.ReleaseMode) 来切换到 "release" 模式以用于生产环境。
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
