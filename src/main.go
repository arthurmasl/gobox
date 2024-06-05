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
	update()
	draw()

	wg.Wait()
}

func initColumns() {
	stash := Column{name: "Stash"}
	stash.AddItem("one")
	stash.AddItem("two")

	stash.items[0].focused = true

	active := Column{name: "Active"}
	active.AddItem("three")

	done := Column{name: "Done"}
	done.AddItem("four")

	columns = append(columns, stash, active, done)
}

func update() {
	go handleInput()

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

		if char == 'w' {
			fmt.Println("hhh")
		}
	}
}

func draw() {
	clearConsole()

	for _, row := range grid {
		fmt.Print(row)
		fmt.Print("\r\n")
	}
}
