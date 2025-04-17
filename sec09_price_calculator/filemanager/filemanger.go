package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// FileManager struct implements the IOMangager interface for file input/output
type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// ReadLines reads a file and returns its lines as a slice of strings.
func (f *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(f.InputFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading file:", scanner.Err())
		return nil, scanner.Err()
	}
	return lines, nil
}

// WriteResult writes the given data to a file in JSON format.
func (f *FileManager) WriteResult(data any) error {
	file, err := os.Create(f.OutputFilePath)
	if err != nil {
		return fmt.Errorf("Error creating file: %w", err)
	}
	defer file.Close()

	time.Sleep(3 * time.Second) // Simulate a long task

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("Error encoding JSON to file: %w", err)
	}
	return nil
}

// New creates a new FileManager instance with the given input and output file paths.
func New(InputFilePath, OutputFilePath string) *FileManager {
	return &FileManager{
		InputFilePath:  InputFilePath,
		OutputFilePath: OutputFilePath,
	}
}
