package redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/gq-leon/sport-backend/bootstrap"
)

var Client *redis.Client

func InitRedis(env *bootstrap.Env) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     env.RedisAddr,
		Password: env.RedisPassword,
		DB:       env.RedisDB,
		PoolSize: env.RedisPoolSize,
	})

	_, err := Client.Ping(context.Background()).Result()
	return err
}
