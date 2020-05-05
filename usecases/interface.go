package usecases

import "github.com/felipehfs/clean-api/entities"

// UserUsecase represent the services of user
type UserUsecase interface {
	Create(*entities.User) (int64, error)
	Authenticate(email string, password string) error
}

// BookUsecase represents the bussiness logic for book
type BookUsecase interface {
	Create(*entities.Book) (int64, error)
	Get() ([]entities.Book, error)
	Update(*entities.Book) error
}
