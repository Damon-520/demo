package server

import (
	"demoapi/internal/conf"
	"demoapi/internal/pkg/middlewares/auth"
	"demoapi/internal/pkg/middlewares/trace"
	"demoapi/internal/router"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
)

func NewGinHttpServer(c *conf.Server, logger log.Logger, r router.GinRouter) *http.Server {

	var opts []http.ServerOption
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", (*gin.Engine)(r))

	return srv
}

func NewGinHttpRouter(logger log.Logger, conf *conf.Config) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	router_ := gin.Default()

	router_.Use(
		// middlewareAtuh.Cross(),
		// middlewareAtuh.CheckLoginMiddleware(conf.Env),
		kgin.Middlewares(
			recovery.Recovery(),
			metadata.Server(metadata.WithPropagatedPrefix("x-")),
			trace.WrapTraceIdForCtx("x-traceid"),
			// middleware.WrapRpcIdForCtxServer("x-rpcid"),
			trace.WrapStressTestingForCtx(),
			// middleware.GinLogServer(logger),
			// middleware.DeadlineServer("x-timeout"),
			auth.ParseHeader(logger),
		))

	return router_
}
