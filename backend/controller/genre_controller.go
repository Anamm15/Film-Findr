package controller

import (
	"FilmFindr/dto"
	"FilmFindr/service"
	"FilmFindr/utils"

	"github.com/gin-gonic/gin"
)

type GenreController interface {
	GetAllGenre(ctx *gin.Context)
	CreateGenre(ctx *gin.Context)
	DeleteGenre(ctx *gin.Context)
}

type genreController struct {
	genreService service.GenreService
}

func NewGenreController(genreService service.GenreService) GenreController {
	return &genreController{
		genreService: genreService,
	}
}

func (s *genreController) GetAllGenre(ctx *gin.Context) {
	genres, err := s.genreService.GetAllGenre(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_GENRE, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_GENRE, genres)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *genreController) CreateGenre(ctx *gin.Context) {
	var genre dto.GenreRequest
	if err := ctx.ShouldBindJSON(&genre); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	createdGenre, err := s.genreService.CreateGenre(ctx, genre)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATED_GENRE, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_GENRE_CREATED, createdGenre)
	ctx.JSON(dto.STATUS_CREATED, res)
}

func (s *genreController) DeleteGenre(ctx *gin.Context) {
	genreId := ctx.Param("id")
	var genre dto.GenreRequest
	if err := ctx.ShouldBindJSON(&genre); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err := s.genreService.DeleteGenre(ctx, genreId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_GENRE, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_GENRE_UPDATED, nil)
	ctx.JSON(dto.STATUS_OK, res)
}
