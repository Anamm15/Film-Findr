package entity

import (
	"time"

	"github.com/google/uuid"
)

type FilmGenre struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FilmID    uuid.UUID `gorm:"type:uuid;not null;"`
	GenreID   uuid.UUID `gorm:"type:uuid;not null;"`
	Film      Film      `gorm:"foreignKey:FilmID"`
	Genre     Genre     `gorm:"foreignKey:GenreID"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
