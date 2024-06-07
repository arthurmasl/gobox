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
	if targetIndex >= 0 && targetIndex < len(columns) {
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
	// TODO
}
