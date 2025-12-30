package models

import (
	"fmt"
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

// var events = []Event{}

func (e *Event) Save() error {
	saveEventQuery := `
		INSERT INTO events (name, description, location, dateTime, userId) 
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(saveEventQuery)

	fmt.Printf("Err : %v", err)

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

func GetAllEvents() ([]Event, error) {
	getEventsQuery := `SELECT * FROM events`
	rows, err := db.DB.Query(getEventsQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close() // close the statement when the function returns, to prevent resource leaks

	events := []Event{}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (E *Event) UpdateEvent() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&E.Name, &E.Description, &E.Location, &E.DateTime, &E.Id)

	return err
}

func (E *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&E.Id)

	return err
}
