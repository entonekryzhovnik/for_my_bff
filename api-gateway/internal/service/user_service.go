package service

import (
	"api-gateway/gen/go/user"
	"api-gateway/internal/repository"
	"context"
)

type UserService struct {
	userRepo *repository.UserRepository
}

// Конструктор
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Заглушка для получения пользователя
func (s *UserService) GetUser(ctx context.Context, id string) (*user.GetUserResponse, error) {
	return s.userRepo.GetUser(ctx, id)
}

// Заглушка для создания пользователя
func (s *UserService) CreateUser(ctx context.Context, name string) (*user.CreateUserResponse, error) {
	return s.userRepo.CreateUser(ctx, name)
}
