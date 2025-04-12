package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// WriteFloatToFile writes a float64 value to a file
// in a string format with two decimal places.
// It creates the file if it doesn't exist and overwrites it if it does.
// The file is created with 0644 permissions.
func WriteFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprintf("%.2f", value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

// GetFloatFromFile reads a float64 value from a file.
// It returns the value as a float64 and an error if any occurs.
// If the file doesn't exist or the value can't be parsed,
// it returns a default value and an error message.
func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {

	data, err := os.ReadFile(fileName)
	if err != nil {
		return defaultValue, errors.New("could not read number from file. setting to default value: " + fmt.Sprint(defaultValue))
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {

		return defaultValue, errors.New("could not parse number from file. setting to default value: " + fmt.Sprint(defaultValue))
	}

	return value, nil
}
