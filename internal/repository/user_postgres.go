package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Vzhrkv/avito_internship/internal/database"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) AddBalance(id uint, funds uint) error {
	u, err := up.getUser(id)
	if err == nil {
		fmt.Println(u)
		u.Balance += funds
		fmt.Println()
		query := "update users set balance=($1) where user_id=($2)"
		return up.db.QueryRow(query, u.Balance, id).Err()
	}
	query := "insert into users (user_id, balance) values ($1, $2)"
	row := up.db.QueryRow(query, id, funds)
	return row.Err()
}

func (up *UserPostgres) GetBalance(id uint) (uint, error) {
	var money uint
	query := "select balance from users where user_id=($1)"
	row := up.db.QueryRow(query, id)
	err := row.Scan(&money)
	if err != nil {
		return 0, err
	}
	return money, nil
}

func (up *UserPostgres) getUser(id uint) (model.User, error) {
	var u model.User
	query := "select * from users where user_id=($1)"
	row := up.db.QueryRow(query, id)
	err := row.Scan(&u.UserID, &u.Balance)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (up *UserPostgres) ReserveFunds(user_id uint, service_id uint, order_id uint, price uint) error {
	user, err := up.getUser(user_id)
	if err != nil {
		return err
	}
	if user.Balance < price {
		return errors.New("User don't required mount of money to book this service")
	}

	user.Balance -= price

	query_users := "update users set balance=($1) where user_id=($2)"
	up.db.QueryRow(query_users, user.Balance, user.UserID)
	query_reserv := "insert into reservedfunds (user_id, service_id, order_id, price) values ($1, $2, $3, $4)"
	row := up.db.QueryRow(query_reserv, user.UserID, service_id, order_id, price)
	return row.Err()
}

func (up *UserPostgres) ConfirmOrder(user_id uint, service_id uint, order_id uint, price uint) error {
	var numRow uint
	existsQuery := "select * from reservedfunds where user_id=($1) and service_id=($2) and order_id=($3) and price=($4)"
	existsRow := up.db.QueryRow(existsQuery, user_id, service_id, order_id, price)
	err := existsRow.Scan(&numRow)
	if err != nil {
		return err
	}
	query := "delete from reservedfunds where user_id=($1) and service_id=($2) and order_id=($3) and price=($4)"
	row := up.db.QueryRow(query, user_id, service_id, order_id, price)
	return row.Err()

}
