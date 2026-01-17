package repository

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) ([]entity.User, error)
	GetUserById(ctx context.Context, userID uuid.UUID) (entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
	UpdateUser(ctx context.Context, user dto.UserUpdateRequest) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	CountUsers(ctx context.Context) (int64, error)
	GetWeeklyUsers(ctx context.Context) ([]dto.WeeklyUser, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUser(ctx context.Context) ([]entity.User, error) {
	var user []entity.User
	if err := r.db.WithContext(ctx).Select("id", "username", "nama", "bio", "photo_profil", "role").
		Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserById(ctx context.Context, userID uuid.UUID) (entity.User, error) {
	var user entity.User

	if err := r.db.WithContext(ctx).
		Select("id", "username", "nama", "bio", "photo_profil", "role").
		Where("id = ?", userID).
		First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Select("id", "username", "nama", "bio", "photo_profil", "password", "role").
		Where("username = ?", username).
		Take(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, userReq dto.UserUpdateRequest) error {
	var user entity.User
	if err := r.db.WithContext(ctx).First(&user, userReq.ID).Error; err != nil {
		return err
	}

	if userReq.Nama != "" {
		user.Nama = userReq.Nama
	}

	if userReq.Username != "" {
		user.Username = userReq.Username
	}

	if userReq.Bio != "" {
		user.Bio = userReq.Bio
	}

	if userReq.OldPhotoProfil != user.PhotoProfil && userReq.OldPhotoProfil != "" {
		user.PhotoProfil = userReq.OldPhotoProfil
	}

	if err := r.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, &userID).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) CountUsers(ctx context.Context) (int64, error) {
	var count int64

	err := r.db.WithContext(ctx).
		Model(entity.User{}).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *userRepository) GetWeeklyUsers(ctx context.Context) ([]dto.WeeklyUser, error) {
	var results []dto.WeeklyUser

	err := r.db.WithContext(ctx).
		Raw("SELECT * FROM weekly_user ORDER BY weekly DESC LIMIT 8").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
