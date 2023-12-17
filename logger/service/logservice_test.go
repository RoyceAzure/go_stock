package logservice

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	repository "github.com/RoyceAzure/go-stockinfo-logger/repository/mongodb"
	"github.com/RoyceAzure/go-stockinfo-logger/shared/util/config"
	"github.com/stretchr/testify/require"
)

func TestWrite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	config, err := config.LoadConfig("../") //表示讀取當前資料夾
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	mongodb, err := repository.ConnectToMongo(ctx, config.MongodbAddress)
	require.NoError(t, err)

	mongoDao := repository.NewMongoDao(mongodb)
	mongoLogger := NewMongoLogger(mongoDao)

	err = SetUpMutiMongoLogger(mongoLogger, config.ServiceID)
	require.NoError(t, err)
	// Write a log.
	Logger.Info().Str("key1", "valu1").Err(errors.New("err")).Msg("This is a log message that will be written to MongoDB")
	Logger.Warn().Msg("This is a log message that will be written to MongoDB")
	Logger.Trace().Msg("This is a log message that will be written to MongoDB")
	Logger.Error().Str("key1", "valu1").Int("int1", 50).Err(fmt.Errorf("test err")).Msg("test logger write")
}
