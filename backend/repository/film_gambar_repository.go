package repository

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilmGambarRepository interface {
	Save(ctx context.Context, tx *gorm.DB, filmGambar entity.FilmGambar) error
	FindFilmGambarByFilmIDs(ctx context.Context, filmIDs []uuid.UUID) ([]dto.FilmGambarResponse, error)
}

type filmGambarRepository struct {
	db *gorm.DB
}

func NewFilmGambarRepository(db *gorm.DB) FilmGambarRepository {
	return &filmGambarRepository{db: db}
}

func (r *filmGambarRepository) Save(ctx context.Context, tx *gorm.DB, filmGambar entity.FilmGambar) error {
	if err := tx.Create(&filmGambar).Error; err != nil {
		return err
	}
	return nil
}

func (r *filmGambarRepository) FindFilmGambarByFilmIDs(ctx context.Context, filmIDs []uuid.UUID) ([]dto.FilmGambarResponse, error) {
	var gambar []dto.FilmGambarResponse
	err := r.db.WithContext(ctx).
		Model(entity.FilmGambar{}).
		Where("film_id IN ?", filmIDs).
		Find(&gambar).Error
	if err != nil {
		return nil, err
	}

	return gambar, nil
}
