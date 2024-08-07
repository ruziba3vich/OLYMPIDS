package redisservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	pb "github.com/ruziba3vich/OLYMPIDS/ATHLETE/genproto/athlete"

	"github.com/go-redis/redis/v8"
)

type (
	RedisService struct {
		redisDb *redis.Client
		logger  *slog.Logger
	}
)

func New(redisDb *redis.Client, logger *slog.Logger) *RedisService {
	return &RedisService{
		logger:  logger,
		redisDb: redisDb,
	}
}

func (r *RedisService) StoreAthleteInRedis(ctx context.Context, athlete *pb.Athlete) (*pb.Athlete, error) {
	key := fmt.Sprintf("athlete:%s", athlete.Id)
	athleteJSON, err := json.Marshal(athlete)
	if err != nil {
		r.logger.Error("Error marshalling athlete:", slog.String("err: ", err.Error()))
		return nil, err
	}

	if err := r.redisDb.Set(ctx, key, athleteJSON, 10*time.Minute).Err(); err != nil {
		r.logger.Error("Error setting athlete in Redis:", slog.String("err: ", err.Error()))
		return nil, err
	}

	return athlete, nil
}

func (r *RedisService) GetAthleteFromRedis(ctx context.Context, id string) (*pb.Athlete, error) {
	key := fmt.Sprintf("athlete:%s", id)
	val, err := r.redisDb.Get(ctx, key).Result()
	if err == redis.Nil {
		// Athlete not found in cache
		return nil, nil
	} else if err != nil {
		r.logger.Error("Error getting athlete from Redis:", slog.String("err: ", err.Error()))
		return nil, err
	}

	var athlete pb.Athlete
	if err := json.Unmarshal([]byte(val), &athlete); err != nil {
		r.logger.Error("Error unmarshalling athlete:", slog.String("err: ", err.Error()))
		return nil, err
	}

	return &athlete, nil
}

func (r *RedisService) DeleteAthleteFromRedis(ctx context.Context, id string) error {
	key := fmt.Sprintf("athlete:%s", id)
	if err := r.redisDb.Del(ctx, key).Err(); err != nil {
		r.logger.Error("Error deleting athlete from Redis:", slog.String("err: ", err.Error()))
		return err
	}

	return nil
}
