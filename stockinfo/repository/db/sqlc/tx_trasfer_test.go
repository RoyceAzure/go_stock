package db

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

/*
測試併發安全  使用goroutine
*/
func TestStockTransTxEach(t *testing.T) {
	store := NewStore(testDB)

	errs := make(chan error)
	results := make(chan TransferStockTxResults)
	ctx := context.Background()

	//測試用原始數據
	testUser, testFund, testStock, testUserStock = CreateRandomUserStockNoTest(100, 200, 0, 1)
	var userStockChange, stockChange int32
	var fundChange decimal.Decimal
	var oriUserStockAmt int32
	if testUserStock == nil {
		oriUserStockAmt = 0
	} else {
		oriUserStockAmt = testUserStock.Quantity
	}
	n := 100
	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			transation := createRandomTransation(t)

			result, err := store.TransferStockTx(ctx, TransferStockTxParams{
				TransationID: transation.TransationID,
				CreateUser:   transation.CrUser,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		result := <-results
		if err != nil && errors.Is(err, constants.ErrInValidatePreConditionOp) {
			continue
		}
		require.NoError(t, err)

		require.NotEmpty(t, result)

		//chack transfer

		transRes := result.StockTrans
		fundRes := result.Fund
		oriFund := result.OriFund

		require.NotEmpty(t, transRes.TransationID)
		_, err = store.GetStockTransaction(context.Background(), transRes.TransationID)
		require.NoError(t, err)

		//檢查stock
		if strings.EqualFold(transRes.TransactionType, "buy") {
			require.Greater(t, result.Stock.MarketCap, int64(0))
			require.Equal(t, int64(transRes.TransationAmt), result.OriStock.MarketCap-result.Stock.MarketCap)

		} else if strings.EqualFold(transRes.TransactionType, "sell") {
			require.Equal(t, int64(transRes.TransationAmt), result.Stock.MarketCap-result.OriStock.MarketCap)
		}
		//檢查fund
		require.NotEmpty(t, fundRes)
		require.Equal(t, fundRes.UserID, transRes.UserID)

		D_amt := decimal.NewFromInt32(transRes.TransationAmt)
		require.NoError(t, err)

		D_stock_cur_price, err := decimal.NewFromString(transRes.TransationPricePerShare)
		require.NoError(t, err)

		priceToHandle := D_amt.Mul(D_stock_cur_price)

		D_ori_balance, err := decimal.NewFromString(oriFund.Balance)
		require.NoError(t, err)

		D_new_balance, err := decimal.NewFromString(fundRes.Balance)
		require.NoError(t, err)

		if strings.EqualFold(transRes.TransactionType, "buy") {
			require.True(t, D_new_balance.GreaterThanOrEqual(decimal.NewFromInt(0)))
			require.True(t, D_ori_balance.Sub(D_new_balance).Equal(priceToHandle))

		} else if strings.EqualFold(transRes.TransactionType, "sell") {
			require.True(t, D_new_balance.GreaterThanOrEqual(decimal.NewFromInt(0)))
			require.True(t, D_new_balance.Sub(D_ori_balance).Equal(priceToHandle))
		}

		//檢查UserStock
		oriUserStock := result.OriUserStock
		UserSotck := result.UserStock

		//原本沒有該股票
		if oriUserStock == nil {
			oriUserStock = &UserStock{}
		}

		//該股票賣完
		if UserSotck == nil {
			UserSotck = &UserStock{}
		}

		D_ori_amt := decimal.NewFromInt32(oriUserStock.Quantity)
		D_new_amt := decimal.NewFromInt32(UserSotck.Quantity)

		if strings.EqualFold(transRes.TransactionType, "buy") {
			require.True(t, D_new_amt.GreaterThanOrEqual(decimal.NewFromInt(0)))
			require.True(t, D_new_amt.Sub(D_ori_amt).Equal(D_amt))
		} else if strings.EqualFold(transRes.TransactionType, "sell") {
			require.True(t, D_ori_amt.Sub(D_new_amt).Equal(D_amt))
		}

		//更新全部資料
		if strings.EqualFold(transRes.TransactionType, "buy") {
			stockChange = stockChange - transRes.TransationAmt

			userStockChange = userStockChange + transRes.TransationAmt
			perPrice, _ := decimal.NewFromString(transRes.TransationPricePerShare)
			amt := decimal.NewFromInt32(transRes.TransationAmt)
			fundChange = fundChange.Sub(amt.Mul(perPrice))
		} else if strings.EqualFold(transRes.TransactionType, "sell") {
			stockChange = stockChange + transRes.TransationAmt
			userStockChange = userStockChange - transRes.TransationAmt
			perPrice, _ := decimal.NewFromString(transRes.TransationPricePerShare)
			amt := decimal.NewFromInt32(transRes.TransationAmt)
			fundChange = fundChange.Add(amt.Mul(perPrice))
		}
	}

	//檢查全部資料

	fund, err := testQueries.GetFund(ctx, testFund.FundID)
	require.NoError(t, err)
	require.NotEmpty(t, fund)

	var finalUserStockAmt int32

	userStock, err := testQueries.GetUserStockByUidandSidForUpdateNoK(ctx, GetUserStockByUidandSidForUpdateNoKParams{
		UserID:  testUser.UserID,
		StockID: testStock.StockID,
	})
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			finalUserStockAmt = 0
		} else {
			require.NoError(t, err)
		}
	} else {
		finalUserStockAmt = userStock.Quantity
	}

	stock, err := testQueries.GetStock(ctx, testStock.StockID)
	require.NoError(t, err)
	require.NotEmpty(t, stock)

	oriBalance, _ := decimal.NewFromString(testFund.Balance)
	finalBalance, _ := decimal.NewFromString(fund.Balance)
	require.Equal(t, finalBalance, oriBalance.Add(fundChange))

	require.Equal(t, stock.MarketCap, testStock.MarketCap+int64(stockChange))
	require.Equal(t, finalUserStockAmt, oriUserStockAmt+userStockChange)
}

func createRandomTransation(t *testing.T) StockTransaction {
	trans, err := testQueries.CreateStockTransaction(context.Background(), CreateStockTransactionParams{
		UserID:                  testUser.UserID,
		StockID:                 testStock.StockID,
		FundID:                  testFund.FundID,
		TransactionType:         util.RandomTransactionType(),
		TransactionDate:         time.Now().UTC(),
		TransationAmt:           int32(util.RandomInt(1, 5)),
		TransationPricePerShare: util.RandomStrInt(100000),
		CrUser:                  "test",
	})
	require.NoError(t, err)
	require.NotEmpty(t, trans)
	return trans
}
