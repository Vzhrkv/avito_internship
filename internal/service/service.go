package service

import (
	model "github.com/Vzhrkv/avito_internship/internal/database"
	"github.com/Vzhrkv/avito_internship/internal/repository"
)

type UserBalance interface {
	CreateBalance(u *model.User) error
	GetBalance(id uint) (uint, error)
}

type Service struct {
	UserBalance
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserBalance: NewBalanceService(repo.UserBalance),
	}
}
