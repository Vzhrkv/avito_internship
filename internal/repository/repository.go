package repository

import (
	"database/sql"
)

type UserBalance interface {
	AddBalance(id uint, funds uint) error
	GetBalance(id uint) (uint, error)
	ReserveFunds(userId uint, serviceId uint, orderId uint, price uint) error
	ConfirmOrder(userId uint, serviceId uint, orderId uint, price uint) error
	SendToOtherUser(userId uint, otherUserId uint, funds uint) error
}

type Repository struct {
	UserBalance
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserBalance: NewUserPostgres(db),
	}
}
