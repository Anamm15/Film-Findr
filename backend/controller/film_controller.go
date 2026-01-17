package controller

import (
	"fmt"

	"FilmFindr/dto"
	"FilmFindr/service"
	"FilmFindr/utils"

	"github.com/gin-gonic/gin"
)

type FilmController interface {
	GetAllFilm(ctx *gin.Context)
	GetFilmById(ctx *gin.Context)
	GetTopFilm(ctx *gin.Context)
	GetTrendingFilm(ctx *gin.Context)
	CreateFilm(ctx *gin.Context)
	UpdateFilm(ctx *gin.Context)
	DeleteFilm(ctx *gin.Context)
	AddFilmGenre(ctx *gin.Context)
	DeleteFilmGenre(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	SearchFilm(ctx *gin.Context)
}

type filmController struct {
	filmService      service.FilmService
	filmGenreService service.FilmGenreService
}

func NewFilmController(
	filmService service.FilmService,
	filmGenreService service.FilmGenreService,
) FilmController {
	return &filmController{
		filmService:      filmService,
		filmGenreService: filmGenreService,
	}
}

func (s *filmController) GetAllFilm(ctx *gin.Context) {
	page := ctx.Query("page")
	films, err := s.filmService.GetAllFilm(ctx, page)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_FILM, films)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) GetFilmById(ctx *gin.Context) {
	filmId := ctx.Param("id")

	film, err := s.filmService.GetFilmByID(ctx, filmId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_FILM, film)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) CreateFilm(ctx *gin.Context) {
	var film dto.CreateFilmRequest
	if err := ctx.ShouldBind(&film); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_INTERNAL_ERROR, err.Error(), nil)
		ctx.JSON(dto.STATUS_INTERNAL_SERVER_ERROR, res)
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_IMAGE_NOT_FOUND, dto.ErrGetImageRequest.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	createdFilm, err := s.filmService.CreateFilm(ctx, film, files)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATED_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_INTERNAL_SERVER_ERROR, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATED_FILM, createdFilm)
	ctx.JSON(dto.STATUS_CREATED, res)
}

func (s *filmController) UpdateFilm(ctx *gin.Context) {
	filmId := ctx.Param("id")

	var film dto.UpdateFilmRequest
	if err := ctx.ShouldBindJSON(&film); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		fmt.Println(err)
		return
	}

	err := s.filmService.UpdateFilm(ctx, filmId, film)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		fmt.Println(err)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_FILM, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) DeleteFilm(ctx *gin.Context) {
	filmId := ctx.Param("id")

	err := s.filmService.DeleteFilm(ctx, filmId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETED_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETED_FILM, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) AddFilmGenre(ctx *gin.Context) {
	var filmGenre dto.FilmGenreRequest
	if err := ctx.ShouldBindJSON(&filmGenre); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	createdFilmGenre, err := s.filmGenreService.CreateFilmGenre(ctx, filmGenre)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATED_FILM_GENRE, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATED_FILM_GENRE, createdFilmGenre)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) DeleteFilmGenre(ctx *gin.Context) {
	var filmGenre dto.FilmGenreRequest
	if err := ctx.ShouldBindJSON(&filmGenre); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err := s.filmGenreService.DeleteFilmGenre(ctx, filmGenre)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETED_FILM_GENRE, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETED_FILM_GENRE, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) UpdateStatus(ctx *gin.Context) {
	filmId := ctx.Param("id")
	var req dto.UpdateStatusFilmRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err := s.filmService.UpdateStatus(ctx, filmId, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_STATUS_FILM, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) SearchFilm(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	// status := ctx.Query("status")
	// genreIDs := ctx.QueryArray("genre_ids")
	page := ctx.Query("page")

	// var genresId []int
	// if len(genreIDs) > 0 {
	// 	genresId = make([]int, len(genreIDs))
	// 	for i, genreID := range genreIDs {
	// 		genresId[i] = utils.StringToInt(genreID)
	// 	}
	// }

	var req dto.SearchFilmRequest

	// if keyword != "" {
	// 	req.Keyword = &keyword
	// }

	// if status != "" {
	// 	req.Status = &status
	// }

	// if len(genresId) > 0 {
	// 	req.Genres = &genresId
	// }

	req.Keyword = keyword
	films, err := s.filmService.SearchFilm(ctx, req, page)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_SEARCH_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	if len(films.Film) == 0 {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_FILM_NOT_FOUND, dto.MESSAGE_FAILED_FILM_NOT_FOUND, nil)
		ctx.JSON(dto.STATUS_NOT_FOUND, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_SEARCH_FILM, films)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) GetTopFilm(ctx *gin.Context) {
	films, err := s.filmService.GetTopFilm(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_FILM, films)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) GetTrendingFilm(ctx *gin.Context) {
	films, err := s.filmService.GetTrendingFilm(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_FILM, films)
	ctx.JSON(dto.STATUS_OK, res)
}
