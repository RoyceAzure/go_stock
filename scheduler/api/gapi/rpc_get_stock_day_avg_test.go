package gapi

import (
	"context"
	"testing"

	"github.com/RoyceAzure/go-stockinfo-schduler/api/pb"
	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"

)

/*
交互測試
TODO : 使用mock
*/
func TestGetStockDayAvg(t *testing.T) {
	config, err := config.LoadConfig("../../")
	require.NoError(t, err)

	ctx := context.Background()
	pgxpool, err := pgxpool.New(ctx, config.DBSource)
	require.NoError(t, err)

	dao := repository.NewSQLDao(pgxpool)

	server := newTestServer(t, config, dao, nil)

	res, err := server.GetStockDayAvg(ctx, &pb.StockDayAvgRequest{})
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Greater(t, len(res.Result), 0)
}
