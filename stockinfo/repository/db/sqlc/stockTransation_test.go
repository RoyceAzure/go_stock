package db

import (
	"context"
	_ "database/sql"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/stretchr/testify/require"
)

func TestCreateStockTransactions(t *testing.T) {
	CreateRandomStockTransactions(t)
}

func CreateRandomStockTransactions(t *testing.T) StockTransaction {
	arg := CreateStockTransactionParams{
		UserID:                  1,
		StockID:                 1,
		FundID:                  1,
		TransactionType:         util.RandomCurrencyTypeStr(),
		TransactionDate:         time.Now().UTC(),
		TransationAmt:           10,
		TransationPricePerShare: "100",
		CrUser:                  "royce",
	}

	stockTrans, err := testQueries.CreateStockTransaction(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, stockTrans)

	require.Equal(t, arg.UserID, stockTrans.UserID)
	require.Equal(t, arg.TransactionType, stockTrans.TransactionType)
	// require.Equal(t, arg.TransactionDate.Time, stockTrans.TransactionDate.Time.UTC())
	require.Equal(t, arg.TransationAmt, stockTrans.TransationAmt)
	require.Equal(t, arg.TransationPricePerShare, stockTrans.TransationPricePerShare)

	//檢查db自動產生  不能為0值
	require.NotZero(t, stockTrans.TransationID)
	require.NotZero(t, stockTrans.CrDate)
	return stockTrans
}

func CreateStockTransactionsSepcific(t *testing.T, userId int64, stockId int64, fund int64, transType string, date time.Time, amt int32, pricePerShare string) StockTransaction {
	arg := CreateStockTransactionParams{
		UserID:                  userId,
		StockID:                 stockId,
		FundID:                  fund,
		TransactionType:         transType,
		TransactionDate:         date,
		TransationAmt:           amt,
		TransationPricePerShare: pricePerShare,
		CrUser:                  "test",
	}

	stockTrans, err := testQueries.CreateStockTransaction(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, stockTrans)

	require.Equal(t, arg.UserID, stockTrans.UserID)
	require.Equal(t, arg.TransactionType, stockTrans.TransactionType)
	require.Equal(t, arg.TransationAmt, stockTrans.TransationAmt)
	require.Equal(t, arg.TransationPricePerShare, stockTrans.TransationPricePerShare)

	require.NotZero(t, stockTrans.TransationID)
	require.NotZero(t, stockTrans.CrDate)
	return stockTrans
}

// func TestGetStockTransactions(t *testing.T) {
// 	fund := CreateRandomStockTransactions(t)
// 	fund2, err := testQueries.GetFund(context.Background(), fund.FundID)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, fund2)

// 	require.Equal(t, fund.UserID, fund2.UserID)
// 	require.Equal(t, fund.Balance, fund2.Balance)
// 	require.Equal(t, fund.CurrencyType, fund2.CurrencyType)
// 	require.Equal(t, fund.CrDate, fund2.CrDate)
// }

// func TestUpdateUser(t *testing.T) {
// 	user := CreateRandomUser(t)

// 	arg := UpdateUserParams{
// 		UserID:   user.UserID,
// 		UserName: utility.RandomString(10),
// 	}
// 	user2, err := testQueries.UpdateUser(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, arg.UserName, user2.UserName)
// 	require.Equal(t, user.Email, user2.Email)
// 	require.Equal(t, user.Password.String, user2.Password.String)
// 	require.Equal(t, user.SsoIdentifer.String, user2.SsoIdentifer.String)

// 	require.WithinDuration(t, user.CrDate.Time, user2.CrDate.Time, time.Second)
// }

// func TestDeleteStockTransactions(t *testing.T) {
// 	res, err := testQueries.GetStockTransactionsFilter(context.Background(), GetStockTransactionsFilterParams{
// 		Offsets: 0,
// 		Limits:  10,
// 	})
// 	require.NoError(t, err)
// 	require.NotEmpty(t, res)
// }