package main

import (
	"fmt"

	"example.com/notes_taking_application/note"
	"example.com/notes_taking_application/scan_user_input"
)

func getNoteData() (string, string) {
	title := scan_user_input.ScanUserInput("Enter note title:")
	content := scan_user_input.ScanUserInput("Enter note content:")

	return title, content
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
}
