package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/term"
)

func handleInput() {
	wg.Add(1)

	oldState, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		fmt.Println("error making terminal raw")
	}
	defer term.Restore(int(syscall.Stdin), oldState)

	buf := make([]byte, 1)

	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println("error readin input")
		}

		char := buf[0]

		if char == 'q' {
			wg.Done()
			break
		}

		c, i := getFocusedIndexes()

		if char == 'w' {
			nextColumnIndex := c + 1

			if nextColumnIndex < len(columns) {
				nextItemIndex := i
				if len(columns[nextColumnIndex].items) <= nextItemIndex {
					nextItemIndex = len(columns[nextColumnIndex].items) - 1
				}

				columns[nextColumnIndex].items[nextItemIndex].focused = true
				columns[c].items[i].focused = false
			}
		}

		if char == 'b' {
			prevColumnIndex := c - 1

			if prevColumnIndex >= 0 {
				prevItemIndex := i
				if len(columns[prevColumnIndex].items) <= prevItemIndex {
					prevItemIndex = len(columns[prevItemIndex].items) - 1
				}

				columns[prevColumnIndex].items[prevItemIndex].focused = true
				columns[c].items[i].focused = false
			}
		}

		if char == 'j' {
			nextItemIndex := i + 1
			if nextItemIndex < len(columns[c].items) {
				columns[c].items[nextItemIndex].focused = true
				columns[c].items[i].focused = false
			}
		}

		if char == 'k' {
			nextItemIndex := i - 1
			if nextItemIndex >= 0 {
				columns[c].items[nextItemIndex].focused = true
				columns[c].items[i].focused = false
			}
		}

		update()
		draw()

	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getFocusedIndexes() (int, int) {
	focusedColumn := 0
	focusedItem := 0

	for c, column := range columns {
		for i, item := range column.items {
			if item.focused {
				focusedColumn = c
				focusedItem = i

				break
			}
		}
	}

	return focusedColumn, focusedItem
}
