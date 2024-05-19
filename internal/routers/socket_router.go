package routers

import (
	"context"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/ws"
	ticker_handler "github.com/WildEgor/e-shop-cdn/internal/handlers/ticker"
	ws_connect_handler "github.com/WildEgor/e-shop-cdn/internal/handlers/ws_connect"
	ws_connect_handler2 "github.com/WildEgor/e-shop-cdn/internal/handlers/ws_disconnect"
	handshake_handler "github.com/WildEgor/e-shop-cdn/internal/handlers/ws_handshake"
	sub_handler "github.com/WildEgor/e-shop-cdn/internal/handlers/ws_subscribe"
	unsub_handler "github.com/WildEgor/e-shop-cdn/internal/handlers/ws_unsubscribe"
	"github.com/WildEgor/e-shop-cdn/internal/models"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type SocketRouter struct {
	whh *handshake_handler.WSHandshakeHandler
	sh  *sub_handler.SubscribeHandler
	ush *unsub_handler.UnsubscribeHandler
	th  *ticker_handler.TickerHandler
	ch  *ws_connect_handler.WSConnectHandler
	dch *ws_connect_handler2.WSDisconnectHandler
	hub ws.IHub
}

func NewSocketRouter(whh *handshake_handler.WSHandshakeHandler, sh *sub_handler.SubscribeHandler, ush *unsub_handler.UnsubscribeHandler, th *ticker_handler.TickerHandler, ch *ws_connect_handler.WSConnectHandler, dch *ws_connect_handler2.WSDisconnectHandler, hub ws.IHub) *SocketRouter {
	return &SocketRouter{whh: whh, sh: sh, ush: ush, th: th, ch: ch, dch: dch, hub: hub}
}

// Setup
func (sr *SocketRouter) Setup(app *fiber.App) {
	app.Get("/ws", websocket.New(sr.whh.Handle))

	sr.hub.On(ws.EVENT_CONNECT, sr.ch.Handle)
	sr.hub.On(ws.EVENT_DISCONNECTED, sr.dch.Handle)
	sr.hub.On(models.EVENT_SUB_TOPIC, sr.sh.Handle)
	sr.hub.On(models.EVENT_UNSUB_TOPIC, sr.ush.Handle)

	go sr.th.Handle(context.TODO())
}
