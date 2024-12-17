package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
	reader := bufio.NewReader(os.Stdin)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string error", err)
			return
		}

		conn.Write([]byte(msg))
	}
}
