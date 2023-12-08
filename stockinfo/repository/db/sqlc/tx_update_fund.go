package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
)

/*
AfterCreate is call back, use it err to check commit or rollback
*/
type UpdateFundTxParams struct {
	UserID int64
	UPN    string
	Amount string
}

type UpdateFundTxResults struct {
	Fund Fund
}

/*
如果遇到db高流量情況，trsation會delay, 所以async操作需要設置delay
*/
func (store *SQLStore) UpdateFundTx(ctx context.Context, arg UpdateFundTxParams) (UpdateFundTxResults, error) {
	var result UpdateFundTxResults

	//
	//後續 這個 func(q *Queries)  應該就使已經寫好且沒有Tx的版本
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		fund, err := q.GetFundByUidandCurForUpdateNoK(ctx, GetFundByUidandCurForUpdateNoKParams{
			UserID:       arg.UserID,
			CurrencyType: string(constants.TW),
		})
		if err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("use has no fund in tw")
			}
			return fmt.Errorf("get fund return err")
		}

		updateFund, err := q.UpdateFund(ctx, UpdateFundParams{
			FundID:  fund.FundID,
			Balance: arg.Amount,
			UpDate:  util.TimeToSqlNiTime(time.Now().UTC()),
			UpUser:  util.StringToSqlNiStr(arg.UPN),
		})
		if err != nil {
			return fmt.Errorf("update found err")
		}
		result.Fund = updateFund
		return nil
	})

	return result, err
}
