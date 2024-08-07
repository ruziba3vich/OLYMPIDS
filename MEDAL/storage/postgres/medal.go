package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"medal-service/models"
	"medal-service/repositroy"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type MedalPostgresImpl struct {
	db           *sqlx.DB
	queryBuilder sq.StatementBuilderType
	redis        *redis.Client
}

func NewMedalPostgres(queryBuilder sq.StatementBuilderType, db *sqlx.DB, redis *redis.Client) repositroy.MedalRepository {
	return &MedalPostgresImpl{
		queryBuilder: queryBuilder,
		db:           db,
		redis:        redis,
	}
}

func (p *MedalPostgresImpl) CreateMedal(ctx context.Context, medal *models.CreateMedal) (*models.Medal, error) {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := p.queryBuilder.Insert("medals").
		Columns("description", "athlete_id", "type").
		Values(medal.Description, medal.AthleteID, medal.Type).
		Suffix("RETURNING id, created_at, updated_at").
		RunWith(tx).
		QueryRowContext(ctx)

	var newMedal models.Medal
	err = query.Scan(&newMedal.ID, &newMedal.CreatedAt, &newMedal.UpdatedAt)
	if err != nil {
		return nil, err
	}
	newMedal.Description = medal.Description
	newMedal.AthleteID = medal.AthleteID
	newMedal.Type = medal.Type

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	cacheKey := "medal:" + newMedal.ID.String()
	medalJSON, err := json.Marshal(newMedal)
	if err != nil {
		return nil, err
	}
	p.redis.Set(ctx, cacheKey, medalJSON, time.Hour*24).Err()

	return &newMedal, nil
}

func (p *MedalPostgresImpl) GetMedalByID(ctx context.Context, id string) (*models.Medal, error) {
	cacheKey := "medal:" + id
	cachedMedal, err := p.redis.Get(ctx, cacheKey).Result()
	if err == nil && cachedMedal != "" {
		var medal models.Medal
		if err := json.Unmarshal([]byte(cachedMedal), &medal); err == nil {
			return &medal, nil
		}
	}

	query := p.queryBuilder.Select("id", "description", "athlete_id", "type", "created_at", "updated_at", "deleted_at").
		From("medals").
		Where(sq.And{sq.Eq{"id": id}, sq.Eq{"deleted_at": nil}}).
		RunWith(p.db).
		QueryRowContext(ctx)

	var medal models.Medal
	err = query.Scan(&medal.ID, &medal.Description, &medal.AthleteID, &medal.Type, &medal.CreatedAt, &medal.UpdatedAt, &medal.DeletedAt)
	if err != nil {
		return nil, err
	}

	medalJSON, err := json.Marshal(medal)
	if err != nil {
		return nil, err
	}
	p.redis.Set(ctx, cacheKey, medalJSON, time.Hour*24).Err()

	return &medal, nil
}

func (p *MedalPostgresImpl) DeleteMedalByID(ctx context.Context, id string) error {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := p.queryBuilder.Update("medals").
		Set("deleted_at", sq.Expr("CURRENT_TIMESTAMP")).
		Where(sq.Eq{"id": id}).
		RunWith(tx).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	cacheKey := "medal:" + id
	p.redis.Del(ctx, cacheKey).Err()

	return nil
}

func (p *MedalPostgresImpl) GetMedalsByTimeRange(ctx context.Context, startDate, endDate time.Time, page, limit int32) ([]*models.Medal, error) {
	offset := (page - 1) * limit
	rows, err := p.queryBuilder.Select("id", "description", "athlete_id", "type", "created_at", "updated_at", "deleted_at").
		From("medals").
		Where(sq.And{sq.GtOrEq{"created_at": startDate}, sq.LtOrEq{"created_at": endDate}, sq.Eq{"deleted_at": nil}}).
		Offset(uint64(offset)).
		Limit(uint64(limit)).
		RunWith(p.db).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medals []*models.Medal
	for rows.Next() {
		var medal models.Medal
		if err := rows.Scan(&medal.ID, &medal.Description, &medal.AthleteID, &medal.Type, &medal.CreatedAt, &medal.UpdatedAt, &medal.DeletedAt); err != nil {
			return nil, err
		}
		medals = append(medals, &medal)
	}
	return medals, nil
}

func (p *MedalPostgresImpl) GetMedalsByCountry(ctx context.Context, country string) ([]*models.Medal, error) {
	rows, err := p.queryBuilder.Select("id", "description", "athlete_id", "type", "created_at", "updated_at", "deleted_at").
		From("medals").
		Where(sq.And{sq.Eq{"country": country}, sq.Eq{"deleted_at": nil}}).
		RunWith(p.db).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medals []*models.Medal
	for rows.Next() {
		var medal models.Medal
		if err := rows.Scan(&medal.ID, &medal.Description, &medal.AthleteID, &medal.Type, &medal.CreatedAt, &medal.UpdatedAt, &medal.DeletedAt); err != nil {
			return nil, err
		}
		medals = append(medals, &medal)
	}
	return medals, nil
}

func (p *MedalPostgresImpl) GetMedalsByAthlete(ctx context.Context, athlete string) ([]*models.Medal, error) {
	rows, err := p.queryBuilder.Select("id", "description", "athlete_id", "type", "created_at", "updated_at", "deleted_at").
		From("medals").
		Where(sq.And{sq.Eq{"athlete_id": athlete}, sq.Eq{"deleted_at": nil}}).
		RunWith(p.db).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medals []*models.Medal
	for rows.Next() {
		var medal models.Medal
		if err := rows.Scan(&medal.ID, &medal.Description, &medal.AthleteID, &medal.Type, &medal.CreatedAt, &medal.UpdatedAt, &medal.DeletedAt); err != nil {
			return nil, err
		}
		medals = append(medals, &medal)
	}
	return medals, nil
}

func (p *MedalPostgresImpl) UpdateMedal(ctx context.Context, medal *models.UpdateMedal) (*models.Medal, error) {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := p.queryBuilder.Update("medals").
		Set("description", medal.Description).
		Set("athlete_id", medal.AthleteID).
		Set("type", medal.Type).
		Set("updated_at", sq.Expr("CURRENT_TIMESTAMP")).
		Where(sq.And{sq.Eq{"id": medal.ID}, sq.Eq{"deleted_at": nil}}).
		Suffix("RETURNING id, created_at, updated_at, deleted_at").
		RunWith(tx).
		QueryRowContext(ctx)

	var updatedMedal models.Medal
	err = query.Scan(&updatedMedal.ID, &updatedMedal.CreatedAt, &updatedMedal.UpdatedAt, &updatedMedal.DeletedAt)
	if err != nil {
		return nil, err
	}
	updatedMedal.Description = medal.Description
	updatedMedal.AthleteID = medal.AthleteID
	updatedMedal.Type = medal.Type

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// Cache the updated medal
	cacheKey := "medal:" + updatedMedal.ID.String()
	medalJSON, err := json.Marshal(updatedMedal)
	if err != nil {
		return nil, err
	}
	p.redis.Set(ctx, cacheKey, medalJSON, time.Hour*24).Err()

	return &updatedMedal, nil
}

func (p *MedalPostgresImpl) ListMedals(ctx context.Context, page, limit int32, sortOrder, typeFilter string) ([]*models.Medal, error) {
	if typeFilter == "" {
		typeFilter = "%"
	}
	if sortOrder == "" {
		sortOrder = "created_at DESC"
	} else if sortOrder != "created_at" && sortOrder != "updated_at" {
		return nil, fmt.Errorf("invalid sort order: %s", sortOrder)
	}
	log.Println(sortOrder, typeFilter, page, limit)
	sortOrder = fmt.Sprintf("%s", sortOrder)

	rows, err := p.queryBuilder.Select("id", "description", "athlete_id", "type", "created_at", "updated_at", "deleted_at").
		From("medals").
		Where(sq.And{sq.Eq{"deleted_at": nil}, sq.Like{"type": typeFilter}}).OrderBy(sortOrder).
		Limit(uint64(limit)).
		Offset(uint64((page - 1) * limit)).
		RunWith(p.db).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medals []*models.Medal
	for rows.Next() {
		var medal models.Medal
		if err := rows.Scan(&medal.ID, &medal.Description, &medal.AthleteID, &medal.Type, &medal.CreatedAt, &medal.UpdatedAt, &medal.DeletedAt); err != nil {
			return nil, err
		}
		medals = append(medals, &medal)
	}
	return medals, nil
}
