package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	utility "github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
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
	TransationID int64  `json:"trans_id"`
	CreateUser   string `json:"cr_uesr"`
}

type TransferStockTxResults struct {
	Stock        Stock            `json:"stock"`
	OriStock     Stock            `json:"ori_stock"`
	Fund         Fund             `json:"fund"`
	OriFund      Fund             `json:"ori_fund"`
	StockTrans   StockTransaction `json:"stock_trans"`
	UserStock    *UserStock       `json:"user_stocks"`
	OriUserStock *UserStock       `json:"ori_user_stocks"`
	User         User             `json:"users"`
}

/*
fund id 會鎖
user stock會鎖
有沒有必要鎖 stock 裡面的數量???

外面做:

	type檢查

裡面做:

	鎖錢包檢查金額
	鎖userstock檢查股票數量
	鎖stock 檢查數量
*/
func (store *SQLStore) TransferStockTx(ctx context.Context, arg TransferStockTxParams) (TransferStockTxResults, error) {
	var result TransferStockTxResults

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		isHasUserStock := true

		stockTrans, err := q.GetStockTransaction(ctx, arg.TransationID)
		if err != nil {
			return utility.InternalError(err)
		}
		perPrice, err := decimal.NewFromString(stockTrans.TransationPricePerShare)
		if err != nil {
			return utility.InternalError(err)
		}

		var isSelled bool
		if strings.EqualFold(stockTrans.TransactionType, constants.SELL) {
			isSelled = true
		} else {
			isSelled = false
		}
		//select for update no key, 受引用關聯的表仍可以做操作
		//由於目前的測試  所有的平行測試都是按照同樣順序  先取fund在取  userstock, 只要第一個人取得fund  後續其他人就會卡在這
		//有就不會有其他人先取得userstock，導致deadlock情況發生
		oriFund, err := q.GetFundByUidandFidForUpdateNoK(ctx, GetFundByUidandFidForUpdateNoKParams{
			UserID: stockTrans.UserID,
			FundID: stockTrans.FundID,
		})
		if err != nil {
			return constants.ErrInValidatePreConditionOp
		}
		//select for update no key, 受引用關聯的表仍可以做操作
		oriUserStock, err := q.GetUserStockByUidandSidForUpdateNoK(ctx, GetUserStockByUidandSidForUpdateNoKParams{
			UserID:  stockTrans.UserID,
			StockID: stockTrans.StockID,
		})
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				isHasUserStock = false
				oriUserStock = UserStock{}
			} else {
				return constants.ErrInternal
			}
		}
		//檢查stock總數量
		var oriStock, newStock Stock
		oriStock, err = q.GetStockForUpdate(ctx, stockTrans.StockID)
		if err != nil {
			return constants.ErrInternal
		}
		if strings.EqualFold(stockTrans.TransactionType, "buy") && int32(oriStock.MarketCap) < stockTrans.TransationAmt {
			return fmt.Errorf("not enough stock %w", constants.ErrInValidatePreConditionOp)
		}
		//計算操作所需金額
		D_priceToHandle := perPrice.Mul(decimal.NewFromInt32(stockTrans.TransationAmt))
		if err != nil {
			return constants.ErrInValidatePreConditionOp
		}

		D_amt := decimal.NewFromInt32(stockTrans.TransationAmt)
		D_ori_balance, err := decimal.NewFromString(oriFund.Balance)
		if err != nil {
			return constants.ErrInValidatePreConditionOp
		}

		if strings.EqualFold(stockTrans.TransactionType, "sell") && oriUserStock.Quantity < stockTrans.TransationAmt {
			return fmt.Errorf("not enough user stock %w", constants.ErrInValidatePreConditionOp)
		} else if strings.EqualFold(stockTrans.TransactionType, "buy") && D_ori_balance.LessThan(D_priceToHandle) {
			return fmt.Errorf("not enough money %w", constants.ErrInValidatePreConditionOp)
		}

		var newUserStock UserStock
		var new_balance, oriTotlaStcokCost decimal.Decimal
		var new_user_stock_quantity int32
		var newMarketCap int64
		//更新操做
		if isSelled {
			//fund
			new_balance = D_ori_balance.Add(D_priceToHandle)

			//total stock 數量更新
			newMarketCap = oriStock.MarketCap + int64(stockTrans.TransationAmt)

			//userStoc
			new_user_stock_quantity = oriUserStock.Quantity - stockTrans.TransationAmt

			if new_user_stock_quantity == 0 {
				//刪除操作
				err := q.DeleteUserStock(ctx, oriUserStock.UserStockID)
				if err != nil {
					return constants.ErrInternal
				}
			} else {
				//賣出情況的平均持有成本不變
				newUserStock, err = q.UpdateUserStock(ctx, UpdateUserStockParams{
					UserID:                stockTrans.UserID,
					StockID:               stockTrans.FundID,
					Quantity:              new_user_stock_quantity,
					PurchasePricePerShare: oriUserStock.PurchasePricePerShare,
					PurchasedDate:         stockTrans.TransactionDate,
				})
				if err != nil {
					return constants.ErrInternal
				}
			}
		} else {
			//fund
			new_balance = D_ori_balance.Sub(D_priceToHandle)

			//total stock 數量更新
			newMarketCap = oriStock.MarketCap - int64(stockTrans.TransationAmt)

			//userStoc
			new_user_stock_quantity = stockTrans.TransationAmt + oriUserStock.Quantity
			oriTotlaStcokCost = D_amt.Mul(perPrice)
			newPricePerShare := oriTotlaStcokCost.Add(D_amt.Mul(perPrice)).Div(D_amt)
			if isHasUserStock {
				newUserStock, err = q.UpdateUserStock(ctx, UpdateUserStockParams{
					UserID:                stockTrans.UserID,
					StockID:               stockTrans.FundID,
					Quantity:              new_user_stock_quantity,
					PurchasePricePerShare: newPricePerShare.String(),
					PurchasedDate:         stockTrans.TransactionDate,
					UpUser:                utility.StringToSqlNiStr(arg.CreateUser),
				})
				if err != nil {
					return constants.ErrInternal
				}
			} else {
				//Insert
				newUserStock, err = q.CreateUserStock(ctx, CreateUserStockParams{
					UserID:                stockTrans.UserID,
					StockID:               stockTrans.FundID,
					Quantity:              new_user_stock_quantity,
					PurchasePricePerShare: newPricePerShare.String(),
					PurchasedDate:         stockTrans.TransactionDate,
					CrUser:                arg.CreateUser,
				})
				if err != nil {
					return constants.ErrInternal
				}
			}
		}

		newFund, err := q.UpdateFund(ctx, UpdateFundParams{
			FundID:  stockTrans.FundID,
			Balance: new_balance.String(),
			UpDate:  utility.TimeToSqlNiTime(time.Now().UTC()),
			UpUser:  utility.StringToSqlNiStr(arg.CreateUser),
		})
		if err != nil {
			return constants.ErrInternal
		}

		//更新total stock
		newStock, err = q.UpdateStock(ctx, UpdateStockParams{
			StockID: oriStock.StockID,
			MarketCap: sql.NullInt64{
				Int64: newMarketCap,
				Valid: true,
			},
		})
		if err != nil {
			return constants.ErrInternal
		}
		result.StockTrans = stockTrans
		result.Stock = newStock
		result.OriStock = oriStock
		result.OriFund = oriFund
		if !isHasUserStock {
			result.OriUserStock = nil
		} else {
			result.OriUserStock = &oriUserStock
		}
		result.Fund = newFund
		result.UserStock = &newUserStock
		log.Println("Test Successed")
		return nil
	})

	return result, err
}
