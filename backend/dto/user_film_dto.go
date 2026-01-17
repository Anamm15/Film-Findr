package dto

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	MESSAGE_FAILED_UPDATED_STATUS_USER_FILM = "Failed to update status user film"
	MESSAGE_FAILED_CREATED_USER_FILM        = "Failed to add film"
	MESSAGE_FAILED_UPDATED_USER_FILM        = "Failed to update status film"
	MESSAGE_FAILED_GET_USER_FILM            = "Failed to get user film"

	MESSAGE_SUCCESS_UPDATED_STATUS_USER_FILM = "Status user film updated successfully"
	MESSAGE_SUCCESS_CREATED_USER_FILM        = "Success add film"
	MESSAGE_SUCCESS_UPDATED_USER_FILM        = "Status film updated successfully"
	MESSAGE_SUCCESS_GET_USER_FILM            = "Success get user film"
)

var (
	ErrGetUserFilm           = errors.New("Failed to get user film")
	ErrCreateUserFilm        = errors.New("Failed to create user film")
	ErrUpdateStatusUserFilm  = errors.New("Failed to update status user film")
	ErrCheckUserFilm         = errors.New("You must add to your watchlist to review")
	ErrStatusFilmNotYetAired = errors.New("Film with not yet aired must be plan to watch")
	ErrUserFilmAlreadyExists = errors.New("User film already exists")
)

type (
	UserFilmCreateRequest struct {
		UserID uuid.UUID `json:"user_id" validate:"required"`
		FilmID uuid.UUID `json:"film_id" validate:"required"`
		Status string    `json:"status" validate:"required"`
	}

	UserFilmUpdateStatusRequest struct {
		Status string    `json:"status" validate:"required"`
		FilmID uuid.UUID `json:"film_id" validate:"required"`
	}

	UserFilmResponse struct {
		ID     uuid.UUID          `json:"id"`
		Status string             `json:"status"`
		Film   FilmDetailResponse `json:"film"`
	}

	PaginatedUserFilmResponse struct {
		UserFilms []UserFilmResponse `json:"user_films"`
		CountPage int                `json:"count_page"`
	}

	WeeklyUser struct {
		Label time.Time `json:"label" gorm:"column:weekly"`
		Value int64     `json:"value" gorm:"column:total_user"`
	}
)
