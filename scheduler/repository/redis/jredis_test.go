package jredis

import (
	"context"
	"os"
	"os/signal"
	"testing"

	"github.com/RoyceAzure/go-stockinfo-scheduler/util/config"
	"github.com/stretchr/testify/require"
)

func TestRedis(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	config, err := config.LoadConfig("../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	jredis := NewJredis(config)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err = jredis.Start(ctx)
	require.NoError(t, err)
}
