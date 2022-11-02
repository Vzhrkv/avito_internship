package service

import "github.com/Vzhrkv/avito_internship/pkg/repository"

type UserBalance interface {
}

type Service struct {
	UserBalance
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
