package unsub_handler

import (
	"encoding/json"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/pubsub"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/ws"
	"github.com/WildEgor/e-shop-cdn/internal/dtos"
	"github.com/WildEgor/e-shop-cdn/internal/repositories"
)

type UnsubscribeHandler struct {
	ps pubsub.IPubSub
	sr repositories.ISubsRepository
}

func NewUnsubscribeHandler(ps pubsub.IPubSub, sr repositories.ISubsRepository) *UnsubscribeHandler {
	return &UnsubscribeHandler{
		ps,
		sr,
	}
}

func (h *UnsubscribeHandler) Handle(p *ws.EventPayload) {
	uid := p.Client.GetUID()

	dto := &dtos.UnsubscribePayloadDto{}
	if err := json.Unmarshal(p.Data, dto); err != nil {
		return
	}

	topic, _ := h.sr.FindTopicById(dto.TopicId.String())
	if topic != nil {
		h.sr.UnsubscribeFromTopicById(topic.Id.String())

		h.ps.Unsub([]string{topic.Topic}, &pubsub.SubscriberConnectionOpts{
			CID: p.UUID,
			UID: uid,
		})
	}
}
