package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type Component struct {
	from       int
	to         int
	name       string
	data       string
	insertions *[]int
}

const (
	COMPONENT_START    = "<component"
	COMPONENT_END      = "</component>"
	COMPONENT_TEMPLATE = "<%s />"
)

var (
	fileInput  = "src/input/index.ahtml"
	fileOutput = "src/output/index.html"
)

func main() {
	file, err := os.ReadFile(fileInput)
	check(err)

	text := string(file[:len(file)-1])
	lines := strings.Split(text, "\n")

	var components []Component
	var name string
	var from int

	for i, line := range lines {
		// get components
		if strings.Contains(line, COMPONENT_START) {
			from = i
			name = getProperty(line, "name")
			continue
		}

		if strings.Contains(line, COMPONENT_END) {
			insertions := make([]int, 0)
			component := Component{
				name:       name,
				from:       from,
				to:         i,
				data:       strings.Join(lines[from+1:i], "\n"),
				insertions: &insertions,
			}
			components = append(components, component)
		}

		// get insertions
		for _, c := range components {
			if strings.Contains(line, fmt.Sprintf(COMPONENT_TEMPLATE, c.name)) {
				*c.insertions = append(*c.insertions, i)
			}
		}
	}

	// insert
	for _, c := range components {
		for _, i := range *c.insertions {
			bef := lines[:i]
			aft := lines[i+1:]

			lines = append(bef, c.data)
			lines = append(lines, aft...)
		}
	}

	// remove components
	lines = lines[components[len(components)-1].to+2:]

	data := []byte(strings.Join(lines, "\n"))
	writeErr := os.WriteFile(fileOutput, data, 0644)
	check(writeErr)
}

func getProperty(line string, name string) string {
	postfixLen := 2
	propertyLen := utf8.RuneCountInString(name) + postfixLen

	nameStart := strings.Index(line, fmt.Sprintf("%s=\"", name)) + propertyLen
	nameEnd := strings.Index(line[nameStart:], "\"")
	property := line[nameStart : nameStart+nameEnd]

	return property
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
