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

func expandSpaces(columnIndex, itemIndex int) {
	if columnIndex > 0 && len(columns[columnIndex-1].items) < itemIndex+1 {
		grid[itemIndex+1] = strings.Join(
			[]string{strings.Repeat(" ", 20*columnIndex), grid[itemIndex+1]},
			strings.Repeat(" ", columnIndex),
		)
	}
}
