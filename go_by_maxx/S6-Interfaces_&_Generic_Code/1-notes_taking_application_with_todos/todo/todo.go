package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// these are struct tags which will help in json marshalling & unmarshalling
type Todo struct {
	Text string `json:"text"`
}

func New(content string) (*Todo, error) {

	if content == "" {
		return &Todo{}, errors.New("invalid input")
	}

	return &Todo{
		Text: content,
	}, nil
}

func (t *Todo) Display() {
	fmt.Printf("\nText: %s\n\n", t.Text)
}

// We can add the copy of Todo as we are not changing anything inside the struct, so no need of pointer receiver
func (t Todo) Save() error {
	fileName := "todo.json"
	jsonData, err := json.Marshal(t)

	if err != nil {
		return err
	}

	// os.WriteFile(fileName, jsonData, 0644)
	// return nil

	return os.WriteFile(fileName, jsonData, 0644)
}
