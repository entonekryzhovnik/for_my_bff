package handlers

import (
	"api-gateway/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BalanceHandler struct {
	service *service.BalanceService
}

func NewBalanceHandler(service *service.BalanceService) *BalanceHandler {
	return &BalanceHandler{service: service}
}

func (h *BalanceHandler) GetBalanceByUserID(c echo.Context) error {
	// TODO: Вызывать gRPC GetBalanceByUserID
	return c.JSON(http.StatusOK, map[string]string{"balance": "100"})
}
