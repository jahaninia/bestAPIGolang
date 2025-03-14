package repositories

import (
    "context"
    "gorm.io/gorm"
    "your_project/models"
)

type UserRepository interface {
    Create(ctx context.Context, user models.User) (uint, error)
    FindByID(ctx context.Context, id uint) (models.User, error)
    FindByUsername(ctx context.Context, username string) (models.User, error)
    Delete(ctx context.Context, id uint) error
    List(ctx context.Context) ([]models.User, error)
}

type userRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{DB: db}
}

func (r *userRepository) Create(ctx context.Context, user models.User) (uint, error) {
    result := r.DB.WithContext(ctx).Create(&user)
    return user.ID, result.Error
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (models.User, error) {
    var user models.User
    result := r.DB.WithContext(ctx).First(&user, id)
    return user, result.Error
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (models.User, error) {
    var user models.User
    result := r.DB.WithContext(ctx).Where("username = ?", username).First(&user)
    return user, result.Error
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
    result := r.DB.WithContext(ctx).Delete(&models.User{}, id)
    return result.Error
}

func (r *userRepository) List(ctx context.Context) ([]models.User, error) {
    var users []models.User
    result := r.DB.WithContext(ctx).Find(&users)
    return users, result.Error
}