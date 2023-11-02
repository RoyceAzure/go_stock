package db

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestStockTransTx(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	store := NewStore(testDB)

	errs := make(chan error)
	results := make(chan TransferStockTxResults)

	n := 5
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferStockTx(context.Background(), TransferStockTxParams{
				UserID:          testUser.UserID,
				StockID:         testStock.StockID,
				FundID:          testFund.FundID,
				TransType:       utility.RandomTransactionType(),
				TransactionDate: time.Now().UTC(),
				Amt:             int32(utility.RandomInt(1, 5)),
				PerPrice:        decimal.NewFromInt(utility.RandomInt(1000, 10000)),
				CreateUser:      "royce",
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)

		//chack transfer

		transRes := result.StockTrans
		fundRes := result.Fund
		oriFund := result.OriFund

		require.NotEmpty(t, transRes.TransationId)
		_, err = store.GetStockTransaction(context.Background(), transRes.TransationId)
		require.NoError(t, err)

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
		oriUserStock := result.OriUserStocks
		UserSotck := result.UserStocks

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
	}
}
