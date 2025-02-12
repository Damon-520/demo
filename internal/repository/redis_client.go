package repository

import (
	"demoapi/internal/conf"
	"demoapi/internal/core/redisx"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

// type RdbClient redis.Client

// func NewAnchorRedisClient(conf *conf.Data, logger log.Logger) *RdbClient {

// 	client := redisx.NewRedis(conf.AnchorRedis, logger)

// 	return (*RdbClient)(client)
// }

type ApiRdbClient redis.Client

func NewApiRedisClient(conf *conf.Data, logger log.Logger) *ApiRdbClient {

	client := redisx.NewRedis(conf.RedisForApi, logger)

	return (*ApiRdbClient)(client)
}
