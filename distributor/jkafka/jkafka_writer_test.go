package jkafka

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/random"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/require"
)

func TestNewJKafkaWriter(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	CreateNewWriter(t)
}

func CreateNewWriter(t *testing.T) KafkaWriter {
	config, err := config.LoadConfig("../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// defer cancel()

	w := NewJKafkaWriter(config.KafkaDistributorAddress)
	require.NotEmpty(t, w)
	return w
}

func TestWriteMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	jw := CreateNewWriter(t)
	topic := "test"
	messages := []kafka.Message{
		{
			Topic: topic,
			Key:   []byte("stock1"),
			Value: []byte(random.RandomString(5)),
		},
		{
			Topic: topic,
			Key:   []byte("stock2"),
			Value: []byte(random.RandomString(5)),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := jw.WriteMessages(ctx, messages)
	require.NoError(t, err)
}
