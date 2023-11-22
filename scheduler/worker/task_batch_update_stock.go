package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	QueueCritical        = "critical"
	QueueDefault         = "default"
	TaskBatchUpdateStock = "task:batch_update_stock"
)

type BatchUpdateStockPayload struct {
	StockCode    string `json:"stock_code"`
	StockName    string `json:"stock_name"`
	CurrentPrice string `json:"current_price"`
	MarketCap    int64  `json:"market_cap"`
}

func (distributor *RedisTaskDistributor) DistributeTaskBatchUpdateStock(
	ctx context.Context,
	payload []BatchUpdateStockPayload,
	opts ...asynq.Option,
) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Errorf("failed to marshal task payload %w", err)
	}

	task := asynq.NewTask(TaskBatchUpdateStock, jsonData, opts...)
	taskInfo, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task %w", err)
	}

	log.Info().
		Str("type", task.Type()).
		Bytes("body", task.Payload()).
		Str("queue", taskInfo.Queue).
		Int("max_retry", taskInfo.MaxRetry).
		Msg("enqueued task")
	return nil
}