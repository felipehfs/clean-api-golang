package entities

import "time"

// User represents the client
type User struct {
	ID        int64     `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	FirstName string    `json:"firstname,omitempty"`
	LastName  string    `json:"lastname,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
