package repository

import (
	"context"
	"math"

	"FilmFindr/dto"
	"FilmFindr/entity"
	"FilmFindr/helpers"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetReviewByFilmId(ctx context.Context, filmId uuid.UUID, offset int) ([]entity.Review, int64, error)
	GetReviewByUserId(ctx context.Context, id uuid.UUID, offset int) ([]entity.Review, int64, error)
	GetRatingFromMaterializedView(ctx context.Context, filmId uuid.UUID) (float64, error)
	CreateReview(ctx context.Context, review *entity.Review) error
	UpdateReview(ctx context.Context, reviewId uuid.UUID, review dto.UpdateReviewRequest) error
	DeleteReview(ctx context.Context, id uuid.UUID) error
	CountReviews(ctx context.Context) (int64, error)
	GetWeeklyReviews(ctx context.Context) ([]dto.WeeklyReview, error)
	GetListRatingAndCount(ctx context.Context) ([]dto.RatingListAndCountResponse, error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) GetReviewByFilmId(ctx context.Context, filmId uuid.UUID, offset int) ([]entity.Review, int64, error) {
	var reviews []entity.Review
	var countReview int64

	if err := r.db.WithContext(ctx).
		Model(&entity.Review{}).
		Where("film_id = ?", filmId).
		Count(&countReview).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "username", "photo_profil") }).
		Select("id", "komentar", "rating", "created_at", "user_id").
		Where("film_id = ?", filmId).
		Order("created_at DESC").
		Limit(helpers.LIMIT_REVIEW).
		Offset(offset).
		Find(&reviews).Error; err != nil {
		return nil, 0, err
	}

	totalPage := math.Ceil(float64(countReview) / float64(helpers.LIMIT_REVIEW))
	return reviews, int64(totalPage), nil
}

func (r *reviewRepository) GetReviewByUserId(ctx context.Context, id uuid.UUID, offset int) ([]entity.Review, int64, error) {
	var review []entity.Review
	var countReview int64

	if err := r.db.WithContext(ctx).
		Model(&entity.Review{}).
		Where("user_id = ?", id).
		Count(&countReview).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Select("id", "komentar", "rating", "created_at", "user_id").
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "username", "photo_profil") }).
		Where("user_id = ?", id).
		Order("created_at DESC").
		Limit(helpers.LIMIT_REVIEW).
		Offset(offset).
		Find(&review).Error; err != nil {
		return nil, 0, err
	}

	totalPage := math.Ceil(float64(countReview) / float64(helpers.LIMIT_REVIEW))
	return review, int64(totalPage), nil
}

func (r *reviewRepository) GetRatingFromMaterializedView(ctx context.Context, filmId uuid.UUID) (float64, error) {
	var rating dto.RatingFilm

	err := r.db.WithContext(ctx).
		Raw("SELECT * FROM rating_film WHERE film_id = ?", filmId).
		Scan(&rating).Error
	if err != nil {
		return 0, err
	}

	return rating.Rating, nil
}

func (r *reviewRepository) CreateReview(ctx context.Context, review *entity.Review) error {
	if err := r.db.WithContext(ctx).Create(&review).Error; err != nil {
		return err
	}
	return nil
}

func (r *reviewRepository) UpdateReview(ctx context.Context, reviewId uuid.UUID, req dto.UpdateReviewRequest) error {
	var review entity.Review
	if err := r.db.First(&review, reviewId).Error; err != nil {
		return err
	}

	if req.Komentar != "" {
		review.Komentar = req.Komentar
	}
	if req.Rating != 0 {
		review.Rating = req.Rating
	}

	if err := r.db.WithContext(ctx).Save(&review).Error; err != nil {
		return err
	}

	return nil
}

func (r *reviewRepository) DeleteReview(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Review{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *reviewRepository) CountReviews(ctx context.Context) (int64, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Model(entity.Review{}).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *reviewRepository) GetWeeklyReviews(ctx context.Context) ([]dto.WeeklyReview, error) {
	var results []dto.WeeklyReview

	err := r.db.WithContext(ctx).
		Raw("SELECT * FROM weekly_review ORDER BY weekly DESC LIMIT 8").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *reviewRepository) GetListRatingAndCount(ctx context.Context) ([]dto.RatingListAndCountResponse, error) {
	var results []dto.RatingListAndCountResponse
	err := r.db.WithContext(ctx).
		Table("reviews AS r").
		Select("r.rating, COUNT(r.id) AS count").
		Group("r.rating").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
