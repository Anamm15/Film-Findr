package repository

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilmGenreRepository interface {
	CreateFilmGenre(ctx context.Context, tx *gorm.DB, genre entity.FilmGenre) (entity.FilmGenre, error)
	DeleteFilmGenre(ctx context.Context, filmGenre entity.FilmGenre) error
	FindGenreByFilmIDs(ctx context.Context, filmIDs []uuid.UUID) ([]dto.GenreResponse, error)
}

type filmGenreRepository struct {
	db *gorm.DB
}

func NewFilmGenreRepository(db *gorm.DB) FilmGenreRepository {
	return &filmGenreRepository{db: db}
}

func (r *filmGenreRepository) CreateFilmGenre(ctx context.Context, tx *gorm.DB, genre entity.FilmGenre) (entity.FilmGenre, error) {
	if err := tx.Create(&genre).Error; err != nil {
		return entity.FilmGenre{}, err
	}

	return genre, nil
}

func (r *filmGenreRepository) DeleteFilmGenre(ctx context.Context, genre entity.FilmGenre) error {
	var newFilmGenre entity.FilmGenre
	if err := r.db.WithContext(ctx).
		Delete(&newFilmGenre, "film_id = ? AND genre_id = ?", genre.FilmID, genre.GenreID).Error; err != nil {
		return err
	}

	return nil
}

func (r *filmGenreRepository) FindGenreByFilmIDs(ctx context.Context, filmIDs []uuid.UUID) ([]dto.GenreResponse, error) {
	var genres []dto.GenreResponse
	err := r.db.WithContext(ctx).
		Raw(`
		SELECT g.id, g.nama, fg.film_id
		FROM film_genres fg
		JOIN genres g ON g.id = fg.genre_id
		WHERE fg.film_id IN ?
	`, filmIDs).Scan(&genres).Error
	if err != nil {
		return nil, err
	}

	return genres, nil
}
