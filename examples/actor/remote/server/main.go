package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Message struct {
	To, From, Body string
}

type RemoteActor struct{}

func (actor *RemoteActor) ReceiveMessage(message Message, reply *string) error {
	fmt.Printf("Remote actor %s received message: %s\n", message.To, message.Body)
	*reply = "Message received"
	return nil
}

func main() {
	actor := new(RemoteActor)
	rpc.Register(actor)

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		fmt.Printf("Error listening: %s\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("Remote actor listening on localhost:1234")

	rpc.Accept(listener)
}
