package entity

import (
	"time"

	"github.com/google/uuid"
)

type ReaksiReview struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Reaksi    string    `gorm:"type:varchar(50);not null"`
	UserID    uuid.UUID `gorm:"type:uuid;not null; index:idx_user_review"`
	ReviewID  uuid.UUID `gorm:"type:uuid;not null; index:idx_user_review"`
	User      User      `gorm:"foreignKey:UserID"`
	Review    Review    `gorm:"foreignKey:ReviewID"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
