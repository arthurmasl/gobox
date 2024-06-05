package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"syscall"

	"golang.org/x/term"
)

type Column struct {
	name  string
	items []Item
}

type Item struct {
	value   string
	focused bool
}

func (column *Column) AddItem(value string) {
	column.items = append(column.items, Item{value: value})
}

var (
	separator     = strings.Repeat(" ", 1)
	emptyLine     = strings.Repeat(" ", 80)
	format        = "%-20v"
	formatFocused = "%-19v"
	grid          = [20]string{}
)

var (
	columns = []Column{}
	wg      = sync.WaitGroup{}
)

func main() {
	initColumns()

	go handleInput()

	update()
	draw()

	wg.Wait()
}

func initColumns() {
	stash := Column{name: "Stash"}
	stash.AddItem("one")
	stash.AddItem("two")
	stash.AddItem("three")

	stash.items[0].focused = true

	active := Column{name: "Active"}
	active.AddItem("four")
	active.AddItem("five")

	done := Column{name: "Done"}
	done.AddItem("six")

	columns = append(columns, stash, active, done)
}

func update() {
	grid = [20]string{}

	for _, column := range columns {
		title := fmt.Sprintf(format, strings.Join([]string{" ", column.name}, ""))
		grid[0] = strings.Join([]string{grid[0], title}, separator)

		for i, item := range column.items {
			value := fmt.Sprintf(format, strings.Join([]string{" ", item.value}, ""))

			if item.focused {
				value = strings.Replace(value, " ", ">", 1)
			}

			grid[1+i] = strings.Join([]string{grid[1+i], value}, separator)
		}
	}
}

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
				columns[nextColumnIndex].items[0].focused = true
				columns[c].items[i].focused = false
			}
		}

		if char == 'b' {
			prevIndex := c - 1
			if prevIndex >= 0 {
				columns[prevIndex].items[0].focused = true
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

func draw() {
	clearConsole()

	fmt.Print("  [b]left [w]right", "\r\n")
	fmt.Print("  [j]down  [k]up", "\r\n")
	fmt.Print(emptyLine, "\r\n")
	for _, row := range grid {
		fmt.Print(row, "\r\n")
	}
}
