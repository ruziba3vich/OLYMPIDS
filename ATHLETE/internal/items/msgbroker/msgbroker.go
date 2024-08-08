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
	genprotos "github.com/ruziba3vich/OLYMPIDS/ATHLETE/genproto/athlete"
	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/service"
	"google.golang.org/protobuf/proto"
)

type (
	MsgBroker struct {
		service          *service.Service
		channel          *amqp.Channel
		createAthlete    <-chan amqp.Delivery
		updateAthlete    <-chan amqp.Delivery
		deleteAthlete    <-chan amqp.Delivery
		logger           *slog.Logger
		wg               *sync.WaitGroup
		numberOfServices int
	}
)

func New(service *service.Service,
	channel *amqp.Channel,
	logger *slog.Logger,
	createAthlete <-chan amqp.Delivery,
	updateAthlete <-chan amqp.Delivery,
	deleteAthlete <-chan amqp.Delivery,
	wg *sync.WaitGroup,
	numberOfServices int) *MsgBroker {
	return &MsgBroker{
		service:          service,
		channel:          channel,
		createAthlete:    createAthlete,
		updateAthlete:    updateAthlete,
		deleteAthlete:    deleteAthlete,
		logger:           logger,
		wg:               wg,
		numberOfServices: numberOfServices,
	}
}

func (m *MsgBroker) StartToConsume(ctx context.Context, contentType string) {
	m.wg.Add(m.numberOfServices)
	consumerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go m.consumeMessages(consumerCtx, m.createAthlete, "create_athlete")
	go m.consumeMessages(consumerCtx, m.updateAthlete, "update_athlete")
	go m.consumeMessages(consumerCtx, m.deleteAthlete, "delete_athlete")

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
			case "create_athlete":
				var req genprotos.CreateAthleteRequest
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					continue
				}
				response, err = m.service.CreateAthlete(ctx, &req)
			case "update_athlete":
				var req genprotos.UpdateAthleteRequest
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					continue
				}
				response, err = m.service.UpdateAthlete(ctx, &req)
			case "delete_athlete":
				var req genprotos.DeleteAthleteRequest
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					continue
				}
				response, err = m.service.DeleteAthlete(ctx, &req)
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
