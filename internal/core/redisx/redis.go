package redisx

import (
	"demoapi/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

func NewRedis(conf *conf.Redis, logger log.Logger) *redis.Client {

	log_ := log.NewHelper(log.With(logger, "x_module", "data/new-redisx"))

	options := redis.Options{
		Addr:         conf.Address,
		Password:     conf.Password,      // no password set
		DB:           int(conf.Database), // use default DB
		DialTimeout:  conf.DialTimeout.AsDuration(),
		WriteTimeout: conf.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.ReadTimeout.AsDuration(),
	}
	rdb := redis.NewClient(&options)
	if rdb == nil {
		log_.Fatalf("failed opening connection to redisx")
	}

	return rdb
}
