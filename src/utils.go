package main

import (
	"os"
	"os/exec"
	"strings"
)

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// move item right by spaces if there is no item in left column
func expandSpaces(columnIndex, itemIndex int) {
	if columnIndex > 0 && len(columns[columnIndex-1].items) < itemIndex+1 {
		grid[itemIndex+1] = strings.Join(
			[]string{strings.Repeat(" ", 20*columnIndex), grid[itemIndex+1]},
			strings.Repeat(" ", columnIndex),
		)
	}
}

func getPrefix(item *Item) string {
	prefix := separator
	if item.focused {
		prefix = ">"
	}

	return prefix
}
