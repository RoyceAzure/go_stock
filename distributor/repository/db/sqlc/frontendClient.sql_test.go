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

func TestCreateFrontendClient(t *testing.T) {
	CreateFrontendClient(t)
}

func CreateFrontendClient(t *testing.T) FrontendClient {
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

	ip := random.RandomIP()
	region := "tw"
	// CreatedAt := time.Now().UTC()
	res, err := dao.CreateFrontendClient(ctx, CreateFrontendClientParams{
		Ip:     ip,
		Region: region,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, ip, res.Ip)
	require.Equal(t, region, res.Region)
	return res
}

func TestDeleteFrontendClient(t *testing.T) {
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

	fc := CreateFrontendClient(t)

	// CreatedAt := time.Now().UTC()
	err = dao.DeleteFrontendClient(ctx, fc.ClientUid)
	require.NoError(t, err)
}

func TestGetFrontendClientByID(t *testing.T) {
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

	fc := CreateFrontendClient(t)

	// CreatedAt := time.Now().UTC()
	res, err := dao.GetFrontendClientByID(ctx, fc.ClientUid)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, fc.ClientUid, res.ClientUid)
	require.Equal(t, fc.Ip, res.Ip)
	require.Equal(t, fc.Region, res.Region)
	require.Equal(t, fc.CreatedAt, res.CreatedAt)
}

func TestGetFrontendClientByIP(t *testing.T) {
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

	fc := CreateFrontendClient(t)

	// CreatedAt := time.Now().UTC()
	res, err := dao.GetFrontendClientByIP(ctx, fc.Ip)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, fc.ClientUid, res.ClientUid)
	require.Equal(t, fc.Ip, res.Ip)
	require.Equal(t, fc.Region, res.Region)
	require.Equal(t, fc.CreatedAt, res.CreatedAt)
}

func TestGetFrontendClients(t *testing.T) {
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
		CreateFrontendClient(t)
	}
	// CreatedAt := time.Now().UTC()
	arg := GetFrontendClientsParams{
		Limit:  5,
		Offset: 0,
	}

	res, err := dao.GetFrontendClients(ctx, arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 5)
}
