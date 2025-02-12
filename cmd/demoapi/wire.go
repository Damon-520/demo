//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"demoapi/internal/conf"
	httpProvider "demoapi/internal/http/providers"
	pkg "demoapi/internal/pkg/providers"
	repoProvider "demoapi/internal/repository/providers"
	"demoapi/internal/router"
	"demoapi/internal/server"
	serviceProvider "demoapi/internal/service/providers"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Config, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		pkg.ProviderSet,
		server.ProviderSet,
		router.ProviderSet,
		repoProvider.ProviderSet,
		serviceProvider.ProviderSet,
		httpProvider.ProviderSet,
		newZipkinTracer,
		newElasticsearchClient,
		newApp))
}

// newElasticsearchClient creates a new Elasticsearch client.
func newElasticsearchClient(config *conf.Config) (*elasticsearch.Client, error) {
	esURL := "https://localhost:9200" // 从配置中获取 URL
	cfg := elasticsearch.Config{
		Addresses: []string{esURL},
		Username:  "elastic", // 配置账号
		Password:  "123456",  // 配置密码
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// newZipkinTracer creates a new Zipkin tracer.
func newZipkinTracer(config *conf.Config) (*zipkin.Tracer, func(), error) {
	zipkinURL := "http://localhost:9411/api/v2/spans"
	reporter := zipkinhttp.NewReporter(zipkinURL)
	endpoint, _ := zipkin.NewEndpoint("demo-api", "localhost:8081")
	tracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		reporter.Close()
	}
	return tracer, cleanup, nil
}
