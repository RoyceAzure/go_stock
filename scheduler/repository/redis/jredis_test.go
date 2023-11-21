package jredis

import (
	"context"
	"os"
	"os/signal"
	"testing"

	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/stretchr/testify/require"
)

func TestRedis(t *testing.T) {
	config, err := config.LoadConfig("../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	jredis := NewJredis(config)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err = jredis.Start(ctx)
	require.NoError(t, err)
}
