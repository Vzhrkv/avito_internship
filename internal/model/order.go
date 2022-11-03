package model

type Order struct {
	UserID    uint `json:"user_id"`
	ServiceID uint `json:"service_id"`
	OrderID   uint `json:"order_id"`
	Price     uint `json:"price"`
}
