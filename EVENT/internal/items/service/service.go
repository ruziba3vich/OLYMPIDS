package service

import (
	"context"
	"log/slog"

	pb "github.com/ruziba3vich/OLYMPIDS/EVENT/genproto/event"

	"github.com/ruziba3vich/OLYMPIDS/EVENT/internal/items/repository"
)

type (
	Service struct {
		pb.UnimplementedEventServiceServer
		storage repository.IEventRepo
		logger  *slog.Logger
	}
)

func New(storage repository.IEventRepo, logger *slog.Logger) *Service {
	return &Service{
		storage: storage,
		logger:  logger,
	}
}

func (s *Service) CreateEvent(ctx context.Context, in *pb.CreateEventRequest) (*pb.Event, error) {
	s.logger.Info("Received request to create event")
	return s.storage.CreateEvent(ctx, in)
}

func (s *Service) GetEvent(ctx context.Context, in *pb.GetEventRequest) (*pb.Event, error) {
	s.logger.Info("Received request to get event")
	return s.storage.GetEvent(ctx, in)
}

func (s *Service) UpdateEvent(ctx context.Context, in *pb.UpdateEventRequest) (*pb.Event, error) {
	s.logger.Info("Received request to update event")
	return s.storage.UpdateEvent(ctx, in)
}

func (s *Service) DeleteEvent(ctx context.Context, in *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	s.logger.Info("Received request to delete event")
	return s.storage.DeleteEvent(ctx, in)
}

func (s *Service) GetEventBySport(ctx context.Context, in *pb.GetEventBySportRequest) (*pb.GetEventBySportResponse, error) {
	s.logger.Info("Received request to get event by sport")
	return s.storage.GetEventBySport(ctx, in)
}
