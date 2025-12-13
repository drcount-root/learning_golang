package main

import (
	"fmt"
	"time"

	writetofile "example.com/structs/write_to_file"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// To attach a method to a struct, we define a function with a receiver argument
func (u *User) outputData(data string) {
	fmt.Println("This is from outputData method of User struct", data)
	fmt.Printf("User Data: %s %s, born on %s, created at %v\n",
		u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) clearUserFirstName() {
	// if we try to do this without a pointer receiver, it will not modify the original struct as it woul then deal with the copy of the struct
	u.firstName = ""
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

	saveUserData(&newUser)

	fmt.Println(firstName, lastName, birthdate)

	dataToWrite := fmt.Sprintf("First Name: %s\tLast Name: %s\tBirthdate: %s\tCreated At: %v",
		newUser.firstName, newUser.lastName, newUser.birthdate, newUser.createdAt)

	writetofile.WriteUserDataToFile("users.txt", dataToWrite)

	newUser.outputData("hi")

	newUser.clearUserFirstName()
	newUser.outputData("after clearing first name")

	newUser.updateUserLing("motiwala")

	fmt.Printf("Lastly updated user %v", newUser);
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

func saveUserData(user *User) {
	fmt.Printf("user pointer: %v\n", user)
	// Saving user data from the user pointer address
	fmt.Printf("Saving user data for: %v\n", *user)

	// So noramally we need to dereference the pointer to access the fields
	fmt.Printf("User First Name: %s\n", (*user).firstName)
	// But Go allows us to access struct fields directly using the pointer
	fmt.Printf("User Last Name: %s\n", user.lastName)
}

func (u *User) updateUserLing(name string) {
	u.lastName = name
}
