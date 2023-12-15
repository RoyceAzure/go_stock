package service

import (
	"context"
	"database/sql"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
)

type TransferStockServiceParams struct {
	TransationID int64  `json:"trans_id"`
	Operator     string `json:"operator"`
}

/*
檢查stock code合法
要找出fundid
這裡要去redis取得 stock current price
固定使用tw fund

這段會使用queue方式執行


檢查tw fund,  stock  spr，塞選並找到所需spr  建立交易紀錄

這個有可能會在queue裡面重複執行
*/

func (service *TransferService) StockTransfer(ctx context.Context, arg TransferStockServiceParams) error {

	//交易開始時先寫入交易紀錄  最後在更新成功與失敗  要有crdate  update
	_, err := service.store.TransferStockTx(ctx, db.TransferStockTxParams{
		TransationID: arg.TransationID,
		CreateUser:   arg.Operator,
	})
	var msg sql.NullString
	msg.Valid = false
	var result db.TransationResult
	if err != nil {
		msg.String = err.Error()
		msg.Valid = true
		result = db.TransationResultFailed
	} else {
		result = db.TransationResultSuccessed
	}
	service.store.UpdateStockTransationResult(ctx, db.UpdateStockTransationResultParams{
		Result:       result,
		TransationID: arg.TransationID,
		Msg:          msg,
	})
	return err
}
