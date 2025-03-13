package handlers

import (
	"context"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"

	"github.com/entonekryzhovnik/user-service/gen/go/userpb" // gRPC-прото
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userClient userpb.UserServiceClient
}

// Конструктор для UserHandler
func NewUserHandler(userClient userpb.UserServiceClient) *UserHandler {
	return &UserHandler{userClient: userClient}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	// Парсим userID из строки в int64
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user ID"})
	}

	// Создаём контекст
	ctx := context.Background()

	// Вызываем gRPC метод
	res, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: userID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Возвращаем JSON
	return c.JSON(http.StatusOK, map[string]string{"name": res.User.Email})
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req userpb.CreateUserRequest // Обнови тип на CreateUserRequest

	// Приводим тело запроса к нужному формату
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request body")
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Вызов метода gRPC клиента для создания пользователя
	createdUser, err := h.userClient.CreateUser(c.Request().Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create user via gRPC")
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, createdUser)
}
