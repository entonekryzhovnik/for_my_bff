package repository

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser() error {
	// TODO: gRPC вызов к user-service
	return nil
}

func (r *UserRepository) GetUser() (string, error) {
	// TODO: gRPC вызов к user-service
	return "mock_user", nil
}
