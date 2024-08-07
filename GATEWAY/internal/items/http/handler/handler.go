package handler

import (
	"log"
	"log/slog"

	athlete_pb "github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/athlete"
	auth_pb "github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/auth"
	event_pb "github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/event"
	medals_pb "github.com/ruziba3vich/OLYMPIDS/GATEWAY/genproto/medals"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/config"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/http/handler/athlete"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/http/handler/auth"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/http/handler/event"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/http/handler/medals"
	"github.com/ruziba3vich/OLYMPIDS/GATEWAY/internal/items/redisservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	Handler struct {
		AthleteRepo *athlete.AthleteHandler
		AuthRepo    *auth.AuthHandler
		EventRepo   *event.EventHandler
		MedalsRepo  *medals.MedalsHandler
	}
)

func connect(port string) *grpc.ClientConn {
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func New(redis *redisservice.RedisService, logger *slog.Logger, config *config.Config) *Handler {
	athleteClient := athlete.NewAthleteHandler(logger, athlete_pb.NewAthleteServiceClient(connect(config.Server.AthletePort)), redis)
	authClient := auth.NewAthleteHandler(logger, auth_pb.NewAuthServiceClient(connect(config.Server.AuthPort)), redis)
	eventClient := event.NewAthleteHandler(logger, event_pb.NewEventServiceClient(connect(config.Server.EventPort)), redis)
	medalsClient := medals.NewAthleteHandler(logger, medals_pb.NewMedalServiceClient(connect(config.Server.MedalPort)), redis)

	return &Handler{
		AthleteRepo: athleteClient,
		AuthRepo:    authClient,
		EventRepo:   eventClient,
		MedalsRepo:  medalsClient,
	}
}
