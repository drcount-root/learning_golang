package models

import "time"

// `binding:"required"` means that this field is required and cannot be empty
type Event struct {
	Id          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e *Event) Save() {
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events
}
