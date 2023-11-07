package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetSDAVGALLs(t *testing.T) {
	startTime := time.Now()
	limit := 20
	page := 1500
	offset := (page - 1) * limit
	data, err := testQueries.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			Limit:  500,
			Offset: int32(offset),
		})
	require.NoError(t, err)
	require.NotEmpty(t, data)
	require.Len(t, data, 500)
	require.WithinDuration(t, time.Now(), startTime, time.Millisecond)
}
