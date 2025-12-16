package main

import (
	"fmt"

	"example.com/struct_embedding/user"
)

func main() {

	user, err := user.New("Subham", 25)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nUser: %v\n", user)

	extraInfo, err := user.NewExtraInfo(10000.0, 2200.0)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\nExtra Info: %v\n", extraInfo)

	admin, err := user.NewAdmin(extraInfo, "8M7l7@example.com", "admin123")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nAdmin: %v\n", admin)

	admin.DeleteAdminExtraInfo()

	fmt.Printf("\nAdmin after deleting extra info: %v\n", admin)

	admin.User.UpdateUserDetails("Subham Sahu", 27)

	// as user is embedded in admin, we can directly access the user methods
	admin.UpdateUserDetails("Subham Sahu", 27)

	fmt.Printf("\nAdmin after updating user details: %v\n", admin)
}
