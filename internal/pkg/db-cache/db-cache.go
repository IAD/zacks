package db_cache

import (
	"context"
	"fmt"

	"github/IAD/zacks/internal/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewDBCache(mongoDB *mongo.Database) *DBCache {
	return &DBCache{
		mongoDB: mongoDB,
	}
}

type DBCache struct {
	mongoDB *mongo.Database
}

func (c *DBCache) GetRating(ctx context.Context, ticker string) (*models.Rating, error) {
	return nil, fmt.Errorf("not implemented")
}

func (c *DBCache) GetHistory(ctx context.Context, ticker string) ([]models.Rating, error) {
	return make([]models.Rating, 0), fmt.Errorf("not implemented")
}

func (c *DBCache) AddRating(ctx context.Context, rating models.Rating) error {
	return fmt.Errorf("not implemented")
}
