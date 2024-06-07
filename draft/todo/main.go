package main

import (
	"fmt"
	"strings"
	"sync"
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
	column.items = append(column.items, Item{value, false})
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
		title := fmt.Sprintf(format, strings.Join([]string{"#", column.name}, ""))
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

func draw() {
	clearConsole()

	fmt.Print("[q]quit", "\r\n")
	fmt.Print("[b]left   [w]right", "\r\n")
	fmt.Print("[j]down   [k]up", "\r\n")
	fmt.Print(emptyLine, "\r\n")
	for _, row := range grid {
		fmt.Print(row, "\r\n")
	}
}
