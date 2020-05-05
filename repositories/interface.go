package repositories

import "github.com/felipehfs/clean-api/entities"

// SQLUserRepository is a contract of user repo
type SQLUserRepository interface {
	Create(*entities.User) (int64, error)
	SearchEmail(string) (*entities.User, error)
}

// SQLBookRepository is a contract of the data provider
type SQLBookRepository interface {
	Create(*entities.Book) (int64, error)
	Get() ([]entities.Book, error)
	Update(*entities.Book) error
	Remove(int64) error
}
