package main

import "fmt"

func main() {

	messages := make(chan string)
	go func() {
		messages <- "Hello from first GO"
		messages <- "Second Hello from first GO"
	}()
	go func() { messages <- "Hello from second GO" }()
	msg1 := <-messages
	msg2 := <-messages
	msg3 := <-messages
	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)
}
