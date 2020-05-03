package pg_test

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/felipehfs/clean-api/entities"
	"github.com/felipehfs/clean-api/repositories/pg"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestShouldSearchUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	repo := &pg.UserRepository{
		DB: db,
	}

	user := entities.User{
		ID:        1,
		FirstName: "José",
		LastName:  "Teste",
		Email:     "joseteste@example.com",
		Password:  "123456",
	}

	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "email", "password"}).
		AddRow(user.ID, user.FirstName, user.LastName, user.Email, user.Password)

	mock.ExpectQuery("^SELECT id, firstname, lastname, email, password FROM users").
		WillReturnRows(rows)

	_, err = repo.SearchEmail(user.Email)
	if err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestShouldCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	repo := &pg.UserRepository{
		DB: db,
	}

	defer db.Close()
	user := &entities.User{}
	user.Email = "felipe@example.com"
	user.Password = "1234"
	user.FirstName = "Felipe"
	user.LastName = "Henrique"

	mock.ExpectExec("^INSERT INTO users").
		WithArgs(user.FirstName, user.LastName, user.Email, user.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Create(user)

	if err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func setup() {

}

func tearDown() {

}
