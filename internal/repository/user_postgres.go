package repository

import (
	"database/sql"
	"errors"
	"github.com/Vzhrkv/avito_internship/internal/model"
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
		u.Balance += funds
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

func (up *UserPostgres) SendToOtherUser(userId uint, otherUserId uint, funds uint) error {
	u, err := up.getUser(userId)
	if err != nil {
		return err
	}
	other_u, err := up.getUser(otherUserId)
	if err != nil {
		return err
	}
	if u.Balance >= funds {
		u.Balance -= funds
		other_u.Balance += funds
		query := "update users set balance=($1) where user_id=($2)"
		if err = up.db.QueryRow(query, u.Balance, u.UserID).Err(); err != nil {
			return err
		}
		if err = up.db.QueryRow(query, other_u.Balance, other_u.UserID).Err(); err != nil {
			return err
		}
		return nil
	}
	return errors.New("User don't have such amount of money")
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

func (up *UserPostgres) ReserveFunds(userId uint, serviceId uint, orderId uint, price uint) error {
	user, err := up.getUser(userId)
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
	row := up.db.QueryRow(query_reserv, user.UserID, serviceId, orderId, price)
	return row.Err()
}

func (up *UserPostgres) ConfirmOrder(userId uint, serviceId uint, orderId uint, price uint) error {
	_, err := up.getOrder(userId, serviceId, orderId, price)
	if err != nil {
		return err
	}
	query := "delete from reservedfunds where user_id=($1) and service_id=($2) and order_id=($3) and price=($4)"
	row := up.db.QueryRow(query, userId, serviceId, orderId, price)
	return row.Err()
}

func (up *UserPostgres) getOrder(userId uint, serviceId uint, orderId uint, price uint) (model.Order, error) {
	var ord model.Order
	query := "select * from reservedfunds where user_id=($1) and service_id=($2) and order_id=($3) and price=($4)"
	row := up.db.QueryRow(query, userId, serviceId, orderId, price)
	err := row.Scan(&ord.UserID, &ord.ServiceID, &ord.OrderID, &ord.Price)
	if err != nil {
		return model.Order{}, err
	}
	return ord, nil
}
