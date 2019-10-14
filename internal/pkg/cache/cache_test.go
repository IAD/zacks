package cache

import (
	"context"
	"testing"
	"time"

	"github/IAD/zacks/internal/pkg/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRating(t *testing.T) {
	cache := NewCache()
	ctx := context.Background()

	rating, err := cache.GetRating(ctx, "AAPL")
	assert.Nil(t, err)
	assert.Nil(t, rating)

	added := models.Rating{
		Ticker:            "AAPL",
		Name:              "AAPL Name",
		Rank:              1,
		RankText:          "A",
		ScoreValue:        2,
		ScoreValueText:    "B",
		ScoreGrowth:       3,
		ScoreGrowthText:   "C",
		ScoreMomentum:     4,
		ScoreMomentumText: "D",
		ScoreVGM:          5,
		ScoreVGMText:      "E",
		DividendAmount:    1.11,
		DividendPercent:   20,
		Beta:              1.44,
		ForwardPE:         6.77,
		PEGRatio:          2.33,
		DateReceived:      time.Now(),
	}

	err = cache.AddRating(ctx, added)
	assert.Nil(t, err)

	rating, err = cache.GetRating(ctx, "AAPL")
	assert.Nil(t, err)
	require.NotNil(t, rating)

	assert.Equal(t, added, *rating)
}

func TestGetHistory(t *testing.T) {
	cache := NewCache()
	ctx := context.Background()

	rating, err := cache.GetRating(ctx, "AAPL")
	assert.Nil(t, err)
	assert.Nil(t, rating)

	added1 := models.Rating{
		Ticker:            "AAPL",
		Name:              "AAPL Name",
		Rank:              1,
		RankText:          "A",
		ScoreValue:        2,
		ScoreValueText:    "B",
		ScoreGrowth:       3,
		ScoreGrowthText:   "C",
		ScoreMomentum:     4,
		ScoreMomentumText: "D",
		ScoreVGM:          5,
		ScoreVGMText:      "E",
		DividendAmount:    1.11,
		DividendPercent:   20,
		Beta:              1.44,
		ForwardPE:         6.77,
		PEGRatio:          2.33,
		DateReceived:      time.Now(),
	}

	err = cache.AddRating(ctx, added1)
	assert.Nil(t, err)

	added2 := models.Rating{
		Ticker:            "AAPL",
		Name:              "AAPL Name",
		Rank:              1,
		RankText:          "A",
		ScoreValue:        3,   // different
		ScoreValueText:    "C", // different
		ScoreGrowth:       3,
		ScoreGrowthText:   "C",
		ScoreMomentum:     4,
		ScoreMomentumText: "D",
		ScoreVGM:          5,
		ScoreVGMText:      "E",
		DividendAmount:    1.11,
		DividendPercent:   20,
		Beta:              1.44,
		ForwardPE:         6.77,
		PEGRatio:          2.33,
		DateReceived:      time.Now(),
	}

	err = cache.AddRating(ctx, added2)
	assert.Nil(t, err)

	ratings, err := cache.GetHistory(ctx, "AAPL")
	assert.Nil(t, err)
	require.NotNil(t, ratings)
	require.Len(t, ratings, 2)

	assert.Equal(t, []models.Rating{added1, added2}, ratings)
}
