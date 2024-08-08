package msgbroker

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
	genprotos "github.com/ruziba3vich/OLYMPIDS/EVENT/genproto/event"
	"github.com/ruziba3vich/OLYMPIDS/EVENT/internal/items/service"
	"google.golang.org/protobuf/proto"
)

type (
	MsgBroker struct {
		service          *service.Service
		channel          *amqp.Channel
		eventCreate      <-chan amqp.Delivery
		eventUpdates     <-chan amqp.Delivery
		eventDeletions   <-chan amqp.Delivery
		logger           *slog.Logger
		wg               *sync.WaitGroup
		numberOfServices int
	}
)

func New(service *service.Service,
	channel *amqp.Channel,
	logger *slog.Logger,
	eventCreate <-chan amqp.Delivery,
	eventUpdates <-chan amqp.Delivery,
	eventDeletions <-chan amqp.Delivery,
	wg *sync.WaitGroup,
	numberOfServices int) *MsgBroker {
	return &MsgBroker{
		service:          service,
		channel:          channel,
		eventCreate:      eventCreate,
		eventUpdates:     eventUpdates,
		eventDeletions:   eventDeletions,
		logger:           logger,
		wg:               wg,
		numberOfServices: numberOfServices,
	}
}

func (m *MsgBroker) StartToConsume(ctx context.Context, contentType string) {
	m.wg.Add(m.numberOfServices)
	consumerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go m.consumeMessages(consumerCtx, m.eventCreate, "create_event")
	go m.consumeMessages(consumerCtx, m.eventUpdates, "update_event")
	go m.consumeMessages(consumerCtx, m.eventDeletions, "delete_event")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	m.logger.Info("Shutting down, waiting for consumers to finish")
	cancel()
	m.wg.Wait()
	m.logger.Info("All consumers have stopped")
}

func (m *MsgBroker) consumeMessages(ctx context.Context, messages <-chan amqp.Delivery, logPrefix string) {
	defer m.wg.Done()
	for {
		select {
		case val := <-messages:
			var response proto.Message
			var err error

			switch logPrefix {
			case "create_event":
				var req genprotos.CreateEventRequest
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					continue
				}
				response, err = m.service.CreateEvent(ctx, &req)
			case "update_event":
				var req genprotos.UpdateEventRequest
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					continue
				}
				response, err = m.service.UpdateEvent(ctx, &req)
			case "delete_event":
				var req genprotos.DeleteEventRequest
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					continue
				}
				response, err = m.service.DeleteEvent(ctx, &req)
			}

			if err != nil {
				m.logger.Error("Failed in %s: %s\n", logPrefix, err.Error())
				val.Nack(false, false)
				continue
			}

			val.Ack(false)

			_, err = proto.Marshal(response)
			if err != nil {
				m.logger.Error("Failed to marshal response", "error", err)
				continue
			}

		case <-ctx.Done():
			m.logger.Info("Context done, stopping consumer", "consumer", logPrefix)
			return
		}
	}
}
