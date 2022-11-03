package service

import (
	model "github.com/Vzhrkv/avito_internship/internal/database"
	"github.com/Vzhrkv/avito_internship/internal/repository"
)

type UserBalance interface {
	AddBalance(u *model.User) error
	GetBalance(id uint) (uint, error)
	ReserveBalance(user_id uint, service_id uint, order_id uint, price uint) error
	ConfirmOrder(user_id uint, service_id uint, order_id uint, price uint) error
}

type Service struct {
	UserBalance
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserBalance: NewBalanceService(repo.UserBalance),
	}
}
