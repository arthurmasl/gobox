package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic("Error starting server" + err.Error())
	}
	defer listener.Close()

	fmt.Println("server started on port 8000")
}
