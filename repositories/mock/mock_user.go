package mock

import (
	"time"

	"github.com/felipehfs/clean-api/entities"
)

// MockedUserRepository mocks the user Repository
type MockedUserRepository struct {
	Expectation map[string]interface{}
}

// Create inserts a new user
func (m MockedUserRepository) Create(*entities.User) (int64, error) {
	result, ok := m.Expectation["Create"]
	if ok {
		id, _ := result.(int64)

		return id, nil
	}
	return 1, nil
}

// SearchEmail returns a row with specified email
func (m MockedUserRepository) SearchEmail(email string) (*entities.User, error) {
	if email == "alreadyexists@example.com" {
		return &entities.User{
			ID:        2,
			FirstName: "Admin",
			LastName:  "",
			Email:     "alreadyexists@example.com",
			Password:  "1234",
			CreatedAt: time.Now(),
		}, nil
	}
	return nil, nil
}
