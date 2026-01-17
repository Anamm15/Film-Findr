package dto

import (
	"errors"

	"FilmFindr/entity"
	"FilmFindr/helpers"

	"github.com/google/uuid"
)

const (
	// failed
	MESSAGE_USER_NOT_FOUND        = "User not found"
	MESSAGE_FAILED_GET_LIST_USER  = "Failed get list user"
	MESSAGE_FAILED_GET_USER       = "Failed get user"
	MESSAGE_FAILED_CREATED_USER   = "Failed created user"
	MESSAGE_FAILED_UPDATED_USER   = "Failed updated user"
	MESSAGE_FAILED_DELETED_USER   = "Failed deleted user"
	MESSAGE_FAILED_LOGOUT         = "Failed to logout"
	MESSAGE_FAILED_USER_NOT_LOGIN = "User not login"

	// success message
	MESSAGE_SUCCESS_GET_LIST_USER = "Success get list of user"
	MESSAGE_SUCCESS_GET_USER      = "Success get user"
	MESSAGE_SUCCESS_LOGIN         = "Login successfully"
	MESSAGE_SUCCESS_CREATED_USER  = "User created successfully"
	MESSAGE_SUCCESS_UPDATED_USER  = "User updated successfully"
	MESSAGE_SUCCESS_DELETED_USER  = "User deleted successfully"
	MESSAGE_SUCCESS_LOGOUT        = "Logout successfully"
)

var (
	ErrGetAllUser        = errors.New("Failed to get all user")
	ErrGetUserByID       = errors.New("Failed to get user by id")
	ErrCreateUser        = errors.New("Failed to create user")
	ErrCheckUsername     = errors.New("Username is not found")
	ErrUpdateUser        = errors.New("Failed to update user")
	ErrDeleteUser        = errors.New("Failed to delete user")
	ErrFailedLogin       = errors.New("Failed to login")
	ErrPasswordNotMatch  = errors.New("Password not match")
	ErrEmailOrPassword   = errors.New("Email or password not match")
	ErrUserNotFound      = errors.New("User not found")
	ErrUserNotLogin      = errors.New("User not login")
	ErrFailedSaveSession = errors.New("Failed to save session")
	ErrFailedUploadFile  = errors.New("Failed to upload file")
)

type (
	UserCreateRequest struct {
		Nama     string `json:"nama" form:"nama" binding:"required"`
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
		Bio      string `json:"bio" form:"bio"`
	}

	UserLoginRequest struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserFilm struct {
		ID     uuid.UUID    `json:"id"`
		Status string       `json:"status"`
		Film   FilmResponse `json:"film"`
	}

	UserRequest struct {
		Nama        string `json:"nama" form:"nama" binding:"required"`
		Username    string `json:"username" form:"username" binding:"required"`
		Bio         string `json:"bio" form:"bio"`
		PhotoProfil string `json:"photo_profil" form:"photo_profil"`
	}

	UserResponse struct {
		ID          uuid.UUID `json:"id"`
		Nama        string    `json:"nama"`
		Username    string    `json:"username"`
		Bio         string    `json:"bio"`
		PhotoProfil string    `json:"photo_profil"`
		Role        string    `json:"role"`
	}

	UserUpdateRequest struct {
		ID             uuid.UUID `json:"id"`
		Nama           string    `json:"nama" form:"nama"`
		Bio            string    `json:"bio" form:"bio"`
		Username       string    `json:"username" form:"username"`
		OldPhotoProfil string    `json:"old_photo_profil"`
	}

	UserResponseWithFilm struct {
		ID          uuid.UUID  `json:"id"`
		Nama        string     `json:"nama"`
		Username    string     `json:"username"`
		Bio         string     `json:"bio"`
		PhotoProfil string     `json:"photo_profil"`
		UserFilm    []UserFilm `json:"user_film"`
	}
)

func EntityToUserResponse(user entity.User) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Nama:        user.Nama,
		Username:    user.Username,
		Bio:         user.Bio,
		PhotoProfil: user.PhotoProfil,
		Role:        user.Role,
	}
}

func (r *UserCreateRequest) ToModel(user *entity.User, profileURL string) {
	user.Nama = r.Nama
	user.Username = r.Username
	user.Password = r.Password
	user.Bio = r.Bio
	user.PhotoProfil = profileURL
	user.Role = helpers.ENUM_ROLE_USER
}
