package worker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	logger "github.com/RoyceAzure/go-stockinfo/repository/logger_distributor"
	"github.com/RoyceAzure/go-stockinfo/service"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
	"github.com/hibiken/asynq"
)

const TaskStockTransation = "task:stock_transation"

type PayloadTransation struct {
	TransationID int64  `json:"trans_id"`
	Operator     string `json:"operator"`
}

/*
製作task  這裡把所有資訊都包進task裡面  包括retry delay 甚至是要使用甚麼優先級的queue

使用client enqueue task
*/
func (processer *RedisTaskProcessor) ProcessTaskStockTransfer(
	ctx context.Context,
	task *asynq.Task,
) error {
	var payload PayloadTransation
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal task payload %w", asynq.SkipRetry)
	}

	arg := service.TransferStockServiceParams{
		TransationID: payload.TransationID,
		Operator:     payload.Operator,
	}

	err := processer.stockTranService.StockTransfer(ctx, arg)
	if err != nil {
		if errors.Is(err, constants.ErrInValidatePreConditionOp) || errors.Is(err, constants.ErrInvalidArgument) {
			logger.Logger.Warn().Err(err).Msg("task stock transfer get err")
			return fmt.Errorf("failed to process task stock transfer %w", asynq.SkipRetry)
		}
		logger.Logger.Error().Err(err).Msg("task stock transfer get internal err")
		return fmt.Errorf("failed to process task stock transfer")
	}
	logger.Logger.Info().Str("type", task.Type()).
		Int64("transation_id", payload.TransationID).
		Msg("processed task successed")
	return nil
}
