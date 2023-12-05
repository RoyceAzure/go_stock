package repository

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/random"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func TestCreateClientRegister(t *testing.T) {
	CreateClientRegister(t)
}

func CreateClientRegister(t *testing.T) ClientRegister {
	config, err := config.LoadConfig("../../../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := pgxpool.New(ctx, config.DBSource)
	require.NoError(t, err)
	require.NotEmpty(t, conn)

	dao := NewSQLDistributorDao(conn)
	require.NotEmpty(t, dao)

	stockCode := random.RandomString(5)
	res, err := dao.CreateClientRegister(ctx, CreateClientRegisterParams{
		ClientUid: testClient.ClientUid,
		StockCode: stockCode,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, testClient.ClientUid, res.ClientUid)
	require.Equal(t, stockCode, res.StockCode)
	return res
}

func TestGetClientRegisterByClientUID(t *testing.T) {
	config, err := config.LoadConfig("../../../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := pgxpool.New(ctx, config.DBSource)
	require.NoError(t, err)
	require.NotEmpty(t, conn)

	dao := NewSQLDistributorDao(conn)
	require.NotEmpty(t, dao)

	cr := CreateClientRegister(t)

	res, err := dao.GetClientRegisterByClientUID(ctx, cr.ClientUid)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, item := range res {
		require.Equal(t, cr.ClientUid, item.ClientUid)
	}
}

func TestClientRegisters(t *testing.T) {
	config, err := config.LoadConfig("../../../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := pgxpool.New(ctx, config.DBSource)
	require.NoError(t, err)
	require.NotEmpty(t, conn)

	dao := NewSQLDistributorDao(conn)
	require.NotEmpty(t, dao)

	cr := CreateClientRegister(t)

	err = dao.DeleteClientRegister(ctx, DeleteClientRegisterParams{
		ClientUid: cr.ClientUid,
		StockCode: cr.StockCode,
	})
	require.NoError(t, err)
}

func TestGetClientRegisters(t *testing.T) {
	config, err := config.LoadConfig("../../../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := pgxpool.New(ctx, config.DBSource)
	require.NoError(t, err)
	require.NotEmpty(t, conn)

	dao := NewSQLDistributorDao(conn)
	require.NotEmpty(t, dao)

	for i := 0; i < 5; i++ {
		CreateClientRegister(t)
	}

	arg := GetClientRegistersParams{
		Limit:  5,
		Offset: 0,
	}

	crs, err := dao.GetClientRegisters(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, crs, 5)

	for _, cr := range crs {
		require.NotEmpty(t, cr)
	}
}
