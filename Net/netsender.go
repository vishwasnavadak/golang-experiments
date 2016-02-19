package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.4:8000")
	if err != nil {
		fmt.Println(conn)
	} else {
		fmt.Println("Successfully Connected!")
	}

	fmt.Fprintf(conn, "Hello From the other side\n")
	fmt.Println("Message Sent")
}
