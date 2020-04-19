package cache

import (
	"context"
	"sync"

	"github/IAD/zacks/internal/pkg/models"
)

func NewCache() *Cache {
	return &Cache{
		history: make(map[string][]models.Rating),
	}
}

type Cache struct {
	lock    sync.RWMutex
	history map[string][]models.Rating
}

func (c *Cache) GetRating(ctx context.Context, ticker string) (*models.Rating, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	ratings, ok := c.history[ticker]
	if !ok {
		return nil, nil
	}
	if len(ratings) > 0 {
		rating := ratings[len(ratings)-1]
		return &rating, nil
	}

	return nil, nil
}

func (c *Cache) GetHistory(ctx context.Context, ticker string) ([]models.Rating, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	ratings, ok := c.history[ticker]
	if !ok {
		return make([]models.Rating, 0), nil
	}

	result := make([]models.Rating, 0)
	result = append(result, ratings...)

	return result, nil
}

func (c *Cache) AddRating(ctx context.Context, rating models.Rating) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	ratings, ok := c.history[rating.Ticker]
	if !ok {
		ratings = make([]models.Rating, 0)
	}

	// skip duplicates
	if len(ratings) > 0 {
		lastRating := ratings[len(ratings)-1]
		if lastRating.Equals(rating) {
			return nil
		}
	}

	ratings = append(ratings, rating)

	c.history[rating.Ticker] = ratings
	return nil
}
