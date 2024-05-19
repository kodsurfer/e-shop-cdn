package adapters

import (
	eShopAuth "github.com/WildEgor/e-shop-cdn/internal/adapters/auth"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/pubsub"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/storage"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/ws"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/auth"
	"github.com/google/wire"
)

// AdaptersSet represent integration with external services
var AdaptersSet = wire.NewSet(
	storage.NewStorageAdapter,
	wire.Bind(new(storage.IFileStorage), new(*storage.StorageAdapter)),
	pubsub.NewPubSub,
	wire.Bind(new(pubsub.IPubSub), new(*pubsub.PubSub)),
	ws.NewHub,
	wire.Bind(new(ws.IHub), new(*ws.Hub)),
	eShopAuth.NewClient,
	wire.Bind(new(auth.IClient), new(*eShopAuth.Client)),
)
