package redisservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	pb "github.com/ruziba3vich/OLYMPIDS/EVENT/genproto/event"

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

func (r *RedisService) StoreEventInRedis(ctx context.Context, event *pb.Event) (*pb.Event, error) {
	key := fmt.Sprintf("event:%s", event.Id)
	eventJSON, err := json.Marshal(event)
	if err != nil {
		r.logger.Error("Error marshalling event:", slog.String("err: ", err.Error()))
		return nil, err
	}

	if err := r.redisDb.Set(ctx, key, eventJSON, 10*time.Minute).Err(); err != nil {
		r.logger.Error("Error setting event in Redis:", slog.String("err: ", err.Error()))
		return nil, err
	}

	return event, nil
}

func (r *RedisService) GetEventFromRedis(ctx context.Context, id string) (*pb.Event, error) {
	key := fmt.Sprintf("event:%s", id)
	eventJSON, err := r.redisDb.Get(ctx, key).Result()
	if err == redis.Nil {
		// Event not found in Redis
		return nil, nil
	} else if err != nil {
		r.logger.Error("Error getting event from Redis:", slog.String("err: ", err.Error()))
		return nil, err
	}

	event := &pb.Event{}
	if err := json.Unmarshal([]byte(eventJSON), event); err != nil {
		r.logger.Error("Error unmarshalling event:", slog.String("err: ", err.Error()))
		return nil, err
	}

	return event, nil
}

func (r *RedisService) DeleteEventFromRedis(ctx context.Context, id string) error {
	key := fmt.Sprintf("event:%s", id)
	if err := r.redisDb.Del(ctx, key).Err(); err != nil {
		r.logger.Error("Error deleting event from Redis:", slog.String("err: ", err.Error()))
		return err
	}

	return nil
}

func (r *RedisService) GetEventById(ctx context.Context, id string) error {
	key := fmt.Sprintf("event:%s", id)
	if err := r.redisDb.Get(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
