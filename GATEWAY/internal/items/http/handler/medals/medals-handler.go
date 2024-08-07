package medals

import (
	"log/slog"

	pb "github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/medals"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/redisservice"
)

type (
	MedalsHandler struct {
		medals pb.MedalServiceClient
		logger *slog.Logger
		redis  *redisservice.RedisService
	}
)

func NewAthleteHandler(logger *slog.Logger, medals pb.MedalServiceClient, redis *redisservice.RedisService) *MedalsHandler {
	return &MedalsHandler{
		medals: medals,
		logger: logger,
		redis:  redis,
	}
}
