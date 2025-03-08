package handlers

import (
	"api-gateway/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	// TODO: Вызывать gRPC CreateUser
	return c.JSON(http.StatusCreated, map[string]string{"message": "User created"})
}

func (h *UserHandler) GetUser(c echo.Context) error {
	// TODO: Вызывать gRPC GetUser
	return c.JSON(http.StatusOK, map[string]string{"user": "mock_user"})
}
