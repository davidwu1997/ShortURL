package redis

import (
	"context"
	"fmt"
	"shortURL/config"

	"github.com/pkg/errors"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Client *redis.Client
}

func InitRedisClient(config *config.Config) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password:     config.Redis.Password,
		DB:           config.Redis.DB,
		PoolSize:     config.Redis.MaxPoolSize,
		MinIdleConns: config.Redis.MinIdleConns,
	})
	ctx, cancel := context.WithTimeout(context.Background(), config.Redis.DialTimeout)
	defer cancel()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.Wrap(err, "redis error:")
	}

	return &Cache{Client: client}, nil
}
