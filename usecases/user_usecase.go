package usecases

import (
	"errors"

	"github.com/felipehfs/clean-api/entities"
	"github.com/felipehfs/clean-api/repositories"
)

var (
	// ErrorPasswordRequired the field password must be filled
	ErrorPasswordRequired = errors.New("The password is required")
	// ErrorEmailRequired the field password must be filled
	ErrorEmailRequired = errors.New("The email is required")
	// ErrorEmailAlreadyExists the field must be unique
	ErrorEmailAlreadyExists = errors.New("The email must be unique")
)

// UserService implements the user usecase
type UserService struct {
	Repository repositories.SQLUserRepository
}

// Create inserts the new user
func (service UserService) Create(user *entities.User) (int64, error) {
	if len(user.Password) == 0 {
		return -1, ErrorPasswordRequired
	}

	if len(user.Email) == 0 {
		return -1, ErrorEmailRequired
	}

	userFounded, err := service.Repository.SearchEmail(user.Email)
	if err != nil {
		return -1, err
	}

	if userFounded != nil {
		return -1, ErrorEmailAlreadyExists
	}

	return service.Repository.Create(user)
}

// Authenticate executes the login with api
func (service UserService) Authenticate(email string, password string) error {
	userFounded, err := service.Repository.SearchEmail(email)
	if err != nil {
		return err
	}

	if userFounded.Password != password {
		return errors.New("The password is wrong")
	}

	return nil
}
