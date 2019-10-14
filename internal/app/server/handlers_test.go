package main

import (
	"context"
	"testing"

	"github/IAD/zacks/internal/app/server/gen/client/zacksclient"
	"github/IAD/zacks/internal/app/server/gen/client/zacksclient/operations"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlers_GetTickerHandler(t *testing.T) {
	client := zacksclient.NewClientWithBasePath("localhost:8080", "")
	ok, not, ise, err := client.Operations.GetTicker(&operations.GetTickerParams{
		Context: context.Background(),
		Ticker:  "M",
	})
	assert.Nil(t, err)
	assert.Nil(t, ise)
	assert.Nil(t, not)

	require.NotNil(t, ok)
	assert.NotNil(t, ok.Payload)
	assert.Equal(t, "M", ok.Payload.Ticker)
}

func TestHandlers_GetTickerHistoryHandler(t *testing.T) {
	client := zacksclient.NewClientWithBasePath("localhost:8080", "")

	{
		//warm up cache
		ok, not, ise, err := client.Operations.GetTicker(&operations.GetTickerParams{
			Context: context.Background(),
			Ticker:  "M",
		})
		assert.Nil(t, err)
		assert.Nil(t, ise)
		assert.Nil(t, not)
		require.NotNil(t, ok)
	}

	ok, not, ise, err := client.Operations.GetTickerHistory(&operations.GetTickerHistoryParams{
		Context: context.Background(),
		Ticker:  "M",
	})
	assert.Nil(t, err)
	assert.Nil(t, ise)
	assert.Nil(t, not)

	require.NotNil(t, ok)
	assert.NotNil(t, ok.Payload)
	require.True(t, len(ok.Payload) > 0)
	assert.Equal(t, "M", ok.Payload[0].Ticker)
}
