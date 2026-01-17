package service

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"
	"FilmFindr/helpers"
	"FilmFindr/repository"
	"FilmFindr/utils"
)

type UserFilmService interface {
	GetUserFilmByUserId(ctx context.Context, userId string, page string) (dto.PaginatedUserFilmResponse, error)
	CreateUserFilm(ctx context.Context, userFilm dto.UserFilmCreateRequest) (entity.UserFilm, error)
	UpdateStatusUserFilm(ctx context.Context, userFilmId string, userFilm dto.UserFilmUpdateStatusRequest) error
}

type userFilmService struct {
	userFilmRepository repository.UserFilmRepository
	filmRepository     repository.FilmRepository
}

func NewUserFilmService(
	userFilmRepository repository.UserFilmRepository,
	filmRepository repository.FilmRepository,
) UserFilmService {
	return &userFilmService{
		userFilmRepository: userFilmRepository,
		filmRepository:     filmRepository,
	}
}

func (s *userFilmService) GetUserFilmByUserId(ctx context.Context, userIdParam string, pageQuery string) (dto.PaginatedUserFilmResponse, error) {
	userID, err := utils.StringToUUID(userIdParam)
	if err != nil {
		return dto.PaginatedUserFilmResponse{}, err
	}

	page, err := utils.StringToInt(pageQuery)
	if err != nil || page < 0 {
		page = 1
	}

	offset := (page - 1) % helpers.LIMIT_FILM
	userFilms, countUserFilm, err := s.userFilmRepository.GetUserFilmByUserId(ctx, userID, offset)
	if err != nil {
		return dto.PaginatedUserFilmResponse{}, dto.ErrGetUserFilm
	}

	var userFilmResponses []dto.UserFilmResponse
	for _, userFilm := range userFilms {
		var FilmGambarResponse []dto.FilmGambarResponse
		var FilmGenreResponse []dto.GenreResponse
		var FilmResponse dto.FilmDetailResponse

		for _, FilmGambar := range userFilm.Film.FilmGambar {
			FilmGambarResponse = append(FilmGambarResponse, dto.FilmGambarResponse{
				ID:  FilmGambar.ID,
				Url: FilmGambar.Url,
			})
		}

		for _, FilmGenre := range userFilm.Film.FilmGenre {
			FilmGenreResponse = append(FilmGenreResponse, dto.GenreResponse{
				ID:   FilmGenre.Genre.ID,
				Nama: FilmGenre.Genre.Nama,
			})
		}

		formattedDate := utils.FormatDate(userFilm.Film.TanggalRilis)
		FilmResponse = dto.FilmDetailResponse{
			ID:           userFilm.Film.ID,
			Judul:        userFilm.Film.Judul,
			TanggalRilis: formattedDate,
			Durasi:       userFilm.Film.Durasi,
			Status:       userFilm.Film.Status,
			Gambar:       FilmGambarResponse,
			Genres:       FilmGenreResponse,
		}

		userFilmResponses = append(userFilmResponses, dto.UserFilmResponse{
			ID:     userFilm.ID,
			Status: userFilm.Status,
			Film:   FilmResponse,
		})
	}

	var GetUserFilmResponse dto.PaginatedUserFilmResponse
	GetUserFilmResponse.UserFilms = userFilmResponses
	GetUserFilmResponse.CountPage = int(countUserFilm)
	return GetUserFilmResponse, nil
}

func (s *userFilmService) CreateUserFilm(ctx context.Context, userFilmReq dto.UserFilmCreateRequest) (entity.UserFilm, error) {
	userFilm := entity.UserFilm{
		Status: userFilmReq.Status,
		UserID: userFilmReq.UserID,
		FilmID: userFilmReq.FilmID,
	}

	film, err := s.filmRepository.CheckStatusFilm(ctx, userFilmReq.FilmID)
	if err != nil {
		return entity.UserFilm{}, dto.ErrCheckUserFilm
	}

	if film.Status == helpers.ENUM_FILM_NOT_YET_AIRED && userFilm.Status != helpers.ENUM_LIST_FILM_PLAN_TO_WATCH {
		return entity.UserFilm{}, dto.ErrStatusFilmNotYetAired
	}

	isUserFilmExist, err := s.userFilmRepository.CheckUserFilm(ctx, userFilm.UserID, userFilm.FilmID)
	if isUserFilmExist {
		return entity.UserFilm{}, dto.ErrUserFilmAlreadyExists
	}

	userFilmRes, err := s.userFilmRepository.CreateUserFilm(ctx, userFilm)
	if err != nil {
		return entity.UserFilm{}, dto.ErrCreateUserFilm
	}

	return userFilmRes, nil
}

func (s *userFilmService) UpdateStatusUserFilm(ctx context.Context, idParam string, userFilm dto.UserFilmUpdateStatusRequest) error {
	userFilmId, err := utils.StringToUUID(idParam)
	if err != nil {
		return err
	}

	film, err := s.filmRepository.CheckStatusFilm(ctx, userFilm.FilmID)
	if err != nil {
		return err
	}

	if film.Status == helpers.ENUM_FILM_NOT_YET_AIRED && userFilm.Status != helpers.ENUM_LIST_FILM_PLAN_TO_WATCH {
		return dto.ErrUpdateStatusUserFilm
	}

	if err := s.userFilmRepository.UpdateStatusUserFilm(ctx, userFilmId, userFilm.Status); err != nil {
		return dto.ErrUpdateStatusUserFilm
	}

	return nil
}
