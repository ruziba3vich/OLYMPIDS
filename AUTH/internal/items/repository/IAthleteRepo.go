package repository

import (
	"context"

	pb "github.com/ruziba3vich/OLYMPIDS/AUTH/genproto/auth"
)

type (
	IAuthRepo interface {
		Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error)
		Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error)
		Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error)
		CreateAdmin(ctx context.Context, in *pb.CreateAdminRequest) (*pb.CreateAdminResponse, error)
	}
)
