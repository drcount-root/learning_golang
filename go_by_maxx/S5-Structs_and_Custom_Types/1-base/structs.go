package main

import (
	"fmt"
	"time"

	"example.com/structs/write_to_file"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var newUser User

	// newUser = User{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthdate: birthdate,
	// }

	// newUser := User{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthdate: birthdate,
	// 	createdAt: time.Now(),
	// }

	createdAt := time.Now()

	newUser := User{
		firstName,
		lastName,
		birthdate,
		createdAt,
	}

	newUser2 := User{}
	newUser2.firstName = "Amit"
	newUser2.lastName = "Kumar"
	newUser2.birthdate = "01/01/2000"
	newUser2.createdAt = time.Now()

	fmt.Println(newUser)
	fmt.Println(newUser2)

	fmt.Println(firstName, lastName, birthdate)

	dataToWrite := fmt.Sprintf("First Name: %s\tLast Name: %s\tBirthdate: %s\tCreated At: %v",
		newUser.firstName, newUser.lastName, newUser.birthdate, newUser.createdAt)

	writetofile.WriteUserDataToFile("users.txt", dataToWrite)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
