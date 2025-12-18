package main

import (
	"fmt"

	"example.com/notes_taking_application_with_todos/note"
	"example.com/notes_taking_application_with_todos/scan_user_input"
	"example.com/notes_taking_application_with_todos/todo"
)

func getNoteData() (string, string) {
	title := scan_user_input.ScanUserInput("Enter note title:")
	content := scan_user_input.ScanUserInput("\nEnter note content:")

	return title, content
}

func main() {
	title, content := getNoteData()
	todoText := scan_user_input.ScanUserInput("\nEnter todo text:")
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	// todo.Display()

	// err = todo.Save()

	// if err != nil {
	// 	fmt.Println("Error saving todo:", err)
	// 	return
	// }

	// fmt.Println("Todo saved successfully!")

	// err = saveData(todo)

	err = outputData(todo)

	if err != nil {
		return
	}

	// userNote.Display()

	// err = userNote.Save()

	// if err != nil {
	// 	fmt.Println("Error saving note:", err)
	// 	return
	// }

	// fmt.Println("Note saved successfully!")

	// err = saveData(userNote)

	err = outputData(userNote)

	if err != nil {
		return
	}
}

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

// We can use saver interface to save both Todo and Note types as both structs are implementing the Save method.
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving data:")
		return err
	}

	fmt.Println("Data saved successfully!")

	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}
