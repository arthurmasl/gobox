package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.ReadFile("src/input/index.ahtml")
	check(err)

	text := string(file[:len(file)-1])
	lines := strings.Split(text, "\n")

	var start int
	var end int
	var name string

	var insertions []int

	result := lines

	for i, line := range lines {
		if strings.Contains(line, "<component name=") {
			start = i
			nameStart := strings.Index(line, "name=\"") + 6
			nameEnd := strings.Index(line[nameStart:], "\"")
			name = line[nameStart : nameStart+nameEnd]
			continue
		}

		if strings.Contains(line, "</component>") {
			end = i
			continue
		}

		if strings.Contains(line, fmt.Sprintf("<%s />", name)) {
			insertions = append(insertions, i)

			continue
		}
	}

	component := strings.Join(lines[start+1:end], "\n")

	for _, i := range insertions {
		bef := result[:i]
		aft := result[i+1:]

		result = append(bef, component)
		result = append(result, aft...)
	}

	result = result[end+1:]

	data := []byte(strings.Join(result, "\n"))
	writeErr := os.WriteFile("./src/output/index.html", data, 0644)
	check(writeErr)
}
