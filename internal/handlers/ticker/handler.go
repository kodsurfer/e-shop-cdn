package ticker_handler

import (
	"context"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/pubsub"
	"time"
)

// TickerHandler dummy notification for test
type TickerHandler struct {
	ps pubsub.IPubSub
}

func NewTickerHandler(ps pubsub.IPubSub) *TickerHandler {
	return &TickerHandler{
		ps,
	}
}

func (h *TickerHandler) Handle(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)

	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			// h.ps.Publish("/a/b/*", "some changes in your dir")
		case <-ctx.Done():
			break
		}
	}
}
