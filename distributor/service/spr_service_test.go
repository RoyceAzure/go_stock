package service

import (
	"context"
	"testing"
	"time"

	jkafka "github.com/RoyceAzure/go-stockinfo-distributor/jkafka"
	sqlc "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	remote_repo "github.com/RoyceAzure/go-stockinfo-distributor/repository/remote_repo"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func TestGetFilterSPRByIP(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
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

	s := NewDistributorService(remoteDao, sqlDao, nil)

	res, err := s.GetFilterSPRByIP(ctx, "127.0.0.1")
	require.NoError(t, err)
	require.NotEmpty(t, res)
}

func TestGetFilterSPRByIPAndSendToKa(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
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

	kwriter := jkafka.NewJKafkaWriter(config.KafkaDistributorAddress)

	s := NewDistributorService(remoteDao, sqlDao, kwriter)

	err = s.GetAllRegisStockAndSendToKa(context.Background())
	require.NoError(t, err)
}
