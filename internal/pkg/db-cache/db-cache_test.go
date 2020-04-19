package dbcache

import (
	"context"
	"testing"
	"time"

	"github/IAD/zacks/internal/pkg/models"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestDBCache_AddRating(t *testing.T) {
	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	require.Nil(t, err)

	mongoDB := mongoClient.Database("zacks-test")
	dbCache := NewDBCache(mongoDB)

	rating := models.Rating{
		Ticker:            "Test",
		Name:              "T",
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
		DividendPercent:   2.22,
		Beta:              3.33,
		ForwardPE:         4.44,
		PEGRatio:          5.55,
		DateReceived:      time.Now(),
	}
	err = dbCache.AddRating(ctx, rating)
	require.Nil(t, err)
}
