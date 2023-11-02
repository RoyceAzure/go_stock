package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/shopspring/decimal"
)

/*
不模擬賣家買家搓合的交易
指模擬單方的買賣

一次買賣一個股票

為何要json tag??

在進入Repo, 轉成entity前  都應該用可計算的type形式保存  以便AP做運算
*/
type TransferStockTxParams struct {
	UserID          int64           `json:"user_id"`
	StockID         int64           `json:"stock_id"`
	FundID          int64           `json:"fund_id"`
	TransType       string          `json:"trans_type"`
	Amt             int32           `json:"amt"`
	PerPrice        decimal.Decimal `json:"per_price"`
	TransactionDate time.Time       `json:"transaction_date"`
	CreateUser      string          `json:"cr_uesr"`
}

type TransferStockTxResults struct {
	Fund          Fund             `json:"fund"`
	OriFund       Fund             `json:"ori_fund"`
	StockTrans    StockTransaction `json:"stock_trans"`
	UserStocks    *UserStock       `json:"user_stocks"`
	OriUserStocks *UserStock       `json:"ori_user_stocks"`
	User          User             `json:"users"`
}

// 個別克制化使用trans 的 func
// 做參數與enity轉換  定義call back  執行exec
func (store *SQLStore) TransferStockTx(ctx context.Context, arg TransferStockTxParams) (TransferStockTxResults, error) {
	var result TransferStockTxResults

	//
	//後續 這個 func(q *Queries)  應該就使已經寫好且沒有Tx的版本
	store.execTx(ctx, func(q *Queries) error {
		//開始寫Trans步驟
		var err error
		isHasUserStock := true
		if !strings.EqualFold(arg.TransType, "buy") && !strings.EqualFold(arg.TransType, "sell") {
			return errors.New("TransType is invalidate")
		}

		//不用鎖stock  雖然sotck價格時時變動  但是這裡是成交後  把當下價格傳遞近來  所以不需要擔心stock 價格變動問題
		// stock, err := q.GetStock(ctx, arg.StockID)
		// if err != nil {
		// 	return err
		// }

		//select for update no key, 受引用關聯的表仍可以做操作
		//由於目前的測試  所有的平行測試都是按照同樣順序  先取fund在取  userstock, 只要第一個人取得fund  後續其他人就會卡在這
		//有就不會有其他人先取得userstock，導致deadlock情況發生
		oriFund, err := q.GetfundByUidandFidForUpdateNoK(ctx, GetfundByUidandFidForUpdateNoKParams{
			UserID: arg.UserID,
			FundID: arg.FundID,
		})
		if err != nil {
			return err
		}
		//select for update no key, 受引用關聯的表仍可以做操作
		oriUserStock, err := q.GetserStockByUidandSidForUpdateNoK(ctx, GetserStockByUidandSidForUpdateNoKParams{
			UserID:  arg.UserID,
			StockID: arg.StockID,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				isHasUserStock = false
				oriUserStock = UserStock{}
			} else {
				return err
			}
		}

		//先固定抓取第一個Fund
		//後續操作應該都要針對一種Fund
		D_priceToHandle := arg.PerPrice.Mul(decimal.NewFromInt32(arg.Amt))
		if err != nil {
			return err
		}

		D_amt := decimal.NewFromInt32(arg.Amt)
		D_ori_balance, err := decimal.NewFromString(oriFund.Balance)
		if err != nil {
			return err
		}

		if strings.EqualFold(arg.TransType, "sell") && oriUserStock.Quantity < arg.Amt {
			return errors.New("Not Enough stock")
		} else if strings.EqualFold(arg.TransType, "buy") && D_ori_balance.LessThan(D_priceToHandle) {
			return errors.New("Not Enough money")
		}

		//會對參照的表進行鎖
		result.StockTrans, err = q.CreateStockTransaction(ctx, CreateStockTransactionParams{
			UserID:                  arg.UserID,
			StockID:                 arg.FundID,
			TransactionType:         arg.TransType,
			TransactionDate:         arg.TransactionDate,
			TransationAmt:           arg.Amt,
			TransationPricePerShare: arg.PerPrice.String(),
			CrUser:                  "royce",
		})

		if err != nil {
			return err
		}
		var newUserStock UserStock
		var new_balance, oriTotlaStcokCost decimal.Decimal
		var new_user_stock_quantity int32
		//更新操做
		if strings.EqualFold(arg.TransType, "sell") {
			//fund
			new_balance = D_ori_balance.Add(D_priceToHandle)

			//userStoc
			new_user_stock_quantity = oriUserStock.Quantity - arg.Amt

			if new_user_stock_quantity == 0 {
				//刪除操作
				err := q.DeleteUserStock(ctx, oriUserStock.UserStockID)
				if err != nil {
					return err
				}
			} else {
				//賣出情況的平均持有成本不變
				newUserStock, err = q.UpdateUserStock(ctx, UpdateUserStockParams{
					UserID:                arg.UserID,
					StockID:               arg.FundID,
					Quantity:              new_user_stock_quantity,
					PurchasePricePerShare: oriUserStock.PurchasePricePerShare,
					PurchasedDate:         arg.TransactionDate,
				})
				if err != nil {
					return err
				}
			}
		} else if strings.EqualFold(arg.TransType, "buy") {
			//fund
			new_balance = D_ori_balance.Sub(D_priceToHandle)

			//userStoc
			new_user_stock_quantity = arg.Amt + oriUserStock.Quantity
			oriTotlaStcokCost = D_amt.Mul(arg.PerPrice)
			newPricePerShare := oriTotlaStcokCost.Add(D_amt.Mul(arg.PerPrice)).Div(D_amt)
			if isHasUserStock {
				newUserStock, err = q.UpdateUserStock(ctx, UpdateUserStockParams{
					UserID:                arg.UserID,
					StockID:               arg.FundID,
					Quantity:              new_user_stock_quantity,
					PurchasePricePerShare: newPricePerShare.String(),
					PurchasedDate:         arg.TransactionDate,
					UpUser:                utility.StringToSqlNiStr(arg.CreateUser),
				})
				if err != nil {
					return err
				}
			} else {
				//Insert
				newUserStock, err = q.CreateUserStock(ctx, CreateUserStockParams{
					UserID:                arg.UserID,
					StockID:               arg.FundID,
					Quantity:              new_user_stock_quantity,
					PurchasePricePerShare: newPricePerShare.String(),
					PurchasedDate:         arg.TransactionDate,
					CrUser:                arg.CreateUser,
				})
				if err != nil {
					return err
				}
			}

		}
		newFund, err := q.UpdateFund(ctx, UpdateFundParams{
			FundID:  arg.FundID,
			Balance: new_balance.String(),
			UpDate:  utility.TimeToSqlNiTime(time.Now().UTC()),
			UpUser:  utility.StringToSqlNiStr(arg.CreateUser),
		})
		if err != nil {
			return err
		}

		result.OriFund = oriFund
		result.OriUserStocks = &oriUserStock
		result.Fund = newFund
		result.UserStocks = &newUserStock
		log.Println("Test Successed")
		return nil
	})
	return result, nil
}