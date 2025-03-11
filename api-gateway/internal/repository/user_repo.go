package repository

import (
	"api-gateway/gen/go/user"
	"context"
)

type UserRepository struct{}

// Конструктор
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Заглушка для получения пользователя
func (r *UserRepository) GetUser(ctx context.Context, id string) (*user.GetUserResponse, error) {
	return &user.GetUserResponse{Name: "John Doe"}, nil
}

// Заглушка для создания пользователя
func (r *UserRepository) CreateUser(ctx context.Context, name string) (*user.CreateUserResponse, error) {
	return &user.CreateUserResponse{Id: "12345"}, nil
}
