package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const fileToOperateFrom = "balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank!")

	accountBalance, err := fileops.ReadBalanceFromFile(fileToOperateFrom)

	if err != nil {
		fmt.Println("⚠️ ERROR ⚠️")
		fmt.Println(err)
		fmt.Println("--------------------------------")
		// returns with a message
		panic("Can't continue sorry!")
	}

	for {
		communicationFunc()

		var choice int
		fmt.Print("Enter the choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("$$$$ Your balance is $%v $$$$", accountBalance)
		case 2:
			fmt.Println("Enter the amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}
			accountBalance += amount
			fileops.WriteBalanceToFile(accountBalance, fileToOperateFrom)
			fmt.Printf("You have deposited $%.2f & current balance is $%.2f\n", amount, accountBalance)
		case 3:
			fmt.Println("Enter the amount to withdraw: ")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount")
				continue
			}

			if amount > accountBalance {
				fmt.Println("Insufficient balance")
				continue
			}

			accountBalance -= amount
			fileops.WriteBalanceToFile(accountBalance, fileToOperateFrom)
			fmt.Printf("You have withdrawn $%.2f & current balance is $%.2f\n", amount, accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
