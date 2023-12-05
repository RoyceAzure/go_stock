package repoisitory

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

const (
	TaskWriteLog = "task:write_log"
	LogQueue     = "logQueue"
)

type TaskDistributor interface {
	Write(p []byte) (n int, err error)
}

type LoggerDistributor struct {
	client *asynq.Client
}

func NewLoggerDistributor(client *asynq.Client) TaskDistributor {
	//asynq.NewClient() 這裡就有提for redis
	return &LoggerDistributor{
		client: client,
	}
}

func SetUpLoggerDistributor(logger TaskDistributor) error {
	if logger == nil {
		return fmt.Errorf("logger distributor is not init")
	}
	multiLogger := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, logger)
	Logger = zerolog.New(multiLogger).With().Str("service_name", config.AppConfig.ServiceID).Timestamp().Logger()
	return nil
}

func (LoggerDistributor LoggerDistributor) Write(p []byte) (n int, err error) {
	// Insert the record into the collection.

	if LoggerDistributor.client == nil {
		return 0, fmt.Errorf("logger distributor is not init")
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(LogQueue),
	}
	task := asynq.NewTask(TaskWriteLog, p, opts...)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err = LoggerDistributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return 0, fmt.Errorf("failed to enqueue task %w", err)
	}

	return len(p), nil
}
