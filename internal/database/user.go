package model

type User struct {
	UserID        uint `json:"user_id"`
	Balance       uint `json:"balance"`
	ReservedFunds uint `json:"reserved_funds"`
}
