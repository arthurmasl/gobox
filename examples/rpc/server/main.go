package main

import (
	"fmt"
	"net"
	"net/rpc"
)

const address = "localhost:4444"

type API struct{}

func (api *API) GetMessage(message string, reply *string) error {
	fmt.Printf("Received message from client: %s\n", message)
	*reply = "hello from server"
	return nil
}

func main() {
	api := new(API)
	rpc.Register(api)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	defer listener.Close()

	fmt.Println("Listening RPC on", address)
	rpc.Accept(listener)
}
