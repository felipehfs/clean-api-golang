package pg

import (
	"database/sql"

	"github.com/felipehfs/clean-api/entities"
)

// UserRepository implements SQLUserRepository
type UserRepository struct {
	DB *sql.DB
}

// Create inserts the new data into database
func (repo UserRepository) Create(user *entities.User) (int64, error) {
	sql := `INSERT INTO users (firstname, lastname, email, password) VALUES ($1, $2, $3, $4)`
	res, err := repo.DB.Exec(sql, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// SearchEmail returns the user by email
func (repo UserRepository) SearchEmail(email string) (*entities.User, error) {
	sql := "SELECT id, firstname, lastname, email, " +
		"password FROM users WHERE email=$1"

	var result entities.User
	err := repo.DB.QueryRow(sql, email).
		Scan(&result.ID, &result.FirstName,
			&result.LastName, &result.Email,
			&result.Password)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
