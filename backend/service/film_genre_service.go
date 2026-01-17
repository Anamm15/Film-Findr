package service

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"
	"FilmFindr/repository"

	"gorm.io/gorm"
)

type FilmGenreService interface {
	CreateFilmGenre(ctx context.Context, filmGenre dto.FilmGenreRequest) (dto.FilmGenreResponse, error)
	DeleteFilmGenre(ctx context.Context, filmGenre dto.FilmGenreRequest) error
}

type filmGenreService struct {
	db                  *gorm.DB
	filmGenreRepository repository.FilmGenreRepository
}

func NewFilmGenreService(filmGenreRepository repository.FilmGenreRepository, db *gorm.DB) FilmGenreService {
	return &filmGenreService{filmGenreRepository: filmGenreRepository, db: db}
}

func (s *filmGenreService) CreateFilmGenre(ctx context.Context, filmGenreReq dto.FilmGenreRequest) (dto.FilmGenreResponse, error) {
	filmGenre := entity.FilmGenre{
		FilmID:  filmGenreReq.FilmId,
		GenreID: filmGenreReq.GenreId,
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return dto.FilmGenreResponse{}, dto.ErrCreateFilmGenre
	}

	filmGenreRes, err := s.filmGenreRepository.CreateFilmGenre(ctx, tx, filmGenre)
	if err != nil {
		tx.Rollback()
		return dto.FilmGenreResponse{}, dto.ErrCreateFilmGenre
	}

	tx.Commit()
	return dto.FilmGenreResponse{
		FilmId:  filmGenreRes.FilmID,
		GenreId: filmGenreRes.GenreID,
	}, nil
}

func (s *filmGenreService) DeleteFilmGenre(ctx context.Context, filmGenreReq dto.FilmGenreRequest) error {
	filmGenre := entity.FilmGenre{
		FilmID:  filmGenreReq.FilmId,
		GenreID: filmGenreReq.GenreId,
	}

	err := s.filmGenreRepository.DeleteFilmGenre(ctx, filmGenre)
	if err != nil {
		return dto.ErrDeleteFilmGenre
	}

	return nil
}
