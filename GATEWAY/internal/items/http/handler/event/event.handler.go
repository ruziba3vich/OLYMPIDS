package event

import (
	"encoding/json"
	"log/slog"

	"github.com/gin-gonic/gin"
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
		event:    event,
		logger:   logger,
		redis:    redis,
		msbroker: msbroker,
	}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	h.logger.Info("CreateEvent")

	var req pb.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	body, err := json.Marshal(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to marshal request"})
		return
	}

	err = h.msbroker.CreateEvent(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to publish message"})
		return
	}

	c.JSON(200, gin.H{"message": "Event created successfully"})
}

func (h *EventHandler) UpdateEvent(c *gin.Context) {
	h.logger.Info("UpdateEvent")
	var req pb.UpdateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	body, err := json.Marshal(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to marshal request"})
		return
	}

	err = h.msbroker.UpdateEvent(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to publish message"})
		return
	}

	c.JSON(200, gin.H{"message": "Event updated successfully"})
}

func (h *EventHandler) DeleteEvent(c *gin.Context) {
	h.logger.Info("DeleteEvent")

	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Missing event ID"})
		return
	}

	body, err := json.Marshal(&pb.DeleteEventRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to marshal request"})
		return
	}

	err = h.msbroker.DeleteEvent(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to publish message"})
		return
	}

	c.JSON(200, gin.H{"message": "Event deleted successfully"})
}

func (h *EventHandler) GetEvent(c *gin.Context) {
	h.logger.Info("GetEvent")
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{"error": "Missing event ID"})
		return
	}

	event, err := h.event.GetEvent(c, &pb.GetEventRequest{Id: id})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get event"})
		return
	}

	c.JSON(200, event)
}

func (h *EventHandler) GetEventBySport(c *gin.Context) {
	h.logger.Info("GetAllEvents")

	var req pb.GetEventBySportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	events, err := h.event.GetEventBySport(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get events"})
		return
	}

	c.JSON(200, events)
}