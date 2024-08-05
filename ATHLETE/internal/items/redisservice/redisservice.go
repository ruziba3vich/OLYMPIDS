package redisservice

import (
	"log"

	"github.com/go-redis/redis/v8"
)

type (
	RedisService struct {
		redisDb *redis.Client
		logger  *log.Logger
	}
)

func New(redisDb *redis.Client, logger *log.Logger) *RedisService {
	return &RedisService{
		logger:  logger,
		redisDb: redisDb,
	}
}
