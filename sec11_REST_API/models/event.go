package models

import "time"

// Event represents an event in the system
type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var events []Event = []Event{}

// Save the event to the database
func (e *Event) Save() {
	events = append(events, *e)
}

// GetAllEvents returns all events
func GetAllEvents() []Event {
	return events
}
