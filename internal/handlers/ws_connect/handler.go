package ws_connect_handler

import (
	"github.com/WildEgor/e-shop-cdn/internal/adapters/pubsub"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/ws"
	"github.com/WildEgor/e-shop-cdn/internal/repositories"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"log/slog"
)

type WSConnectHandler struct {
	pubsub  pubsub.IPubSub
	subRepo repositories.ISubsRepository
}

func NewWSConnectHandler(pubsub pubsub.IPubSub, subRepo repositories.ISubsRepository) *WSConnectHandler {
	return &WSConnectHandler{pubsub, subRepo}
}

func (h *WSConnectHandler) Handle(p *ws.EventPayload) {
	uid := p.Client.GetUID()

	topics, err := h.subRepo.FindUserTopics(uid)
	if err != nil {
		slog.Error("fail get topics", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
		return
	}

	subConn := &pubsub.SubscriberConnectionOpts{
		CID: p.Client.GetID(),
		UID: uid,
	}
	h.pubsub.Sub(topics, subConn)
}
