package main

import (
	"fmt"
	"net/rpc"
)

type Actor struct {
	Name     string
	Messages []Message
}

type Message struct {
	To, From, Body string
}

func (actor *Actor) SendMessage(message Message) {
	actor.Messages = append(actor.Messages, message)
}

func (actor *Actor) ProcessMessages(message Message) {
	for _, message := range actor.Messages {
		fmt.Printf("Actor %s received message: %s\n", actor.Name, message.Body)
	}

	actor.Messages = nil
}

type ActorManager struct {
	Actors map[string]*Actor
}

func NewActorManager() *ActorManager {
	return &ActorManager{Actors: make(map[string]*Actor)}
}

func (manager *ActorManager) RegisterActor(name string) {
	manager.Actors[name] = &Actor{Name: name}
}

func (manager *ActorManager) SendMessage(message Message) {
	if actor, ok := manager.Actors[message.To]; ok {
		client, err := rpc.Dial("tcp", message.To)
		if err != nil {
			fmt.Printf("Error connecting to remote actor %s: %s\n", message.To, err)
			return
		}
		defer client.Close()

		var reply string
		err = client.Call("RemoteActor.ReceiveMessage", message, &reply)
		if err != nil {
			fmt.Printf("Error sending message to remote actor %s: %s\n", message.To, err)
			return
		}
	}
}

func main() {
}
