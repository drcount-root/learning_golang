package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	age       int
	createdAt time.Time
}

// Constuructor function for user struct which returns a pointer to a user not a copy of user
func newUser(firstName, lastName string, age int) (*user, error) {
	if firstName == "" || lastName == "" || age < 0 {
		return nil, errors.New("invalid input data")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
		createdAt: time.Now(),
	}, nil
}

func main() {
	fmt.Println("\nEnter first name :");
	var fName string 
	fmt.Scanln(&fName)

	fmt.Println("Enter last name :");
	var lName string 
	fmt.Scanln(&lName)

	fmt.Println("Enter age :")
	var age int
	fmt.Scanln(&age)

	appUser, err := newUser(fName, lName, age)

	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return
	}
	fmt.Printf("appUser: %v\n", appUser)
}
