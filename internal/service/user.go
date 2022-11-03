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

func (bs *BalanceService) CreateBalance(u *model.User) error {
	return bs.repo.CreateBalance(u)
}

func (bs *BalanceService) GetBalance(id uint) (uint, error) {
	return bs.repo.GetBalance(id)
}
