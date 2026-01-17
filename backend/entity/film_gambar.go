package entity

import (
	"time"

	"github.com/google/uuid"
)

type FilmGambar struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Url       string    `gorm:"type:text;not null"`
	FilmID    uuid.UUID `gorm:"type:uuid;not null; index:idx_film_id"`
	Film      Film      `gorm:"foreignKey:FilmID"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
