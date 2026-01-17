package entity

import (
	"time"

	"FilmFindr/helpers"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Nama     string `gorm:"type:varchar(100);not null"`
	Username string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Password string `gorm:"type:text;not null"`
	Role     string `gorm:"type:varchar(20);not null"`

	Bio         string `gorm:"type:text"`
	PhotoProfil string `gorm:"type:text"`

	UserFilms []UserFilm `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Reviews   []Review   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
