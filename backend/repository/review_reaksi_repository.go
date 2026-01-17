package repository

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReaksiReviewRepository interface {
	GetReaksiReviewByUserId(ctx context.Context, reviewId uuid.UUID, userId uuid.UUID) (dto.UserReaksiReview, error)
	UpdateOrCreateUserReaksi(ctx context.Context, review entity.ReaksiReview) error
}

type reviewReaksiRepository struct {
	db *gorm.DB
}

func NewReaksiReviewRepository(db *gorm.DB) ReaksiReviewRepository {
	return &reviewReaksiRepository{db: db}
}

func (r *reviewReaksiRepository) GetReaksiReviewByUserId(ctx context.Context, reviewId uuid.UUID, userId uuid.UUID) (dto.UserReaksiReview, error) {
	var reaksiReview entity.ReaksiReview
	if err := r.db.WithContext(ctx).
		Select("id", "reaksi", "user_id").
		Where("user_id = ?", userId).
		Where("review_id = ?", reviewId).
		Find(&reaksiReview).Error; err != nil {
		return dto.UserReaksiReview{}, err
	}

	return dto.UserReaksiReview{
		ID:     reaksiReview.ID,
		Reaksi: reaksiReview.Reaksi,
		UserID: reaksiReview.UserID,
	}, nil
}

func (r *reviewReaksiRepository) UpdateOrCreateUserReaksi(ctx context.Context, review entity.ReaksiReview) error {
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND review_id = ?", review.UserID, review.ReviewID).
		Assign(entity.ReaksiReview{Reaksi: review.Reaksi}).
		FirstOrCreate(&entity.ReaksiReview{
			UserID:   review.UserID,
			ReviewID: review.ReviewID,
			Reaksi:   review.Reaksi,
		}).Error

	return err
}
