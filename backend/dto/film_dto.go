package dto

import (
	"errors"
	"math"
	"time"

	"FilmFindr/entity"
	"FilmFindr/utils"

	"github.com/google/uuid"
)

const (
	// Failed messages
	MESSAGE_FAILED_FILM_NOT_FOUND      = "Film not found"
	MESSAGE_FAILED_GET_ALL_FILM        = "Failed get all film"
	MESSAGE_FAILED_GET_FILM            = "Failed get film"
	MESSAGE_FAILED_CREATED_FILM        = "Failed created film"
	MESSAGE_FAILED_UPDATED_FILM        = "Failed updated film"
	MESSAGE_FAILED_DELETED_FILM        = "Failed deleted film"
	MESSAGE_FAILED_UPDATED_STATUS_FILM = "Failed to update status film"
	MESSAGE_FAILED_SEARCH_FILM         = "Failed to search film"

	// Success messages
	MESSAGE_SUCCESS_GET_LIST_FILM       = "Success get list film"
	MESSAGE_SUCCESS_GET_FILM            = "Success get film"
	MESSAGE_SUCCESS_CREATED_FILM        = "Film created successfully"
	MESSAGE_SUCCESS_UPDATED_FILM        = "Film updated successfully"
	MESSAGE_SUCCESS_DELETED_FILM        = "Film deleted successfully"
	MESSAGE_SUCCESS_UPDATED_STATUS_FILM = "Status film updated successfully"
	MESSAGE_SUCCESS_SEARCH_FILM         = "Success search film"
)

var (
	ErrGetFilm          = errors.New("failed to get film")
	ErrCreateFilm       = errors.New("failed to create film")
	ErrUpdateFilm       = errors.New("failed to update film")
	ErrDeleteFilm       = errors.New("failed to delete film")
	ErrUpdateStatusFilm = errors.New("failed to update status film")
	ErrCheckStatusFilm  = errors.New("failed to check status film")
	ErrGetImageRequest  = errors.New("no image files found")
	ErrSearchFilm       = errors.New("failed to search film")
)

type (
	CreateFilmRequest struct {
		Judul        string   `form:"judul"`
		Sinopsis     string   `form:"sinopsis"`
		Sutradara    string   `form:"sutradara"`
		Status       string   `form:"status"`
		Durasi       int      `form:"durasi"`
		TotalEpisode int      `form:"total_episode"`
		TanggalRilis string   `form:"tanggal_rilis"`
		Genre        []string `form:"genres"`
	}

	UpdateFilmRequest struct {
		Judul        string    `json:"judul"`
		Sinopsis     string    `json:"sinopsis"`
		Sutradara    string    `json:"sutradara"`
		Status       string    `json:"status"`
		Durasi       int       `json:"durasi"`
		TotalEpisode int       `json:"total_episode"`
		TanggalRilis time.Time `json:"tanggal_rilis" time_format:"2006-01-02"`
	}

	UpdateStatusFilmRequest struct {
		Status string `json:"status" binding:"required"`
	}

	SearchFilmRequest struct {
		Keyword string `json:"keyword"`
		// Status  *string `json:"status"`
		// Genres  *[]int  `json:"genres"`
	}
)

type (
	FilmGambarResponse struct {
		ID     uuid.UUID `json:"id" gorm:"column:id"`
		FilmID uuid.UUID `json:"film_id" gorm:"column:film_id"`
		Url    string    `json:"url" gorm:"column:url"`
	}

	FilmDetailResponse struct {
		ID           uuid.UUID            `json:"id" gorm:"column:id"`
		Judul        string               `json:"judul" gorm:"column:judul"`
		Sinopsis     string               `json:"sinopsis,omitempty" gorm:"column:sinopsis"`
		Sutradara    string               `json:"sutradara,omitempty" gorm:"column:sutradara"`
		Status       string               `json:"status" gorm:"column:status"`
		Durasi       int                  `json:"durasi" gorm:"column:durasi"`
		TotalEpisode int                  `json:"total_episode,omitempty" gorm:"column:total_episode"`
		TanggalRilis string               `json:"tanggal_rilis" gorm:"column:tanggal_rilis"`
		Rating       float64              `json:"rating" gorm:"column:rating"`
		Gambar       []FilmGambarResponse `json:"film_gambar" gorm:"column:film_gambar;foreignKey:FilmID"`
		Genres       []GenreResponse      `json:"genres" gorm:"column:genres"`
	}

	FilmCompactResponse struct {
		ID           uuid.UUID `json:"id" gorm:"column:id"`
		Judul        string    `json:"judul" gorm:"column:judul"`
		Sinopsis     string    `json:"sinopsis,omitempty" gorm:"column:sinopsis"`
		Sutradara    string    `json:"sutradara,omitempty" gorm:"column:sutradara"`
		Status       string    `json:"status" gorm:"column:status"`
		Durasi       int       `json:"durasi" gorm:"column:durasi"`
		TotalEpisode int       `json:"total_episode,omitempty" gorm:"column:total_episode"`
		TanggalRilis time.Time `json:"tanggal_rilis" gorm:"column:tanggal_rilis"`
		Rating       float64   `json:"rating" gorm:"column:rating"`
	}

	PaginatedFilmResponse struct {
		CountPage int                  `json:"count_page"`
		Film      []FilmDetailResponse `json:"films"`
	}

	RatingFilm struct {
		FilmID uuid.UUID `gorm:"column:film_id"`
		Rating float64   `gorm:"column:rating"`
	}

	FilmWithMostReviews struct {
		ID           uuid.UUID `json:"id" gorm:"column:id"`
		Judul        string    `json:"judul" gorm:"column:judul"`
		CountReviews int       `json:"count_reviews" gorm:"column:count_reviews"`
	}

	FilmResponse struct {
		ID           uuid.UUID `json:"film_id"`
		Judul        string    `json:"judul"`
		Status       string    `json:"status"`
		Durasi       int       `json:"durasi"`
		TanggalRilis string    `json:"tanggal_rilis"`
		Rating       float64   `json:"rating"`
	}
)

func (dto *CreateFilmRequest) ToModel(film *entity.Film) {
	film.Judul = dto.Judul
	film.Sinopsis = dto.Sinopsis
	film.Sutradara = dto.Sutradara
	film.Status = dto.Status
	film.Durasi = dto.Durasi
	film.TotalEpisode = dto.TotalEpisode
	film.TanggalRilis = utils.ParseDate(dto.TanggalRilis)
}

func EntityToDetailFilmResponse(film entity.Film) FilmDetailResponse {
	return FilmDetailResponse{
		ID:           film.ID,
		Judul:        film.Judul,
		Sinopsis:     film.Sinopsis,
		Sutradara:    film.Sutradara,
		Status:       film.Status,
		Durasi:       film.Durasi,
		TotalEpisode: film.TotalEpisode,
		TanggalRilis: film.TanggalRilis.Format("2006-01-02"),
	}
}

func AddingGenresAndGambarToDetailFilmResponse(
	film FilmCompactResponse,
	genres map[uuid.UUID][]GenreResponse,
	gambar map[uuid.UUID][]FilmGambarResponse,
) FilmDetailResponse {
	return FilmDetailResponse{
		ID:           film.ID,
		Judul:        film.Judul,
		Sinopsis:     film.Sinopsis,
		Sutradara:    film.Sutradara,
		Status:       film.Status,
		Durasi:       film.Durasi,
		TotalEpisode: film.TotalEpisode,
		TanggalRilis: utils.FormatDate(film.TanggalRilis),
		Rating:       math.Round(film.Rating*100) / 100,
		Genres:       genres[film.ID],
		Gambar:       gambar[film.ID],
	}
}

func AddingGenresAndGambarToListDetailFilmResponse(
	films []FilmCompactResponse,
	genres map[uuid.UUID][]GenreResponse,
	gambar map[uuid.UUID][]FilmGambarResponse,
) []FilmDetailResponse {
	var results []FilmDetailResponse

	for _, film := range films {
		results = append(results, AddingGenresAndGambarToDetailFilmResponse(film, genres, gambar))
	}

	return results
}

func MapGambarByFilmID(gambar []FilmGambarResponse) map[uuid.UUID][]FilmGambarResponse {
	result := make(map[uuid.UUID][]FilmGambarResponse)

	for _, img := range gambar {
		result[img.FilmID] = append(result[img.FilmID], img)
	}

	return result
}
