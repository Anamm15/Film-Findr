package service

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"
	"FilmFindr/helpers"
	"FilmFindr/repository"
	"FilmFindr/utils"

	"github.com/google/uuid"
)

type ReviewService interface {
	CreateReview(ctx context.Context, review dto.CreateReviewRequest, userId uuid.UUID) (dto.ReviewResponse, error)
	GetReviewByUserId(ctx context.Context, id string, userId uuid.UUID, page string) (dto.PaginatedReview, error)
	GetReviewByFilmId(ctx context.Context, filmId string, userId uuid.UUID, page string) (dto.PaginatedReview, error)
	UpdateReview(ctx context.Context, reviewId string, review dto.UpdateReviewRequest) error
	UpdateReaksiReview(ctx context.Context, reviewId string, userId uuid.UUID, reaksi dto.UpdateReaksiReviewRequest) error
	DeleteReview(ctx context.Context, id string) error
}

type reviewService struct {
	reviewRepository       repository.ReviewRepository
	reaksiReviewRepository repository.ReaksiReviewRepository
	userFilmRepository     repository.UserFilmRepository
	filmRepository         repository.FilmRepository
}

func NewReviewService(
	reviewRepository repository.ReviewRepository,
	reaksiReviewRepository repository.ReaksiReviewRepository,
	userFilmRepository repository.UserFilmRepository,
	filmRepository repository.FilmRepository,
) ReviewService {
	return &reviewService{
		reviewRepository:       reviewRepository,
		reaksiReviewRepository: reaksiReviewRepository,
		userFilmRepository:     userFilmRepository,
		filmRepository:         filmRepository,
	}
}

func (s *reviewService) CreateReview(ctx context.Context, reviewReq dto.CreateReviewRequest, userId uuid.UUID) (dto.ReviewResponse, error) {
	review := &entity.Review{}
	reviewReq.ToModel(review, userId)

	film, err := s.filmRepository.CheckStatusFilm(ctx, reviewReq.FilmID)
	if err != nil {
		return dto.ReviewResponse{}, dto.ErrCheckStatusFilm
	}
	if film.Status == helpers.ENUM_FILM_NOT_YET_AIRED {
		return dto.ReviewResponse{}, dto.ErrCreateReviewWithStatus
	}

	checkUserFilm, err := s.userFilmRepository.CheckUserFilm(ctx, userId, reviewReq.FilmID)
	if err != nil {
		return dto.ReviewResponse{}, dto.ErrCheckUserFilm
	}

	if !checkUserFilm {
		return dto.ReviewResponse{}, dto.ErrCreateReviewWithNoWatchlist
	}

	err = s.reviewRepository.CreateReview(ctx, review)
	if err != nil {
		return dto.ReviewResponse{}, dto.ErrCreateReview
	}

	result := dto.EntityToReviewResponse(*review, nil)
	return result, nil
}

func (s *reviewService) GetReviewByUserId(ctx context.Context, idParam string, userId uuid.UUID, pageQuery string) (dto.PaginatedReview, error) {
	id, _ := utils.StringToUUID(idParam)
	page, err := utils.StringToInt(pageQuery)
	if err != nil || page < 0 {
		page = 1
	}

	offset := (page - 1) % helpers.LIMIT_REVIEW
	reviews, countReview, err := s.reviewRepository.GetReviewByUserId(ctx, id, offset)
	if err != nil {
		return dto.PaginatedReview{}, dto.ErrGetReviewByUserId
	}

	var reviewsResponse dto.PaginatedReview
	reviewsResponse.CountPage = int(countReview)

	for _, review := range reviews {
		userReaksiReview, _ := s.reaksiReviewRepository.GetReaksiReviewByUserId(ctx, review.ID, userId)
		reviewResponse := dto.EntityToReviewResponse(review, &userReaksiReview)
		reviewsResponse.Reviews = append(reviewsResponse.Reviews, reviewResponse)
	}

	return reviewsResponse, nil
}

func (s *reviewService) GetReviewByFilmId(ctx context.Context, filmIdParam string, userId uuid.UUID, pageQuery string) (dto.PaginatedReview, error) {
	filmId, _ := utils.StringToUUID(filmIdParam)

	page, err := utils.StringToInt(pageQuery)
	if err != nil || page < 0 {
		page = 1
	}

	offset := (page - 1) % helpers.LIMIT_REVIEW
	reviews, countPage, err := s.reviewRepository.GetReviewByFilmId(ctx, filmId, offset)
	if err != nil {
		return dto.PaginatedReview{}, dto.ErrGetReviewFilmById
	}

	var reviewsResponse dto.PaginatedReview
	reviewsResponse.CountPage = int(countPage)

	for _, review := range reviews {
		userReaksiReview, _ := s.reaksiReviewRepository.GetReaksiReviewByUserId(ctx, review.ID, userId)
		reviewResponse := dto.EntityToReviewResponse(review, &userReaksiReview)
		reviewsResponse.Reviews = append(reviewsResponse.Reviews, reviewResponse)
	}

	return reviewsResponse, nil
}

func (s *reviewService) UpdateReview(ctx context.Context, reviewIDParam string, review dto.UpdateReviewRequest) error {
	reviewID, err := utils.StringToUUID(reviewIDParam)
	if err != nil {
		return err
	}

	err = s.reviewRepository.UpdateReview(ctx, reviewID, review)
	if err != nil {
		return dto.ErrUpdateReview
	}

	return nil
}

func (s *reviewService) UpdateReaksiReview(ctx context.Context, reviewIDParam string, userId uuid.UUID, reaksi dto.UpdateReaksiReviewRequest) error {
	reviewId, err := utils.StringToUUID(reviewIDParam)
	if err != nil {
		return err
	}

	reaksiReview := entity.ReaksiReview{
		Reaksi:   reaksi.Reaksi,
		UserID:   userId,
		ReviewID: reviewId,
	}

	err = s.reaksiReviewRepository.UpdateOrCreateUserReaksi(ctx, reaksiReview)
	if err != nil {
		return dto.ErrUpdateReaksiReview
	}

	return nil
}

func (s *reviewService) DeleteReview(ctx context.Context, idParam string) error {
	id, _ := utils.StringToUUID(idParam)
	err := s.reviewRepository.DeleteReview(ctx, id)
	if err != nil {
		return dto.ErrDeleteReview
	}
	return nil
}
