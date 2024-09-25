package main

import (
	"fmt"
	"log"
)

type ServerState int

const (
	Idle ServerState = iota
	Connected
	Error
	Retrying
)

var stateName = map[ServerState]string{
	Idle:      "idle",
	Connected: "conn",
	Error:     "err",
	Retrying:  "ret",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	transition(Error)
	fmt.Println(Idle)
}

func transition(s ServerState) {
	switch s {
	case Idle:
		log.Println(Idle)
	case Connected:
		log.Println(Connected)
	case Error:
		log.Println(Error)
	case Retrying:
		log.Println(Retrying)
	}
}
