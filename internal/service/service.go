package service

import (
	"github.com/Vzhrkv/avito_internship/internal/repository"
)

type UserBalance interface {
}

type Service struct {
	UserBalance
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
