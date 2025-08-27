package redisx

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"volunteer-team/backend/internal/infrastructure/configs"
)

func InitRedis() redis.Cmdable {
	if !configs.Conf.Redis.Enable {
		return nil
	}
	client := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               configs.Conf.Redis.DSN(),
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           configs.Conf.Redis.Password,
		DB:                 configs.Conf.Redis.DB,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolFIFO:           false,
		PoolSize:           1000,
		MinIdleConns:       1,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		panic(fmt.Sprintf("redis init error:%v", err))
	}
	return client
}
