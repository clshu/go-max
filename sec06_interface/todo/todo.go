package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

// Todo represents a simple note-taking application.
type Todo struct {
	Text string `json:"text"`
}

// New creates a new Todo instance with the given title and content.
func New(text string) (*Todo, error) {
	if text == "" {
		return nil, fmt.Errorf("text cannot be empty")
	}

	return &Todo{
		Text: text,
	}, nil
}

// Display prints the note's details to the console.
func (t *Todo) Display() {

	fmt.Println(t.Text)
}

// Save the todo to the file
func (t *Todo) Save() error {

	fileName := "todo.json"

	json, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}
