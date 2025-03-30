package clients

import (
	"context"
	"errors"
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	userpb "github.com/entonekryzhovnik/user-service/gen/go/userpb"
	// Заглушки для будущих сервисов будут добавлены позже
	// taskpb "api-gateway/gen/go/task"
	// balancepb "api-gateway/gen/go/balance"
)

// Определяем ошибки
var (
	ErrServiceNotAvailable = errors.New("service is not available")
)

// GRPCClientManager управляет всеми gRPC соединениями
type GRPCClientManager struct {
	mu sync.Mutex

	// Соединения
	userConn *grpc.ClientConn
	// taskConn     *grpc.ClientConn     // Будет реализовано позже
	// balanceConn  *grpc.ClientConn     // Будет реализовано позже

	// Клиенты
	userClient userpb.UserServiceClient
	// taskClient    taskpb.TaskServiceClient    // Будет реализовано позже
	// balanceClient balancepb.BalanceServiceClient // Будет реализовано позже

	// Адреса сервисов
	userServiceAddr string
	// taskServiceAddr    string    // Будет реализовано позже
	// balanceServiceAddr string    // Будет реализовано позже
}

// NewGRPCClientManager создает новый экземпляр менеджера клиентов
func NewGRPCClientManager(userServiceAddr string) *GRPCClientManager {
	return &GRPCClientManager{
		userServiceAddr: userServiceAddr,
		// taskServiceAddr: taskServiceAddr,       // Будет реализовано позже
		// balanceServiceAddr: balanceServiceAddr, // Будет реализовано позже
	}
}

// Start устанавливает все соединения с gRPC сервисами
func (m *GRPCClientManager) Start(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var err error

	// Подключение к сервису пользователей
	m.userConn, err = grpc.DialContext(ctx, m.userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Не удалось подключиться к user-service: %v", err)
		return err
	}
	m.userClient = userpb.NewUserServiceClient(m.userConn)

	// Здесь будут подключения к будущим сервисам
	// ...

	return nil
}

// Close закрывает все соединения
func (m *GRPCClientManager) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.userConn != nil {
		if err := m.userConn.Close(); err != nil {
			log.Printf("Ошибка при закрытии соединения с user-service: %v", err)
			return err
		}
		m.userConn = nil
	}

	// Здесь будет закрытие соединений с будущими сервисами
	// ...

	return nil
}

// UserService возвращает клиент для работы с сервисом пользователей
func (m *GRPCClientManager) UserService() userpb.UserServiceClient {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.userClient
}

// Будущие методы для доступа к другим сервисам будут реализованы позже
// TaskService() - для task-service
// BalanceService() - для balance-service
