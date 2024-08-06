package main

import (
	"fmt"
	configs "medal-service/config"
	pb "medal-service/genproto/medals"
	pkgPost "medal-service/pkg/postgres"
	redisConn "medal-service/pkg/redis"
	"medal-service/service"
	"medal-service/storage/postgres"
	"net"

	sq "github.com/Masterminds/squirrel"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

func main() {
	config, err := configs.InitConfig(".") // Ensure this is correct
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Println(config.DatabaseConfig, config.RedisConfig)
	redisClient, err := redisConn.NewRedisDB(config.RedisConfig)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	log.Infof("Connected to Redis: %s", config.RedisConfig.Host+":"+config.RedisConfig.Port)
	postgresConn, err := pkgPost.ConnectPostgres(config.DatabaseConfig)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer postgresConn.Close()
	log.Infof("Connected to PostgreSQL: %s", config.DatabaseConfig.Host+":"+config.DatabaseConfig.Port)

	sqrl := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// redisImp := redisstorage.NewMedalRedisImpl(redisClient)
	postgresImp := postgres.NewMedalPostgres(sqrl, postgresConn, redisClient)

	service := service.NewMedalService(postgresImp)

	listener, err := net.Listen("tcp", config.Host+":"+config.GrpcServerPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Infof("Medal gRPC server listening on: %s", config.GrpcServerPort)
	gRPCServer := grpc.NewServer()

	defer gRPCServer.GracefulStop()

	pb.RegisterMedalServiceServer(gRPCServer, service)

	if err := gRPCServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Infof("Medal gRPC server stopped")
}
