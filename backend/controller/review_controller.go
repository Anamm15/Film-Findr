package controller

import (
	"FilmFindr/dto"
	"FilmFindr/service"
	"FilmFindr/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReviewController interface {
	GetReviewByUserId(ctx *gin.Context)
	GetReviewByFilmId(ctx *gin.Context)
	CreateReview(ctx *gin.Context)
	UpdateReview(ctx *gin.Context)
	UpdateReaksiReview(ctx *gin.Context)
	DeleteReview(ctx *gin.Context)
}

type reviewController struct {
	reviewService service.ReviewService
	jwtService    service.JWTService
}

func NewReviewController(
	reviewService service.ReviewService,
	jwtService service.JWTService,
) ReviewController {
	return &reviewController{
		reviewService: reviewService,
		jwtService:    jwtService,
	}
}

func (c *reviewController) GetReviewByUserId(ctx *gin.Context) {
	userReviewId := ctx.Param("id")
	page := ctx.Query("page")

	tokenStr, _ := ctx.Cookie("access_token")
	userId, _, _ := c.jwtService.GetDataByToken(tokenStr)
	reviews, err := c.reviewService.GetReviewByUserId(ctx, userReviewId, userId, page)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_REVIEW, reviews)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *reviewController) GetReviewByFilmId(ctx *gin.Context) {
	filmId := ctx.Param("id")
	page := ctx.Query("page")

	tokenStr, _ := ctx.Cookie("access_token")
	userId, _, _ := c.jwtService.GetDataByToken(tokenStr)
	reviews, err := c.reviewService.GetReviewByFilmId(ctx, filmId, userId, page)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_REVIEW, reviews)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *reviewController) CreateReview(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uuid.UUID)

	var review dto.CreateReviewRequest
	if err := ctx.ShouldBindJSON(&review); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	createdReview, err := c.reviewService.CreateReview(ctx, review, userID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATED_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATED_REVIEW, createdReview)
	ctx.JSON(dto.STATUS_CREATED, res)
}

func (c *reviewController) UpdateReview(ctx *gin.Context) {
	reviewId := ctx.Param("id")

	var review dto.UpdateReviewRequest
	if err := ctx.ShouldBindJSON(&review); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err := c.reviewService.UpdateReview(ctx, reviewId, review)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_REVIEW, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *reviewController) UpdateReaksiReview(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)
	reviewId := ctx.Param("id")

	var reaksi dto.UpdateReaksiReviewRequest
	if err := ctx.ShouldBindJSON(&reaksi); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err := c.reviewService.UpdateReaksiReview(ctx, reviewId, userId, reaksi)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_REVIEW, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *reviewController) DeleteReview(ctx *gin.Context) {
	reviewId := ctx.Param("id")

	err := c.reviewService.DeleteReview(ctx, reviewId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETED_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETED_REVIEW, nil)
	ctx.JSON(dto.STATUS_OK, res)
}
