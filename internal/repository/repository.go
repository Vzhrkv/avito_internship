package repository

import (
	"database/sql"
)

type UserBalance interface {
	AddBalance(id uint, funds uint) error
	GetBalance(id uint) (uint, error)
	ReserveFunds(user_id uint, service_id uint, order_id uint, price uint) error
	ConfirmOrder(user_id uint, service_id uint, order_id uint, price uint) error
}

type Repository struct {
	UserBalance
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserBalance: NewUserPostgres(db),
	}
}
