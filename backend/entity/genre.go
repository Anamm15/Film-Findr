package entity

import (
	"time"

	"github.com/google/uuid"
)

type Genre struct {
	ID        uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Nama      string      `gorm:"type:varchar(255);not null;"`
	FilmGenre []FilmGenre `json:"film_genre" gorm:"foreignKey:GenreID"`
	CreatedAt time.Time   `gorm:"type:timestamptz;default:now()"`
	UpdatedAt time.Time   `gorm:"type:timestamptz;default:now()"`
}
