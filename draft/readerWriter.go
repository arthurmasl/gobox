package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// writer
	buffer := new(bytes.Buffer)
	buffer.WriteString("hello")
	writeTo(buffer, []byte(", world"))

	// reader
	networkData := strings.NewReader("some string one two")
	readFrom(networkData)
}

func writeTo(writer io.Writer, msg []byte) error {
	_, err := writer.Write(msg)
	return err
}

func readFrom(reader io.Reader) {
	buffer := make([]byte, 8)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
	}
}
