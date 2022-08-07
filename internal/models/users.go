package models

import (
	"database/sql"
	"time"
)

// User is a new User type.
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// User Model wraps a database connection pool.
type UserModel struct {
	DB *sql.DB
}

// Insert method adds a new record to the "users" table.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate method verifys whether a user exitsts with the provided emal address and password. It returns
// the relevant user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Exists method checks if a user exists with a specific ID.
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
