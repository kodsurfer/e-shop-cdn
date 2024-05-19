package ws_disconnect_handler

import (
	"github.com/WildEgor/e-shop-cdn/internal/adapters/pubsub"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/ws"
)

type WSDisconnectHandler struct {
	pubsub pubsub.IPubSub
}

func NewWSDisconnectHandler(pubsub pubsub.IPubSub) *WSDisconnectHandler {
	return &WSDisconnectHandler{pubsub}
}

func (h *WSDisconnectHandler) Handle(p *ws.EventPayload) {
	uid := p.Client.GetUID()

	h.pubsub.RemoveSubscriber(&pubsub.SubscriberConnectionOpts{
		CID: p.UUID,
		UID: uid,
	})
}
