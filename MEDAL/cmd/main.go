package main

import (
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
	cnf, err := configs.New()
	if err := cnf.Load(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// fmt.Println(config.DatabaseConfig, config.RedisConfig)
	redisClient, err := redisConn.NewRedisDB(cnf)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	log.Infof("Connected to Redis: %s", cnf.Redis.Host+":"+cnf.Redis.Port)
	postgresConn, err := pkgPost.ConnectPostgres(cnf.Database)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer postgresConn.Close()
	log.Infof("Connected to PostgreSQL: %s", cnf.Database.Host+":"+cnf.Database.Port)

	sqrl := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	countryMedals := postgres.NewCountryMedals(sqrl, postgresConn, redisClient)
	// redisImp := redisstorage.NewMedalRedisImpl(redisClient)
	postgresImp := postgres.NewMedalPostgres(sqrl, postgresConn, redisClient, countryMedals)

	service := service.NewMedalService(postgresImp)

	listener, err := net.Listen("tcp", cnf.Server.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Infof("Medal gRPC server listening on: %s", cnf.Server.Port)
	gRPCServer := grpc.NewServer()

	defer gRPCServer.GracefulStop()

	pb.RegisterMedalServiceServer(gRPCServer, service)

	if err := gRPCServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Infof("Medal gRPC server stopped")
}
