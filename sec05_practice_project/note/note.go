package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Note represents a simple note with a title, content, and creation date.
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// New creates a new Note instance with the given title and content.
func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, fmt.Errorf("title or content cannot be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

// Display prints the note's details to the console.
func (n *Note) Display() {
	fmt.Println("=======================================")
	fmt.Printf("Title: %s\n", n.Title)
	fmt.Printf("Content: %s\n", n.Content)
	fmt.Printf("Created At: %s\n", n.CreatedAt.Local())
	fmt.Println("=======================================")
}

// Save the note to the file
func (n *Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ReplaceAll(fileName, "/", "_")
	fileName = strings.ToLower(fileName)
	fileName += ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}
