package fetcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	html := testParseData

	data, err := parse(html)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	assert.Equal(t, "Lam Research Corporation: (LRCX)", data.Name)
	assert.Equal(t, "3-Hold", data.RankText)
	assert.Equal(t, int64(3), data.Rank)
	assert.Equal(t, "C", data.ScoreValueText)
	assert.Equal(t, int64(3), data.ScoreValue)
	assert.Equal(t, "B", data.ScoreGrowthText)
	assert.Equal(t, int64(2), data.ScoreGrowth)
	assert.Equal(t, "A", data.ScoreMomentumText)
	assert.Equal(t, int64(1), data.ScoreMomentum)
	assert.Equal(t, "B", data.ScoreVGMText)
	assert.Equal(t, int64(2), data.ScoreVGM)
	assert.Equal(t, float64(4.40), data.DividendAmount)
	assert.Equal(t, float64(2.02), data.DividendPercent)
	assert.Equal(t, float64(1.65), data.Beta)
	assert.Equal(t, float64(16.77), data.ForwardPE)
	assert.Equal(t, float64(1.40), data.PEGRatio)
}

func TestParseEmptyRank(t *testing.T) {
	html := testParseEmptyRankData

	data, err := parse(html)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	assert.Equal(t, "", data.RankText)
	assert.Equal(t, int64(0), data.Rank)
}

func TestNotFound(t *testing.T) {
	html := testNotFoundData

	data, err := parse(html)
	assert.NotNil(t, err)
	assert.Equal(t, "TickerNotExists", err.Error())
	assert.Nil(t, data)
}
