package clients

import (
	"context"

	userpb "github.com/entonekryzhovnik/user-service/gen/go/userpb" //Подключаем сгенерированные файлы из user-service
)

// NewUserClient создает новый клиент для user-service
func NewUserClient(userServiceAddr string) (userpb.UserServiceClient, error) {
	// Создаем менеджер клиентов
	manager := NewGRPCClientManager(userServiceAddr)

	// Запускаем соединения
	if err := manager.Start(context.Background()); err != nil {
		return nil, err
	}

	// Возвращаем клиент
	return manager.UserService(), nil
}
