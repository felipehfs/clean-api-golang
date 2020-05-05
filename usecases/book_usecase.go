package usecases

import (
	"errors"

	"github.com/felipehfs/clean-api/entities"
	"github.com/felipehfs/clean-api/repositories"
)

var (
	// ErrorBookNameRequired must be used for validation
	ErrorBookNameRequired = errors.New("The name of book is required")
	// ErrorBookPriceRequired must be used for validation
	ErrorBookPriceRequired = errors.New("The price of book must be greater than zero")
)

// BookService represents the book bussiness logic
type BookService struct {
	Repository repositories.SQLBookRepository
}

// Create inserts the new book into database
func (bs BookService) Create(book *entities.Book) (int64, error) {

	if book.Name == "" {
		return -1, ErrorBookNameRequired
	}

	if book.Price <= 0 {
		return -1, ErrorBookPriceRequired
	}

	return bs.Repository.Create(book)
}

// Get retrieves the every books
func (bs BookService) Get() ([]entities.Book, error) {
	return bs.Repository.Get()
}

// Update changes the book
func (bs BookService) Update(book *entities.Book) error {

	if book.Name == "" {
		return ErrorBookNameRequired
	}

	if book.Price <= 0 {
		return ErrorBookPriceRequired
	}

	return bs.Repository.Update(book)
}

// Remove a book
func (bs BookService) Remove(id int64) error {
	return bs.Repository.Remove(id)
}
