package repository

import (
	"database/sql"
	"errors"
	"github.com/Vzhrkv/avito_internship/internal/database"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (up *UserPostgres) AddBalance(u *model.User) error {
	query := "insert into users (user_id, balance) values ($1, $2)"
	row := up.db.QueryRow(query, u.UserID, u.Balance)
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
	query := "delete from reservedfunds where user_id=($1) and service_id=($2) and order_id=($3) and price=($4)"
	row := up.db.QueryRow(query, user_id, service_id, order_id, price)
	return row.Err()
}
