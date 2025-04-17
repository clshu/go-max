package cmdmanager

import (
	"encoding/json"
	"fmt"
	"os"
)

// CMDManager struct implements the IOMangager interface for command line input/output
type CMDManager struct {
}

// New creates a new CMDManager instance
func New() *CMDManager {
	return &CMDManager{}
}

// ReadLines reads lines from the command line
func (c *CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please eneter the prices. Confirm every price with ENTER. Press ENTER to finish.")
	var lines []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scanln(&price)
		if price == "" {
			break
		}
		lines = append(lines, price)
	}
	return lines, nil
}

// WriteResult writes the result to the command line
func (c *CMDManager) WriteResult(data any) error {
	c.displayJSON(data)

	return nil
}

func (c *CMDManager) displayJSON(data any) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println()
}
