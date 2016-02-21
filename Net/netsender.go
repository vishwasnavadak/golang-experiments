package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.3:8000")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Connected!")
	}
	defer conn.Close()
	for {
		fmt.Println("Enter a message to send: ")
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Fprintf(conn, text)
	}

}
