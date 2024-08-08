package models

import (
	"time"

	"github.com/google/uuid"
)

type Medal struct {
	ID          uuid.UUID  `db:"id"`
	Description string     `db:"description"`
	AthleteID   uuid.UUID  `db:"athlete_id"`
	Type        string     `db:"type"`
	Country     string     `db:"country"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

type CreateMedal struct {
	Description string    `db:"description"`
	AthleteID   uuid.UUID `db:"athlete_id"`
	Type        string    `db:"type"`
	Country     string    `db:"country"`
}

type UpdateMedal struct {
	ID          uuid.UUID `db:"id"`
	Description string    `db:"description"`
	AthleteID   uuid.UUID `db:"athlete_id"`
	Type        string    `db:"type"`
	Country     string    `db:"country"`
}
