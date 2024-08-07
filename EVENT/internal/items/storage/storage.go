package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	pb "github.com/ruziba3vich/OLYMPIDS/EVENT/genproto/event"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ruziba3vich/OLYMPIDS/EVENT/internal/items/config"
	"github.com/ruziba3vich/OLYMPIDS/EVENT/internal/items/redisservice"
	"github.com/ruziba3vich/OLYMPIDS/EVENT/internal/items/repository"
)

type Storage struct {
	redis        *redisservice.RedisService
	postgres     *sql.DB
	queryBuilder sq.StatementBuilderType
	cfg          *config.Config
	logger       *slog.Logger
}

func New(redis *redisservice.RedisService, postgres *sql.DB, queryBuilder sq.StatementBuilderType, cfg *config.Config, logger *slog.Logger) repository.IEventRepo {
	return &Storage{
		redis:        redis,
		postgres:     postgres,
		queryBuilder: queryBuilder,
		cfg:          cfg,
		logger:       logger,
	}
}

func (s *Storage) CreateEvent(ctx context.Context, in *pb.CreateEventRequest) (*pb.Event, error) {
	id := uuid.New().String()
	created_at := time.Now()

	tx, err := s.postgres.BeginTx(ctx, nil)
	if err != nil {
		s.logger.Error("Error while starting a transaction")
		return nil, err
	}
	defer tx.Rollback()

	query, args, err := s.queryBuilder.Insert("athletes").
		Columns(
			"id",
			"name",
			"sport_type",
			"start_time",
			"end_time",
			"location",
			"description",
			"created_at",
		).Values(
		id,
		in.Name,
		in.SportType,
		in.StartTime.AsTime(),
		in.EndTime.AsTime(),
		in.Location,
		in.Description,
		created_at,
	).ToSql()

	if err != nil {
		s.logger.Error("Error while building query")
		return nil, err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Error while executing query")
		return nil, err
	}

	event := &pb.Event{
		Id:          id,
		Name:        in.Name,
		SportType:   in.SportType,
		StartTime:   in.StartTime,
		EndTime:     in.EndTime,
		Location:    in.Location,
		Description: in.Description,
		CreatedAt:   timestamppb.New(created_at),
	}

	if _, err := s.redis.StoreEventInRedis(ctx, event); err != nil {
		s.logger.Error("Error while storing event in Redis")
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		s.logger.Error("Error while committing transaction:", slog.String("err: ", err.Error()))
		return nil, err
	}

	return event, nil
}

func (s *Storage) GetEvent(ctx context.Context, in *pb.GetEventRequest) (*pb.Event, error) {
	// Try to get event from Redis
	event, err := s.redis.GetEventFromRedis(ctx, in.Id)
	if err != nil {
		s.logger.Error("Error getting event from Redis:", slog.String("err: ", err.Error()))
		return nil, err
	}
	if event != nil {
		s.logger.Info("Event found in Redis")
		return event, nil
	}

	// If not in Redis, get from PostgreSQL
	query, args, err := s.queryBuilder.Select(
		"id",
		"name",
		"sport_type",
		"start_time",
		"end_time",
		"location",
		"description",
		"created_at",
	).
		From("events").
		Where("id = ?", in.Id).
		ToSql()
	if err != nil {
		s.logger.Error("Error generating SQL:", slog.String("err: ", err.Error()))
		return nil, err
	}

	row := s.postgres.QueryRow(query, args...)
	event = &pb.Event{}
	var startTime, endTime, createdAt time.Time
	err = row.Scan(
		&event.Id,
		&event.Name,
		&event.SportType,
		&startTime,
		&endTime,
		&event.Location,
		&event.Description,
		&createdAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.Error("Event not found")
			return nil, fmt.Errorf("event not found")
		}
		s.logger.Error("Error executing SQL:", slog.String("err: ", err.Error()))
		return nil, err
	}

	event.StartTime = timestamppb.New(startTime)
	event.EndTime = timestamppb.New(endTime)
	event.CreatedAt = timestamppb.New(createdAt)

	// Store event in Redis
	if _, err := s.redis.StoreEventInRedis(ctx, event); err != nil {
		s.logger.Error("Error storing event in Redis:", slog.String("err: ", err.Error()))
	}

	return event, nil
}

func (s *Storage) UpdateEvent(ctx context.Context, in *pb.UpdateEventRequest) (*pb.Event, error) {
	updatedAt := time.Now()
	tx, err := s.postgres.BeginTx(ctx, nil)
	if err != nil {
		s.logger.Error("Error while starting a transaction")
		return nil, err
	}
	defer tx.Rollback()

	queryBuilder := s.queryBuilder.Update("events")

	if len(in.Name) > 0 {
		queryBuilder = queryBuilder.Set("name", in.Name)
	}
	if len(in.SportType) > 0 {
		queryBuilder = queryBuilder.Set("sport_type", in.SportType)
	}
	if !in.StartTime.AsTime().IsZero() {
		queryBuilder = queryBuilder.Set("start_time", in.StartTime.AsTime())
	}
	if !in.EndTime.AsTime().IsZero() {
		queryBuilder = queryBuilder.Set("end_time", in.EndTime.AsTime())
	}
	if len(in.Location) > 0 {
		queryBuilder = queryBuilder.Set("location", in.Location)
	}
	if len(in.Description) > 0 {
		queryBuilder = queryBuilder.Set("description", in.Description)
	}

	queryBuilder = queryBuilder.Set("updated_at", updatedAt)
	queryBuilder = queryBuilder.Where("id = ?", in.Id)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		s.logger.Error("Error generating SQL:", slog.String("err: ", err.Error()))
		return nil, err
	}

	result, err := s.postgres.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Error("Error executing SQL:", slog.String("err: ", err.Error()))
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Error("Error getting rows affected:", slog.String("err: ", err.Error()))
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	event, err := s.GetEvent(ctx, &pb.GetEventRequest{Id: in.Id})
	if err != nil {
		return nil, err
	}

	// Update event in Redis
	if _, err := s.redis.StoreEventInRedis(ctx, event); err != nil {
		s.logger.Error("Error storing event in Redis:", slog.String("err: ", err.Error()))
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		s.logger.Error("Error while committing transaction:", slog.String("err: ", err.Error()))
		return nil, err
	}

	return event, nil
}

func (s *Storage) DeleteEvent(ctx context.Context, in *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	query, args, err := s.queryBuilder.Delete("events").
		Where("id = ?", in.Id).
		ToSql()
	if err != nil {
		s.logger.Error("Error generating SQL:", slog.String("err: ", err.Error()))
		return nil, err
	}

	_, err = s.postgres.Exec(query, args...)
	if err != nil {
		s.logger.Error("Error executing SQL:", slog.String("err: ", err.Error()))
		return nil, err
	}

	// Delete event from Redis
	if err := s.redis.DeleteEventFromRedis(ctx, in.Id); err != nil {
		s.logger.Error("Error deleting event from Redis:", slog.String("err: ", err.Error()))
		return nil, err
	}

	return &pb.DeleteEventResponse{Message: "Event deleted successfully"}, nil
}

func (s *Storage) GetEventBySport(ctx context.Context, in *pb.GetEventBySportRequest) (*pb.GetEventBySportResponse, error) {
	query, args, err := s.queryBuilder.Select(
		"id",
		"name",
		"sport_type",
		"start_time",
		"end_time",
		"location",
		"description",
		"created_at",
	).
		From("events").
		Where("sport_type = ?", in.Sport).
		Limit(uint64(in.Limit)).
		Offset(uint64((in.Page - 1) * in.Limit)).
		ToSql()
	if err != nil {
		s.logger.Error("Error generating SQL:", slog.String("err: ", err.Error()))
		return nil, err
	}

	rows, err := s.postgres.Query(query, args...)
	if err != nil {
		s.logger.Error("Error executing SQL:", slog.String("err: ", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var events []*pb.Event
	for rows.Next() {
		event := &pb.Event{}
		var startTime, endTime, createdAt time.Time
		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.SportType,
			&startTime,
			&endTime,
			&event.Location,
			&event.Description,
			&createdAt,
		)
		if err != nil {
			s.logger.Error("Error scanning row:", slog.String("err: ", err.Error()))
			return nil, err
		}

		event.StartTime = timestamppb.New(startTime)
		event.EndTime = timestamppb.New(endTime)
		event.CreatedAt = timestamppb.New(createdAt)
		events = append(events, event)

		// Store each event in Redis
		if _, err := s.redis.StoreEventInRedis(ctx, event); err != nil {
			s.logger.Error("Error storing event in Redis:", slog.String("err: ", err.Error()))
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		s.logger.Error("Error iterating over rows:", slog.String("err: ", err.Error()))
		return nil, err
	}

	return &pb.GetEventBySportResponse{Events: events}, nil
}
