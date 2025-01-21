package main

import (
	"fmt"
	"net/rpc"

	"gobox/internal/utils"
)

type Client struct {
	rpc *rpc.Client
}

func NewClient(address string) (*Client, error) {
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		return &Client{client}, err
	}

	return &Client{client}, nil
}

func (client *Client) PostMessage(message string) {
	var reply string
	err := client.rpc.Call("API.GetMessage", "hello from client", &reply)
	utils.CheckPanic(err)

	fmt.Println("reply:", reply)
}

func main() {
	client, err := NewClient("localhost:4444")
	utils.CheckPanic(err)
	defer client.rpc.Close()

	client.PostMessage("hello from client")
}
