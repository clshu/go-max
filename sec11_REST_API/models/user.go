package models

import (
	"fmt"
	"strings"
	"time"

	"example.com/restapi/db"
	"example.com/restapi/util"
)

const MIN_PASSOWRD_LENGTH = 8

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Save the user to the database
func (u *User) Save() (int64, error) {
	if u.Email == "" {
		return -1, fmt.Errorf("email is required")
	}
	if u.Password == "" {
		return -1, fmt.Errorf("password is required")
	}
	if len(u.Password) < MIN_PASSOWRD_LENGTH {
		return -1, fmt.Errorf("password must be at least 8 characters long")
	}

	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return -1, err
	}

	// email is saved in lower case
	result, err := stmt.Exec(strings.ToLower(u.Email), hashedPassword)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	u.ID = id
	return id, err
}

// ValidateCredentials checks if the provided email and password are valid
func (u *User) ValidateCredentials() error {

	query := `SELECT password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, strings.ToLower(u.Email))
	var hashedPassword string
	err := row.Scan(&hashedPassword)
	if err != nil {
		return fmt.Errorf("invalid email or password")
	}
	same, err := util.ComparePasswordAndHash(u.Password, hashedPassword)
	if err != nil {
		return fmt.Errorf("invalid email or password")
	}
	if !same {
		return fmt.Errorf("invalid email or password")
	}

	return nil
}

// GetUserByID retrieves a user by ID from the database
func GetUserByID(id int64) (*User, error) {
	query := `SELECT * FROM users WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var user User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
