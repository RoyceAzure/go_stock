package service

import (
	"context"
	"testing"
	"time"

	sqlc "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	remote_repo "github.com/RoyceAzure/go-stockinfo-distributor/repository/remote_repo"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func TestGetFilterSPRByIP(t *testing.T) {
	config, err := config.LoadConfig("../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := pgxpool.New(ctx, config.DBSource)
	require.NoError(t, err)
	require.NotEmpty(t, config)

	sqlDao := sqlc.NewSQLDistributorDao(conn)
	remoteDao, err := remote_repo.NewJSchdulerInfoDao(config.GrpcSchedulerAddress)
	require.NoError(t, err)
	require.NotEmpty(t, sqlDao)

	s := NewDistributorService(remoteDao, sqlDao)

	res, err := s.GetFilterSPRByIP(ctx, "127.0.0.1")
	require.NoError(t, err)
	require.NotEmpty(t, res)
}
