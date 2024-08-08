// Package api API.
//
// @title Olympy API
// @version 1.0
// @description API Endpoints for LocalEats
// @termsOfService http://swagger.io/terms/
//
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
//
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host localhost:8080
// @BasePath /
// @schemes http https
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package app

import (
	"log/slog"

	"github.com/gin-contrib/cors"
	_ "github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/http/app/docs"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/middleware"

	casbin "github.com/casbin/casbin/v2"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/config"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/http/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(handler *handler.Handler, logger *slog.Logger, config *config.Config, enforcer *casbin.Enforcer) error {
	router := gin.Default()

	// CORS konfiguratsiyasi
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Swagger dokumentatsiyasi uchun
	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url, ginSwagger.PersistAuthorization(true)))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	superadmin := router.Group("superadmin")
	superadmin.Use(middleware.AuthzMiddleware("/superadmin", enforcer))
	{
		superadmin.POST("/login", handler.AuthRepo.SuperAdminLoginHandler)
		superadmin.POST("/logout", handler.AuthRepo.SuperAdminLogoutHandler)
		superadmin.POST("/createadmin", handler.AuthRepo.SuperAdminCreateAdminHandler)
	}

	user := router.Group("user")
	user.Use(middleware.AuthzMiddleware("/user", enforcer))
	{
		auth := user.Group("auth")
		{
			auth.POST("/register", handler.AuthRepo.RegisterHandler)
			auth.POST("/login", handler.AuthRepo.LoginHandler)
			auth.POST("/logout", handler.AuthRepo.LogoutHandler)
		}
	}

	admin := router.Group("admin")
	admin.Use(middleware.AuthzMiddleware("/admin", enforcer))
	{
		auth := admin.Group("auth")
		{
			auth.POST("/login", handler.AuthRepo.AdminLoginHandler)
			auth.POST("/logout", handler.AuthRepo.AdminLogoutHandler)
			auth.PUT("/update/:id", handler.AuthRepo.UpdateUserHandler)
			auth.DELETE("/delete/:id", handler.AuthRepo.DeleteUserHandler)
		}

		medals := admin.Group("medals")
		{
			medals.GET("/ranking", handler.MedalsRepo.Ranking)
			medals.GET("", handler.MedalsRepo.GetAllMedals)
			medals.PUT("/:id", handler.MedalsRepo.UpdatesMedalByID)
			medals.DELETE("/:id", handler.MedalsRepo.DeleteMedalByID)
		}

		athletes := admin.Group("athletes")
		{
			athletes.POST("/", handler.AthleteRepo.CreateAthleteHandler)
			athletes.GET("/:id", handler.AthleteRepo.GetAthleteHandler)
			athletes.PUT("/:id", handler.AthleteRepo.UpdateAthleteHandler)
			athletes.DELETE("/:id", handler.AthleteRepo.DeleteAthleteHandler)
		}

		events := admin.Group("events")
		{
			events.POST("/", handler.EventRepo.CreateEventHandler)
			events.GET("/:id", handler.EventRepo.GetEventHandler)
			events.PUT("/:id", handler.EventRepo.UpdateEventHandler)
			events.DELETE("/:id", handler.EventRepo.DeleteEventHandler)
			events.GET("/sport", handler.EventRepo.GetEventBySportHandler)
		}
	}

	return router.Run(config.Server.ServerPort)
}
