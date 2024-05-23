package handshake_handler

import (
	"context"
	"encoding/json"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/pubsub"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/ws"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/auth"
	"github.com/gofiber/contrib/websocket"
	"log/slog"
)

// WSHandshakeHandler act like handshake
type WSHandshakeHandler struct {
	hub    *ws.Hub
	pubsub pubsub.IPubSub
	auth   auth.IClient
}

func NewWSHandshakeHandler(hub *ws.Hub, pubsub pubsub.IPubSub, auth auth.IClient) *WSHandshakeHandler {
	return &WSHandshakeHandler{hub, pubsub, auth}
}

func (h *WSHandshakeHandler) Handle(conn *websocket.Conn) {
	token := conn.Headers("Authorization")

	user, err := h.auth.Validate(context.TODO(), token)
	if err != nil {
		conn.Close()
		return
	}

	client := ws.NewClient(user.Id, h.hub, conn)
	h.hub.Register(client)

	sub := h.pubsub.AddSubscriber(&pubsub.SubscriberConnectionOpts{
		CID: client.GetID(),
		UID: user.Id,
	})

	sub.RegisterNotifyHandler(func(conn *pubsub.SubscriberConnectionOpts, msg *pubsub.Message) {
		// TODO: convert payload
		// for test only
		data := &ws.MessagePayload{
			Type: "changes",
			Payload: struct {
				Topic   string `json:"test"`
				Message string `json:"msg"`
			}{
				Topic:   msg.GetTopic(),
				Message: msg.GetMessagePayload(),
			},
		}

		b, _ := json.Marshal(data)

		content, _ := ws.NewMessageData(string(b))

		h.hub.Send(ws.Message{
			ID:      conn.CID,
			From:    conn.UID,
			Content: content,
		})
	})

	go sub.Listen()
	go client.WriteMessage()
	go client.MessageListener()

	client.Run()

	slog.Debug("client disconnected")
}
