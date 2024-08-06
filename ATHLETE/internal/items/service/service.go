package service

import (
	"context"

	pb "github.com/ruziba3vich/OLYMPIDS/ATHLETE/genproto/athlete"

	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/repository"
)

type (
	Service struct {
		pb.UnimplementedAthleteServiceServer
		storage repository.IAthleteRepo
	}
)

func New(storage repository.IAthleteRepo) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateAthlete(ctx context.Context, in *pb.CreateAthleteRequest) (*pb.Athlete, error) {
	return s.storage.CreateAthlete(ctx, in)
}
func (s *Service) GetAthlete(ctx context.Context, in *pb.GetAthleteRequest) (*pb.Athlete, error) {
	return s.storage.GetAthlete(ctx, in)
}
func (s *Service) UpdateAthlete(ctx context.Context, in *pb.UpdateAthleteRequest) (*pb.Athlete, error) {
	return s.storage.UpdateAthlete(ctx, in)
}
func (s *Service) DeleteAthlete(ctx context.Context, in *pb.DeleteAthleteRequest) (*pb.Athlete, error) {
	return s.storage.DeleteAthlete(ctx, in)
}
func (s *Service) GetAthleteBySport(ctx context.Context, in *pb.GetAthleteBySportRequest) (*pb.GetAthleteResponse, error) {
	return s.storage.GetAthleteBySport(ctx, in)
}
