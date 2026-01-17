package dto

import (
	"errors"
	"time"

	"FilmFindr/entity"

	"github.com/google/uuid"
)

const (
	// failed
	MESSAGE_FAILED_GET_REVIEW     = "Failed get review"
	MESSAGE_FAILED_CREATED_REVIEW = "Failed created review"
	MESSAGE_FAILED_UPDATED_REVIEW = "Failed updated review"
	MESSAGE_FAILED_DELETED_REVIEW = "Failed deleted review"

	// success
	MESSAGE_SUCCESS_REVIEW_NOT_FOUND = "Review not found"
	MESSAGE_SUCCESS_CREATED_REVIEW   = "Review created successfully"
	MESSAGE_SUCCESS_UPDATED_REVIEW   = "Review updated successfully"
	MESSAGE_SUCCESS_DELETED_REVIEW   = "Review deleted successfully"
	MESSAGE_SUCCESS_GET_LIST_REVIEW  = "Success get list review"
	MESSAGE_SUCCESS_GET_REVIEW       = "Success get review"
)

var (
	ErrGetReviewByUserId           = errors.New("Failed to get review")
	ErrGetReviewFilmById           = errors.New("Failed to get review in this film")
	ErrGetReviewByID               = errors.New("Failed to get review")
	ErrCreateReview                = errors.New("Failed to create review")
	ErrUpdateReview                = errors.New("Failed to update review")
	ErrUpdateReaksiReview          = errors.New("Failed to update reaksi review")
	ErrDeleteReview                = errors.New("Failed to delete review")
	ErrCreateReviewWithStatus      = errors.New("Review with status not yet aired can't be created")
	ErrCreateReviewWithNoWatchlist = errors.New("You must add this film to your watchlist first")
)

type (
	UserReview struct {
		ID          uuid.UUID `json:"id"`
		Username    string    `json:"username"`
		PhotoProfil string    `json:"photo_profil"`
	}

	UserReaksiReview struct {
		ID     uuid.UUID `json:"id"`
		Reaksi string    `json:"reaksi"`
		UserID uuid.UUID `json:"user_id"`
	}

	ReviewResponse struct {
		ID         uuid.UUID         `json:"id"`
		Komentar   string            `json:"komentar"`
		Rating     int               `json:"rating"`
		User       UserReview        `json:"user"`
		UserReaksi *UserReaksiReview `json:"user_reaksi"`
	}

	PaginatedReview struct {
		CountPage int              `json:"count_page"`
		Reviews   []ReviewResponse `json:"reviews"`
	}

	CreateReviewRequest struct {
		FilmID   uuid.UUID `json:"film_id" binding:"required"`
		Komentar string    `json:"komentar" binding:"required"`
		Rating   int       `json:"rating" binding:"required"`
	}

	UpdateReviewRequest struct {
		Komentar string `json:"komentar"`
		Rating   int    `json:"rating"`
	}

	UpdateReaksiReviewRequest struct {
		Reaksi string `json:"reaksi" validate:"required"`
	}

	WeeklyReview struct {
		Label time.Time `json:"label" gorm:"column:weekly"`
		Value int64     `json:"value" gorm:"column:total_review"`
	}

	RatingListAndCountResponse struct {
		Rating int `json:"rating"`
		Count  int `json:"count"`
	}
)

func (r *CreateReviewRequest) ToModel(review *entity.Review, userID uuid.UUID) {
	review.FilmID = r.FilmID
	review.UserID = userID
	review.Komentar = r.Komentar
	review.Rating = r.Rating
}

func EntityToReviewResponse(review entity.Review, userReaksiReview *UserReaksiReview) ReviewResponse {
	return ReviewResponse{
		ID:         review.ID,
		Komentar:   review.Komentar,
		Rating:     review.Rating,
		UserReaksi: userReaksiReview,
		User: UserReview{
			ID:          review.User.ID,
			Username:    review.User.Username,
			PhotoProfil: review.User.PhotoProfil,
		},
	}
}
