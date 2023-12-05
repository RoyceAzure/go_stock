package jkafka

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/require"
)

func TestCreateNewReader(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	CreateNewReader(t)
}

func CreateNewReader(t *testing.T) KafkaReader {
	config, err := config.LoadConfig("../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// defer cancel()
	// group_id := "2"
	topic := "test"
	r := NewJKafkaReader(&kafka.ReaderConfig{
		Brokers: []string{config.KafkaDistributorAddress},
		// GroupID: group_id,
		// 使用 FirstOffset 使得消费者从每个分区的最早消息开始读取
		// 使用 LastOffset 使得消费者从每个分区的最新消息开始读取
		// 如果需要从特定偏移量开始读取，可以直接设置 Offset 字段
		StartOffset: kafka.LastOffset, // 或 kafka.LastOffset 或者具体的偏移量数字
		Topic:       topic,
	})
	require.NotEmpty(t, r)
	return r
}

func TestReadMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	jr := CreateNewReader(t)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var resultMsgs []string
	var resultErrs []error
	msgch := make(chan kafka.Message)
	errch := make(chan error)
	var wg sync.WaitGroup
	wg.Add(3)
	go jr.ReadMessageAsync(ctx, msgch, errch, &wg)

	go func() {
		defer wg.Done()
		for msg := range msgch {
			val := string(msg.Value)
			resultMsgs = append(resultMsgs, val)
			fmt.Println(val)
		}
	}()

	go func() {
		defer wg.Done()
		for er := range errch {
			resultErrs = append(resultErrs, er)
		}
	}()

	wg.Wait()
	require.NotEmpty(t, resultMsgs)
	require.Empty(t, resultErrs)
}
