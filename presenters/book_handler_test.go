package presenters_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
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

func TestShouldUpdateBook(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/books", bytes.NewReader([]byte(`
		{
			"name": "Pequeno Príncipe",
			"price": 13.90,
			"isbn": "EOREQO-VMVMVSOS-QWEOWEIE"
		}
	`)))

	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	bookHandler.Update(rec, req)
	response := rec.Result()
	if response.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %v but found %v", http.StatusNoContent, response.StatusCode)
	}
}

func TestShouldRemoveBook(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/books", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	bookHandler.Remove(rec, req)
	response := rec.Result()
	if response.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status code %v but found %v", http.StatusNoContent, response.StatusCode)
	}
}
