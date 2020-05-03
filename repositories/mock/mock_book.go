package mock

import (
	"github.com/felipehfs/clean-api/entities"
)

// MockedBookRepository represents the repository for test propurse
type MockedBookRepository struct{}

// Create inserts the new book into database
func (m MockedBookRepository) Create(book *entities.Book) (int64, error) {
	return 1, nil
}

// Get returns every book
func (m MockedBookRepository) Get() ([]entities.Book, error) {
	return []entities.Book{}, nil
}
