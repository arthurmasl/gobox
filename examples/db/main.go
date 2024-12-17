package main

import (
	"fmt"
	"os"
	"strings"
)

type Data map[string]string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.ReadFile("src/data/users.amdb")
	check(err)

	text := string(file[:len(file)-1])
	lines := strings.Split(text, "\n")
	rows := strings.Split(lines[0], " ")
	data := lines[1:]

	dataSlice := make([]Data, len(data))

	for colIndex, col := range data {
		colData := strings.Split(col, " ")
		colMap := make(Data)

		for rowIndex, row := range rows {
			colMap[row] = colData[rowIndex]
		}

		dataSlice[colIndex] = colMap
	}

	fmt.Printf("%+v\n", dataSlice)
}
