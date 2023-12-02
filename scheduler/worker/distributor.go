package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskBatchUpdateStock(
		ctx context.Context,
		payload []BatchUpdateStockPayload,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

/*
asynq.NewClient 產生的client可以共用  支持異步
*/
func NewRedisTaskDistributor(client *asynq.Client) TaskDistributor {
	return &RedisTaskDistributor{
		client: client,
	}
}
