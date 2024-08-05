package storage

import (
	"database/sql"
	"log"

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
