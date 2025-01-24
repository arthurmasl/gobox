package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:2222")
	if err != nil {
		fmt.Println("Error to create server", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started at port 2222")

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Client error", err)
			continue
		}

		fmt.Println("clinet connected")
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Client disconnected")
			return
		}

		fmt.Printf("Received: %s\n", buffer[:n])
	}
}
