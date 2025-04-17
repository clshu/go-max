package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// ReadLines reads a file and returns its lines as a slice of strings.
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
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

// WriteJSONToFile writes the given data to a file in JSON format.
func WriteJSONToFile(filename string, data any) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Error creating file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("Error encoding JSON to file: %w", err)
	}
	return nil
}
