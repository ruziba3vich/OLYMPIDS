package handler

import (
	"log/slog"

	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/redisservice"
)

type (
	Handler struct {
		logger *slog.Logger
		redis  *redisservice.RedisService
	}
)

func New(redis *redisservice.RedisService, logger *slog.Logger) *Handler {
	return &Handler{
		logger: logger,
		redis:  redis,
	}
}
