package main

import (
	"api-gateway/config"
	"api-gateway/internal/clients"
	"api-gateway/internal/handlers"
	"api-gateway/internal/repository"
	"api-gateway/internal/service"
	"context"
	"os"
	"os/signal"
	"syscall"

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

	// Создаем контекст с отменой для graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Инициализация grpc клиент-менеджера
	grpcManager := clients.NewGRPCClientManager(cfg.GRPCUserService)

	if err := grpcManager.Start(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to start gRPC client manager")
	}
	defer func() {
		if err := grpcManager.Close(); err != nil {
			log.Error().Err(err).Msg("Error while closing gRPC connections")
		}
	}()

	// Инициализация репозиториев
	//userRepo := repository.NewUserRepository()
	taskRepo := repository.NewTaskRepository()
	balanceRepo := repository.NewBalanceRepository()

	// Инициализация сервисов
	//userService := service.NewUserService(userRepo)
	taskService := service.NewTaskService(taskRepo)
	balanceService := service.NewBalanceService(balanceRepo)

	// Инициализация хэндлеров
	userHandler := handlers.NewUserHandler(grpcManager.UserService())
	taskHandler := handlers.NewTaskHandler(taskService)
	balanceHandler := handlers.NewBalanceHandler(balanceService)

	// Роутинг
	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/users", userHandler.CreateUser)

	e.POST("/task/create", taskHandler.CreateTask)
	e.GET("/task/get/:taskID", taskHandler.GetTaskByID)
	e.PUT("/task/update", taskHandler.UpdateTask)
	e.GET("/task/get/:userID", taskHandler.GetTaskByUserID)
	e.PUT("/task/complete/:taskID", taskHandler.CompleteTask)

	e.GET("/balance/:userID", balanceHandler.GetBalanceByUserID)

	// Graceful shutdown
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Info().Msg("Shutting down API Gateway...")
		cancel()
		if err := e.Shutdown(ctx); err != nil {
			log.Fatal().Err(err).Msg("Error during server shutdown")
		}
	}()

	log.Info().Msg("API Gateway is running on :" + cfg.ServerPort)
	if err := e.Start(":" + cfg.ServerPort); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
