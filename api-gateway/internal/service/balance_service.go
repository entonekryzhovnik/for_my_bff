package service

import "api-gateway/internal/repository"

type BalanceService struct {
	balanceRepo *repository.BalanceRepository
}

func NewBalanceService(balanceRepo *repository.BalanceRepository) *BalanceService {
	return &BalanceService{balanceRepo: balanceRepo}
}

func (s *BalanceService) GetBalanceByUserID() (string, error) {
	// TODO: Вызывать gRPC GetBalanceByUserID
	return s.balanceRepo.GetBalanceByUserID()
}
