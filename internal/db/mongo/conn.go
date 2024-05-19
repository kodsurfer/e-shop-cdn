package mongo

import (
	"context"
	"github.com/WildEgor/e-shop-cdn/internal/configs"
	"github.com/WildEgor/e-shop-gopack/pkg/libs/logger/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

// Connection represent connection to mongo
type Connection struct {
	Client        *mongo.Client
	mongoDbConfig *configs.MongoConfig
}

func NewMongoConnection(
	mongoDbConfig *configs.MongoConfig,
) *Connection {
	conn := &Connection{
		nil,
		mongoDbConfig,
	}

	conn.Connect()

	return conn
}

// Connect make connection client
func (mc *Connection) Connect() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mc.mongoDbConfig.URI))
	if err != nil {
		slog.Error("fail connect mongo", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
		panic(err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		slog.Error("fail connect mongo", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
	}

	slog.Info("success connect to mongodb")

	mc.Client = client
}

// Disconnect close connection
func (mc *Connection) Disconnect() {
	if mc.Client == nil {
		return
	}

	if err := mc.Client.Disconnect(context.TODO()); err != nil {
		slog.Error("fail disconnect mongo", models.LogEntryAttr(&models.LogEntry{
			Err: err,
		}))
		return
	}

	slog.Info("connection to mongodb closed")
}

// Db return db source from config
func (mc *Connection) Db() *mongo.Database {
	return mc.Client.Database(mc.mongoDbConfig.DbName)
}
