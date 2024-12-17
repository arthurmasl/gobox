package main

import "fmt"

type Kinematic struct {
	vel   int
	speed int
	acc   int
}

type Position struct {
	x, y int
}

type Stats struct {
	hp int
}

type Ai struct {
	state bool
}

type Player struct {
	*Kinematic
	*Position
	*Stats
}

type Monster struct {
	*Stats
	*Ai
}

func createPlayer() *Player {
	return &Player{
		Kinematic: &Kinematic{},
		Position:  &Position{50, 50},
		Stats:     &Stats{200},
	}
}

func createMonster() *Monster {
	return &Monster{
		Stats: &Stats{100},
		Ai:    &Ai{},
	}
}

func main() {
	player := createPlayer()
	monster := createMonster()

	fmt.Printf("%+v\n", player.Position)
	fmt.Printf("%+v\n", monster.hp)
}
