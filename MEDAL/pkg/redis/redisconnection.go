package pkg

import (
	"errors"
	configs "medal-service/config"

	"github.com/go-redis/redis/v8"
)

func NewRedisDB(cfg *configs.Config) (*redis.Client, error) {
	if cfg == nil {
		return nil, errors.New("Redis configuration is nil")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: "",
		DB:       0,
	})

	return client, nil
}
