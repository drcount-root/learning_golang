package user

import (
	"errors"
	"time"
)

type User struct {
	firstName string
	lastName  string
	age       int
	createdAt time.Time
}

// Constuructor function for user struct which returns a pointer to a user not a copy of user
func NewUser(firstName, lastName string, age int) (*User, error) {
	if firstName == "" || lastName == "" || age < 0 {
		return nil, errors.New("invalid input data")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
		createdAt: time.Now(),
	}, nil
}

// We can only export functions, variable, structs, etc. but we cannot export struct fields directly.
// So, if we want to export struct fields, we need to create the struct with exportable fields and map the values accordingly.
// This is just an example to show how to export struct fields.
type DifferentUserWithExportableFields struct {
	FirstName string
	LastName  string
	Age       int
	CreatedAt time.Time
}