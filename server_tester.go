package main

import (
	"fmt"

	"github.com/mateun/network_dispo/messaging"
	"github.com/mateun/network_dispo/tcp"
)

type MyHandler struct {
}

func (mh *MyHandler) Handle() error {
	fmt.Println("handling the message")
	return nil
}

func main() {
	// Make sure our struct implements the MessageHandlerPlugin interface
	var _ messaging.MessageHandlerPlugin = (*MyHandler)(nil)
	handler := new(MyHandler)
	tcp.Start_tcp_server(20100, handler)
}
