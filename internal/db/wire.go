package db

import (
	"github.com/WildEgor/e-shop-cdn/internal/db/mongo"
	"github.com/google/wire"
)

// DbSet provide db connections
var DbSet = wire.NewSet(
	mongo.NewMongoConnection,
)
