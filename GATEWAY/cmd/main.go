package main

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"

	casbin "github.com/casbin/casbin/v2"

	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/config"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/http/app"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/http/handler"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/redisservice"
	redisCl "github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/pkg/redis"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	redis, err := redisCl.NewRedisDB(config)
	if err != nil {
		log.Fatal(err)
	}

	modelPath := filepath.Join("internal", "items", "casbin", "model.conf")
	policyPath := filepath.Join("internal", "items", "casbin", "policy.csv")

	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		log.Fatal(err)
	}

	handler := handler.New(redisservice.New(redis, logger), logger, config)

	log.Fatal(app.Run(handler, logger, config, enforcer))
}
