package service

import "api-gateway/internal/repository"

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser() error {
	// TODO: Вызывать CreateUser у репозитория
	return s.userRepo.CreateUser()
}

func (s *UserService) GetUser() (string, error) {
	// TODO: Вызывать GetUser у репозитория
	return s.userRepo.GetUser()
}
