package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"medal-service/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type CountryMedals struct {
	db           *sqlx.DB
	queryBuilder sq.StatementBuilderType
	redis        *redis.Client
}

func NewCountryMedals(queryBuilder sq.StatementBuilderType, db *sqlx.DB, redis *redis.Client) *CountryMedals {
	return &CountryMedals{
		queryBuilder: queryBuilder,
		db:           db,
		redis:        redis,
	}
}

func (c *CountryMedals) GetCountryMedals(ctx context.Context, country string) (*models.CountryMedals, error) {
	query := c.queryBuilder.Select("name", "gold_count", "silver_count", "bronze_count").
		From("country_medals").
		Where(sq.Eq{"name": country}).
		RunWith(c.db)

	countryMedals := &models.CountryMedals{}
	err := query.QueryRowContext(ctx).Scan(&countryMedals.Name, &countryMedals.GoldCount, &countryMedals.SilverCount, &countryMedals.BronzeCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("country not found: %s", country)
		}
		return nil, err
	}

	return countryMedals, nil
}

func (c *CountryMedals) GetTopCountries(ctx context.Context, limit int) ([]*models.CountryRanking, error) {
	// Debugging: Check if queryBuilder is initialized
	// if c.queryBuilder == nil {
	// 	return nil, fmt.Errorf("queryBuilder is not initialized")
	// }

	// Debugging: Check if db is initialized
	if c.db == nil {
		return nil, fmt.Errorf("db is not initialized")
	}

	query := c.queryBuilder.Select("name",
		"gold_count",
		"silver_count",
		"bronze_count",
		"(gold_count + silver_count + bronze_count) AS total_medals").
		From("country_medals").
		OrderBy("gold_count DESC", "silver_count DESC", "bronze_count DESC").
		Limit(uint64(limit))

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %w", err)
	}

	log.Println("Executing query:", sqlQuery, "with args:", args)
	rows, err := c.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		log.Println("Error querying country")
		return nil, err
	}
	defer rows.Close()

	var countryRankings []*models.CountryRanking
	rank := 1
	for rows.Next() {
		cr := &models.CountryRanking{}
		if err := rows.Scan(&cr.Name, &cr.GoldCount, &cr.SilverCount, &cr.BronzeCount, &cr.TotalMedals); err != nil {
			return nil, err
		}
		cr.Rank = rank
		countryRankings = append(countryRankings, cr)
		rank++
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return countryRankings, nil
}

func (c *CountryMedals) CreateOrUpdate(ctx context.Context, country string, medalType string) (bool, error) {
	query := c.queryBuilder.Select("name").
		From("country_medals").
		Where(sq.Eq{"name": country}).
		RunWith(c.db).
		QueryRowContext(ctx)
	var name string
	err := query.Scan(&name)

	if err == nil {
		query := c.queryBuilder.Update("country_medals").
			Set(medalType+"_count", sq.Expr(medalType+"_count + 1")).
			Where(sq.Eq{"name": country}).
			RunWith(c.db)
		_, err := query.ExecContext(ctx)
		if err != nil {
			return false, err
		}
	} else if err == sql.ErrNoRows {
		query := c.queryBuilder.Insert("country_medals").
			Columns("name", medalType+"_count").
			Values(country, 1).
			RunWith(c.db)

		_, err = query.ExecContext(ctx)
		if err != nil {
			return false, err
		}
	} else {
		return false, err
	}

	countryMedals, err := c.GetCountryMedals(ctx, country)
	if err != nil {
		return false, err
	}

	cacheKey := fmt.Sprintf("country_medals:%s", country)
	cacheData, err := json.Marshal(countryMedals)
	if err == nil {
		c.redis.Set(ctx, cacheKey, cacheData, time.Minute*10)
	}

	return true, nil
}
