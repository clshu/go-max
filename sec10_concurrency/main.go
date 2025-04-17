package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true // signal that the task is done
}

func main() {
	// go greet("Nice to meet you!")
	// go greet("How are you?")
	done := make(chan bool) // create a channel
	go slowGreet("How ... are ... you ...?", done)
	// go greet("I hope you're liking the course!")

	<-done
}
