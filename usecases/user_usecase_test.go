package usecases_test

import (
	"os"
	"testing"
	"time"

	"github.com/felipehfs/clean-api/entities"
	"github.com/felipehfs/clean-api/repositories/mock"
	"github.com/felipehfs/clean-api/usecases"
)

var mockedRepository *mock.MockedUserRepository
var userService *usecases.UserService

var mockedBookRepository *mock.MockedBookRepository
var bookService *usecases.BookService

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	// Inicializando o que precisa para os testes
	mockedRepository = new(mock.MockedUserRepository)
	userService = &usecases.UserService{
		Repository: mockedRepository,
	}

	// Inicializando o que precisa para os testes do livro
	mockedBookRepository = new(mock.MockedBookRepository)
	bookService = &usecases.BookService{
		Repository: mockedBookRepository,
	}
}

func tearDown() {
	// Desacoplando o que precisa para finalizar
}

func TestCreateUser(t *testing.T) {

	testTable := []struct {
		Description   string
		User          entities.User
		ExpectedError error
		ExpectedID    int64
	}{
		{
			Description: "Should require a password",
			User: entities.User{
				ID:        1,
				FirstName: "admin",
				LastName:  "",
				Email:     "admin@example.com",
				Password:  "",
				CreatedAt: time.Now(),
			},
			ExpectedError: usecases.ErrorPasswordRequired,
			ExpectedID:    -1,
		},
		{
			Description: "Should require a email",
			User: entities.User{
				ID:        1,
				FirstName: "admin",
				LastName:  "",
				Email:     "",
				Password:  "1111",
				CreatedAt: time.Now(),
			},
			ExpectedError: usecases.ErrorEmailRequired,
			ExpectedID:    -1,
		},
		{
			Description: "Should store the data",
			User: entities.User{
				ID:        1,
				FirstName: "admin",
				LastName:  "",
				Email:     "admin@example.com",
				Password:  "1235",
				CreatedAt: time.Now(),
			},
			ExpectedError: nil,
			ExpectedID:    1,
		},
		{
			Description: "Should not pass duplicated email",
			User: entities.User{
				ID:        1,
				FirstName: "Admin",
				LastName:  "",
				Email:     "alreadyexists@example.com",
				Password:  "1111",
				CreatedAt: time.Now(),
			},
			ExpectedError: usecases.ErrorEmailAlreadyExists,
			ExpectedID:    -1,
		},
	}

	for _, test := range testTable {
		t.Run(test.Description, func(t *testing.T) {
			id, err := userService.Create(&test.User)
			if err != test.ExpectedError {
				t.Errorf("Expected error: %v but got %v", test.ExpectedError, err)
			}

			if id != test.ExpectedID {
				t.Errorf("Expected id: %v but got %v", test.ExpectedID, id)
			}
		})
	}
}

func TestShouldAuthenticateUser(t *testing.T) {
	user := entities.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "",
		Email:     "alreadyexists@example.com",
		Password:  "1234",
		CreatedAt: time.Now(),
	}

	err := userService.Authenticate(user.Email, user.Password)
	if err != nil {
		t.Error(err)
	}
}
