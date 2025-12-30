package models

import (
	"time"

	"example.com/rest_api_with_db/db"
)

// `binding:"required"` means that this field is required and cannot be empty
type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e *Event) Save() error {
	saveEventQuery := `
		INSERT INTO events (name, description, location, dateTime, user_id) 
		VALUES (?, ?, ?, ?, user_id)
	`

	stmt, err := db.DB.Prepare(saveEventQuery)

	if err != nil {
		return err
	}

	defer stmt.Close() // close the statement when the function returns, to prevent resource leaks
	// this closes regardless of whether the function returns successfully or not

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.Id = id

	// events = append(events, *e)

	return err
}

func GetAllEvents() []Event {
	return events
}
