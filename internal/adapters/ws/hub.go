package ws

import (
	"github.com/google/uuid"
)

var _ IHub = (*Hub)(nil)

// Hub represent and keep websocket connections
type Hub struct {
	/* id uuid */
	id string
	/* send channel for messages */
	send chan Message
	/* clients keep clients map */
	clients safePool
	/* listeners keep listeners map */
	listeners safeListeners
	/* done stop listen */
	done chan struct{}
}

// NewHub create default hub with uuid
func NewHub() *Hub {
	return &Hub{
		id: uuid.New().String(),
		clients: safePool{
			conn: make(map[string]IClient),
		},
		listeners: safeListeners{
			list: make(map[string][]HubHandlerFn),
		},
		send: make(chan Message),
	}
}

// GetID remove hub uuid
func (h *Hub) GetID() string {
	return h.id
}

// Send receive notification messages
func (h *Hub) Send(msg Message) {
	h.send <- msg
}

// Register add new client
func (h *Hub) Register(client IClient) {
	h.clients.set(client)
}

// Unregister remove existed client
func (h *Hub) Unregister(client IClient) {
	h.clients.delete(client.GetID())
	client.Close()
}

// On add event listener
func (h *Hub) On(event string, callback HubHandlerFn) {
	h.listeners.set(event, callback)
}

// fireEvent execute event
func (h *Hub) fireEvent(event string, client *Client, data []byte, error error) {
	callbacks := h.listeners.get(event)

	for _, callback := range callbacks {
		callback(&EventPayload{
			Name:   event,
			Client: client,
			Data:   data,
			Error:  error,
		})
	}
}

// Run keep listen messages
func (h *Hub) Run() {
	for {
		select {
		case message := <-h.send:
			for _, client := range h.clients.all() {
				client.Send(message)
			}
		case <-h.done:
			close(h.send)
			break
		}
	}
}

// Stop listen
func (h *Hub) Stop() {
	h.done <- struct{}{}
}
