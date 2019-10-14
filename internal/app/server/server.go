package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github/IAD/zacks/internal/app/server/gen/server/restapi"
	"github/IAD/zacks/internal/app/server/gen/server/restapi/operations"
	db_cache "github/IAD/zacks/internal/pkg/db-cache"
	"github/IAD/zacks/internal/pkg/zacks"

	"github.com/go-openapi/loads"
	"github.com/joeshaw/envdecode"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	ServerPort   int  `env:"SERVER_PORT,default=8080"`
	CacheEnabled bool `env:"CACHE_ENABLED,default=true"`

	DBCacheEnabled             bool   `env:"DBCACHE_ENABLED,default=true"`
	DBCacheMongodbURL          string `env:"DBCACHE_MONGODB_URL,default=mongodb://localhost:27017"`
	DBCacheMongodbDatabaseName string `env:"DBCACHE_MONGODB_DATABASE_NAME,default=zacks"`

	FetcherEnabled       bool  `env:"FETCHER_ENABLED,default=true"`
	FetcherTimoutSeconds int64 `env:"FETCHER_TIMEOUT_SECONDS,default=5"`

	RefresherEnabled       bool  `env:"REFRESHER_ENABLED,default=true"`
	RefresherRescanSeconds int64 `env:"REFRESHER_RESCAN_SECONDS,default=600"`
}

func main() {

	var config Config
	err := envdecode.Decode(&config)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logger := logrus.NewEntry(logrus.StandardLogger())

	opts := make([]zacks.Option, 0)

	if config.CacheEnabled {
		opts = append(opts, zacks.WithCache())
	}

	if config.DBCacheEnabled {
		mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DBCacheMongodbURL))
		if err != nil {
			log.Fatal(err.Error())
		}
		mongoDB := mongoClient.Database(config.DBCacheMongodbDatabaseName)
		opts = append(opts, zacks.WithDBCache(db_cache.NewDBCache(mongoDB)))
	}

	if config.FetcherEnabled {
		opts = append(opts, zacks.WithFetcher(time.Duration(time.Second)*time.Duration(config.FetcherTimoutSeconds)))
	}

	if config.RefresherEnabled {
		opts = append(opts, zacks.WithRefresher(time.Duration(time.Second)*time.Duration(config.RefresherRescanSeconds)))
	}

	beautyConfig, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(beautyConfig))

	z := zacks.NewZacks(ctx, opts...)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	handlers := NewHandlers(z)

	api := operations.NewZacksAPI(swaggerSpec)

	// Handler for GET /{ticker}
	api.GetTickerHandler = operations.GetTickerHandlerFunc(handlers.GetTickerHandler)
	// Handler for GET /{ticker}/history
	api.GetTickerHistoryHandler = operations.GetTickerHistoryHandlerFunc(handlers.GetTickerHistoryHandler)

	server := restapi.NewServerWithMiddleware(api, "zacks", logger)
	server.Port = config.ServerPort

	log.Printf("Starting server on the port %v", config.ServerPort)
	err = server.Serve()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
