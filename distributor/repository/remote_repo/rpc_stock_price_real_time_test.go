package repository

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/stretchr/testify/require"
)

func TestGetStockPriceRealTime(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	config, err := config.LoadConfig("../../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	dao, err := NewJSchdulerInfoDao(config.GrpcSchedulerAddress)
	require.NoError(t, err)
	require.NotEmpty(t, dao)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := dao.GetStockPriceRealTime(ctx)

	require.NoError(t, err)
	require.NotEmpty(t, res)

}
