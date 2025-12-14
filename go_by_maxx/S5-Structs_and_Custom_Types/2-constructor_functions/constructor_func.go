package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/constructor_functions/product"
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

	newAppUser := user.UserWithExportableFields{
		FirstName: "ABC",
		LastName:  "XYZ",
		Age:       25,
		CreatedAt: time.Now(),
	}

	fmt.Printf("newAppUser: %v\n", newAppUser)

	// Product struct and constructor function usage

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nEnter product name:")
	productName, _ := reader.ReadString('\n')
	productName = strings.TrimSpace(productName)

	fmt.Println("Enter product description:")
	productDescription, _ := reader.ReadString('\n')
	productDescription = strings.TrimSpace(productDescription)

	fmt.Println("Enter original price :")
	var originalPrice float64
	fmt.Scanln(&originalPrice)

	fmt.Println("Enter discounted price :")
	var discountedPrice float64
	fmt.Scanln(&discountedPrice)

	product1, err := product.NewProduct(productName, productDescription, originalPrice, discountedPrice)

	if err != nil {
		fmt.Printf("Error creating product: %v\n", err)
		return
	}

	fmt.Printf("product1: %v\n", product1)
}
