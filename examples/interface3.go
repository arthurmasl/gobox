package main

import (
	"fmt"
	"reflect"
)

type Renderable interface {
	render()
	setIndex(int)
}

type Entity struct {
	x, y  int
	index int
}

type Player struct {
	Entity
	hp int
}

type Tree struct {
	Entity
}

func (e Entity) render() {
	fmt.Printf("entity %v rendered\n", reflect.TypeOf(e))
}

func (e *Entity) setIndex(index int) {
	e.index = index
}

func main() {
	player := Player{}
	tree := Tree{}

	renderList := make([]Renderable, 0)
	renderList = append(renderList, &player)
	renderList = append(renderList, &tree)

	fmt.Printf("%#v\n", renderList)
	// renderList[0].render()
}
