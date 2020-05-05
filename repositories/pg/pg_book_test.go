package pg_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/felipehfs/clean-api/entities"
	"github.com/felipehfs/clean-api/repositories/pg"
)

func TestCreateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	book := &entities.Book{
		ID:    1,
		Name:  "Pequeno Príncipe",
		ISBN:  "RSZERERE-2FGSQWEW-IEOREX",
		Price: 34.03,
	}

	repo := pg.BookRepository{
		DB: db,
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("^INSERT INTO book").
		WithArgs(book.Name, book.ISBN, book.Price).
		WillReturnRows(rows)

	_, err = repo.Create(book)
	if err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestUpdateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := pg.BookRepository{
		DB: db,
	}

	changedBook := &entities.Book{
		ID:    1,
		Name:  "Pequeno Príncipe (updated)",
		ISBN:  "RIESXS-EIEIEJFM-AXXDDF",
		Price: 45.60,
	}

	mock.ExpectExec("^UPDATE books").
		WithArgs(changedBook.ID,
			changedBook.Name, changedBook.ISBN,
			changedBook.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Update(changedBook)
	if err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestRemoveBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	repo := pg.BookRepository{
		DB: db,
	}

	mock.ExpectExec("^DELETE FROM books").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err := repo.Remove(1); err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := pg.BookRepository{
		DB: db,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "isbn", "price"})

	mock.ExpectQuery("^SELECT id, name, isbn, price FROM books").
		WillReturnRows(rows)

	_, err = repo.Get()
	if err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
