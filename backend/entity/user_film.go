package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserFilm struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Status    string    `gorm:"type:varchar(50);not null"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index:idx_user_id"`
	FilmID    uuid.UUID `gorm:"type:uuid;not null;index:idx_film_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Film      Film      `gorm:"foreignKey:FilmID" json:"film"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
