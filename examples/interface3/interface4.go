package main

import (
	"fmt"
)

type entity interface {
	update()
	draw()
}

type player struct {
	pos complex128
	hp  int
}

func (this *player) draw()   { fmt.Printf("drawed at (%v:%v)\n", real(this.pos), imag(this.pos)) }
func (this *player) update() { this.pos += 1 }

type creature struct {
	pos   complex128
	hp    int
	count int
}

func (this *creature) draw()   { fmt.Printf("drawed at (%v:%v)\n", real(this.pos), imag(this.pos)) }
func (this *creature) update() { this.pos += 2 }

func main() {
	entities := make([]entity, 0)
	entities = append(entities, &player{hp: 100}, &creature{hp: 50, count: 10})

	for _, entity := range entities {
		fmt.Printf("%T\n", entity)
		entity.update()
		entity.draw()

		switch entity := entity.(type) {
		case *player:
			fmt.Printf("player %vhp\n", entity.hp)
		case *creature:
			fmt.Printf("enemy %vhp\n", entity.hp)
		}

		fmt.Println()
	}
}
