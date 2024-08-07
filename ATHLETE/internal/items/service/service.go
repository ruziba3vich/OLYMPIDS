package service

import (
	"context"
	"log/slog"

	pb "github.com/ruziba3vich/OLYMPIDS/ATHLETE/genproto/athlete"

	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/repository"
)

type (
	Service struct {
		pb.UnimplementedAthleteServiceServer
		storage repository.IAthleteRepo
		logger  *slog.Logger
	}
)

func New(storage repository.IAthleteRepo, logger *slog.Logger) *Service {
	return &Service{
		storage: storage,
		logger: logger,
	}
}

func (s *Service) CreateAthlete(ctx context.Context, in *pb.CreateAthleteRequest) (*pb.Athlete, error) {
	s.logger.Info("Received request to create athlete")
	return s.storage.CreateAthlete(ctx, in)
}
func (s *Service) GetAthlete(ctx context.Context, in *pb.GetAthleteRequest) (*pb.Athlete, error) {
	s.logger.Info("Received request to get athlete")
	return s.storage.GetAthlete(ctx, in)
}
func (s *Service) UpdateAthlete(ctx context.Context, in *pb.UpdateAthleteRequest) (*pb.Athlete, error) {
	s.logger.Info("Received request to update athlete")
	return s.storage.UpdateAthlete(ctx, in)
}
func (s *Service) DeleteAthlete(ctx context.Context, in *pb.DeleteAthleteRequest) (*pb.Athlete, error) {
	s.logger.Info("Received request to delete athlete")
	return s.storage.DeleteAthlete(ctx, in)
}
func (s *Service) GetAthleteBySport(ctx context.Context, in *pb.GetAthleteBySportRequest) (*pb.GetAthleteResponse, error) {
	s.logger.Info("Received request to get athlete by sport")
	return s.storage.GetAthleteBySport(ctx, in)
}
