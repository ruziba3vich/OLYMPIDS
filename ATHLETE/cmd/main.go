package main

import (
	"log"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/api"
	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/config"
	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/redisservice"
	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/service"
	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/storage"
	redisCl "github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/pkg/redis"
)

func main() {
	config, err := config.New()
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		logger.Fatalln(err)
	}

	db, err := storage.ConnectDB(config)
	if err != nil {
		logger.Fatalln(err)
	}

	redis, err := redisCl.NewRedisDB(config)
	if err != nil {
		logger.Fatalln(err)
	}

	sqrl := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	service := service.New(storage.New(
		redisservice.New(redis, logger),
		db,
		sqrl,
		config,
		logger,
	))

	api := api.New(service)
	logger.Fatalln(api.RUN(config, service))
}
