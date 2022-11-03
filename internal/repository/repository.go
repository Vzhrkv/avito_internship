package repository

import (
	"database/sql"
	model "github.com/Vzhrkv/avito_internship/internal/database"
)

type UserBalance interface {
	CreateBalance(u *model.User) error
	GetBalance(id uint) (uint, error)
}

type Repository struct {
	UserBalance
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserBalance: NewUserPostgres(db),
	}
}
