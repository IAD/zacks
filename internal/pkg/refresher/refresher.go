package refresher

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github/IAD/zacks/internal/pkg/cache"
	"github/IAD/zacks/internal/pkg/fetcher"
	"github/IAD/zacks/internal/pkg/models"
)

type persistable interface {
	GetRating(ctx context.Context, ticker string) (*models.Rating, error)
	GetHistory(ctx context.Context, ticker string) ([]models.Rating, error)
	AddRating(ctx context.Context, rating models.Rating) error
}

func NewRefresher(period time.Duration) *Refresher {
	return &Refresher{
		period:          period,
		nextRefreshTime: make(map[string]time.Time),
		blockedTickers:  make(map[string]struct{}),
	}
}

type Refresher struct {
	lock            sync.RWMutex
	period          time.Duration
	nextRefreshTime map[string]time.Time
	blockedTickers  map[string]struct{}
}

func (r *Refresher) Watch(ticker string) {
	r.lock.Lock()
	defer r.lock.Unlock()

	_, ok := r.nextRefreshTime[ticker]
	if !ok {
		r.nextRefreshTime[ticker] = time.Now()
	}
}

func (r *Refresher) Start(
	ctx context.Context,
	cache *cache.Cache,
	dbCache persistable,
	fetcher *fetcher.Fetcher,
) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	if fetcher == nil {
		return fmt.Errorf("fetcher required")
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			_ = r.do(ctx, cache, dbCache, fetcher)
		}
	}
}

func (r *Refresher) do(
	ctx context.Context,
	cache *cache.Cache,
	dbCache persistable,
	fetcher *fetcher.Fetcher,
) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			ticker := r.getNextTickerForUpdate(ctx)
			if ticker == nil {
				log.Println("0=1", ticker)
				return nil
			}

			rating, err := fetcher.GetRating(ctx, *ticker)
			if err != nil {
				if err.Error() == "TickerNotExists" {
					r.lock.Lock()
					r.blockedTickers[*ticker] = struct{}{}
					r.lock.Unlock()
				} else {
					return err
				}

			}

			if rating != nil {
				if cache != nil {
					go cache.AddRating(ctx, *rating)
				}
				if dbCache != nil {
					go dbCache.AddRating(ctx, *rating)
				}
			}
		}
	}
}

func (r *Refresher) getNextTickerForUpdate(ctx context.Context) *string {
	r.lock.Lock()
	defer r.lock.Unlock()

	var nextTicker *string
	lowestTime := time.Now()

	for ticker, t := range r.nextRefreshTime {
		if _, exists := r.blockedTickers[ticker]; exists {
			continue
		}

		if t.Before(lowestTime) && t.Before(time.Now()) {
			nextTicker = &ticker
			lowestTime = t
		}
	}

	return nextTicker
}
