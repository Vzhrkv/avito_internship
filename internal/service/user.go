package service

import (
	"github.com/Vzhrkv/avito_internship/internal/repository"
)

type BalanceService struct {
	repo repository.UserBalance
}

func NewBalanceService(repo repository.UserBalance) *BalanceService {
	return &BalanceService{repo: repo}
}

func (bs *BalanceService) AddBalance(id uint, funds uint) error {
	return bs.repo.AddBalance(id, funds)
}

func (bs *BalanceService) GetBalance(id uint) (uint, error) {
	return bs.repo.GetBalance(id)
}

func (bs *BalanceService) ReserveBalance(userId uint, serviceId uint, orderId uint, price uint) error {
	return bs.repo.ReserveFunds(userId, serviceId, orderId, price)
}

func (bs *BalanceService) ConfirmOrder(userId uint, serviceId uint, orderId uint, price uint) error {
	return bs.repo.ConfirmOrder(userId, serviceId, orderId, price)
}

func (bs *BalanceService) SendToOtherUser(userId uint, otherUserId uint, funds uint) error {
	return bs.repo.SendToOtherUser(userId, otherUserId, funds)
}
