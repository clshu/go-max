package models

import (
	"time"

	"example.com/restapi/db"
)

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Save the user to the database
func (u *User) Save() (int64, error) {
	query := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}


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