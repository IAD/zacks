package main

import (
	"context"
	"fmt"
	"log"
	"time"

	db_cache "github/IAD/zacks/internal/pkg/db-cache"
	"github/IAD/zacks/internal/pkg/zacks"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx := context.Background()

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
		zacks.WithRefresher(time.Minute*20),
	)

	rating, err := z.GetRating(ctx, "AAPL")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v", rating)
}
