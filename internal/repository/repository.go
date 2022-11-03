package repository

import (
	"database/sql"
	model "github.com/Vzhrkv/avito_internship/internal/database"
)

type UserBalance interface {
	AddBalance(u *model.User) error
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
