package redisservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pb "github.com/ruziba3vich/OLYMPIDS/ATHLETE/genproto/athlete"

	"github.com/go-redis/redis/v8"
)

type (
	RedisService struct {
		redisDb *redis.Client
		logger  *log.Logger
	}
)

func New(redisDb *redis.Client, logger *log.Logger) *RedisService {
	return &RedisService{
		logger:  logger,
		redisDb: redisDb,
	}
}

func (r *RedisService) StoreAthleteInRedis(ctx context.Context, athlete *pb.Athlete) (*pb.Athlete, error) {
	key := fmt.Sprintf("athlete:%s", athlete.Id)
	athleteJSON, err := json.Marshal(athlete)
	if err != nil {
		r.logger.Println("Error marshalling athlete:", err)
		return nil, err
	}

	if err := r.redisDb.Set(ctx, key, athleteJSON, 10*time.Minute).Err(); err != nil {
		r.logger.Println("Error setting athlete in Redis:", err)
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
		r.logger.Println("Error getting athlete from Redis:", err)
		return nil, err
	}

	var athlete pb.Athlete
	if err := json.Unmarshal([]byte(val), &athlete); err != nil {
		r.logger.Println("Error unmarshalling athlete:", err)
		return nil, err
	}

	return &athlete, nil
}

func (r *RedisService) DeleteAthleteFromRedis(ctx context.Context, id string) error {
	key := fmt.Sprintf("athlete:%s", id)
	if err := r.redisDb.Del(ctx, key).Err(); err != nil {
		r.logger.Println("Error deleting athlete from Redis:", err)
		return err
	}

	return nil
}
