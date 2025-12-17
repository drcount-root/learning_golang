package main

import (
	"fmt"

	"example.com/notes_taking_application/note"
	"example.com/notes_taking_application/scan_user_input"
)

func getNoteData() (string, string) {
	title := scan_user_input.ScanUserInput("Enter note title:")
	content := scan_user_input.ScanUserInput("\nEnter note content:")

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

	err = userNote.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	fmt.Println("Note saved successfully!")
}
