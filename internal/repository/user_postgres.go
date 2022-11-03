package repository

import (
	"database/sql"
	"fmt"
	"github.com/Vzhrkv/avito_internship/internal/database"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) CreateBalance(u *model.User) error {
	query := fmt.Sprintf("insert into %s (user_id, balance) values ($1, $2)", usersTable)
	row := up.db.QueryRow(query, u.UserID, u.Balance)
	return row.Err()
}

func (up *UserPostgres) GetBalance(id uint) (uint, error) {
	var money uint
	query := fmt.Sprintf("select balance from %s where user_id=($1)", usersTable)
	row := up.db.QueryRow(query, id)
	err := row.Scan(&money)
	if err != nil {
		return 0, err
	}
	return money, nil
}
