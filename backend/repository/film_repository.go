package repository

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"
	"FilmFindr/helpers"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilmRepository interface {
	GetAllFilm(ctx context.Context, offset int) ([]dto.FilmCompactResponse, error)
	CreateFilm(ctx context.Context, tx *gorm.DB, film *entity.Film) error
	UpdateFilm(ctx context.Context, filmId uuid.UUID, film dto.UpdateFilmRequest) (entity.Film, error)
	DeleteFilm(ctx context.Context, id uuid.UUID) error
	GetFilmByID(ctx context.Context, id uuid.UUID) (dto.FilmCompactResponse, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
	CheckStatusFilm(ctx context.Context, id uuid.UUID) (entity.Film, error)
	SearchFilm(ctx context.Context, req dto.SearchFilmRequest, offset int) ([]entity.Film, error)
	CountFilm(ctx context.Context) (int64, error)
	GetTopFilm(ctx context.Context) ([]dto.FilmCompactResponse, error)
	GetTrendingFilm(ctx context.Context) ([]dto.FilmCompactResponse, error)
	GetFilmWithMostReviews(ctx context.Context) ([]dto.FilmWithMostReviews, error)
}

type filmRepository struct {
	db *gorm.DB
}

func NewFilmRepository(db *gorm.DB) FilmRepository {
	return &filmRepository{db: db}
}

func (r *filmRepository) GetAllFilm(ctx context.Context, offset int) ([]dto.FilmCompactResponse, error) {
	var filmFlats []dto.FilmCompactResponse

	if err := r.db.WithContext(ctx).
		Raw(`
		SELECT f.id, f.judul, f.sinopsis, f.sutradara, f.status, f.durasi,
		       f.total_episode, f.tanggal_rilis,
		       COALESCE(rf.rating, 0) AS rating
		FROM films f
		LEFT JOIN rating_film rf ON rf.film_id = f.id
		ORDER BY f.created_at DESC
		LIMIT ? OFFSET ?
	`, helpers.LIMIT_FILM, offset).Scan(&filmFlats).Error; err != nil {
		return nil, err
	}

	return filmFlats, nil
}

func (r *filmRepository) GetFilmByID(ctx context.Context, id uuid.UUID) (dto.FilmCompactResponse, error) {
	var film dto.FilmCompactResponse
	if err := r.db.WithContext(ctx).
		Raw(`
		SELECT f.id, f.judul, f.sinopsis, f.sutradara, f.status, f.durasi,
		       f.total_episode, f.tanggal_rilis,
		       COALESCE(rf.rating, 0) AS rating
		FROM films f
		LEFT JOIN rating_film rf ON rf.film_id = f.id
		WHERE f.id = ?
	`, id).Scan(&film).Error; err != nil {
		return dto.FilmCompactResponse{}, err
	}

	return film, nil
}

func (r *filmRepository) CreateFilm(ctx context.Context, tx *gorm.DB, film *entity.Film) error {
	err := tx.WithContext(ctx).Create(&film).Error
	return err
}

func (r *filmRepository) UpdateFilm(ctx context.Context, filmId uuid.UUID, reqFilm dto.UpdateFilmRequest) (entity.Film, error) {
	var film entity.Film

	err := r.db.WithContext(ctx).Model(&entity.Film{}).Where("id = ?", filmId).First(&film).Error

	if reqFilm.Judul != "" {
		film.Judul = reqFilm.Judul
	}

	if reqFilm.Sinopsis != "" {
		film.Sinopsis = reqFilm.Sinopsis
	}

	if reqFilm.Sutradara != "" {
		film.Sutradara = reqFilm.Sutradara
	}

	if reqFilm.Status != "" {
		film.Status = reqFilm.Status
	}

	if !reqFilm.TanggalRilis.IsZero() {
		film.TanggalRilis = reqFilm.TanggalRilis
	}

	if reqFilm.Durasi != 0 {
		film.Durasi = reqFilm.Durasi
	}

	if reqFilm.TotalEpisode != 0 {
		film.TotalEpisode = reqFilm.TotalEpisode
	}

	err = r.db.WithContext(ctx).Save(&film).Error
	return film, err
}

func (r *filmRepository) DeleteFilm(ctx context.Context, id uuid.UUID) error {
	err := r.db.WithContext(ctx).Delete(&entity.Film{}, id).Error
	return err
}

func (r *filmRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	err := r.db.WithContext(ctx).Model(&entity.Film{}).Where("id = ?", id).Update("status", status).Error
	return err
}

func (r *filmRepository) CheckStatusFilm(ctx context.Context, id uuid.UUID) (entity.Film, error) {
	var film entity.Film
	if err := r.db.WithContext(ctx).Select("id", "status").Where("id = ?", id).First(&film).Error; err != nil {
		return entity.Film{}, err
	}

	return film, nil
}

func (r *filmRepository) SearchFilm(ctx context.Context, req dto.SearchFilmRequest, offset int) ([]entity.Film, error) {
	var films []entity.Film

	baseQuery := r.db.WithContext(ctx).Model(&entity.Film{}).
		Select("id", "judul", "tanggal_rilis", "durasi", "status")

	kw := "%" + req.Keyword + "%"
	baseQuery = baseQuery.Where("judul ILIKE ? ", kw)
	// if req.Status != nil {
	// 	baseQuery = baseQuery.Where("status = ?", *req.Status)
	// }

	// if req.Genres != nil && len(*req.Genres) > 0 {
	// 	baseQuery = baseQuery.Joins("JOIN film_genre ON film_genre.film_id = films.id").
	// 		Where("film_genre.genre_id IN ?", *req.Genres)
	// }

	if err := baseQuery.
		Preload("FilmGambar", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "url", "film_id")
		}).
		Preload("FilmGenre.Genre").
		Order("created_at DESC").
		Limit(helpers.LIMIT_FILM).
		Offset(offset).
		Find(&films).Error; err != nil {
		return nil, err
	}

	return films, nil
}

func (r *filmRepository) CountFilm(ctx context.Context) (int64, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Model(entity.Film{}).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *filmRepository) GetTopFilm(ctx context.Context) ([]dto.FilmCompactResponse, error) {
	var topFilmsFlat []dto.FilmCompactResponse
	err := r.db.WithContext(ctx).
		Raw(`
		SELECT f.id AS id, f.judul, f.status, f.durasi, f.tanggal_rilis, rf.rating
		FROM top_film_watchlist twc
		JOIN films f ON f.id = twc.film_id
		LEFT JOIN rating_film rf ON f.id = rf.film_id
		ORDER BY twc.total_add DESC
		LIMIT 10
	`).Scan(&topFilmsFlat).Error
	if err != nil {
		return nil, err
	}

	return topFilmsFlat, nil
}

func (r *filmRepository) GetTrendingFilm(ctx context.Context) ([]dto.FilmCompactResponse, error) {
	var trendingFilmFlat []dto.FilmCompactResponse
	err := r.db.WithContext(ctx).
		Raw(`
		SELECT f.id AS id, f.judul, f.status, f.durasi, f.tanggal_rilis, rf.rating
		FROM trending_film_weekly tfw
		JOIN films f ON f.id = tfw.film_id
		LEFT JOIN rating_film rf ON f.id = rf.film_id
		ORDER BY tfw.total_added DESC
		LIMIT 10
	`).Scan(&trendingFilmFlat).Error
	if err != nil {
		return nil, err
	}

	return trendingFilmFlat, nil
}

func (r *filmRepository) GetFilmWithMostReviews(ctx context.Context) ([]dto.FilmWithMostReviews, error) {
	var films []dto.FilmWithMostReviews
	err := r.db.WithContext(ctx).
		Raw(`
		SELECT f.id, f.judul, COUNT(r.id) AS count_reviews
		FROM films f
		LEFT JOIN reviews r ON f.id = r.film_id
		GROUP BY f.id
		HAVING COUNT(r.id) > 0
		ORDER BY count_reviews DESC
		LIMIT 10
	`).Scan(&films).Error
	if err != nil {
		return nil, err
	}

	return films, nil
}
