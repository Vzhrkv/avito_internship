package repository

import "database/sql"

type UserBalance interface {
}

type Repository struct {
	UserBalance
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
