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

	// returns pointer of user struct
	// If we return a copy of user struct, then the changes made to the user struct will not be reflected in the original user struct.
	// So, we return a pointer to the user struct.
	// This is useful when we want to modify the user struct later.
	// Also, it is more efficient to return a pointer to the user struct instead of a copy of the user struct.
	// This is because, when we return a pointer to the user struct, we are not copying the entire struct, but only the address of the struct.
	// This saves memory and time.
	// So, we use pointer to user struct instead of user struct.

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
type UserWithExportableFields struct {
	FirstName string
	LastName  string
	Age       int
	CreatedAt time.Time
}