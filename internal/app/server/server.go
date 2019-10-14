package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github/IAD/zacks/internal/app/server/gen/server/restapi"
	"github/IAD/zacks/internal/app/server/gen/server/restapi/operations"
	db_cache "github/IAD/zacks/internal/pkg/db-cache"
	"github/IAD/zacks/internal/pkg/zacks"

	"github.com/go-openapi/loads"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx := context.Background()
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logger := logrus.NewEntry(logrus.StandardLogger())

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err.Error())
	}
	mongoDB := mongoClient.Database("zacks")

	z := zacks.NewZacks(
		ctx,
		zacks.WithFetcher(time.Second*5),
		zacks.WithCache(),
		zacks.WithDBCache(db_cache.NewDBCache(mongoDB)),
		zacks.WithRefresher(time.Minute),
	)

	rating, err := z.GetRating(ctx, "AAPL")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v", rating)

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
	server.Port = 8080

	err = server.Serve()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
