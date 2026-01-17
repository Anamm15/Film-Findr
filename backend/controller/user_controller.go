package controller

import (
	"FilmFindr/dto"
	"FilmFindr/service"
	"FilmFindr/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	GetAllUser(ctx *gin.Context)
	GetUserByUsername(ctx *gin.Context)
	Me(ctx *gin.Context)
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	LogoutUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(
	userService service.UserService,
	jwtService service.JWTService,
) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) GetAllUser(ctx *gin.Context) {
	users, err := c.userService.GetAllUser(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_USER, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_USER, users)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *userController) GetUserByUsername(ctx *gin.Context) {
	usernameParam := ctx.Query("username")
	userResponse, err := c.userService.GetUserByUsername(ctx, usernameParam)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, userResponse)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *userController) Me(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)
	userResponse, err := c.userService.GetUserById(ctx, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, userResponse)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *userController) RegisterUser(ctx *gin.Context) {
	var user dto.UserCreateRequest
	if err := ctx.ShouldBind(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	photoProfil, err := ctx.FormFile("photo_profil")
	var userResponse dto.UserResponse
	if err != nil {
		userResponse, err = c.userService.RegisterUser(ctx, user, nil)
	} else {
		userResponse, err = c.userService.RegisterUser(ctx, user, photoProfil)
	}

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATED_USER, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATED_USER, userResponse)
	ctx.JSON(dto.STATUS_CREATED, res)
}

func (c *userController) LoginUser(ctx *gin.Context) {
	var req dto.UserLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.AbortWithStatusJSON(dto.STATUS_BAD_REQUEST, response)
		return
	}

	user, err := c.userService.LoginUser(ctx.Request.Context(), req)
	if err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		ctx.AbortWithStatusJSON(dto.STATUS_BAD_REQUEST, response)
		return
	}

	token := c.jwtService.GenerateToken(user.ID, user.Role)
	userResponse := dto.AuthorizationRequest{
		Token: token,
		Role:  user.Role,
	}

	ctx.SetCookie(
		"access_token",
		token,
		3600*12,
		"/",
		"",
		false,
		true,
	)

	response := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, userResponse)
	ctx.JSON(dto.STATUS_OK, response)
}

func (c *userController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "", false, true)

	response := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGOUT, nil)
	ctx.JSON(dto.STATUS_OK, response)
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uuid.UUID)

	var user dto.UserUpdateRequest
	user.ID = userId
	if err := ctx.ShouldBindJSON(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	photoProfil, err := ctx.FormFile("photo_profil")
	if err == nil {
		err = c.userService.UpdateUser(ctx, user, photoProfil)
	} else {
		err = c.userService.UpdateUser(ctx, user, nil)
	}

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_USER, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_USER, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	err := c.userService.DeleteUser(ctx, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETED_USER, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETED_USER, nil)
	ctx.JSON(dto.STATUS_OK, res)
}
