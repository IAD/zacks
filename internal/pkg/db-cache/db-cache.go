package db_cache

import (
	"context"
	"sync"

	"github/IAD/zacks/internal/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func NewDBCache(mongoDB *mongo.Database) *DBCache {
	return &DBCache{
		saveCache: make(map[string]models.Rating),
		mongoDB:   mongoDB,
	}
}

type DBCache struct {
	lock      sync.RWMutex
	saveCache map[string]models.Rating
	mongoDB   *mongo.Database
}

func (c *DBCache) GetRating(ctx context.Context, ticker string) (*models.Rating, error) {
	cursor, err := c.mongoDB.Collection("zacks").Find(
		ctx,
		bson.D{{"ticker", ticker}},
		options.Find().SetSort(bsonx.Doc{{"date_received", bsonx.Int32(1)}}).SetLimit(1),
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		err = cursor.Err()
		if err != nil {
			return nil, err
		}

		result := &models.Rating{}
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	return nil, nil
}

func (c *DBCache) GetHistory(ctx context.Context, ticker string) ([]models.Rating, error) {
	cursor, err := c.mongoDB.Collection("zacks").Find(
		ctx,
		bson.D{{"ticker", ticker}},
		options.Find().SetSort(bsonx.Doc{{"date_received", bsonx.Int32(1)}}),
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := make([]models.Rating, 0)
	for cursor.Next(ctx) {
		err = cursor.Err()
		if err != nil {
			return nil, err
		}

		item := models.Rating{}
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}

func (c *DBCache) AddRating(ctx context.Context, rating models.Rating) error {
	c.lock.Lock()
	currentRating, exists := c.saveCache[rating.Ticker]
	c.lock.Unlock()
	if exists {
		if currentRating.Equals(rating) {
			return nil
		}
	}

	_, err := c.mongoDB.Collection("zacks").InsertOne(ctx, rating)
	if err != nil {
		return err
	}

	c.lock.Lock()
	c.saveCache[rating.Ticker] = rating
	c.lock.Unlock()

	return nil
}
