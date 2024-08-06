package pkg

import (
	"errors"
	configs "medal-service/config"

	"github.com/go-redis/redis/v8"
)

func NewRedisDB(cfg *configs.RedisConfigs) (*redis.Client, error) {
	if cfg == nil {
		return nil, errors.New("Redis configuration is nil")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       0,
	})
	return client, nil
}
