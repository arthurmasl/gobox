package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInput(reader *bufio.Reader) string {
	fmt.Print("Enter threads: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic("input error")
	}

	return input
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	input := getInput(reader)

	fmt.Println(input)
}
