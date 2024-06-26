package main

import (
	"fmt"
	"os"
	"slices"
	"syscall"

	"golang.org/x/term"
)

var (
	focusedColumn = 0
	focusedItem   = 0
	inputMode     = false
)

func handleInput() {
	wg.Add(1)

	oldState, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		fmt.Println("error making terminal raw")
	}
	defer term.Restore(int(syscall.Stdin), oldState)

	buf := make([]byte, 1024)

	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println("error reading input")
		}

		if inputMode {
			addItem(string(buf))
			oldState, _ = term.MakeRaw(int(syscall.Stdin))
			inputMode = false
		}

		char := buf[0]

		if char == 'q' {
			wg.Done()
			break
		}

		focusedColumn, focusedItem = getFocusedIndexes()

		switch char {
		case 'w':
			moveToColumn(focusedColumn + 1)
		case 'b':
			moveToColumn(focusedColumn - 1)
		case 'j':
			moveToItem(focusedItem + 1)
		case 'k':
			moveToItem(focusedItem - 1)
		case 'h':
			moveItem(focusedColumn - 1)
		case 'l':
			moveItem(focusedColumn + 1)
		case 'x':
			deleteItem()

		case 'a':
			term.Restore(int(syscall.Stdin), oldState)
			inputMode = true

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

func moveToColumn(targetIndex int) {
	if targetIndex >= 0 && targetIndex < len(columns) &&
		len(columns[targetIndex].items) > 0 {
		targetItem := focusedItem

		if len(columns[targetIndex].items) <= targetItem {
			targetItem = len(columns[targetIndex].items) - 1
		}

		columns[targetIndex].items[targetItem].focused = true
		columns[focusedColumn].items[focusedItem].focused = false
	}
}

func moveToItem(targetIndex int) {
	if targetIndex >= 0 && targetIndex < len(columns[focusedColumn].items) {
		columns[focusedColumn].items[targetIndex].focused = true
		columns[focusedColumn].items[focusedItem].focused = false
	}
}

func deleteItem() {
	items := &columns[focusedColumn].items

	*items = slices.Delete(
		*items,
		focusedItem,
		focusedItem+1,
	)

	updateCursor()
}

func updateCursor() {
	items := &columns[focusedColumn].items

	if len(*items) == 0 {
		return
	}
	if focusedItem >= len(*items) {
		(*items)[len(*items)-1].focused = true
		return
	}
	(*items)[focusedItem].focused = true
}

func moveItem(targetIndex int) {
	if targetIndex > 2 && focusedColumn == 2 || targetIndex < 0 && focusedColumn == 0 {
		return
	}

	columns[focusedColumn].items[focusedItem].focused = false
	columns[targetIndex].items = append(
		columns[targetIndex].items,
		columns[focusedColumn].items[focusedItem],
	)
	columns[focusedColumn].items = slices.Delete(
		columns[focusedColumn].items,
		focusedItem,
		focusedItem+1,
	)

	updateCursor()
}

func addItem(value string) {
	columns[focusedColumn].items[focusedItem].focused = false
	columns[focusedColumn].items = append(columns[focusedColumn].items, Item{value, true})
}
