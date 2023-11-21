package redisService

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLatestSPR(t *testing.T) {
	sprs, err := testRedisService.GetLatestSPR(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, sprs)
}
