package repositroy

import (
	"context"
	"medal-service/models"
	"time"
)

type MedalRepository interface {
	CreateMedal(ctx context.Context, medal *models.CreateMedal) (*models.Medal, error)
	GetMedalByID(ctx context.Context, id string) (*models.Medal, error)
	DeleteMedalByID(ctx context.Context, id string) error
	GetMedalsByTimeRange(ctx context.Context, startDate, endDate time.Time, page, limit int32) ([]*models.Medal, error)
	GetMedalsByCountry(ctx context.Context, country string) ([]*models.Medal, error)
	GetMedalsByAthlete(ctx context.Context, athlete string) ([]*models.Medal, error)
	UpdateMedal(ctx context.Context, medal *models.UpdateMedal) (*models.Medal, error)
	ListMedals(ctx context.Context, page, limit int32, sortOrder, typefilter string) ([]*models.Medal, error)
	// GetTopCountries(ctx context.Context, limit int) ([]*models.CountryRanking, error)
}
