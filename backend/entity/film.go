package entity

import (
	"time"

	"github.com/google/uuid"
)

type Film struct {
	ID           uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Judul        string       `gorm:"type:varchar(255);not null; index:idx_film_judul"`
	Status       string       `gorm:"type:varchar(255);not null; index:idx_film_status"`
	Sinopsis     string       `gorm:"type:text;not null;"`
	Durasi       int          `gorm:"type:int;not null;"`
	TotalEpisode int          `gorm:"type:int;not null;"`
	Sutradara    string       `gorm:"type:varchar(255);not null;"`
	TanggalRilis time.Time    `gorm:"type:date; not null"`
	FilmGambar   []FilmGambar `gorm:"foreignKey:FilmID"`
	FilmGenre    []FilmGenre  `gorm:"foreignKey:FilmID"`
	Review       []Review     `gorm:"foreignKey:FilmID"`
	CreatedAt    time.Time    `gorm:"type:timestamptz;default:now()"`
	UpdatedAt    time.Time    `gorm:"type:timestamptz;default:now()"`
}
