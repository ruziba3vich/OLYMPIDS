package main

import (
	"log"
	"log/slog"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/ruziba3vich/OLYMPIDS/EVENT/api"
	"github.com/ruziba3vich/OLYMPIDS/EVENT/internal/items/config"
	"github.com/ruziba3vich/OLYMPIDS/EVENT/internal/items/redisservice"
	"github.com/ruziba3vich/OLYMPIDS/EVENT/internal/items/service"
	"github.com/ruziba3vich/OLYMPIDS/EVENT/internal/items/storage"
	redisCl "github.com/ruziba3vich/OLYMPIDS/EVENT/internal/pkg/redis"
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

	api := api.New(service)
	log.Fatalln(api.RUN(config, service))
}
