package repository

import (
	"context"

	pb "github.com/ruziba3vich/OLYMPIDS/EVENT/genproto/event"
)

type (
	IEventRepo interface {
		CreateEvent(ctx context.Context, in *pb.CreateEventRequest) (*pb.Event, error)
		GetEvent(ctx context.Context, in *pb.GetEventRequest) (*pb.Event, error)
		UpdateEvent(ctx context.Context, in *pb.UpdateEventRequest) (*pb.Event, error)
		DeleteEvent(ctx context.Context, in *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error)
		GetEventBySport(ctx context.Context, in *pb.GetEventBySportRequest) (*pb.GetEventBySportResponse, error)
	}
)
