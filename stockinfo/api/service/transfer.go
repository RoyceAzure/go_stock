package service

import (
	"context"
	"fmt"
	"time"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	pb "github.com/RoyceAzure/go-stockinfo/shared/pb"
	utility "github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
	"github.com/shopspring/decimal"
)

type TransferStockServiceParams struct {
	UserID    int64  `json:"user_id"`
	StockCode string `json:"stock_id"`
	TransType string `json:"trans_type"`
	Amt       int32  `json:"amt"`
}

/*
檢查stock code合法
要找出fundid
這裡要去redis取得 stock current price
固定使用tw fund

這段會使用queue方式執行


這個有可能會在queue裡面重複執行
*/

func (service *TransferService) StockTransfer(ctx context.Context, arg TransferStockServiceParams) error {
	if !utility.IsSupportedTransationType(arg.TransType) {
		return fmt.Errorf("unsupported trans type")
	}

	fund, err := service.store.GetFundByUidandCurForUpdateNoK(ctx, db.GetFundByUidandCurForUpdateNoKParams{
		UserID:       arg.UserID,
		CurrencyType: string(constants.TW),
	})
	if err != nil {
		return fmt.Errorf("user get no fund")
	}

	stock, err := service.store.GetStockByCode(ctx, arg.StockCode)
	if err != nil {
		return fmt.Errorf("invalid stock code")
	}

	sprCache := service.schdulerDao.GetSprData(ctx)
	if sprCache.DataTime == "" {
		return fmt.Errorf("sprCache is empty")
	}

	var targetSPR *pb.StockPriceRealTime

	for _, spr := range sprCache.Data {
		if spr.StockCode == stock.StockCode {
			targetSPR = spr
			break
		}
	}

	if targetSPR == nil {
		return fmt.Errorf("can't fetch %s current price")
	}

	//不需要查看是否失敗，重試情況下再建立一個新的
	//成功要寫到以實現損益  成交回報   stockTrans要記錄交易成功與失敗

	stockTrans, err := service.store.CreateStockTransaction(ctx, db.CreateStockTransactionParams{
		UserID:                  arg.UserID,
		StockID:                 stock.StockID,
		TransactionType:         arg.TransType,
		TransactionDate:         time.Now().UTC(),
		TransationAmt:           arg.Amt,
		TransationPricePerShare: targetSPR.OpenPrice,
	})
	if err != nil {
		return fmt.Errorf("create transation failed")
	}

	perPrice, err := decimal.NewFromString(targetSPR.OpenPrice)
	if err != nil {
		return fmt.Errorf("error  convert price")
	}

	//交易開始時先寫入交易紀錄  最後在更新成功與失敗  要有crdate  update
	_, err = service.store.TransferStockTx(ctx, db.TransferStockTxParams{
		UserID:   arg.UserID,
		StockID:  stock.StockID,
		FundID:   fund.FundID,
		Amt:      arg.Amt,
		PerPrice: perPrice,
	})
	if err != nil {
		service.store.UpdateStockTransationResult(ctx, db.UpdateStockTransationResultParams{
			Result:       false,
			TransationID: stockTrans.TransationID,
		})
		return err
	}

	return nil
}
