// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package pkg

import (
	"github.com/WildEgor/e-shop-cdn/internal/adapters/auth"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/pubsub"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/storage"
	"github.com/WildEgor/e-shop-cdn/internal/adapters/ws"
	"github.com/WildEgor/e-shop-cdn/internal/configs"
	"github.com/WildEgor/e-shop-cdn/internal/db/mongo"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/delete"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/download"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/errors"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/get_files"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/health_check"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/metadata"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/ready_check"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/ticker"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/upload"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/ws_connect"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/ws_disconnect"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/ws_handshake"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/ws_subscribe"
	"github.com/WildEgor/e-shop-cdn/internal/handlers/ws_unsubscribe"
	"github.com/WildEgor/e-shop-cdn/internal/repositories"
	"github.com/WildEgor/e-shop-cdn/internal/routers"
	"github.com/WildEgor/e-shop-cdn/internal/services"
	"github.com/google/wire"
)

// Injectors from server.go:

func NewServer() (*Server, error) {
	configurator := configs.NewConfigurator()
	appConfig := configs.NewAppConfig(configurator)
	loggerConfig := configs.NewLoggerConfig(configurator)
	errorsHandler := error_handler.NewErrorsHandler()
	mongoConfig := configs.NewMongoConfig(configurator)
	connection := mongo.NewMongoConnection(mongoConfig)
	fileRepository := repositories.NewFileRepository(connection)
	storageConfig := configs.NewStorageConfig(configurator, appConfig)
	storageAdapter := storage.NewStorageAdapter(storageConfig)
	pubSub := pubsub.NewPubSub()
	uploadHandler := upload_handler.NewUploadHandler(fileRepository, storageAdapter, pubSub)
	deleteHandler := delete_handler.NewDeleteHandler(fileRepository, storageAdapter, pubSub)
	getFilesHandler := get_files_handler.NewGetFilesHandler(fileRepository)
	metadataHandler := metadata_handler.NewMetadataHandler(storageAdapter)
	apiKeyConfig := configs.NewApiKeyConfig(configurator)
	apiKeyValidator := services.NewApiKeyValidator(apiKeyConfig)
	privateRouter := routers.NewPrivateRouter(uploadHandler, deleteHandler, getFilesHandler, metadataHandler, apiKeyValidator)
	healthCheckHandler := health_check_handler.NewHealthCheckHandler()
	readyCheckHandler := ready_check_handler.NewReadyCheckHandler()
	downloadHandler := download_handler.NewDownloadHandler(storageAdapter)
	publicRouter := routers.NewPublicRouter(healthCheckHandler, readyCheckHandler, downloadHandler)
	swaggerRouter := routers.NewSwaggerRouter()
	hub := ws.NewHub()
	client := auth.NewClient()
	wsHandshakeHandler := handshake_handler.NewWSHandshakeHandler(hub, pubSub, client)
	subsRepository := repositories.NewSubsRepository(connection)
	subscribeHandler := sub_handler.NewSubscribeHandler(pubSub, subsRepository)
	unsubscribeHandler := unsub_handler.NewUnsubscribeHandler(pubSub, subsRepository)
	tickerHandler := ticker_handler.NewTickerHandler(pubSub)
	wsConnectHandler := ws_connect_handler.NewWSConnectHandler(pubSub, subsRepository)
	wsDisconnectHandler := ws_disconnect_handler.NewWSDisconnectHandler(pubSub)
	socketRouter := routers.NewSocketRouter(wsHandshakeHandler, subscribeHandler, unsubscribeHandler, tickerHandler, wsConnectHandler, wsDisconnectHandler, hub)
	server := NewApp(appConfig, loggerConfig, errorsHandler, privateRouter, publicRouter, swaggerRouter, socketRouter, connection, hub)
	return server, nil
}

// server.go:

var ServerSet = wire.NewSet(AppSet)
