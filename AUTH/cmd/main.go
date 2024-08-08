package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"sync"

	sq "github.com/Masterminds/squirrel"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ruziba3vich/OLYMPIDS/AUTH/api"
	"github.com/ruziba3vich/OLYMPIDS/AUTH/internal/items/config"
	"github.com/ruziba3vich/OLYMPIDS/AUTH/internal/items/msgbroker"
	"github.com/ruziba3vich/OLYMPIDS/AUTH/internal/items/redisservice"
	"github.com/ruziba3vich/OLYMPIDS/AUTH/internal/items/service"
	"github.com/ruziba3vich/OLYMPIDS/AUTH/internal/items/storage"
	redisCl "github.com/ruziba3vich/OLYMPIDS/AUTH/internal/pkg/redis"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	ctx := context.Background()

	db, err := storage.ConnectDB(config)
	if err != nil {
		logger.Error("error while connecting postgres:", slog.String("err:", err.Error()))
	}

	redis, err := redisCl.NewRedisDB(config)
	if err != nil {
		logger.Error("error while connecting redis:", slog.String("err:", err.Error()))
	}

	sqrl := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	service := service.New(storage.New(
		redisservice.New(redis, logger),
		db,
		sqrl,
		config,
		logger,
	), logger)

	conn, err := amqp.Dial(config.RabbitMQ.RabbitMQ)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	regQueue, err := getQueue(ch, "register_user")
	if err != nil {
		log.Fatal(err)
	}
	regMsgs, err := getMessageQueue(ch, regQueue)
	if err != nil {
		log.Fatal(err)
	}

	updQueue, err := getQueue(ch, "update_user")
	if err != nil {
		log.Fatal(err)
	}
	updMsgs, err := getMessageQueue(ch, updQueue)
	if err != nil {
		log.Fatal(err)
	}

	delQueue, err := getQueue(ch, "delete_user")
	if err != nil {
		log.Fatal(err)
	}
	delMsgs, err := getMessageQueue(ch, delQueue)
	if err != nil {
		log.Fatal(err)
	}

	setQueue, err := getQueue(ch, "set_admin")
	if err != nil {
		log.Fatal(err)
	}
	setMsgs, err := getMessageQueue(ch, setQueue)
	if err != nil {
		log.Fatal(err)
	}

	msgBroker := msgbroker.New(service, ch, logger, regMsgs, updMsgs, delMsgs, setMsgs, &sync.WaitGroup{}, 3)

	api := api.New(service)

	go func() {
		log.Fatalln(api.RUN(config, service))
	}()

	msgBroker.StartToConsume(ctx, "application/json")
}

func getQueue(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
}

func getMessageQueue(ch *amqp.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
}
