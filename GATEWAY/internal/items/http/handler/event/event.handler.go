package event

import (
	"log/slog"

	pb "github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/event"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/msgbroker/event"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/redisservice"
)

type (
	EventHandler struct {
		event    pb.EventServiceClient
		logger   *slog.Logger
		redis    *redisservice.RedisService
		msbroker *event.EventMsgBroker
	}
)

func NewEventHandler(logger *slog.Logger, event pb.EventServiceClient, redis *redisservice.RedisService, msbroker *event.EventMsgBroker) *EventHandler {
	return &EventHandler{
		event:  event,
		logger: logger,
		redis:  redis,
		msbroker: msbroker,
	}
}

// func (e *EventH/andler) Publish()
