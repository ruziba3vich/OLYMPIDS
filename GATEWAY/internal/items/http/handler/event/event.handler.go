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

// CreateEventHandler godoc
// @Summary Create a new event
// @Description Create a new event by providing the event details
// @Tags Admin Events
// @Accept  json
// @Produce  json
// @Param event body pb.CreateEventRequest true "Event details"
// @Success 200 {object} gin.H 
// @Failure 400 {object} gin.H 
// @Failure 500 {object} gin.H 
// @Router /admin/events/ [post]
func (h *EventHandler) CreateEventHandler(c *gin.Context) {
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

// UpdateEventHandler godoc
// @Summary Update an existing event
// @Description Update an event by providing the event ID and details
// @Tags Admin Events
// @Accept  json
// @Produce  json
// @Param id path string true "Event ID"
// @Param event body pb.UpdateEventRequest true "Event details"
// @Success 200 {object} gin.H 
// @Failure 400 {object} gin.H 
// @Failure 500 {object} gin.H 
// @Router /admin/events/{id} [put]
func (h *EventHandler) UpdateEventHandler(c *gin.Context) {
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

// DeleteEventHandler godoc
// @Summary Delete an event
// @Description Delete an event by providing the event ID
// @Tags Admin Events
// @Produce  json
// @Param id path string true "Event ID"
// @Success 200 {object} gin.H 
// @Failure 400 {object} gin.H 
// @Failure 500 {object} gin.H 
// @Router /admin/events/{id} [delete]
func (h *EventHandler) DeleteEventHandler(c *gin.Context) {
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

// GetEventHandler godoc
// @Summary Get an event by ID
// @Description Retrieve an event by providing the event ID
// @Tags Admin Events
// @Produce  json
// @Param id path string true "Event ID"
// @Success 200 {object} pb.Event "Event details"
// @Failure 400 {object} gin.H 
// @Failure 500 {object} gin.H 
// @Router /admin/events/{id} [get]
func (h *EventHandler) GetEventHandler(c *gin.Context) {
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

// GetEventBySportHandler godoc
// @Summary Get events by sport
// @Description Retrieve events based on the sport by providing sport details
// @Tags Admin Events
// @Accept  json
// @Produce  json
// @Param sport body pb.GetEventBySportRequest true "Sport details"
// @Success 200 {object} pb.GetEventBySportResponse "List of events"
// @Failure 400 {object} gin.H 
// @Failure 500 {object} gin.H 
// @Router /admin/events/sport [get]
func (h *EventHandler) GetEventBySportHandler(c *gin.Context) {
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