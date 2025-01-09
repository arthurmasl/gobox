package main

import (
	"fmt"
	"reflect"
)

type State rune

const (
	Idle State = iota
	Run
	Attack
)

var states = map[State]string{Idle: "idle", Run: "run", Attack: "attack"}

func (state State) String() string {
	return states[state]
}

func main() {
	for k, v := range states {
		fmt.Println(k, v)
	}

	printType(Idle)
	printType("aa")
	printType('a')
	printType(complex(11, 55))
	printType(1<<63 - 1)
	printType(int64(1<<63 - 1))
}

func printType(element any) {
	fmt.Print("element type is ")

	switch element.(type) {
	case State:
		fmt.Println("state")
	default:
		fmt.Println(reflect.TypeOf(element))
	}
}
