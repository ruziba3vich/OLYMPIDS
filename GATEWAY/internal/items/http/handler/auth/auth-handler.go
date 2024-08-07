package auth

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	pb "github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/auth"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/msgbroker/auth"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/redisservice"
)

type (
	AuthHandler struct {
		auth      pb.AuthServiceClient
		logger    *slog.Logger
		redis     *redisservice.RedisService
		msgbroker *auth.AuthMsgBroker
	}
)

func NewAthleteHandler(logger *slog.Logger, auth pb.AuthServiceClient, redis *redisservice.RedisService, msgbroker *auth.AuthMsgBroker) *AuthHandler {
	return &AuthHandler{
		auth:      auth,
		logger:    logger,
		redis:     redis,
		msgbroker: msgbroker,
	}
}

func (h *AuthHandler) Rigister(c *gin.Context) {

}

func (h *AuthHandler) Login(c *gin.Context) {

}

func (h *AuthHandler) Refresh(c *gin.Context) {

}
