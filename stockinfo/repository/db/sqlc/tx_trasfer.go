package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	logger "github.com/RoyceAzure/go-stockinfo/repository/logger_distributor"
	util "github.com/RoyceAzure/go-stockinfo/shared/util"
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
stock, userstock, fund 都會鎖

外面做:

	type檢查

裡面做:

	交易紀錄必須先存在
	fund, stock必須存在
	userStock可不必先存在
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
			logger.Logger.Error().Err(err).Msg("get stock transation failed")
			return util.InternalError(err)
		}

		stockTrans, err = q.UpdateStockTransationResult(ctx, UpdateStockTransationResultParams{
			Result:       TransationResultProcessed,
			TransationID: arg.TransationID,
		})
		if err != nil {
			logger.Logger.Error().Err(err).Msg("update stock transation failed")
			return util.InternalError(err)
		}

		currentPerPrice, err := decimal.NewFromString(stockTrans.TransationPricePerShare)
		if err != nil {
			logger.Logger.Error().Err(err).Msg("convert per price failed")
			return util.InternalError(err)
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
			logger.Logger.Error().Err(err).Msg("get fund by uid and fid err")
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
				oriUserStock.PurchasePricePerShare = "0.00"
			} else {
				logger.Logger.Error().Err(err).Msg("get user stock by uid and sid get err")
				return constants.ErrInternal
			}
		}
		//檢查stock總數量
		var oriStock, newStock Stock
		oriStock, err = q.GetStockForUpdate(ctx, stockTrans.StockID)
		if err != nil {
			logger.Logger.Error().Err(err).Msg("transfer stock get err")
			return constants.ErrInternal
		}
		if strings.EqualFold(stockTrans.TransactionType, "buy") && int32(oriStock.MarketCap) < stockTrans.TransationAmt {
			logger.Logger.Error().Err(err).Msg("transfer stock get err")
			return fmt.Errorf("not enough stock %w", constants.ErrInValidatePreConditionOp)
		}
		//計算操作所需金額
		D_priceToHandle := currentPerPrice.Mul(decimal.NewFromInt32(stockTrans.TransationAmt)).Mul(decimal.NewFromInt(1000))
		if err != nil {
			logger.Logger.Error().Err(err).Msg("failed to compute total price")
			return constants.ErrInternal
		}

		D_amt := decimal.NewFromInt32(stockTrans.TransationAmt)
		D_ori_balance, err := decimal.NewFromString(oriFund.Balance)
		if err != nil {
			logger.Logger.Error().Err(err).Msg("failed to convert balance")
			return constants.ErrInternal
		}

		if strings.EqualFold(stockTrans.TransactionType, "sell") && oriUserStock.Quantity < stockTrans.TransationAmt {
			logger.Logger.Error().Err(err).Msg("not enough user stock")
			return fmt.Errorf("not enough user stock %w", constants.ErrInValidatePreConditionOp)
		} else if strings.EqualFold(stockTrans.TransactionType, "buy") && D_ori_balance.LessThan(D_priceToHandle) {
			logger.Logger.Error().Err(err).Msg("not enough money")
			return fmt.Errorf("not enough money %w", constants.ErrInValidatePreConditionOp)
		}

		var newUserStock UserStock
		var new_balance decimal.Decimal
		var new_user_stock_quantity int32
		var newMarketCap int64

		//已花費成本價錢
		costPerPrice, err := decimal.NewFromString(oriUserStock.PurchasePricePerShare)
		if err != nil {
			logger.Logger.Error().Err(err).Msg("transfer stock get err")
			return constants.ErrInternal
		}
		costAmt := decimal.NewFromInt32(oriUserStock.Quantity)

		costTotalPrice := costPerPrice.Mul(costAmt)

		//更新操做
		if isSelled {
			//sell

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
					logger.Logger.Error().Err(err).Msg("failed to delete user stock")
					return constants.ErrInternal
				}
			} else {
				//賣出情況的平均持有成本不變
				newUserStock, err = q.UpdateUserStock(ctx, UpdateUserStockParams{
					UserID:                stockTrans.UserID,
					StockID:               stockTrans.StockID,
					Quantity:              new_user_stock_quantity,
					PurchasePricePerShare: oriUserStock.PurchasePricePerShare,
					PurchasedDate:         stockTrans.TransactionDate,
				})
				if err != nil {
					logger.Logger.Error().Err(err).Msg("failed to update user stock")
					return constants.ErrInternal
				}
			}

			//inssert RealizedProfitLoss

			var realizedPrecent decimal.Decimal
			if costTotalPrice.Equal(decimal.NewFromInt32(0)) {
				realizedPrecent = D_priceToHandle.Div(decimal.NewFromInt32(1))
			} else {
				realizedPrecent = D_priceToHandle.Div(costTotalPrice)
			}

			_, err = q.CreateRealizedProfitLoss(ctx, CreateRealizedProfitLossParams{
				TransationID:    stockTrans.TransationID,
				UserID:          stockTrans.UserID,
				ProductName:     fmt.Sprintf("%s %s", oriStock.StockCode, oriStock.StockName),
				CostPerPrice:    oriUserStock.PurchasePricePerShare,
				CostTotalPrice:  costTotalPrice.String(),
				Realized:        D_priceToHandle.String(),
				RealizedPrecent: realizedPrecent.String(),
			})
			if err != nil {
				logger.Logger.Error().Err(err).Msg("failed to create realized profit loss")
				return constants.ErrInternal
			}

		} else {
			//buy

			//fund
			new_balance = D_ori_balance.Sub(D_priceToHandle)

			//total stock 數量更新
			newMarketCap = oriStock.MarketCap - int64(stockTrans.TransationAmt)
			new_user_stock_quantity = oriUserStock.Quantity + stockTrans.TransationAmt
			//userStock
			if isHasUserStock {
				newPricePerShare := (costTotalPrice.Add(D_priceToHandle)).Div(D_amt.Add(costAmt))
				newUserStock, err = q.UpdateUserStock(ctx, UpdateUserStockParams{
					UserID:                stockTrans.UserID,
					StockID:               stockTrans.StockID,
					Quantity:              new_user_stock_quantity,
					PurchasePricePerShare: newPricePerShare.String(),
					PurchasedDate:         stockTrans.TransactionDate,
					UpUser:                util.StringToSqlNiStr(arg.CreateUser),
				})
				if err != nil {
					logger.Logger.Error().Err(err).Msg("failed to update user stock")
					return constants.ErrInternal
				}
			} else {
				//Insert  不需要更新price per share
				newUserStock, err = q.CreateUserStock(ctx, CreateUserStockParams{
					UserID:                stockTrans.UserID,
					StockID:               stockTrans.StockID,
					Quantity:              new_user_stock_quantity,
					PurchasePricePerShare: stockTrans.TransationPricePerShare,
					PurchasedDate:         stockTrans.TransactionDate,
					CrUser:                arg.CreateUser,
				})
				if err != nil {
					logger.Logger.Error().Err(err).Msg("failed to create user stock")
					return constants.ErrInternal
				}
			}
		}
		newFund, err := q.UpdateFund(ctx, UpdateFundParams{
			FundID:  stockTrans.FundID,
			Balance: new_balance.String(),
			UpDate:  util.TimeToSqlNiTime(time.Now().UTC()),
			UpUser:  util.StringToSqlNiStr(arg.CreateUser),
		})
		if err != nil {
			logger.Logger.Error().Err(err).Msg("failed to update fund")
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
			logger.Logger.Error().Err(err).Msg("failed to update stock")
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
		logger.Logger.Info().Msg("transfer stock successed")
		return nil
	})

	return result, err
}
