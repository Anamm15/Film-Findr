package entity

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Rating    int       `gorm:"type:int;not null"`
	Komentar  string    `gorm:"type:text;not null"`
	UserID    uuid.UUID `gorm:"not null index:idx_user_id"`
	FilmID    uuid.UUID `gorm:"not null index:idx_film_id"`
	User      User      `gorm:"foreignKey:UserID"`
	Film      Film      `gorm:"foreignKey:FilmID"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
