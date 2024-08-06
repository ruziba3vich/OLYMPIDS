package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	pb "github.com/ruziba3vich/OLYMPIDS/ATHLETE/genproto/athlete"
	"google.golang.org/protobuf/types/known/timestamppb"

	sq "github.com/Masterminds/squirrel"
	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/config"
	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/redisservice"
	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/repository"
)

type Storage struct {
	redis        *redisservice.RedisService
	postgres     *sql.DB
	queryBuilder sq.StatementBuilderType
	cfg          *config.Config
	logger       *log.Logger
}

func New(redis *redisservice.RedisService, postgres *sql.DB, queryBuilder sq.StatementBuilderType, cfg *config.Config, logger *log.Logger) repository.IAthleteRepo {
	return &Storage{
		redis:        redis,
		postgres:     postgres,
		queryBuilder: queryBuilder,
		cfg:          cfg,
		logger:       logger,
	}
}

func (s *Storage) CreateAthlete(ctx context.Context, in *pb.CreateAthleteRequest) (*pb.Athlete, error) {
	id := uuid.New().String()
	createdAt := time.Now()
	query, args, err := s.queryBuilder.Insert("athletes").
		Columns(
			"id",
			"first_name",
			"last_name",
			"gender",
			"nationality",
			"height",
			"weight",
			"sport",
			"date_of_birth",
			"created_at",
		).Values(
		id,
		in.FirstName,
		in.LastName,
		in.Gender,
		in.Nationality,
		in.Height,
		in.Weight,
		in.Sport,
		in.DateOfBirth.AsTime(),
		createdAt,
	).ToSql()
	if err != nil {
		s.logger.Println("Error generating SQL:", err)
		return nil, err
	}

	_, err = s.postgres.Exec(query, args...)
	if err != nil {
		s.logger.Println("Error executing SQL:", err)
		return nil, err
	}

	athlete := &pb.Athlete{
		Id:          id,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Gender:      in.Gender,
		Nationality: in.Nationality,
		Height:      in.Height,
		Weight:      in.Weight,
		Sport:       in.Sport,
		CreatedAt:   timestamppb.New(createdAt),
	}

	// Store athlete in Redis
	if _, err := s.redis.StoreAthleteInRedis(ctx, athlete); err != nil {
		s.logger.Println("Error storing athlete in Redis:", err)
	}

	return athlete, nil
}

func (s *Storage) GetAthlete(ctx context.Context, in *pb.GetAthleteRequest) (*pb.Athlete, error) {
	// Try to get athlete from Redis
	athlete, err := s.redis.GetAthleteFromRedis(ctx, in.Id)
	if err != nil {
		s.logger.Println("Error getting athlete from Redis:", err)
		return nil, err
	}
	if athlete != nil {
		return athlete, nil
	}

	// If not in Redis, get from PostgreSQL
	query, args, err := s.queryBuilder.Select(
		"id",
		"first_name",
		"last_name",
		"gender",
		"nationality",
		"height",
		"weight",
		"sport",
		"created_at",
	).
		From("athletes").
		Where("id = ?", in.Id).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		s.logger.Println("Error generating SQL:", err)
		return nil, err
	}

	row := s.postgres.QueryRow(query, args...)
	athlete = &pb.Athlete{}
	var createdAt time.Time
	err = row.Scan(
		&athlete.Id,
		&athlete.FirstName,
		&athlete.LastName,
		&athlete.Gender,
		&athlete.Nationality,
		&athlete.Height,
		&athlete.Weight,
		&athlete.Sport,
		&createdAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("athlete not found")
		}
		s.logger.Println("Error scanning row:", err)
		return nil, err
	}

	athlete.CreatedAt = timestamppb.New(createdAt)

	// Store athlete in Redis
	if _, err := s.redis.StoreAthleteInRedis(ctx, athlete); err != nil {
		s.logger.Println("Error storing athlete in Redis:", err)
	}

	return athlete, nil
}

func (s *Storage) UpdateAthlete(ctx context.Context, in *pb.UpdateAthleteRequest) (*pb.Athlete, error) {
	updated_at := time.Now()
	tx, err := s.postgres.BeginTx(ctx, nil)
	if err != nil {
		s.logger.Println("Error while starting a transaction")
		return nil, err
	}
	defer tx.Rollback()

	queryBuilder := s.queryBuilder.Update("athletes")

	if len(in.FirstName) > 0 {
		queryBuilder = queryBuilder.Set("first_name", in.FirstName)
	}
	if len(in.LastName) > 0 {
		queryBuilder = queryBuilder.Set("last_name", in.LastName)
	}
	if len(in.Gender) > 0 {
		queryBuilder = queryBuilder.Set("gender", in.Gender)
	}
	if len(in.Nationality) > 0 {
		queryBuilder = queryBuilder.Set("nationality", in.Nationality)
	}
	if len(in.Height) > 0 {
		queryBuilder = queryBuilder.Set("height", in.Height)
	}
	if len(in.Height) > 0 {
		queryBuilder = queryBuilder.Set("weight", in.Weight)
	}
	if len(in.Sport) > 0 {
		queryBuilder = queryBuilder.Set("sport", in.Sport)
	}

	queryBuilder = queryBuilder.Set("updated_at", updated_at)
	queryBuilder = queryBuilder.Where("id = ?", in.Id).Where("deleted_at IS NULL")

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		s.logger.Println("Error generating SQL:", err)
		return nil, err
	}

	result, err := s.postgres.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.Println("Error executing SQL:", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	athlete, err := s.GetAthlete(ctx, &pb.GetAthleteRequest{Id: in.Id})
	if err != nil {
		return nil, err
	}

	// Update athlete in Redis
	if _, err := s.redis.StoreAthleteInRedis(ctx, athlete); err != nil {
		s.logger.Println("Error storing athlete in Redis:", err)
	}

	if err := tx.Commit(); err != nil {
		s.logger.Println("Error while committing transaction:", err.Error())
		return nil, err
	}

	return athlete, nil
}

func (s *Storage) DeleteAthlete(ctx context.Context, in *pb.DeleteAthleteRequest) (*pb.Athlete, error) {
	deleted_at := time.Now()
	athlete, err := s.GetAthlete(ctx, &pb.GetAthleteRequest{Id: in.Id})
	if err != nil {
		return nil, err
	}

	query, args, err := s.queryBuilder.Update("athletes").
		Where("id = ?", in.Id).
		Where("deleted_at IS NULL").
		Set("deleted_at", deleted_at).
		ToSql()
	if err != nil {
		s.logger.Println("Error generating SQL:", err)
		return nil, err
	}

	_, err = s.postgres.Exec(query, args...)
	if err != nil {
		s.logger.Println("Error executing SQL:", err)
		return nil, err
	}

	// Delete athlete from Redis
	if err := s.redis.DeleteAthleteFromRedis(ctx, in.Id); err != nil {
		s.logger.Println("Error deleting athlete from Redis:", err)
	}

	return athlete, nil
}

func (s *Storage) GetAthleteBySport(ctx context.Context, in *pb.GetAthleteBySportRequest) (*pb.GetAthleteResponse, error) {
	query, args, err := s.queryBuilder.Select(
		"id",
		"first_name",
		"last_name",
		"gender",
		"nationality",
		"height",
		"weight",
		"sport",
		"created_at",
	).
		From("athletes").
		Where("sport = ?", in.Sport).
		Limit(uint64(in.Limit)).
		Offset(uint64((in.Page - 1) * in.Limit)).
		ToSql()
	if err != nil {
		s.logger.Println("Error generating SQL:", err)
		return nil, err
	}

	rows, err := s.postgres.Query(query, args...)
	if err != nil {
		s.logger.Println("Error executing SQL:", err)
		return nil, err
	}
	defer rows.Close()

	var athletes []*pb.Athlete
	for rows.Next() {
		athlete := &pb.Athlete{}
		var createdAt time.Time
		err := rows.Scan(
			&athlete.Id,
			&athlete.FirstName,
			&athlete.LastName,
			&athlete.Gender,
			&athlete.Nationality,
			&athlete.Height,
			&athlete.Weight,
			&athlete.Sport,
			&createdAt,
		)
		if err != nil {
			s.logger.Println("Error scanning row:", err)
			return nil, err
		}

		athlete.CreatedAt = timestamppb.New(createdAt)
		athletes = append(athletes, athlete)

		// Store each athlete in Redis
		if _, err := s.redis.StoreAthleteInRedis(ctx, athlete); err != nil {
			s.logger.Println("Error storing athlete in Redis:", err)
		}
	}

	if err = rows.Err(); err != nil {
		s.logger.Println("Error iterating over rows:", err)
		return nil, err
	}

	return &pb.GetAthleteResponse{Athletes: athletes}, nil
}
