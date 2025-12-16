package note

import (
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (*Note, error) {
	return &Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (n *Note) Display() {
	fmt.Printf("\nTitle: %s\nContent: %s\nCreated At: %s\n\n", n.title, n.content, n.createdAt.Format(time.RFC1123))
}
