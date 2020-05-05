package usecases_test

import (
	"testing"

	"github.com/felipehfs/clean-api/entities"
)

func TestShouldCreateBook(t *testing.T) {
	newBook := &entities.Book{
		ID:    1,
		Name:  "Example",
		ISBN:  "ABSFEFE-QEXFOGKGK-MWIEFJFE",
		Price: 12.0,
	}

	_, err := bookService.Create(newBook)
	if err != nil {
		t.Error(err)
	}

}

func TestShouldUpdateBook(t *testing.T) {
	newBook := &entities.Book{
		ID:    1,
		Name:  "Example",
		ISBN:  "ABSFEFE-QEXFOGKGK-MWIEFJFE",
		Price: 12.0,
	}

	err := bookService.Update(newBook)
	if err != nil {
		t.Error(err)
	}
}

func TestShouldGetUser(t *testing.T) {
	_, err := bookService.Get()
	if err != nil {
		t.Error(err)
	}
}
