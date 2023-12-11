package worker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/RoyceAzure/go-stockinfo/service"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskStockTransation = "task:stock_transation"

type PayloadTransation struct {
	UserID    int64  `json:"user_id"`
	StockCode string `json:"stock_id"`
	TransType string `json:"trans_type"`
	Amt       int32  `json:"amt"`
	Operator  string `json:"operator"`
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
		UserID:    payload.UserID,
		StockCode: payload.StockCode,
		TransType: payload.TransType,
		Amt:       payload.Amt,
		Operator:  payload.Operator,
	}

	err := processer.stockTranService.StockTransfer(ctx, arg)
	if err != nil {
		if errors.Is(err, constants.ErrInValidatePreConditionOp) || errors.Is(err, constants.ErrInvalidArgument) {
			log.Warn().Err(err).Msg("task stock transfer get err")
			return fmt.Errorf("failed to process task stock transfer %w", asynq.SkipRetry)
		}
		log.Error().Err(err).Msg("task stock transfer get internal err")
		return fmt.Errorf("failed to process task stock transfer")
	}
	log.Info().Str("type", task.Type()).
		Str("user_id", fmt.Sprintf("%s", arg.UserID)).
		Str("stock_code", string(arg.StockCode)).
		Str("transation_type", arg.TransType).
		Int64("amt", int64(arg.Amt)).
		Msg("processed task successed")
	return nil
}
