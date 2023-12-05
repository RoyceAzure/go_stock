package logservice

import (
	"context"
	"errors"
	"testing"
	"time"

	repository "github.com/RoyceAzure/go-stockinfo-logger/repository/mongodb"
	"github.com/RoyceAzure/go-stockinfo-logger/shared/util/config"
	"github.com/rs/zerolog"
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

	logger := zerolog.New(mongoLogger).With().Timestamp().Logger()

	// Write a log.
	logger.Info().Str("key1", "valu1").Err(errors.New("err")).Msg("This is a log message that will be written to MongoDB")
	logger.Warn().Msg("This is a log message that will be written to MongoDB")
	logger.Trace().Msg("This is a log message that will be written to MongoDB")
}
