package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Listening")
	}
	defer ln.Close()
	conn, err := ln.Accept()
	if err != nil {
		panic(err)

	}
	for {

		message, _ := bufio.NewReader(conn).ReadString('\n')
		t := time.Now()
		msg := t.Format("Mon Jan _2 2006 15:04:05") + ":" + message
		fmt.Println(msg)

	}
}
