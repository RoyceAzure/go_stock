package cronworker

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/jkafka"
	sqlc "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	remote_repo "github.com/RoyceAzure/go-stockinfo-distributor/repository/remote_repo"
	"github.com/RoyceAzure/go-stockinfo-distributor/service"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func TestWorker(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	config, err := config.LoadConfig("../") //表示讀取當前資料夾
	require.NoError(t, err)
	schedulerDao, err := remote_repo.NewJSchdulerInfoDao(config.GrpcSchedulerAddress)
	require.NoError(t, err)
	kafkaWriter := jkafka.NewJKafkaWriter(config.KafkaDistributorAddress)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	pgxPool, err := pgxpool.New(ctx, config.DBSource)
	require.NoError(t, err)
	defer pgxPool.Close()

	dbDao := sqlc.NewSQLDistributorDao(pgxPool)

	service := service.NewDistributorService(schedulerDao, dbDao, kafkaWriter)
	cronWorker := NewSchdulerWorker(service, time.Local)
	cronWorker.SetUpSchdulerWorker(context.Background())
	cronWorker.Start()
}
