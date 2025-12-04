package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileToOperateFrom = "balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank!")

	accountBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("⚠️ ERROR ⚠️")
		fmt.Println(err)
		fmt.Println("--------------------------------")
		// returns with a message
		panic("Can't continue sorry!")
	}

	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
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

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(fileToOperateFrom)
	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	// returning nil means we have no error
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	// Sprint formats using the default formats for its operands and returns the resulting string
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileToOperateFrom, []byte(balanceText), 0644)
}
