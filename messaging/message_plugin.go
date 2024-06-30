package messaging

type MessageHandlerPlugin interface {
	Handle() error
}
