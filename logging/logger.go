package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type ConfirmedOrder struct {
	UserID    uint `json:"user_id"`
	ServiceID uint `json:"service_id"`
	OrderID   uint `json:"order_id"`
	Price     uint `json:"price"`
}

func LogToFile(in *ConfirmedOrder) {
	file, err := os.OpenFile("logs/confirmation.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		logrus.Println(err)
	}

	msg := fmt.Sprintf("User with user_id=%d confirmed order: service_id=%d, order_id=%d, price=%d",
		in.UserID, in.ServiceID, in.OrderID, in.Price)

	logrus.SetOutput(file)
	logrus.Error(msg)
}