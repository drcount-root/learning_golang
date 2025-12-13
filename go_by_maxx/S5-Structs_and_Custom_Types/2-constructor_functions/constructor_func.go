package main

import (
	"fmt"
	"time"

	"example.com/constructor_functions/user"
)

// type user struct {
// 	firstName string
// 	lastName  string
// 	age       int
// 	createdAt time.Time
// }

// Constuructor function for user struct which returns a pointer to a user not a copy of user
// func newUser(firstName, lastName string, age int) (*user, error) {
// 	if firstName == "" || lastName == "" || age < 0 {
// 		return nil, errors.New("invalid input data")
// 	}

// 	return &user{
// 		firstName: firstName,
// 		lastName:  lastName,
// 		age:       age,
// 		createdAt: time.Now(),
// 	}, nil
// }

func main() {
	fmt.Println("\nEnter first name :")
	var fName string
	fmt.Scanln(&fName)

	fmt.Println("Enter last name :")
	var lName string
	fmt.Scanln(&lName)

	fmt.Println("Enter age :")
	var age int
	fmt.Scanln(&age)

	appUser, err := user.NewUser(fName, lName, age)

	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return
	}
	fmt.Printf("appUser: %v\n", appUser)

	// var newAppUser *user.DifferentUserWithExportableFields

	// newAppUser = &user.DifferentUserWithExportableFields{
	// 	FirstName: "ABC",
	// 	LastName:  "XYZ",
	// 	Age:       25,
	// 	CreatedAt: time.Now(),
	// }

	newAppUser := user.DifferentUserWithExportableFields{
		FirstName: "ABC",
		LastName:  "XYZ",
		Age:       25,
		CreatedAt: time.Now(),
	}

	fmt.Printf("newAppUser: %v\n", newAppUser)
}
