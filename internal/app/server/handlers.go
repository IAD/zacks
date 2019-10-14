package main

import (
	"github/IAD/zacks/internal/app/server/gen/server/models"
	"github/IAD/zacks/internal/app/server/gen/server/restapi/operations"
	models2 "github/IAD/zacks/internal/pkg/models"
	"github/IAD/zacks/internal/pkg/zacks"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func NewHandlers(z *zacks.Zacks) *Handlers {
	return &Handlers{
		z: z,
	}
}

type Handlers struct {
	z *zacks.Zacks
}

// Handler for GET /{ticker}
func (h *Handlers) GetTickerHandler(
	params *operations.GetTickerParams,
	getTickerOK operations.NewGetTickerOKFunc,
	getTickerNotFound operations.NewGetTickerNotFoundFunc,
	getTickerInternalServerError operations.NewGetTickerInternalServerErrorFunc,
) middleware.Responder {
	rating, err := h.z.GetRating(params.Scope.Ctx, params.Ticker)
	if err != nil {
		return getTickerInternalServerError().WithPayload(&models.Message{
			Code:    500,
			Message: err.Error(),
		})
	}

	if rating == nil {
		return getTickerNotFound().WithPayload(&models.Message{
			Code:    404,
			Message: "not found",
		})
	}

	return getTickerOK().WithPayload(convertRatingToSwagger(rating))
}

// Handler for GET /{ticker}/history
func (h *Handlers) GetTickerHistoryHandler(
	params *operations.GetTickerHistoryParams,
	getTickerHistoryOK operations.NewGetTickerHistoryOKFunc,
	getTickerHistoryNotFound operations.NewGetTickerHistoryNotFoundFunc,
	getTickerHistoryInternalServerError operations.NewGetTickerHistoryInternalServerErrorFunc,
) middleware.Responder {
	ratings, err := h.z.GetHistory(params.Scope.Ctx, params.Ticker)
	if err != nil {
		return getTickerHistoryInternalServerError().WithPayload(&models.Message{
			Code:    500,
			Message: err.Error(),
		})
	}

	return getTickerHistoryOK().WithPayload(convertRatingsToSwagger(ratings))
}

func convertRatingToSwagger(rating *models2.Rating) *models.Rank {
	return &models.Rank{
		Beta:              rating.Beta,
		DateReceived:      strfmt.DateTime(rating.DateReceived),
		DividendAmount:    rating.DividendAmount,
		DividendPercent:   rating.DividendPercent,
		ForwardPe:         rating.ForwardPE,
		Name:              rating.Name,
		PegRatio:          rating.PEGRatio,
		Rank:              rating.Rank,
		RankText:          rating.RankText,
		ScoreGrowth:       rating.ScoreGrowth,
		ScoreGrowthText:   rating.ScoreGrowthText,
		ScoreMomentum:     rating.ScoreMomentum,
		ScoreMomentumText: rating.ScoreMomentumText,
		ScoreValue:        rating.ScoreValue,
		ScoreValueText:    rating.ScoreValueText,
		ScoreVgm:          rating.ScoreVGM,
		ScoreVgmText:      rating.ScoreVGMText,
		Ticker:            rating.Ticker,
	}
}

func convertRatingsToSwagger(ratings []models2.Rating) models.RankCollection {
	result := make(models.RankCollection, 0)

	for _, item := range ratings {
		result = append(result, convertRatingToSwagger(&item))
	}

	return result
}
