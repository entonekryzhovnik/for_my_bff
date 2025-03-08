package repository

type BalanceRepository struct{}

func NewBalanceRepository() *BalanceRepository {
	return &BalanceRepository{}
}

func (r *BalanceRepository) GetBalanceByUserID() (string, error) {
	// TODO: gRPC вызов к balance-service
	return "100", nil
}
