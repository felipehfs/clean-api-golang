package presenters

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/felipehfs/clean-api/entities"
	"github.com/felipehfs/clean-api/repositories"
	"github.com/felipehfs/clean-api/usecases"
	"github.com/gorilla/mux"
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

// Update is a controllers to change a book
func (b BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	var newBook entities.Book
	vars := mux.Vars(r)

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newBook.ID = id

	err = b.Service.Update(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}
