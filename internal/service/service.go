package service

import (
	"github.com/Vzhrkv/avito_internship/internal/repository"
)

type UserBalance interface {
	AddBalance(id uint, funds uint) error
	GetBalance(id uint) (uint, error)
	ReserveBalance(userId uint, serviceId uint, orderId uint, price uint) error
	ConfirmOrder(userId uint, serviceId uint, orderId uint, price uint) error
	SendToOtherUser(userId uint, otherUserId uint, funds uint) error
}

type Service struct {
	UserBalance
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserBalance: NewBalanceService(repo.UserBalance),
	}
}
