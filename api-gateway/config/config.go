package config

import (
	"os"
)

type Config struct {
	GRPCUserService    string
	GRPCTaskService    string
	GRPCBalanceService string
	ServerPort         string
}

func LoadConfig() *Config {
	return &Config{
		GRPCUserService:    getEnv("GRPC_USER_SERVICE", "localhost:50051"),
		GRPCTaskService:    getEnv("GRPC_TASK_SERVICE", "localhost:50052"),
		GRPCBalanceService: getEnv("GRPC_BALANCE_SERVICE", "localhost:50053"),
		ServerPort:         getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
