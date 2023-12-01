package repository

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-logger/shared/util/config"
	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	config, err := config.LoadConfig("../../") //表示讀取當前資料夾
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	mongodb, err := ConnectToMongo(ctx, config.MongodbAddress)
	require.NoError(t, err)

	mongoDao := NewMongoDao(mongodb)
	err = mongoDao.Insert(context.Background(), LogEntry{
		ID:          "test",
		ServiceName: "testName",
		Message:     "testjsondata",
		CreatedAt:   time.Now().UTC(),
	})
	require.NoError(t, err)
}

func TestGetAll(t *testing.T) {
	config, err := config.LoadConfig("../../") //表示讀取當前資料夾
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	mongodb, err := ConnectToMongo(ctx, config.MongodbAddress)
	require.NoError(t, err)

	mongoDao := NewMongoDao(mongodb)
	res, err := mongoDao.GetAll(context.Background())
	require.Greater(t, len(res), 5)
	require.NoError(t, err)
}
