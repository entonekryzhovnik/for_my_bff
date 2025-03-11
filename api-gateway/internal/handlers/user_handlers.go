package handlers

import (
	"net/http"

	"api-gateway/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

// Конструктор
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Получение пользователя
func (h *UserHandler) GetUser(c echo.Context) error {
	userID := c.Param("id")

	user, err := h.userService.GetUser(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

// Создание пользователя
func (h *UserHandler) CreateUser(c echo.Context) error {
	req := struct {
		Name string `json:"name"`
	}{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	user, err := h.userService.CreateUser(c.Request().Context(), req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}
