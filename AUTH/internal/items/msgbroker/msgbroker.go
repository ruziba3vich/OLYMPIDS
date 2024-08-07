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
	genprotos "github.com/ruziba3vich/OLYMPIDS/AUTH/genproto/auth"
	"github.com/ruziba3vich/OLYMPIDS/AUTH/internal/items/service"
	"google.golang.org/protobuf/proto"
)

type (
	MsgBroker struct {
		service          *service.Service
		channel          *amqp.Channel
		registrations    <-chan amqp.Delivery
		userUpdates      <-chan amqp.Delivery
		userDeletions    <-chan amqp.Delivery
		setAdmin         <-chan amqp.Delivery
		logger           *slog.Logger
		wg               *sync.WaitGroup
		numberOfServices int
	}
)

func New(service *service.Service,
	channel *amqp.Channel,
	logger *slog.Logger,
	registrations <-chan amqp.Delivery,
	userUpdates <-chan amqp.Delivery,
	userDeletions <-chan amqp.Delivery,
	setAdmin <-chan amqp.Delivery,
	wg *sync.WaitGroup,
	numberOfServices int) *MsgBroker {
	return &MsgBroker{
		service:          service,
		channel:          channel,
		registrations:    registrations,
		userUpdates:      userUpdates,
		userDeletions:    userDeletions,
		logger:           logger,
		wg:               wg,
		numberOfServices: numberOfServices,
	}
}

func (m *MsgBroker) StartToConsume(ctx context.Context, contentType string) {
	m.wg.Add(m.numberOfServices)
	consumerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go m.consumeMessages(consumerCtx, m.registrations, "create")
	go m.consumeMessages(consumerCtx, m.userUpdates, "update")
	go m.consumeMessages(consumerCtx, m.userDeletions, "delete")
	go m.consumeMessages(consumerCtx, m.setAdmin, "setadmin")

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
			case "register":
				var req genprotos.RegisterRequest
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					continue
				}
				response, err = m.service.Register(ctx, &req)
			case "setadmin":
				var req genprotos.CreateAdminRequest
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					continue
				}
				response, err = m.service.CreateAdmin(ctx, &req)
			case "update":
				var req genprotos.UpdateUserRequest
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					continue
				}
				response, err = m.service.UpdateUser(ctx, &req)
			case "delete":
				var req genprotos.DeleteUserRequest
				if err := json.Unmarshal(val.Body, &req); err != nil {
					m.logger.Error("Error while unmarshaling data", "error", err)
					val.Nack(false, false)
					continue
				}
				response, err = m.service.DeleteUser(ctx, &req)
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
