package main

import (
	"os"
	"os/exec"
)

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getPrefix(item *Item) string {
	prefix := separator
	if item.focused {
		prefix = ">"
	}

	return prefix
}
