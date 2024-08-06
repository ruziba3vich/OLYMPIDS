package api

import (
	"log"
	"net"

	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/config"
	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/service"

	pb "github.com/ruziba3vich/OLYMPIDS/ATHLETE/genproto/athlete"

	"google.golang.org/grpc"
)

type API struct {
	service *service.Service
}

func New(service *service.Service) *API {
	return &API{
		service: service,
	}
}

func (a *API) RUN(config *config.Config, service *service.Service) error {
	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()

	pb.RegisterAthleteServiceServer(serverRegisterer, service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
