package presenters_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldCreateBook(t *testing.T) {
	testCases := []struct {
		Description    string
		Body           []byte
		ExpectedStatus int
	}{
		{
			Description: "Should create a book",
			Body: []byte(`
				{
					"name": "Pequeno Príncipe",
					"price": 13.90,
					"isbn": "EOREQO-VMVMVSOS-QWEOWEIE"
				}
			`),
			ExpectedStatus: http.StatusCreated,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.Description, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/books",
				bytes.NewReader(testcase.Body))

			bookHandler.Create(rec, req)
			response := rec.Result()

			if response.StatusCode != testcase.ExpectedStatus {
				t.Errorf("Expected status code %d but got %d",
					testcase.ExpectedStatus, response.StatusCode)
			}
		})
	}
}

func TestShouldGetBook(t *testing.T) {
	testCases := []struct {
		Description    string
		Body           []byte
		ExpectedStatus int
	}{
		{
			Description:    "Should get books",
			Body:           nil,
			ExpectedStatus: http.StatusOK,
		},
	}

	for _, testcase := range testCases {
		t.Run(testcase.Description, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/books",
				bytes.NewReader(testcase.Body))

			bookHandler.Get(rec, req)
			response := rec.Result()

			if response.StatusCode != testcase.ExpectedStatus {
				t.Errorf("Expected status code %d but got %d",
					testcase.ExpectedStatus, response.StatusCode)
			}
		})
	}
}
