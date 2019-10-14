package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRating_Equals(t *testing.T) {
	r := &Rating{
		Ticker:            "A",
		Name:              "NameA",
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

	or := &Rating{
		Ticker:            "A",
		Name:              "NameA",
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
		DateReceived:      time.Now().Add(time.Minute),
	}

	assert.True(t, r.Equals(*or))
}

func TestRating_EqualsFalse(t *testing.T) {
	r := &Rating{
		Ticker:            "A",
		Name:              "NameA",
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

	or := &Rating{
		Ticker:            "A",
		Name:              "NameA",
		Rank:              1,
		RankText:          "B", // different value
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
		DateReceived:      time.Now().Add(time.Minute),
	}

	assert.False(t, r.Equals(*or))
}
