package main

import "fmt"

func clearConsole() {
	// cmd := exec.Command("clear")
	// cmd.Stdout = os.Stdout
	// cmd.Run()
	fmt.Print("\033[H\033[2J")
}

func getPrefix(item *Item) string {
	prefix := separator
	if item.focused {
		prefix = ">"
	}

	return prefix
}
