package worker

import (
	"context"
	"encoding/json"
	"fmt"

	logger "github.com/RoyceAzure/go-stockinfo/repository/logger_distributor"
	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/hibiken/asynq"
)

/*
製作task  這裡把所有資訊都包進task裡面  包括retry delay 甚至是要使用甚麼優先級的queue

使用client enqueue task
*/
func (distributor *RedisTaskDistributor) DistributeTaskStockTransation(
	ctx context.Context,
	payload *PayloadTransation,
	opts ...asynq.Option,
) error {
	md := util.ExtractMetaData(ctx)
	jsonData, err := json.Marshal(payload)
	if err != nil {
		logger.Logger.Error().
			Err(err).
			Any("meta", md).
			Msg("enqueued task failed")
		return fmt.Errorf("failed to marshal task payload %w", err)
	}

	//沒有提到要如何執行task , 只有給payload跟設定
	task := asynq.NewTask(TaskStockTransation, jsonData, opts...)
	taskInfo, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		logger.Logger.Error().
			Err(err).
			Str("type", task.Type()).
			Any("meta", md).
			Bytes("body", task.Payload()).
			Str("queue", taskInfo.Queue).
			Int("max_retry", taskInfo.MaxRetry).
			Msg("enqueued task")
		return fmt.Errorf("failed to enqueue task %w", err)
	}

	logger.Logger.Info().
		Str("type", task.Type()).
		Any("meta", md).
		Bytes("body", task.Payload()).
		Str("queue", taskInfo.Queue).
		Int("max_retry", taskInfo.MaxRetry).
		Msg("enqueued task")
	return nil
}
