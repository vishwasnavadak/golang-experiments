package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error Occured")
	} else {
		fmt.Println("Listening")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Unable to accept")
			break
		}
		message, _ := bufio.NewReader(conn).ReadString('\n')
		msg := string(message)
		fmt.Println("Received: ", msg)
	}
}
