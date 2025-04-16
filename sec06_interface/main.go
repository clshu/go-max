package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-project/note"
	"example.com/go-project/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteDate()
	todoText := getUserInput("Todo text: ")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	outputData(userNote)

	fmt.Println("Note saved successfully!")

	todo, err := todo.New(todoText)
	outputData(todo)
	fmt.Println("Todo saved successfully!")

	printSomething("------------------")
	printSomething("Hello, World!")
	printSomething(123)
	printSomething(45.67)
	printSomething(true)
	printSomething([]string{"apple", "banana", "cherry"})
	printSomething(nil)
	printSomething(todo)
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteDate() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func outputData(o outputable) {
	o.Display()
	err := o.Save()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func saveToFile(s saver) error {
	err := s.Save()
	if err != nil {
		return err
	}
	return nil
}

func printSomething(value interface{}) {
	intValue, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intValue)
		return
	}
	stringValue, ok := value.(string)
	if ok {
		fmt.Println("String:", stringValue)
		return
	}
	floatValue, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatValue)
		return
	}
	boolValue, ok := value.(bool)
	if ok {
		fmt.Println("Boolean:", boolValue)
		return
	}
	sliceValue, ok := value.([]string)
	if ok {
		fmt.Println("Slice of strings:", sliceValue)
		return
	}
	if value == nil {
		fmt.Println("Nil value")
		return
	}
	fmt.Println("Unknown type:", value)

	// switch value.(type) {
	// case string:
	// 	fmt.Println("String:", value)
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case bool:
	// 	fmt.Println("Boolean:", value)
	// case []string:
	// 	fmt.Println("Slice of strings:", value)
	// default:
	// 	fmt.Println("Unknown type:", value)
	// }
}
