package refresher

import (
	"context"
	"testing"
	"time"

	"github/IAD/zacks/internal/pkg/cache"
	fetcher "github/IAD/zacks/internal/pkg/fetcher"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRefresherWithoutFetcher(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	refresher := NewRefresher(time.Minute * 5)
	err := refresher.Start(ctx, nil, nil, nil)
	require.Error(t, err)
}

func TestRefresherWithCache(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	ticker := "AAPL"

	c := cache.NewCache()
	f := fetcher.NewFetcher(time.Second * 10)

	refresher := NewRefresher(time.Minute * 10)
	go func() {
		err := refresher.Start(ctx, c, nil, f)
		require.Nil(t, err)
	}()

	refresher.Watch(ticker)

	for i := 0; i <= 20; i++ {
		rating, _ := c.GetRating(ctx, ticker)
		if rating != nil {
			assert.Equal(t, ticker, rating.Ticker)
			assert.NotEqual(t, int64(0), rating.Rank)
			return
		}

		time.Sleep(time.Millisecond * 500)
	}

	t.Error("Expected to eventually receive data into the cache")
}
