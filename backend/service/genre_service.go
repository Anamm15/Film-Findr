package service

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"
	"FilmFindr/repository"
	"FilmFindr/utils"
)

type GenreService interface {
	GetAllGenre(ctx context.Context) ([]dto.GenreResponse, error)
	CreateGenre(ctx context.Context, genre dto.GenreRequest) (dto.GenreResponse, error)
	DeleteGenre(ctx context.Context, genreId string) error
}

type genreService struct {
	genreRepository repository.GenreRepository
}

func NewGenreService(genreRepository repository.GenreRepository) GenreService {
	return &genreService{
		genreRepository: genreRepository,
	}
}

func (s *genreService) GetAllGenre(ctx context.Context) ([]dto.GenreResponse, error) {
	genres, err := s.genreRepository.GetAllGenre(ctx)
	if err != nil {
		return nil, dto.ErrGetAllGenre
	}

	var response []dto.GenreResponse
	for _, genre := range genres {
		response = append(response, dto.GenreResponse{
			ID:        genre.ID,
			Nama:      genre.Nama,
			CreatedAt: genre.CreatedAt,
		})
	}

	return response, nil
}

func (s *genreService) CreateGenre(ctx context.Context, genre dto.GenreRequest) (dto.GenreResponse, error) {
	GenreRequest := entity.Genre{
		Nama: genre.Nama,
	}

	createdGenre, err := s.genreRepository.CreateGenre(ctx, GenreRequest)
	if err != nil {
		return dto.GenreResponse{}, dto.ErrGetAllGenre
	}

	response := dto.GenreResponse{
		ID:   createdGenre.ID,
		Nama: createdGenre.Nama,
	}

	return response, nil
}

func (s *genreService) DeleteGenre(ctx context.Context, genreIDParam string) error {
	genreID, err := utils.StringToUUID(genreIDParam)
	if err != nil {
		return err
	}

	err = s.genreRepository.DeleteGenre(ctx, genreID)
	if err != nil {
		return dto.ErrDeleteGenre
	}

	return nil
}
