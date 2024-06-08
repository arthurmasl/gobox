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

const (
	separator = " "
	empty     = ""
	format    = "%-20v"
	br        = "\r\n"
	maxItems  = 5
)

var (
	wg      = sync.WaitGroup{}
	columns = []Column{}
	grid    = [20]string{}
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

	stash.items[0].focused = true

	active := Column{name: "Active"}
	active.AddItem("three")
	// active.AddItem("four")

	done := Column{name: "Done"}
	done.AddItem("six")
	done.AddItem("four")
	done.AddItem("five")

	columns = append(columns, stash, active, done)
}

func update() {
	grid = [20]string{}

	for columnIndex, column := range columns {
		// draw titles
		title := fmt.Sprintf(format, strings.Join([]string{"#", column.name}, empty))
		grid[0] = strings.Join([]string{grid[0], title}, separator)

		// draw items
		for itemIndex := range maxItems {
			itemsLen := len(columns[columnIndex].items)

			// placeholder
			if itemsLen < maxItems && itemIndex >= itemsLen {
				value := fmt.Sprintf(format, empty)
				grid[itemIndex+1] = strings.Join([]string{grid[itemIndex+1], value}, separator)
				continue
			}

			// item
			item := columns[columnIndex].items[itemIndex]
			value := fmt.Sprintf(
				format,
				strings.Join([]string{getPrefix(&item), item.value}, empty),
			)

			grid[itemIndex+1] = strings.Join([]string{grid[itemIndex+1], value}, separator)
		}
	}
}

func draw() {
	clearConsole()

	fmt.Print("[q]quit", br)
	fmt.Print("[b]left        [w]right", br)
	fmt.Print("[j]down        [k]up", br)
	fmt.Print("[h]move left   [l]move right", br)
	fmt.Print(separator, br)

	for _, row := range grid {
		fmt.Print(row, br)
	}
}
