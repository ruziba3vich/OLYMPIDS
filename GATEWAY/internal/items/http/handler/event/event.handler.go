package event

import (
	"log/slog"

	pb "github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/event"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/redisservice"
)

type (
	EventHandler struct {
		event  pb.EventServiceClient
		logger *slog.Logger
		redis  *redisservice.RedisService
	}
)

func NewAthleteHandler(logger *slog.Logger, event pb.EventServiceClient, redis *redisservice.RedisService) *EventHandler {
	return &EventHandler{
		event:  event,
		logger: logger,
		redis:  redis,
	}
}
