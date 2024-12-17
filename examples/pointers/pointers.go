package main

import (
	"fmt"
)

type Column struct {
	name  string
	items Items
}

type Items []Item

type Item struct {
	value   string
	focused bool
}

// func (items *Items) Add(value string) {
// 	items = append(items, Item{value, false})
// }

var columns = []Column{
	{name: "col1", items: []Item{
		{value: "one"},
	}},
}

func main() {
	items := &columns[0].items
	*items = append(*items, Item{"NEW ITEM", false})

	fmt.Println(columns)
}
