package app

import (
	"log/slog"

	casbin "github.com/casbin/casbin/v2"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/config"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/http/handler"
)

func Run(handler *handler.Handler, logger *slog.Logger, config *config.Config, enforcer *casbin.Enforcer) error {
	router := gin.Default()
	return router.Run(config.Server.ServerPort)
}
