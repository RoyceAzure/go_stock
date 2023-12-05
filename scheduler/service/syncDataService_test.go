package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

)

func TestDownloadDataSVAA(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	res, err := testService.DownloadAndInsertDataSVAA(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, res)
}

func TestSyncStock(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	res, errs := testService.SyncStock(context.Background())
	require.Len(t, errs, 0)
	require.NotEmpty(t, res)
}

func TestSyncStockPriceRealTime(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	res, errs := testService.SyncStockPriceRealTime(context.Background())
	require.Len(t, errs, 0)
	require.Greater(t, res, int64(0))
	require.NotEmpty(t, res)
}

func TestRedisSyncStockPriceRealTime(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	errs := testService.RedisSyncStockPriceRealTime(context.Background())
	require.Len(t, errs, 0)
}
