package repositories

import "github.com/felipehfs/clean-api/entities"

// SQLUserRepository is a contract of user repo
type SQLUserRepository interface {
	Create(*entities.User) (int64, error)
	SearchEmail(string) (*entities.User, error)
}
