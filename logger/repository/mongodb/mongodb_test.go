package repository

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/random"
	"github.com/RoyceAzure/go-stockinfo-logger/shared/util/config"
	"github.com/stretchr/testify/require"

)

func TestInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	config := config.Config{
		MongodbAddress: "mongodb://localhost:27017",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	mongodb, err := ConnectToMongo(ctx, config.MongodbAddress)
	require.NoError(t, err)

	mongoDao := NewMongoDao(mongodb)
	err = mongoDao.Insert(context.Background(), LogEntry{
		ID:          random.RandomString(10),
		ServiceName: "testName",
		Message:     "testjsondata",
		CreatedAt:   time.Now().UTC(),
	})
	require.NoError(t, err)
}

func TestGetAll(t *testing.T) {
	config := config.Config{
		MongodbAddress: "mongodb://localhost:27017",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	mongodb, err := ConnectToMongo(ctx, config.MongodbAddress)
	require.NoError(t, err)

	mongoDao := NewMongoDao(mongodb)
	res, err := mongoDao.GetAll(context.Background())
	require.Greater(t, len(res), 5)
	require.NoError(t, err)
}
