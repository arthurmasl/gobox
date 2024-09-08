package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("tcp error", err)
		return
	}
	defer conn.Close()

	sendData(conn)
}

func sendData(conn net.Conn) {
	data := []byte("hello, server")

	_, err := conn.Write(data)
	if err != nil {
		fmt.Println("data error", err)
		return
	}
}
