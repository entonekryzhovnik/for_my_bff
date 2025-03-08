package main

import (
	"api-gateway/config"
	"api-gateway/internal/handlers"
	"api-gateway/internal/repository"
	"api-gateway/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	// Загружаем .env
	if err := godotenv.Load(); err != nil {
		log.Warn().Msg("No .env file found, using default values")
	}

	// Загружаем конфиг
	cfg := config.LoadConfig()

	log.Info().Msgf("Starting API Gateway on port %s", cfg.ServerPort)

	// Инициализация логгера
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Starting API Gateway...")

	// Создаем Echo instance
	e := echo.New()

	// Инициализация репозиториев
	userRepo := repository.NewUserRepository()
	taskRepo := repository.NewTaskRepository()
	balanceRepo := repository.NewBalanceRepository()

	// Инициализация сервисов
	userService := service.NewUserService(userRepo)
	taskService := service.NewTaskService(taskRepo)
	balanceService := service.NewBalanceService(balanceRepo)

	// Инициализация хэндлеров
	userHandler := handlers.NewUserHandler(userService)
	taskHandler := handlers.NewTaskHandler(taskService)
	balanceHandler := handlers.NewBalanceHandler(balanceService)

	// Роутинг
	e.POST("/user/create", userHandler.CreateUser)
	e.GET("/user/get", userHandler.GetUser)

	e.POST("/task/create", taskHandler.CreateTask)
	e.GET("/task/get/:taskID", taskHandler.GetTaskByID)
	e.PUT("/task/update", taskHandler.UpdateTask)
	e.GET("/task/get/:userID", taskHandler.GetTaskByUserID)
	e.PUT("/task/complete/:taskID", taskHandler.CompleteTask)

	e.GET("/balance/:userID", balanceHandler.GetBalanceByUserID)

	log.Info().Msg("API Gateway is running on :8080")
	log.Fatal().Err(e.Start(":" + cfg.ServerPort))
}
