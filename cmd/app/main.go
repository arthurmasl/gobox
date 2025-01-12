package main

import (
	"fmt"
)

type player struct {
	pos complex128
	hp  int
}

type enemy struct {
	pos   complex128
	hp    int
	aggro bool
}

func (this *player) draw()   { fmt.Printf("player drawed at %v\n", this.pos) }
func (this *player) update() { this.pos += 1 }

func (this *enemy) draw()   { fmt.Printf("enemy drawed at %v\n", this.pos) }
func (this *enemy) update() { this.pos += 2 }

type entity interface {
	update()
	draw()
}

func main() {
	entities := make([]entity, 0)
	entities = append(entities, &player{}, &enemy{})

	for _, entity := range entities {
		fmt.Printf("%T\n", entity)
		entity.update()
		entity.draw()

		switch entity := entity.(type) {
		case *player:
			fmt.Println(entity.hp, entity.pos)
			fmt.Println("player")
		case *enemy:
			fmt.Println(entity.aggro)
			fmt.Println("enemy")
		}

		fmt.Println()
	}
}
