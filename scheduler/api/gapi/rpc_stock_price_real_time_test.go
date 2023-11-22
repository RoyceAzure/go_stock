package gapi

import (
	"context"
	"testing"

	"github.com/RoyceAzure/go-stockinfo-schduler/api/pb"
	jredis "github.com/RoyceAzure/go-stockinfo-schduler/repository/redis"
	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	redisService "github.com/RoyceAzure/go-stockinfo-schduler/service/redisService"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func TestGetStockPriceRealTime(t *testing.T) {
	config, err := config.LoadConfig("../../")
	require.NoError(t, err)

	ctx := context.Background()
	pgxpool, err := pgxpool.New(ctx, config.DBSource)
	require.NoError(t, err)

	dao := repository.NewSQLDao(pgxpool)

	jr := jredis.NewJredis(config)

	_ = redisService.NewJRedisService(jr)

	server := newTestServer(t, config, dao, nil, nil)

	res, err := server.GetStockPriceRealTime(ctx, &pb.StockPriceRealTimeRequest{})
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Greater(t, len(res.Result), 0)
}
