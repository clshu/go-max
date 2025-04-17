package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true // signal that the task is done
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true // signal that the task is done
	close(doneChan)  // close the channel
}

func main() {
	// dones := make([]chan bool, 4) // create a channel
	// for i := 0; i < 4; i++ {
	// 	dones[i] = make(chan bool) // create a channel
	// }
	done := make(chan bool) // create a channel

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)

	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	for range done {
		// wait for the task to be done
	}

	// for i := 0; i < 4; i++ {
	// 	<-dones[i] // wait for the task to be done
	// }

	// <-done
	// <-done
	// <-done
	// <-done
	//
	// The following error happeded because I forgot to put <-true in greet()
	// 	fatal error: all goroutines are asleep - deadlock!

	// goroutine 1 [chan receive]:
	// main.main()
	// 	/Users/chinlong/else/golang/udemy/go-max/sec10_concurrency/main.go:31 +0x194

	// goroutine 36 [chan send]:
	// main.slowGreet({0x1023b5798, 0x18}, 0x14000110150)
	//
	//	/Users/chinlong/else/golang/udemy/go-max/sec10_concurrency/main.go:15 +0x9c
	// !!!! greet() didn't have <-true !!!! (my own comment, not from the output)
	// created by main.main in goroutine 1
	//
	//	/Users/chinlong/else/golang/udemy/go-max/sec10_concurrency/main.go:27 +0x12c
	//
	// exit status 2
}
