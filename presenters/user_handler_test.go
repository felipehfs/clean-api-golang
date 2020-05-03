package presenters_test

import (
	"io"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
}

func TestCreateUserHandler(t *testing.T) {
	tt := []struct {
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
		}
	}

	rec := httptest.NewRecorder()
	res := httptest.NewRequest()
}

func setup() {

}

func tearDown() {

}
