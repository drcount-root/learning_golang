package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	dummyPhoneNumber := randomdata.PhoneNumber()
	fmt.Println("Dummy phone number: ", dummyPhoneNumber)
}
