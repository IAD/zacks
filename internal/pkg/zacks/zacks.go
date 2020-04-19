package zacks

import (
	"context"
	"fmt"
	"time"

	"github/IAD/zacks/internal/pkg/cache"
	"github/IAD/zacks/internal/pkg/fetcher"
	"github/IAD/zacks/internal/pkg/models"
	"github/IAD/zacks/internal/pkg/refresher"
)

func NewZacks(
	ctx context.Context,
	opts ...Option,
) *Zacks {
	z := &Zacks{}

	for _, opt := range opts {
		opt(z)
	}

	if z.refresher != nil && z.fetcher != nil {
		go func() {
			_ = z.refresher.Start(ctx, z.cache, z.dbCache, z.fetcher)
		}()
	}

	return z
}

type Zacks struct {
	cache     *cache.Cache
	fetcher   *fetcher.Fetcher
	refresher *refresher.Refresher
	dbCache   persistable
}

func (z *Zacks) GetRating(ctx context.Context, ticker string) (*models.Rating, error) {
	var rating *models.Rating
	var err error
	if z.cache != nil {
		rating, err = z.cache.GetRating(ctx, ticker)
		if rating != nil && err == nil {
			return rating, err
		}
	}

	if z.refresher != nil {
		z.refresher.Watch(ticker)
	}

	if z.dbCache != nil {
		rating, err := z.dbCache.GetRating(ctx, ticker)
		if rating != nil {
			if z.cache != nil {
				go func() {
					_ = z.cache.AddRating(ctx, *rating)
				}()
			}
			if err == nil {
				return rating, err
			}
		}
	}

	if z.fetcher != nil {
		rating, err := z.fetcher.GetRating(ctx, ticker)
		if rating != nil {
			if z.cache != nil {
				go func() {
					_ = z.cache.AddRating(ctx, *rating)
				}()
			}
			if z.dbCache != nil {
				go func() {
					_ = z.dbCache.AddRating(ctx, *rating)
				}()
			}
		}

		return rating, err
	}

	return nil, fmt.Errorf("can't process. The ticker %s isn't found in cache and don't have a fetcher", ticker)
}

func (z *Zacks) GetHistory(ctx context.Context, ticker string) ([]models.Rating, error) {
	if z.cache != nil {
		ratings, err := z.cache.GetHistory(ctx, ticker)
		if err == nil && len(ratings) > 0 {
			return ratings, nil
		}
	}

	if z.dbCache != nil {
		ratings, err := z.dbCache.GetHistory(ctx, ticker)
		if err == nil && len(ratings) > 0 {
			return ratings, nil
		}
	}

	if z.fetcher != nil {
		rating, err := z.fetcher.GetRating(ctx, ticker)
		if err == nil && rating != nil {
			return []models.Rating{*rating}, nil
		}
	}

	return []models.Rating{}, fmt.Errorf("can't process. The ticker %s isn't found in cache", ticker)
}

type Option func(*Zacks)

func WithCache() Option {
	return func(z *Zacks) {
		z.cache = cache.NewCache()
	}
}

type persistable interface {
	GetRating(ctx context.Context, ticker string) (*models.Rating, error)
	GetHistory(ctx context.Context, ticker string) ([]models.Rating, error)
	AddRating(ctx context.Context, rating models.Rating) error
}

func WithDBCache(dbCache persistable) Option {
	return func(z *Zacks) {
		z.dbCache = dbCache
	}
}

func WithFetcher(timeout time.Duration) Option {
	return func(z *Zacks) {
		z.fetcher = fetcher.NewFetcher(timeout)
	}
}

func WithRefresher(period time.Duration) Option {
	return func(z *Zacks) {
		z.refresher = refresher.NewRefresher(period)
	}
}
