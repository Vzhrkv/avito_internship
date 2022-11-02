package repository

type UserBalance interface {
}

type Repository struct {
	UserBalance
}

func NewRepository() *Repository {
	return &Repository{}
}
