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
