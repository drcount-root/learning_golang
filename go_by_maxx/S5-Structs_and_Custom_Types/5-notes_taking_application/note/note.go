package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// these are struct tags which will help in json marshalling & unmarshalling
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n *Note) Display() {
	fmt.Printf("\nTitle: %s\nContent: %s\nCreated At: %s\n\n", n.Title, n.Content, n.CreatedAt.Format(time.RFC1123))
}

// We can add the copy of Note as we are not changing anything inside the struct, so no need of pointer receiver
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonData, err := json.Marshal(note)

	if err != nil {
		return err
	}

	// os.WriteFile(fileName, jsonData, 0644)
	// return nil

	return os.WriteFile(fileName, jsonData, 0644)
}
