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

func (actor *Actor) ProcessMessages() {
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
		actor.SendMessage(message)
	} else {
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
	actorManager := NewActorManager()

	actorManager.RegisterActor("actor1")
	actorManager.RegisterActor("actor2")

	// send message to local actor
	actorManager.SendMessage(Message{
		To:   "actor1",
		From: "actor2",
		Body: "Hello from actor2 to actor1",
	})

	// send message to remote actor
	actorManager.SendMessage(Message{
		To:   "localhost:1234",
		From: "actor1",
		Body: "Hello from actor1 to localhost:1234",
	})

	// process messages in all actors
	for _, actor := range actorManager.Actors {
		actor.ProcessMessages()
	}
}
