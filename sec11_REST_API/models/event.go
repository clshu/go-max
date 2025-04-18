package models

import (
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
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Save the event to the database
func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return err
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
