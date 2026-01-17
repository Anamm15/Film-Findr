package repository

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GenreRepository interface {
	GetAllGenre(ctx context.Context) ([]entity.Genre, error)
	CreateGenre(ctx context.Context, genre entity.Genre) (entity.Genre, error)
	DeleteGenre(ctx context.Context, genreId uuid.UUID) error
	GetGenreListAndCount(ctx context.Context) ([]dto.GenreListAndCountResponse, error)
}

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) GenreRepository {
	return &genreRepository{db: db}
}

func (r *genreRepository) GetAllGenre(ctx context.Context) ([]entity.Genre, error) {
	var genres []entity.Genre
	if err := r.db.WithContext(ctx).Select("id", "nama").
		Find(&genres).Error; err != nil {
		return nil, err
	}

	return genres, nil
}

func (r *genreRepository) CreateGenre(ctx context.Context, genre entity.Genre) (entity.Genre, error) {
	err := r.db.WithContext(ctx).Create(&genre).Error
	return genre, err
}

func (r *genreRepository) DeleteGenre(ctx context.Context, genreId uuid.UUID) error {
	err := r.db.WithContext(ctx).Delete(&entity.Genre{}, genreId).Error
	return err
}

func (r *genreRepository) GetGenreListAndCount(ctx context.Context) ([]dto.GenreListAndCountResponse, error) {
	var results []dto.GenreListAndCountResponse
	err := r.db.WithContext(ctx).
		Table("genres AS g").
		Select("g.nama, COUNT(fg.genre_id) AS count").
		Joins("LEFT JOIN film_genres fg ON fg.genre_id = g.id").
		Group("g.nama").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
