package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadBalanceFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
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

func WriteBalanceToFile(balance float64, filename string) {
	// Sprint formats using the default formats for its operands and returns the resulting string
	balanceText := fmt.Sprint(balance)
	os.WriteFile(filename, []byte(balanceText), 0644)
}
