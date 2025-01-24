package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:2222")
	if err != nil {
		fmt.Println("Error connecting", err)
		return
	}
	defer connection.Close()

	for {
		fmt.Println("sending message")
		connection.Write([]byte("hello from client"))
		time.Sleep(time.Second * 2)
	}
}
