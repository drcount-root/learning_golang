package scan_user_input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScanUserInput(input string) string {
	fmt.Println("\nEnter data :")
	// var userInput string
	// fmt.Scanln(&userInput)

	// how to read a big string with spaces
	reader := bufio.NewReader(os.Stdin)

	// at line break it will stop reading input, thats why we use '\n' & it must be in single quotes as its a rune (a special value type in go)

	userInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	} 

	// removing trailing newline character
	texts := strings.TrimSuffix(userInput, "\n")
	texts = strings.TrimSuffix(texts, "\r") // for windows

	return texts
}
