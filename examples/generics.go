package main

import (
	"fmt"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))

	for k := range m {
		r = append(r, k)
	}

	return r
}

type Player struct {
	hp int
}

type Enemy struct {
	mp int
}

func UpdateEntity(entity interface{}) {
	switch t := entity.(type) {
	case Player:
		fmt.Println("pl", t)
	case Enemy:
		fmt.Println("en", t)

	}
	fmt.Printf("%#v\n", entity)
}

func main() {
	player := Player{}
	enemy := Enemy{}

	UpdateEntity(player)
	UpdateEntity(enemy)
}
