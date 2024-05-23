package ws

const (
	EVENT_CONNECT      = "connected"
	EVENT_DISCONNECTED = "disconnected"
	ON_MESSAGE         = "on_message"
)

// HubHandlerFn handler for custom events
type HubHandlerFn = func(payload *EventPayload)

// IClient represent websocket client
type IClient interface {
	MessageListener()
	GetID() string
	GetUID() string
	WriteMessage()
	Send(msg Message)
	Run()
	Close()
}

// IHub represent websocket server
type IHub interface {
	GetID() string
	Send(msg Message)
	Register(client IClient)
	Unregister(client IClient)
	On(event string, callback HubHandlerFn)
	Run()
	Stop()
}
