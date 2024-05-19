package ws

import (
	"encoding/json"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

const (
	maxMessageSize = 512
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
)

var _ IClient = (*Client)(nil)

// Client represent websocket connection
type Client struct {
	/* id as uuid */
	id string
	/* uid as user_id */
	uid string
	/* hub for interconnections */
	hub *Hub
	/* conn websocket connection */
	conn *websocket.Conn
	/* send channel for notifications */
	send chan Message
	/* done for termination */
	done chan struct{}
}

// NewClient create client with (or without) unique id
func NewClient(uid string, hub *Hub, conn *websocket.Conn) *Client {
	return &Client{
		id:   uuid.New().String(),
		uid:  uid,
		hub:  hub,
		conn: conn,
		send: make(chan Message),
		done: make(chan struct{}),
	}
}

// GetID return client unique id
func (c *Client) GetID() string {
	return c.id
}

// GetUID return client uid
func (c *Client) GetUID() string {
	return c.uid
}

// Send route notification
func (c *Client) Send(msg Message) {
	c.send <- msg
}

// Close channels
func (c *Client) Close() {
	close(c.send)
	c.done <- struct{}{}
}

// Run keep main goroutine wait (dummy solution)
func (c *Client) Run() {
	c.hub.fireEvent(EVENT_CONNECT, c, nil, nil)

	<-c.done
}

// MessageListener listen incoming websocket messages
func (c *Client) MessageListener() {
	defer func() {
		slog.Debug("close ws connection", models.LogEntryAttr(&models.LogEntry{
			Props: map[string]interface{}{
				"id": c.GetID(),
			},
		}))

		c.hub.fireEvent(EVENT_DISCONNECTED, c, nil, nil)

		c.hub.Unregister(c)
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		mType, readMessage, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error("ws error", models.LogEntryAttr(&models.LogEntry{
					Err: err,
				}))
			}
			break
		}

		slog.Debug("message", models.LogEntryAttr(&models.LogEntry{
			Props: map[string]interface{}{
				"type": mType,
				"data": string(readMessage),
			},
		}))

		content, err := NewMessageData(string(readMessage))

		data, _ := json.Marshal(content.Payload.Payload)

		if content.Payload != nil {
			c.hub.fireEvent(content.Payload.Type, c, data, err)
		}
	}
}

// WriteMessage send messages back to websocket client
func (c *Client) WriteMessage() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		c.hub.fireEvent(EVENT_DISCONNECTED, c, nil, nil)

		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if message.From != c.GetID() || message.IsBroadcast {
				err := c.conn.WriteJSON(message)
				if err != nil {
					slog.Error("ws write message error", models.LogEntryAttr(&models.LogEntry{
						Err: err,
					}))
				}
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
