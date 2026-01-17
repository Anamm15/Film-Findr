package service

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/repository"
)

type DashboardService interface {
	GetDashboard(ctx context.Context) (dto.GetDashboardResponse, error)
	GetGenreDashboard(ctx context.Context) ([]dto.GenreListAndCountResponse, error)
	GetReviewDashboard(ctx context.Context) (dto.GetReviewDashboardResponse, error)
}

type dashboardService struct {
	filmRepository   repository.FilmRepository
	reviewRepository repository.ReviewRepository
	userRepository   repository.UserRepository
	genreRepository  repository.GenreRepository
	filmService      FilmService
}

func NewDashboardService(
	filmRepository repository.FilmRepository,
	reviewRepository repository.ReviewRepository,
	userRepository repository.UserRepository,
	genreRepository repository.GenreRepository,
	filmService FilmService,
) DashboardService {
	return &dashboardService{
		filmRepository:   filmRepository,
		reviewRepository: reviewRepository,
		userRepository:   userRepository,
		filmService:      filmService,
		genreRepository:  genreRepository,
	}
}

func (s *dashboardService) GetDashboard(ctx context.Context) (dto.GetDashboardResponse, error) {
	topFilms, err := s.filmService.GetTopFilm(ctx)
	if err != nil {
		return dto.GetDashboardResponse{}, err
	}

	trendingFilms, err := s.filmService.GetTrendingFilm(ctx)
	if err != nil {
		return dto.GetDashboardResponse{}, err
	}

	weeklyUsers, err := s.userRepository.GetWeeklyUsers(ctx)
	if err != nil {
		return dto.GetDashboardResponse{}, err
	}

	weeklyReviews, err := s.reviewRepository.GetWeeklyReviews(ctx)
	if err != nil {
		return dto.GetDashboardResponse{}, err
	}

	countUsers, err := s.userRepository.CountUsers(ctx)
	if err != nil {
		return dto.GetDashboardResponse{}, err
	}

	countReviews, err := s.reviewRepository.CountReviews(ctx)

	var results dto.GetDashboardResponse
	results = dto.GetDashboardResponse{
		CountUsers:    int(countUsers),
		CountReview:   int(countReviews),
		TopFilms:      topFilms,
		TrendingFilms: trendingFilms,
		WeeklyUsers:   weeklyUsers,
		WeeklyReviews: weeklyReviews,
	}

	return results, nil
}

func (s *dashboardService) GetGenreDashboard(ctx context.Context) ([]dto.GenreListAndCountResponse, error) {
	genreListAndCount, err := s.genreRepository.GetGenreListAndCount(ctx)
	if err != nil {
		return nil, err
	}

	return genreListAndCount, nil
}

func (s *dashboardService) GetReviewDashboard(ctx context.Context) (dto.GetReviewDashboardResponse, error) {
	reviewListAndCount, err := s.reviewRepository.GetListRatingAndCount(ctx)
	if err != nil {
		return dto.GetReviewDashboardResponse{}, err
	}

	topFilmWithMostReviews, err := s.filmRepository.GetFilmWithMostReviews(ctx)
	if err != nil {
		return dto.GetReviewDashboardResponse{}, err
	}

	var results dto.GetReviewDashboardResponse
	results = dto.GetReviewDashboardResponse{
		RatingDistribution: reviewListAndCount,
		MostReviewedFilms:  topFilmWithMostReviews,
	}

	return results, nil
}
