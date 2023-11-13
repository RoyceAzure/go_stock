package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadDataSVAA(t *testing.T) {
	res, err := testService.DownloadAndInsertDataSVAA(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, res)
}

func TestSyncStock(t *testing.T) {
	res, errs := testService.SyncStock(context.Background())
	require.Len(t, errs, 0)
	require.NotEmpty(t, res)
}
