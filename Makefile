# 🚀 Makefile для управления микросервисами

.PHONY: help start stop restart build clean user-service api-gateway all

# 📋 Цвета для вывода
CYAN=\033[0;36m
YELLOW=\033[0;33m
GREEN=\033[0;32m
RED=\033[0;31m
PURPLE=\033[0;35m
NC=\033[0m # No Color

# 📌 Переменные
USER_SERVICE_DIR=./user-service
API_GATEWAY_DIR=./api-gateway

# 📚 Помощь
help: ## 📜 Показывает доступные команды
	@echo "$(CYAN)🚀 Makefile для управления микросервисами$(NC)"
	@echo ""
	@echo "$(YELLOW)Использование:$(NC)"
	@echo "  $(GREEN)make$(NC) $(PURPLE)<команда>$(NC)"
	@echo ""
	@echo "$(YELLOW)Команды:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(NC) %s\n", $$1, $$2}'

# 🏗️ Сборка
build: ## 🏗️ Собирает все сервисы
	@echo "$(CYAN)🏗️ Сборка всех сервисов...$(NC)"
	@make build-user
	@make build-api
	@echo "$(GREEN)✅ Сборка завершена!$(NC)"

build-user: ## 👤 Собирает user-service
	@echo "$(CYAN)🏗️ Сборка user-service...$(NC)"
	@cd $(USER_SERVICE_DIR) && go build -o user-service
	@echo "$(GREEN)✅ user-service собран!$(NC)"

build-api: ## 🔌 Собирает api-gateway
	@echo "$(CYAN)🏗️ Сборка api-gateway...$(NC)"
	@cd $(API_GATEWAY_DIR) && go build -o api-gateway
	@echo "$(GREEN)✅ api-gateway собран!$(NC)"

# 🚀 Запуск
start: ## 🚀 Запускает все сервисы
	@echo "$(CYAN)🚀 Запуск всех сервисов...$(NC)"
	@make start-user
	@sleep 2
	@make start-api
	@echo "$(GREEN)✅ Все сервисы запущены!$(NC)"

start-user: ## 👤 Запускает user-service
	@echo "$(CYAN)🚀 Запуск user-service...$(NC)"
	@cd $(USER_SERVICE_DIR) && ./user-service > /dev/tty 2>&1 &
	@echo "$(GREEN)✅ user-service запущен!$(NC)"

start-api: ## 🔌 Запускает api-gateway
	@echo "$(CYAN)🚀 Запуск api-gateway...$(NC)"
	@cd $(API_GATEWAY_DIR) && ./api-gateway > /dev/tty 2>&1 &
	@echo "$(GREEN)✅ api-gateway запущен!$(NC)"

# ⏹️ Остановка
stop: ## ⏹️ Останавливает все сервисы
	@echo "$(CYAN)⏹️ Остановка всех сервисов...$(NC)"
	@-pkill -f "user-service"
	@-pkill -f "api-gateway"
	@echo "$(GREEN)✅ Все сервисы остановлены!$(NC)"

# 🔄 Перезапуск
restart: ## 🔄 Перезапускает все сервисы
	@echo "$(CYAN)🔄 Перезапуск всех сервисов...$(NC)"
	@make stop
	@sleep 2
	@make start
	@echo "$(GREEN)✅ Все сервисы перезапущены!$(NC)"

# 🧹 Очистка
clean: ## 🧹 Удаляет бинарные файлы
	@echo "$(CYAN)🧹 Очистка проекта...$(NC)"
	@rm -f $(USER_SERVICE_DIR)/user-service
	@rm -f $(API_GATEWAY_DIR)/api-gateway
	@echo "$(GREEN)✅ Очистка завершена!$(NC)"

# 📊 Статус
status: ## 📊 Показывает статус сервисов
	@echo "$(CYAN)📊 Статус сервисов:$(NC)"
	@if pgrep -f "user-service" > /dev/null; then \
		echo "$(GREEN)✅ user-service запущен (PID: $$(pgrep -f "user-service"))$(NC)"; \
	else \
		echo "$(RED)❌ user-service не запущен$(NC)"; \
	fi
	@if pgrep -f "api-gateway" > /dev/null; then \
		echo "$(GREEN)✅ api-gateway запущен (PID: $$(pgrep -f "api-gateway"))$(NC)"; \
	else \
		echo "$(RED)❌ api-gateway не запущен$(NC)"; \
	fi

# 🌐 Проверка доступности сервисов
check: ## 🌐 Проверяет доступность сервисов
	@echo "$(CYAN)🌐 Проверка доступности сервисов...$(NC)"
	@curl -s -o /dev/null -w "$(GREEN)✅ API Gateway: %{http_code}$(NC)\n" http://localhost:8080/users/1 || echo "$(RED)❌ API Gateway недоступен$(NC)"

# По умолчанию
all: help 