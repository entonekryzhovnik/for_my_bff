<p align="center">
  <img src="https://github.com/user-attachments/assets/dacd08d1-0fc3-4d6e-8c4f-815a4d53e455" alt="pet project" />
</p>

---

<p align="center">
 😇🦕💗 вот такой вот питомец-проектик, который я делаю со своим любимым парнем! 😇🦕💗
    <br><br>
  тут будет много всего прикольного, потому что у меня очень умный (и красивый) парень
</p>

# 🚀 Микросервисный проект с gRPC

Современный микросервисный проект с использованием Go, gRPC и REST API.

## 📋 Структура проекта

Проект состоит из нескольких микросервисов:

- **👤 user-service**: Сервис для управления пользователями
- **🔌 api-gateway**: Основной API Gateway, который соединяет все сервисы через gRPC

Будущие сервисы (в разработке):
- **📝 task-service**: Сервис для управления задачами
- **💰 balance-service**: Сервис для управления балансом пользователей

## 🚀 Запуск проекта

### Локальная разработка с использованием Makefile

```bash
# Сборка всех сервисов
make build

# Запуск всех сервисов
make start

# Проверка статуса сервисов
make status

# Проверка доступности API Gateway
make check

# Остановка всех сервисов
make stop

# Перезапуск всех сервисов
make restart

# Очистка бинарных файлов
make clean
```

## 🔄 Тестирование API

Примеры запросов:

### Получение пользователя
```bash
curl -X GET http://localhost:8080/users/1
```

### Создание нового пользователя
```bash
curl -X POST http://localhost:8080/users -d '{"email":"test@example.com"}'
```

## 🛠️ Разработка

### Обновление Proto-файлов

Для обновления proto-файлов в api-gateway из user-service:

```bash
./scripts/update_protos.sh
```

### Структура директорий

```
├── api-gateway/               # API Gateway
│   ├── config/                # Конфигурация
│   ├── gen/                   # Сгенерированные proto-файлы
│   ├── internal/              # Внутренние пакеты
│   │   ├── clients/           # gRPC клиенты
│   │   ├── handlers/          # HTTP обработчики
│   │   ├── repository/        # Репозитории
│   │   └── service/           # Сервисы
│   ├── main.go                # Точка входа
│   └── .env                   # Переменные окружения
├── user-service/              # Сервис пользователей
│   ├── ...
├── Makefile                   # Makefile для управления проектом
└── scripts/                   # Полезные скрипты
    └── update_protos.sh       # Скрипт для обновления proto-файлов
```

## 🔮 Будущие улучшения

- Добавление сервисов task-service и balance-service
- Добавление аутентификации и авторизации
- Мониторинг и логирование

## 🧪 Технологии

- **Golang**: Основной язык программирования
- **gRPC**: Для коммуникации между сервисами
- **Echo**: Веб-фреймворк для API Gateway
- **SQLite**: База данных для user-service
- **Make**: Для автоматизации задач

<p align="center">
  <img src="https://github.com/user-attachments/assets/dacd08d1-0fc3-4d6e-8c4f-815a4d53e455" alt="pet project" />
</p>

---

<p align="center">
 😇🦕💗 вот такой вот питомец-проектик, который я делаю со своим любимым парнем! 😇🦕💗
    <br><br>
  тут будет много всего прикольного, потому что у меня очень умный (и красивый) парень
</p>
