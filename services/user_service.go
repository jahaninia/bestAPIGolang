package services

import (
    "context"
    "your_project/models"
    "your_project/repositories"
)

type UserService interface {
    Register(ctx context.Context, user models.User) (uint, error)
    GetUser(ctx context.Context, id uint) (models.User, error)
    GetUserByUsername(ctx context.Context, username string) (models.User, error)
    DeleteUser(ctx context.Context, id uint) error
    ListUsers(ctx context.Context) ([]models.User, error)
}

type userService struct {
    repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, user models.User) (uint, error) {
    return s.repo.Create(ctx, user)
}

func (s *userService) GetUser(ctx context.Context, id uint) (models.User, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
    return s.repo.FindByUsername(ctx, username)
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
    return s.repo.Delete(ctx, id)
}

func (s *userService) ListUsers(ctx context.Context) ([]models.User, error) {
    return s.repo.List(ctx)
}