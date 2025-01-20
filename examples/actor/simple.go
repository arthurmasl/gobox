package main

import "fmt"

type Actor struct {
	inbox chan string
}

func NewActor() *Actor {
	return &Actor{inbox: make(chan string)}
}

func (actor *Actor) Receive() {
	for message := range actor.inbox {
		fmt.Println(message)
	}
}

func (actor *Actor) Send(message string) {
	actor.inbox <- message
}

func main() {
	actor1 := NewActor()
	go actor1.Receive()

	actor1.Send("hello")
	actor1.Send("world")
}
