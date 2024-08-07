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

	auth := router.Group("auth")
	{
		auth.POST("/register", handler.AuthRepo.Rigister)
		auth.POST("/login", handler.AuthRepo.Login)
		auth.POST("/refresh", handler.AuthRepo.Refresh)
	}

	medals := router.Group("medals")
	{
		medals.GET("/ranking", handler.MedalsRepo.Ranking)
		medals.GET("", handler.MedalsRepo.GetAllMedals)
		medals.PUT("/:id", handler.MedalsRepo.UpdatesMedalByID)
		medals.DELETE("/:id", handler.MedalsRepo.DeleteMedalByID)
	}

	athletes := router.Group("athletes")
	{
		athletes.GET("", handler.AthleteRepo.GetAllAthletes)
		athletes.GET("/:id", handler.AthleteRepo.GetAthleteByID)
		athletes.POST("", handler.AthleteRepo.CreateAthlete)
		athletes.PUT("/:id", handler.AthleteRepo.UpdateAthleteByID)
		athletes.DELETE("/:id", handler.AthleteRepo.DeleteAthleteByID)
	}
	return router.Run(config.Server.ServerPort)
}
