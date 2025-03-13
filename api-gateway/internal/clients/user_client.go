package clients

import (
	userpb "github.com/entonekryzhovnik/user-service/gen/go/userpb" //Подключаем сгенерированные файлы из user-service
	"google.golang.org/grpc"
	"log"
)

func NewUserClient(userServiceAddr string) (userpb.UserServiceClient, error) {
	// Подключаемся к user-service
	conn, err := grpc.Dial(userServiceAddr, grpc.WithInsecure()) // ⚠️ WithInsecure() только для разработки
	if err != nil {
		log.Fatalf("Не удалось подключиться к user-service: %v", err)
		return nil, err
	}

	// Создаём gRPC-клиента
	client := userpb.NewUserServiceClient(conn)
	return client, nil
}
