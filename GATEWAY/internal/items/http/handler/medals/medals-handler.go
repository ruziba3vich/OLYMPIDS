package medals

import (
	"log/slog"

	"github.com/gin-gonic/gin"
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

func (h *MedalsHandler) Ranking(c *gin.Context) {

}

// admin
func (h *MedalsHandler) GetAllMedals(c *gin.Context) {

}

// admin
func (h *MedalsHandler) UpdatesMedalByID(c *gin.Context) {

}

// admin
func (h *MedalsHandler) DeleteMedalByID(c *gin.Context) {

}
