package service

import (
	model "github.com/Vzhrkv/avito_internship/internal/database"
	"github.com/Vzhrkv/avito_internship/internal/repository"
)

type BalanceService struct {
	repo repository.UserBalance
}

func NewBalanceService(repo repository.UserBalance) *BalanceService {
	return &BalanceService{repo: repo}
}

func (bs *BalanceService) AddBalance(u *model.User) error {
	return bs.repo.AddBalance(u)
}

func (bs *BalanceService) GetBalance(id uint) (uint, error) {
	return bs.repo.GetBalance(id)
}

func (bs *BalanceService) ReserveBalance(user_id uint, service_id uint, order_id uint, price uint) error {
	return bs.repo.ReserveFunds(user_id, service_id, order_id, price)
}

func (bs *BalanceService) ConfirmOrder(user_id uint, service_id uint, order_id uint, price uint) error {
	return bs.repo.ConfirmOrder(user_id, service_id, order_id, price)
}
