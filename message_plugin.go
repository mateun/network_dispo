package main

type MessageHandlerPlugin interface {
	Handle() error
}
