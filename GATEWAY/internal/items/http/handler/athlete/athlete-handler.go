package athlete

import (
	"encoding/json"
	"log/slog"

	"github.com/gin-gonic/gin"
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

// @Summary Create a new athlete
// @Description Create a new athlete record
// @Tags Admin Athletes
// @Accept json
// @Produce json
// @Param athlete body pb.CreateAthleteRequest true "Athlete data"
// @Success 201 {object} gin.H 
// @Failure 400 {object} gin.H 
// @Failure 500 {object} gin.H 
// @Router /admin/athletes [post]
func (h *AthleteHandler) CreateAthleteHandler(c *gin.Context) {
	h.logger.Info("CreateAthleteHandler called")
	var req pb.CreateAthleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	body, err := json.Marshal(&req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := h.msgbroker.CreateAthlete(body); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(201, gin.H{"message": "Athlete created successfully"})
}

// @Summary Get an athlete by ID
// @Description Retrieve an athlete record by its ID
// @Tags Admin Athletes
// @Produce json
// @Param id path string true "Athlete ID"
// @Success 200 {object} pb.Athlete "Athlete record"
// @Failure 400 {object} gin.H 
// @Failure 500 {object} gin.H 
// @Router /admin/athletes/{id} [get]
func (h *AthleteHandler) GetAthleteHandler(c *gin.Context) {
	h.logger.Info("GetAthleteHandler called")

	id := c.Param("id")

	if id == "" {
		c.IndentedJSON(400, gin.H{"error": "Athlete ID is required"})
		return
	}

	resp, err := h.athlete.GetAthlete(c.Request.Context(), &pb.GetAthleteRequest{Id: id})
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, resp)
}

// @Summary Update an existing athlete
// @Description Update the details of an existing athlete record
// @Tags Admin Athletes
// @Accept json
// @Produce json
// @Param athlete body pb.UpdateAthleteRequest true "Athlete data"
// @Success 200 {object} gin.H 
// @Failure 400 {object} gin.H 
// @Failure 500 {object} gin.H 
// @Router /admin/athletes [put]
func (h *AthleteHandler) UpdateAthleteHandler(c *gin.Context) {
	h.logger.Info("UpdateAthleteHandler called")
	var req pb.UpdateAthleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	body, err := json.Marshal(&req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := h.msgbroker.UpdateAthlete(body); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "Athlete updated successfully"})
}

// @Summary Delete an athlete by ID
// @Description Delete an athlete record by its ID
// @Tags Admin Athletes
// @Produce json
// @Param id path string true "Athlete ID"
// @Success 200 {object} gin.H 
// @Failure 400 {object} gin.H 
// @Failure 500 {object} gin.H 
// @Router /admin/athletes/{id} [delete]
func (h *AthleteHandler) DeleteAthleteHandler(c *gin.Context) {
	h.logger.Info("DeleteAthleteHandler called")

	id := c.Param("id")

	if id == "" {
		c.IndentedJSON(400, gin.H{"error": "Athlete ID is required"})
		return
	}

	req := &pb.DeleteAthleteRequest{Id: id}

	body, err := json.Marshal(&req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := h.msgbroker.DeleteAthlete(body); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "Athlete deleted successfully"})
}
