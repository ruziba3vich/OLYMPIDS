package medals

import (
	"context"
	"encoding/json"
	"log/slog"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/medals"
	pb "github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/medals"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/msgbroker/medal"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/redisservice"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	MedalsHandler struct {
		medals    pb.MedalServiceClient
		logger    *slog.Logger
		redis     *redisservice.RedisService
		msgbroker *medal.MedalMsgBroker
	}
)

func NewAthleteHandler(logger *slog.Logger, medals pb.MedalServiceClient, redis *redisservice.RedisService, msgbroker *medal.MedalMsgBroker) *MedalsHandler {
	return &MedalsHandler{
		medals:    medals,
		logger:    logger,
		redis:     redis,
		msgbroker: msgbroker,
	}
}

// @Summary Create a new medal
// @Description Create a new medal with the provided details
// @Tags Admin Medals
// @Accept json
// @Produce json
// @Param medal body pb.CreateMedalRequest true "Medal details"
// @Success 201 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /admin/medals/ [post]
func (h *MedalsHandler) CreateMedalHandler(c *gin.Context) {
	h.logger.Info("CreateMedalHandler")

	var req pb.CreateMedalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	body, err := json.Marshal(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to marshal request"})
		return
	}

	err = h.msgbroker.CreateMedal(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to publish message"})
		return
	}

	c.JSON(201, gin.H{"message": "Medal created successfully"})
}

// @Summary Get a medal by ID
// @Description Retrieve the details of a medal using its ID
// @Tags Admin Medals
// @Produce json
// @Param id path string true "Medal ID"
// @Success 200 {object} pb.GetMedalResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /admin/medals/{id} [get]
func (h *MedalsHandler) GetMedalHandler(c *gin.Context) {
	h.logger.Info("GetMedalHandler")

	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Missing medal ID"})
		return
	}

	medal, err := h.medals.GetMedal(c, &pb.GetMedalRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get medal"})
		return
	}

	c.JSON(200, medal)
}

// @Summary Update a medal
// @Description Update the details of an existing medal
// @Tags Admin Medals
// @Accept json
// @Produce json
// @Param medal body pb.UpdateMedalRequest true "Updated medal details"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /admin/medals/ [put]
func (h *MedalsHandler) UpdateMedalHandler(c *gin.Context) {
	h.logger.Info("UpdateMedalHandler")

	var req pb.UpdateMedalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	body, err := json.Marshal(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to marshal request"})
		return
	}

	err = h.msgbroker.UpdateMedal(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to publish message"})
		return
	}

	c.JSON(200, gin.H{"message": "Medal updated successfully"})
}

// @Summary Delete a medal by ID
// @Description Remove a medal using its ID
// @Tags Admin Medals
// @Produce json
// @Param id path string true "Medal ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /admin/medals/{id} [delete]
func (h *MedalsHandler) DeleteMedalHandler(c *gin.Context) {
	h.logger.Info("DeleteMedalHandler")

	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Missing medal ID"})
		return
	}

	body, err := json.Marshal(&pb.DeleteMedalRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to marshal request"})
		return
	}

	err = h.msgbroker.DeleteMedal(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to publish message"})
		return
	}
}

// @Summary Get Medals Rankings by Country
// @Description Get Medals Rankings by Country
// @Tags Admin Medals
// @Produce json
// @Param limit query string true "limit"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /admin/medals/country-ranking [get]
func (h *MedalsHandler) GetCountryRankings(c *gin.Context) {
	h.logger.Info("GetCountryRankings")

	limit := c.Query("limit")
	if limit == "" {
		limit = "10"
	}
	req := medals.GetRankingByCountryRequest{Limit: limit}

	medals, err := h.medals.RankingByCountry(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get medals rankings"})
		return
	}

	c.JSON(200, medals)
}

// @Summary Get Medals by Time range
// @Description Get Medals Rankings by Country
// @Tags Admin Medals
// @Produce json
// @Param start query string true "start_time"
// @Param end query string true "end_time"
// @Param limit query string true "limit"
// @Param page query string true "page"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /admin/medals/range [get]
func (h *MedalsHandler) GetMedalsByTimeRange(c *gin.Context) {
	h.logger.Info("GetMedalsByTimeRange")

	start := c.Query("start")
	end := c.Query("end")
	limit := c.Query("limit")
	page := c.Query("page")

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		limitInt = 10
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		pageInt = 1
	}

	sParsed, err := time.Parse(time.RFC3339, start)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid start time format"})
		return
	}
	eParsed, err := time.Parse(time.RFC3339, end)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid end time format"})
		return
	}

	startTime := timestamppb.New(sParsed)
	endTime := timestamppb.New(eParsed)
	if start == "" || end == "" {
		c.JSON(400, gin.H{"error": "Missing start or end time"})
		return
	}

	medals, err := h.medals.GetMedalsByTimeRange(context.Background(), &pb.GetMedalsByTimeRangeRequest{
		StartTime: startTime,
		EndTime:   endTime,
		Limit:     int32(limitInt),
		Page:      int32(pageInt),
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get medals by time range"})
		return
	}

	c.JSON(200, medals)
}
