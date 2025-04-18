package models

import (
	"log"
	"time"

	"example.com/restapi/db"
)

// Event represents an event in the system
type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64     `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Save the event to the database
func (e *Event) Save() (int64, error) {
	query := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	e.ID = id
	return id, err
}

// Update the event in the database
func (e *Event) Update() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, date_time = ?, user_id = ?, updated_at = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.UpdatedAt, e.ID)
	return err
}

// Delete the event from the database
func (e *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}

// GetAllEvents returns all events
func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID, &event.CreatedAt, &event.UpdatedAt); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

// GetEventByID returns an event by ID
func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var event Event
	if err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID, &event.CreatedAt, &event.UpdatedAt); err != nil {
		return nil, err
	}
	return &event, nil
}

// Register registers a user for an event
func (e *Event) Register(userID int64) error {
	query := `INSERT INTO registrations (event_id, user_id) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userID)
	return err
}

// CancelRegistration cancels a user's registration for an event
func (e *Event) CancelRegistration(userID int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	log.Println("Event ID:", e.ID)
	log.Println("User ID:", userID)
	_, err = stmt.Exec(e.ID, userID)
	log.Println("Error:", err)
	return err
}

// CheckRegistration checks if a user is registered for an event
func (e *Event) CheckRegistration(userID int64) (int64, error) {
	query := `SELECT id FROM registrations WHERE event_id = ? AND user_id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	var id int64

	err = stmt.QueryRow(e.ID, userID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
