package main

import (
	"fmt"
	"strings"
)

type Column struct {
	name  string
	items []Item
}

type Item struct {
	value string
}

func (column *Column) AddItem(value string) {
	column.items = append(column.items, Item{value: value})
}

var (
	separator = strings.Repeat(" ", 1)
	format    = "%-20v"
	grid      = [20]string{}
)

var columns = []Column{}

func main() {
	initColumns()
	update()
	draw()
}

func initColumns() {
	stash := Column{name: "Stash"}
	stash.AddItem("one")
	stash.AddItem("two")

	active := Column{name: "Active"}
	active.AddItem("three")

	done := Column{name: "Done"}
	done.AddItem("four")

	columns = append(columns, stash, active, done)

	fmt.Println(columns)
}

func update() {
	for _, column := range columns {
		title := fmt.Sprintf(format, column.name)
		grid[0] = strings.Join([]string{grid[0], title}, separator)

		for i, item := range column.items {
			value := fmt.Sprintf(format, item.value)
			grid[1+i] = strings.Join([]string{grid[1+i], value}, separator)
		}
	}
}

func draw() {
	clearConsole()

	for _, row := range grid {
		fmt.Println(row)
	}
}
