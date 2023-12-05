package jkafka

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/require"
)

func TestNewJKafka(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	config := config.Config{
		KafkaDistributorAddress: "127.0.0.1:29092",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	topic := "test"
	partition := 0

	conn, err := kafka.DialLeader(ctx, "tcp", config.KafkaDistributorAddress, topic, partition)
	require.NoError(t, err)
	require.NotEmpty(t, conn)

	jk := NewJKafka(conn)
	require.NotEmpty(t, jk)
}
