package repository

import (
	"context"
	"math"

	"FilmFindr/entity"
	"FilmFindr/helpers"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserFilmRepository interface {
	GetUserFilmByUserId(ctx context.Context, userId uuid.UUID, offset int) ([]entity.UserFilm, int64, error)
	CreateUserFilm(ctx context.Context, userFilm entity.UserFilm) (entity.UserFilm, error)
	UpdateStatusUserFilm(ctx context.Context, userFilmId uuid.UUID, status string) error
	CheckUserFilm(ctx context.Context, userId uuid.UUID, filmId uuid.UUID) (bool, error)
}

type userFilmRepository struct {
	db *gorm.DB
}

func NewUserFilmRepository(db *gorm.DB) UserFilmRepository {
	return &userFilmRepository{db: db}
}

func (r *userFilmRepository) GetUserFilmByUserId(ctx context.Context, userId uuid.UUID, offset int) ([]entity.UserFilm, int64, error) {
	var userFilms []entity.UserFilm
	var userFilmsCount int64

	if err := r.db.WithContext(ctx).
		Model(&entity.UserFilm{}).
		Where("user_id = ?", userId).
		Count(&userFilmsCount).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Model(&entity.UserFilm{}).
		Select("id", "status", "user_id", "film_id").
		Where("user_id = ?", userId).
		Preload("Film", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "judul", "tanggal_rilis", "durasi", "status")
		}).
		Preload("Film.FilmGambar", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "url", "film_id")
		}).
		Preload("Film.FilmGenre.Genre", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "nama")
		}).
		Order("created_at DESC").
		Limit(helpers.LIMIT_FILM).
		Offset(offset).
		Find(&userFilms).Error; err != nil {
		return nil, 0, err
	}

	totalPage := math.Ceil(float64(userFilmsCount) / float64(helpers.LIMIT_FILM))
	return userFilms, int64(totalPage), nil
}

func (r *userFilmRepository) CreateUserFilm(ctx context.Context, userFilm entity.UserFilm) (entity.UserFilm, error) {
	if err := r.db.Create(&userFilm).Error; err != nil {
		return entity.UserFilm{}, err
	}

	return userFilm, nil
}

func (r *userFilmRepository) UpdateStatusUserFilm(ctx context.Context, userFilmId uuid.UUID, status string) error {
	if err := r.db.WithContext(ctx).Table("user_films").Where("id = ?", userFilmId).Update("status", status).Error; err != nil {
		return err
	}

	return nil
}

func (r *userFilmRepository) CheckUserFilm(ctx context.Context, userId uuid.UUID, filmId uuid.UUID) (bool, error) {
	var userFilm entity.UserFilm
	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND film_id = ?", userId, filmId).
		First(&userFilm).Error; err != nil {
		return false, err
	}

	return userFilm.ID != uuid.Nil, nil
}
