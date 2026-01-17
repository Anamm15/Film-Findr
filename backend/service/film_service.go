package service

import (
	"context"
	"math"
	"mime/multipart"

	"FilmFindr/dto"
	"FilmFindr/entity"
	"FilmFindr/helpers"
	"FilmFindr/repository"
	"FilmFindr/utils"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilmService interface {
	GetAllFilm(ctx context.Context, page string) (dto.PaginatedFilmResponse, error)
	GetFilmByID(ctx context.Context, id string) (dto.FilmDetailResponse, error)
	CreateFilm(ctx context.Context, filmReq dto.CreateFilmRequest, files []*multipart.FileHeader) (dto.FilmDetailResponse, error)
	UpdateFilm(ctx context.Context, filmId string, film dto.UpdateFilmRequest) error
	DeleteFilm(ctx context.Context, id string) error
	UpdateStatus(ctx context.Context, id string, req dto.UpdateStatusFilmRequest) error
	SearchFilm(ctx context.Context, req dto.SearchFilmRequest, page string) (dto.PaginatedFilmResponse, error)
	GetTopFilm(ctx context.Context) ([]dto.FilmDetailResponse, error)
	GetTrendingFilm(ctx context.Context) ([]dto.FilmDetailResponse, error)
}

type filmService struct {
	filmRepository       repository.FilmRepository
	filmGambarRepository repository.FilmGambarRepository
	filmGenreRepository  repository.FilmGenreRepository
	reviewRepository     repository.ReviewRepository
	cloudinary           *cloudinary.Cloudinary
	db                   *gorm.DB
}

func NewFilmService(
	db *gorm.DB,
	cloudinary *cloudinary.Cloudinary,
	filmRepository repository.FilmRepository,
	filmGambarRepository repository.FilmGambarRepository,
	filmGenreRepository repository.FilmGenreRepository,
	reviewRepository repository.ReviewRepository,
) FilmService {
	return &filmService{
		db:                   db,
		cloudinary:           cloudinary,
		filmRepository:       filmRepository,
		filmGambarRepository: filmGambarRepository,
		filmGenreRepository:  filmGenreRepository,
		reviewRepository:     reviewRepository,
	}
}

func (s *filmService) GetAllFilm(ctx context.Context, pageQuery string) (dto.PaginatedFilmResponse, error) {
	page, err := utils.StringToInt(pageQuery)
	if err != nil || page < 0 {
		page = 1
	}

	offset := (page - 1) % helpers.LIMIT_FILM
	countFilm, err := s.filmRepository.CountFilm(ctx)
	if err != nil {
		return dto.PaginatedFilmResponse{}, err
	}

	films, err := s.filmRepository.GetAllFilm(ctx, offset)
	if err != nil {
		return dto.PaginatedFilmResponse{}, err
	}

	filmIDs := make([]uuid.UUID, len(films))
	for i, f := range films {
		filmIDs[i] = f.ID
	}

	genres, err := s.filmGenreRepository.FindGenreByFilmIDs(ctx, filmIDs)
	if err != nil {
		return dto.PaginatedFilmResponse{}, err
	}

	gambar, err := s.filmGambarRepository.FindFilmGambarByFilmIDs(ctx, filmIDs)
	if err != nil {
		return dto.PaginatedFilmResponse{}, err
	}

	genreMappedByFilmIDs := dto.MapGenresByFilmID(genres)
	gambarMappedByFilmIDs := dto.MapGambarByFilmID(gambar)
	results := dto.AddingGenresAndGambarToListDetailFilmResponse(films, genreMappedByFilmIDs, gambarMappedByFilmIDs)

	totalPage := int(math.Ceil(float64(countFilm) / float64(helpers.LIMIT_FILM)))
	GetFilmResponses := dto.PaginatedFilmResponse{
		Film:      results,
		CountPage: totalPage,
	}
	return GetFilmResponses, nil
}

func (s *filmService) GetFilmByID(ctx context.Context, idParam string) (dto.FilmDetailResponse, error) {
	id, err := utils.StringToUUID(idParam)
	if err != nil {
		return dto.FilmDetailResponse{}, dto.ErrGetFilm
	}

	filmFlat, err := s.filmRepository.GetFilmByID(ctx, id)
	if err != nil {
		return dto.FilmDetailResponse{}, dto.ErrGetFilm
	}

	filmIDs := []uuid.UUID{filmFlat.ID}
	genres, err := s.filmGenreRepository.FindGenreByFilmIDs(ctx, filmIDs)
	if err != nil {
		return dto.FilmDetailResponse{}, err
	}

	gambar, err := s.filmGambarRepository.FindFilmGambarByFilmIDs(ctx, filmIDs)
	if err != nil {
		return dto.FilmDetailResponse{}, err
	}

	genreMappedByFilmIDs := dto.MapGenresByFilmID(genres)
	gambarMappedByFilmIDs := dto.MapGambarByFilmID(gambar)
	filmResponse := dto.AddingGenresAndGambarToDetailFilmResponse(filmFlat, genreMappedByFilmIDs, gambarMappedByFilmIDs)
	return filmResponse, nil
}

func (s *filmService) CreateFilm(ctx context.Context, filmReq dto.CreateFilmRequest, files []*multipart.FileHeader) (dto.FilmDetailResponse, error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		return dto.FilmDetailResponse{}, dto.ErrCreateFilm
	}

	film := entity.Film{}
	filmReq.ToModel(&film)

	err := s.filmRepository.CreateFilm(ctx, tx, &film)
	var filmGambarResponse []dto.FilmGambarResponse
	var genreResponse []dto.GenreResponse

	if err != nil {
		tx.Rollback()
		return dto.FilmDetailResponse{}, dto.ErrCreateFilm
	}

	for _, genreID := range filmReq.Genre {
		genreUUID, err := utils.StringToUUID(genreID)
		if err != nil {
			tx.Rollback()
			return dto.FilmDetailResponse{}, dto.ErrCreateFilm
		}

		filmGenre := entity.FilmGenre{
			FilmID:  film.ID,
			GenreID: genreUUID,
		}
		genre, err := s.filmGenreRepository.CreateFilmGenre(ctx, tx, filmGenre)
		genreResponse = append(genreResponse, dto.GenreResponse{
			ID:   genre.Genre.ID,
			Nama: genre.Genre.Nama,
		})
		if err != nil {
			tx.Rollback()
			return dto.FilmDetailResponse{}, dto.ErrCreateFilm
		}
	}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			tx.Rollback()
			return dto.FilmDetailResponse{}, dto.ErrCreateFilm
		}

		uniqueName := utils.GenerateUniqueImageName(film.Judul, file.Filename)
		uploadResult, err := s.cloudinary.Upload.Upload(ctx, src, uploader.UploadParams{
			Folder:   "ReviewFilem",
			PublicID: uniqueName,
		})
		src.Close()

		if err != nil {
			tx.Rollback()
			return dto.FilmDetailResponse{}, dto.ErrFailedUploadFile
		}

		filmGambar := entity.FilmGambar{
			FilmID: film.ID,
			Url:    uploadResult.SecureURL,
		}

		if err := s.filmGambarRepository.Save(ctx, tx, filmGambar); err != nil {
			tx.Rollback()
			return dto.FilmDetailResponse{}, dto.ErrCreateFilm
		}

		filmGambarResponse = append(filmGambarResponse, dto.FilmGambarResponse{
			ID:  filmGambar.ID,
			Url: filmGambar.Url,
		})
	}

	if err := tx.Commit().Error; err != nil {
		return dto.FilmDetailResponse{}, dto.ErrCreateFilm
	}

	result := dto.EntityToDetailFilmResponse(film)

	result.Genres = genreResponse
	result.Gambar = filmGambarResponse
	return result, nil
}

func (s *filmService) UpdateFilm(ctx context.Context, filmIdParam string, film dto.UpdateFilmRequest) error {
	filmId, err := utils.StringToUUID(filmIdParam)
	if err != nil {
		return dto.ErrUpdateFilm
	}

	_, err = s.filmRepository.UpdateFilm(ctx, filmId, film)
	if err != nil {
		return dto.ErrUpdateFilm
	}
	return nil
}

func (s *filmService) DeleteFilm(ctx context.Context, idParam string) error {
	id, err := utils.StringToUUID(idParam)
	if err != nil {
		return dto.ErrUpdateFilm
	}

	err = s.filmRepository.DeleteFilm(ctx, id)
	if err != nil {
		return dto.ErrDeleteFilm
	}

	return nil
}

func (s *filmService) UpdateStatus(ctx context.Context, idParam string, req dto.UpdateStatusFilmRequest) error {
	id, err := utils.StringToUUID(idParam)
	if err != nil {
		return dto.ErrUpdateFilm
	}

	err = s.filmRepository.UpdateStatus(ctx, id, req.Status)
	if err != nil {
		return dto.ErrUpdateStatusFilm
	}

	return nil
}

func (s *filmService) SearchFilm(ctx context.Context, req dto.SearchFilmRequest, pageQuery string) (dto.PaginatedFilmResponse, error) {
	page, err := utils.StringToInt(pageQuery)
	if err != nil || page < 0 {
		page = 1
	}

	offset := (page - 1) % helpers.LIMIT_FILM
	countFilm, err := s.filmRepository.CountFilm(ctx)
	if err != nil {
		return dto.PaginatedFilmResponse{}, err
	}

	films, err := s.filmRepository.SearchFilm(ctx, req, offset)
	if err != nil {
		return dto.PaginatedFilmResponse{}, err
	}

	var filmResponses []dto.FilmDetailResponse

	for _, film := range films {
		var fileResponses []dto.FilmGambarResponse
		var genreResponses []dto.GenreResponse

		rating, _ := s.reviewRepository.GetRatingFromMaterializedView(ctx, film.ID)
		rating = math.Round(rating*100) / 100

		for _, file := range film.FilmGambar {
			fileResponses = append(fileResponses, dto.FilmGambarResponse{
				ID:  file.ID,
				Url: file.Url,
			})
		}

		for _, genre := range film.FilmGenre {
			genreResponses = append(genreResponses, dto.GenreResponse{
				ID:   genre.Genre.ID,
				Nama: genre.Genre.Nama,
			})
		}

		formattedDate := utils.FormatDate(film.TanggalRilis)
		filmResponses = append(filmResponses, dto.FilmDetailResponse{
			ID:           film.ID,
			Judul:        film.Judul,
			TanggalRilis: formattedDate,
			Durasi:       film.Durasi,
			Status:       film.Status,
			Rating:       rating,
			Gambar:       fileResponses,
			Genres:       genreResponses,
		})
	}

	totalPage := int(math.Ceil(float64(countFilm) / float64(helpers.LIMIT_FILM)))
	getFilmResponses := dto.PaginatedFilmResponse{
		CountPage: totalPage,
		Film:      filmResponses,
	}
	return getFilmResponses, nil
}

func (s *filmService) GetTopFilm(ctx context.Context) ([]dto.FilmDetailResponse, error) {
	topFilms, err := s.filmRepository.GetTopFilm(ctx)
	if err != nil {
		return nil, err
	}

	filmIDs := make([]uuid.UUID, len(topFilms))
	for i, f := range topFilms {
		filmIDs[i] = f.ID
	}

	genres, err := s.filmGenreRepository.FindGenreByFilmIDs(ctx, filmIDs)
	if err != nil {
		return nil, err
	}

	gambar, err := s.filmGambarRepository.FindFilmGambarByFilmIDs(ctx, filmIDs)
	if err != nil {
		return nil, err
	}

	genreMappedByFilmIDs := dto.MapGenresByFilmID(genres)
	gambarMappedByFilmIDs := dto.MapGambarByFilmID(gambar)
	results := dto.AddingGenresAndGambarToListDetailFilmResponse(topFilms, genreMappedByFilmIDs, gambarMappedByFilmIDs)
	return results, nil
}

func (s *filmService) GetTrendingFilm(ctx context.Context) ([]dto.FilmDetailResponse, error) {
	trendingFilm, err := s.filmRepository.GetTrendingFilm(ctx)
	if err != nil {
		return nil, err
	}

	filmIDs := make([]uuid.UUID, len(trendingFilm))
	for i, f := range trendingFilm {
		filmIDs[i] = f.ID
	}

	genres, err := s.filmGenreRepository.FindGenreByFilmIDs(ctx, filmIDs)
	if err != nil {
		return nil, err
	}

	gambar, err := s.filmGambarRepository.FindFilmGambarByFilmIDs(ctx, filmIDs)
	if err != nil {
		return nil, err
	}

	genreMappedByFilmIDs := dto.MapGenresByFilmID(genres)
	gambarMappedByFilmIDs := dto.MapGambarByFilmID(gambar)
	results := dto.AddingGenresAndGambarToListDetailFilmResponse(trendingFilm, genreMappedByFilmIDs, gambarMappedByFilmIDs)
	return results, nil
}
