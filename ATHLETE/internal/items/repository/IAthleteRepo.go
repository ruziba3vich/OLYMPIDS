package repository

import (
	"context"

	pb "github.com/ruziba3vich/OLYMPIDS/ATHLETE/genproto/athlete"
)

type (
	IAthleteRepo interface {
		CreateAthlete(ctx context.Context, in *pb.CreateAthleteRequest) (*pb.Athlete, error)
		GetAthlete(ctx context.Context, in *pb.GetAthleteRequest) (*pb.Athlete, error)
		UpdateAthlete(ctx context.Context, in *pb.UpdateAthleteRequest) (*pb.Athlete, error)
		DeleteAthlete(ctx context.Context, in *pb.DeleteAthleteRequest) (*pb.Athlete, error)
		GetAthleteBySport(ctx context.Context, in *pb.GetAthleteBySportRequest) (*pb.GetAthleteResponse, error)
	}
)
