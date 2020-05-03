package presenters_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/felipehfs/clean-api/presenters"
	"github.com/felipehfs/clean-api/repositories/mock"
	"github.com/felipehfs/clean-api/usecases"
)

var mockedRepository *mock.MockedUserRepository
var userService *usecases.UserService
var userHandler *presenters.UserHandler

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestCreateUserHandler(t *testing.T) {
	tableTest := []struct {
		Description    string
		Body           []byte
		ExpectedStatus int
	}{
		{
			Description: "Should return 201 created",
			Body: []byte(`
				{
					"firstname":"admin",
					"lastname": "",
					"email": "admin@example.com",
					"password": "123456"
				}
			`),
			ExpectedStatus: http.StatusCreated,
		},
	}

	for _, test := range tableTest {
		rec := httptest.NewRecorder()
		res := httptest.NewRequest("POST", "/register", bytes.NewReader(test.Body))
		userHandler.Register(rec, res)

		response := rec.Result()
		if response.StatusCode != test.ExpectedStatus {
			t.Errorf("Expected status code %v but got the code %v", response.StatusCode, test.ExpectedStatus)
		}
	}
}

func TestLoginUserHandler(t *testing.T) {
	tableTest := []struct {
		Description    string
		Body           []byte
		ExpectedStatus int
	}{
		{
			Description: "Should login successfully",
			Body: []byte(`
				{
					"email": "alreadyexists@example.com",
					"password": "1234"
				}
			`),
			ExpectedStatus: http.StatusOK,
		},
	}

	for _, test := range tableTest {
		rec := httptest.NewRecorder()
		res := httptest.NewRequest("POST", "/login", bytes.NewReader(test.Body))
		userHandler.Login(rec, res)

		response := rec.Result()
		if response.StatusCode != test.ExpectedStatus {
			t.Errorf("Expected status code %v but got the code %v", response.StatusCode, test.ExpectedStatus)
		}
	}
}

func setup() {
	mockedRepository = new(mock.MockedUserRepository)

	userService = &usecases.UserService{
		Repository: mockedRepository,
	}

	userHandler = &presenters.UserHandler{
		Service: userService,
	}
}

func tearDown() {

}
