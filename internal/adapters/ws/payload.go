package ws

import "github.com/gofiber/contrib/websocket"

// EventPayload handlers payload
type EventPayload struct {
	/* Name of event */
	Name string
	/* UUID of connection */
	UUID   string
	Client IClient
	Conn   *websocket.Conn
	/* Data raw bytes */
	Data []byte
	/* Error instance */
	Error error
}
