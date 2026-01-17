package dto

const (
	MESSAGE_SUCCESS_GET_DASHBOARD        = "succesfully get dashboard statistics"
	MESSAGE_SUCCESS_GET_GENRE_DASHBOARD  = "succesfully get genre dashboard statistics"
	MESSAGE_SUCCESS_GET_REVIEW_DASHBOARD = "succesfully get review dashboard statistics"
	MESSAGE_FAILED_GET_DASHBOARD         = "Failed to get dashboard"
	MESSAGE_FAILED_GET_GENRE_DASHBOARD   = "Failed to get genre dashboard"
	MESSAGE_FAILED_GET_REVIEW_DASHBOARD  = "Failed to get review dashboard"
)

type (
	GetDashboardResponse struct {
		CountUsers    int                  `json:"count_users"`
		CountReview   int                  `json:"count_review"`
		TopFilms      []FilmDetailResponse `json:"top_films"`
		TrendingFilms []FilmDetailResponse `json:"trending_films"`
		WeeklyUsers   []WeeklyUser         `json:"weekly_users"`
		WeeklyReviews []WeeklyReview       `json:"weekly_reviews"`
	}

	GetReviewDashboardResponse struct {
		RatingDistribution []RatingListAndCountResponse `json:"rating_distribution"`
		MostReviewedFilms  []FilmWithMostReviews        `json:"most_reviewed_films"`
	}
)
