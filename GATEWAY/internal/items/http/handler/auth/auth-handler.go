package auth

import (
	"log/slog"

	pb "github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/auth"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/redisservice"
)

type (
	AuthHandler struct {
		auth   pb.AuthServiceClient
		logger *slog.Logger
		redis  *redisservice.RedisService
	}
)

func NewAthleteHandler(logger *slog.Logger, auth pb.AuthServiceClient, redis *redisservice.RedisService) *AuthHandler {
	return &AuthHandler{
		auth:   auth,
		logger: logger,
		redis:  redis,
	}
}
