package dto

import (
	"errors"

	"github.com/google/uuid"
)

const (
	MESSAGE_FAILED_CREATED_FILM_GENRE = "Failed to create film genre"
	MESSAGE_FAILED_DELETED_FILM_GENRE = "Failed to delete film genre"

	MESSAGE_SUCCESS_CREATED_FILM_GENRE = "Film genre created successfully"
	MESSAGE_SUCCESS_DELETED_FILM_GENRE = "Film genre deleted successfully"
)

var (
	ErrCreateFilmGenre = errors.New("Failed to create film genre")
	ErrDeleteFilmGenre = errors.New("Failed to delete film genre")
)

type (
	FilmGenreRequest struct {
		FilmId  uuid.UUID `json:"film_id"`
		GenreId uuid.UUID `json:"genre_id"`
	}

	FilmGenreResponse struct {
		FilmId  uuid.UUID `json:"film_id"`
		GenreId uuid.UUID `json:"genre_id"`
	}
)
