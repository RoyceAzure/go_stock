package repository

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func TestCreateFrontendClient(t *testing.T) {
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

	ip := "127.0.0.1"
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
}
