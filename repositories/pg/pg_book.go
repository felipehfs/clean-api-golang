package pg

import (
	"database/sql"

	"github.com/felipehfs/clean-api/entities"
)

// BookRepository represent the data provider
type BookRepository struct {
	DB *sql.DB
}

// Create inserts a new book
func (repo BookRepository) Create(book *entities.Book) (int64, error) {
	sql := "INSERT INTO books (name, isbn, price) VALUES ($1, $2, $3) RETURNING id"

	var id int64
	err := repo.DB.QueryRow(sql, book.Name, book.ISBN, book.Price).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

// Update changes the book
func (repo BookRepository) Update(book *entities.Book) error {
	sql := `UPDATE books SET name=$2, isbn=$3, price=$4 WHERE id=$1`
	_, err := repo.DB.Exec(sql, book.ID, book.Name, book.ISBN, book.Price)
	if err != nil {
		return err
	}

	return nil
}

// Get retrieves all books
func (repo BookRepository) Get() ([]entities.Book, error) {
	sql := "SELECT id, name, isbn, price FROM books"

	res, err := repo.DB.Query(sql)
	if err != nil {
		return nil, err
	}

	var books []entities.Book

	for res.Next() {
		var book entities.Book
		err := res.Scan(&book.ID, &book.Name, &book.ISBN, &book.Price)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
