package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("tcp error", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener error", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Println("client connected", conn)

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("client handler error", err)
			return
		}

		fmt.Printf("Received: %s", buffer[:n])
	}
}
