package main

import "fmt"
import "time"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(100)
	}
}

func main() {

	f("direct")

	go f("goroutine")
	go fmt.Println("Second GO Routine")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
