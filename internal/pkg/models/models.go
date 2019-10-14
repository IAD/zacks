package models

import (
	"reflect"
	"time"
)

type Rating struct {
	Ticker            string    `json:"ticker"`
	Name              string    `json:"name"`
	Rank              int64     `json:"rank"`
	RankText          string    `json:"rank_text"`
	ScoreValue        int64     `json:"score_value"`
	ScoreValueText    string    `json:"score_value_text"`
	ScoreGrowth       int64     `json:"score_growth"`
	ScoreGrowthText   string    `json:"score_growth_text"`
	ScoreMomentum     int64     `json:"score_momentum"`
	ScoreMomentumText string    `json:"score_momentum_text"`
	ScoreVGM          int64     `json:"score_vgm"`
	ScoreVGMText      string    `json:"score_vgm_text"`
	DividendAmount    float64   `json:"dividend_amount"`
	DividendPercent   float64   `json:"dividend_percent"`
	Beta              float64   `json:"beta"`
	ForwardPE         float64   `json:"forward_pe"`
	PEGRatio          float64   `json:"peg_ratio"`
	DateReceived      time.Time `json:"date_received"`
}

func (r Rating) Equals(or Rating) bool {
	r.DateReceived = or.DateReceived
	return reflect.DeepEqual(r, or)
}
