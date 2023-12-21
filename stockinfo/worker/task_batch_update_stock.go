package worker

import (
	"context"
	"encoding/json"
	"fmt"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	logger "github.com/RoyceAzure/go-stockinfo/repository/logger_distributor"
	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/hibiken/asynq"
)

const TaskBatchUpdateStock = "task:batch_update_stock"

type BatchUpdateStockPayload struct {
	StockCode    string `json:"stock_code"`
	StockName    string `json:"stock_name"`
	CurrentPrice string `json:"current_price"`
	MarketCap    int64  `json:"market_cap"`
}

/*
TODO : 失敗的資料要寫入errorList , mongoDB??
*/
func (processer RedisTaskProcessor) ProcessTaskBatchUpdateStock(ctx context.Context, task *asynq.Task) error {
	var payloads []BatchUpdateStockPayload
	md := util.ExtractMetaData(ctx)
	err := json.Unmarshal(task.Payload(), &payloads)
	if err != nil {
		return fmt.Errorf("failed to unmarshal task payload %w", asynq.SkipRetry)
	}
	//no transation
	var successCouunt int
	for _, payload := range payloads {
		successData, err := processer.store.UpdateStockCPByCode(ctx, db.UpdateStockCPByCodeParams{
			CurrentPrice: util.StringToSqlNiStr(payload.CurrentPrice),
			StockCode:    payload.StockCode,
		})
		if err != nil {
			logger.Logger.Warn().
				Err(err).
				Str("type", task.Type()).
				Any("meta", md).
				Msg("task batch update stock get err")
		} else {
			logger.Logger.Info().Str("type", task.Type()).
				Bytes("task payload", task.Payload()).
				Any("meta", md).
				Str("stock code", successData.StockCode).
				Msg("processed task successed")
		}
	}
	logger.Logger.Info().Str("type", task.Type()).
		Any("meta", md).
		Int("success rows", successCouunt).
		Msg("processed task end")
	return err
}
