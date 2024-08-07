package athlete

import (
	"log/slog"

	pb "github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/athlete"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/msgbroker/athlete"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/redisservice"
)

type (
	AthleteHandler struct {
		athlete   pb.AthleteServiceClient
		logger    *slog.Logger
		redis     *redisservice.RedisService
		msgbroker *athlete.AthleteMsgBroker
	}
)

func NewAthleteHandler(logger *slog.Logger, athlete pb.AthleteServiceClient, redis *redisservice.RedisService, msgbroker *athlete.AthleteMsgBroker) *AthleteHandler {
	return &AthleteHandler{
		athlete:   athlete,
		logger:    logger,
		redis:     redis,
		msgbroker: msgbroker,
	}
}
