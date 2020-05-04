package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/felipehfs/clean-api/entities"
	"github.com/felipehfs/clean-api/repositories"
	"github.com/felipehfs/clean-api/usecases"
)

// BookHandler represents the controllers for book
type BookHandler struct {
	Service usecases.BookUsecase
}

// NewBookHandler instantiates a new BookHandler
func NewBookHandler(repository repositories.SQLBookRepository) *BookHandler {
	service := usecases.BookService{
		Repository: repository,
	}

	return &BookHandler{
		Service: service,
	}
}

// Create inserts the new book
func (b BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var book entities.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := b.Service.Create(&book)
	book.ID = id

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// Get retrieves every book
func (b BookHandler) Get(w http.ResponseWriter, r *http.Request) {
	books, err := b.Service.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}
