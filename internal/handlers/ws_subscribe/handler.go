package sub_handler

import (
	"encoding/json"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/pubsub"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/ws"
	"github.com/WildEgor/e-shop-cdn/internal/dtos"
	"github.com/WildEgor/e-shop-cdn/internal/repositories"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"log/slog"
)

type SubscribeHandler struct {
	ps pubsub.IPubSub
	sr repositories.ISubsRepository
}

func NewSubscribeHandler(ps pubsub.IPubSub, sr repositories.ISubsRepository) *SubscribeHandler {
	return &SubscribeHandler{
		ps,
		sr,
	}
}

func (h *SubscribeHandler) Handle(p *ws.EventPayload) {
	dto := &dtos.SubscribePayloadDto{}
	if err := json.Unmarshal(p.Data, dto); err != nil {
		return
	}

	uid := p.Client.GetUID()

	_, err := h.sr.SubscribeToTopic(uid, dto.Topic)
	if err != nil {
		slog.Error("fail sub topics", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
		return
	}

	h.ps.Sub([]string{dto.Topic}, &pubsub.SubscriberConnectionOpts{
		CID: p.UUID,
		UID: uid,
	})
}
