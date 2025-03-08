package handlers

import (
	"api-gateway/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	// TODO: Вызывать gRPC CreateTask
	return c.JSON(http.StatusCreated, map[string]string{"message": "Task created"})
}

func (h *TaskHandler) GetTaskByID(c echo.Context) error {
	// TODO: Вызывать gRPC GetTaskByID
	return c.JSON(http.StatusOK, map[string]string{"task": "mock_task"})
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Task updated"})
}

func (h *TaskHandler) GetTaskByUserID(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"task": "mock_task"})
}

func (h *TaskHandler) CompleteTask(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Task completed"})
}
