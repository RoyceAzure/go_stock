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
	testUser, testFund, testStock, testUserStock = CreateRandomUserStockNoTest(util.RandomInt(1, 10),
		util.RandomInt(10, 20),
		util.RandomInt(1, 10),
		util.RandomInt(10, 100))
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

		priceToHandle := D_amt.Mul(D_stock_cur_price).Mul(decimal.NewFromInt(1000))

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
			fundChange = fundChange.Sub(priceToHandle)
		} else if strings.EqualFold(transRes.TransactionType, "sell") {
			stockChange = stockChange + transRes.TransationAmt
			userStockChange = userStockChange - transRes.TransationAmt
			fundChange = fundChange.Add(priceToHandle)
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

func TestTransation(t *testing.T) {

	ctx := context.Background()
	testStore := NewStore(testDB)
	//測試用原始數據
	testCase := []struct {
		name                  string //子測試名稱
		setUserStock          func(*CreateUserStockParams)
		setFund               func(*CreateFundParams)
		setStock              func(*CreateStockParams)
		setStockTransation    func(*CreateStockTransactionParams)
		checkFund             func(t *testing.T, res *Fund)
		checkUserStock        func(t *testing.T, res *UserStock)
		checkRealized         func(t *testing.T, res *RealizedProfitLoss)
		checkUnRealized       func(t *testing.T, res *GetRealizedProfitLosssByUserIdDetialRow)
		checkCreateTransation func(t *testing.T, res *StockTransaction)
		checkTransferTx       func(t *testing.T, res *TransferStockTxResults, err error)
	}{
		{
			name: "buy, not enough money",
			setFund: func(arg *CreateFundParams) {
				arg.Balance = "1000.00"
				arg.CurrencyType = string(constants.TW)
			},
			setStock: func(arg *CreateStockParams) {
				arg.StockCode = util.RandomString(5)
				arg.StockName = util.RandomString(5)
				arg.CurrentPrice = "1000"
			},
			setUserStock: func(arg *CreateUserStockParams) {
				arg.Quantity = 1
				arg.PurchasePricePerShare = "1.1111"
			},
			setStockTransation: func(arg *CreateStockTransactionParams) {
				arg.TransactionType = string(constants.BUY)
				arg.TransationAmt = 1
				arg.TransationPricePerShare = "1000.00"
			},
			checkCreateTransation: func(t *testing.T, res *StockTransaction) {
				require.NotEmpty(t, res)
			},
			checkTransferTx: func(t *testing.T, res *TransferStockTxResults, err error) {
				require.Empty(t, res)
				require.Error(t, err)
				require.True(t, errors.Is(err, constants.ErrInValidatePreConditionOp))
			},
		},
		{
			name: "sell, not enough user stock",
			setFund: func(arg *CreateFundParams) {
				arg.Balance = "1000.00"
				arg.CurrencyType = string(constants.TW)
			},
			setStock: func(arg *CreateStockParams) {
				arg.StockCode = util.RandomString(5)
				arg.StockName = util.RandomString(5)
				arg.CurrentPrice = "1000"
			},
			setUserStock: func(arg *CreateUserStockParams) {
				arg.Quantity = 5
				arg.PurchasePricePerShare = "1.1111"
			},
			setStockTransation: func(arg *CreateStockTransactionParams) {
				arg.TransactionType = string(constants.SELL)
				arg.TransationAmt = 10
				arg.TransationPricePerShare = "1000.00"
			},
			checkCreateTransation: func(t *testing.T, res *StockTransaction) {
				require.NotEmpty(t, res)
			},
			checkTransferTx: func(t *testing.T, res *TransferStockTxResults, err error) {
				require.Empty(t, res)
				require.Error(t, err)
				require.True(t, errors.Is(err, constants.ErrInValidatePreConditionOp))
			},
		},
		{
			name: "buy, success buy new stock",
			setFund: func(arg *CreateFundParams) {
				arg.Balance = "10000001.00"
				arg.CurrencyType = string(constants.TW)
			},
			setStock: func(arg *CreateStockParams) {
				arg.StockCode = util.RandomString(5)
				arg.StockName = util.RandomString(5)
				arg.CurrentPrice = "1000"
				arg.MarketCap = 10000
			},
			setUserStock: func(arg *CreateUserStockParams) {
				arg.Quantity = 5
				arg.PurchasePricePerShare = "1.1111"
			},
			setStockTransation: func(arg *CreateStockTransactionParams) {
				arg.TransactionType = string(constants.BUY)
				arg.TransationAmt = 10
				arg.TransationPricePerShare = "1000.00"
			},
			checkCreateTransation: func(t *testing.T, res *StockTransaction) {
				require.NotEmpty(t, res)
			},
			checkTransferTx: func(t *testing.T, res *TransferStockTxResults, err error) {
				require.NotEmpty(t, res)
				require.NoError(t, err)
			},
			checkFund: func(t *testing.T, res *Fund) {
				require.NotNil(t, res)
				balance, err := decimal.NewFromString(res.Balance)
				require.NoError(t, err)
				require.Equal(t, balance, decimal.NewFromInt(1))
			},
			checkUserStock: func(t *testing.T, res *UserStock) {
				require.NotNil(t, res)
				require.Equal(t, int32(15), res.Quantity)
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {

			user := CreateRandomUserNoTest()

			fundArg := CreateFundParams{}
			fundArg.UserID = user.UserID
			if tc.setFund != nil {
				tc.setFund(&fundArg)
			}
			fund := CreateFundSpecify(t, fundArg.UserID, fundArg.Balance, fundArg.CurrencyType)

			stockArg := CreateStockParams{}
			if tc.setStock != nil {
				tc.setStock(&stockArg)
			}

			stock := CreateStockSpecific(t, stockArg.StockCode, stockArg.StockName, stockArg.CurrentPrice, stockArg.MarketCap)

			userStockArg := CreateUserStockParams{}
			userStockArg.UserID = user.UserID
			userStockArg.StockID = stock.StockID
			if tc.setUserStock != nil {
				tc.setUserStock(&userStockArg)
			}

			userStock := CreateUserStockSpecific(t, userStockArg.UserID, userStockArg.StockID, userStockArg.Quantity, userStockArg.PurchasePricePerShare)

			transationArg := CreateStockTransactionParams{
				UserID:  user.UserID,
				StockID: stock.StockID,
				FundID:  fund.FundID,
			}
			if tc.setStockTransation != nil {
				tc.setStockTransation(&transationArg)
			}

			stockTransation := CreateStockTransactionsSepcific(t, transationArg.UserID,
				transationArg.StockID,
				transationArg.FundID,
				transationArg.TransactionType,
				transationArg.TransactionDate,
				transationArg.TransationAmt,
				transationArg.TransationPricePerShare)

			//grpc server可以直接call func
			tc.checkCreateTransation(t, &stockTransation)
			res, err := testStore.TransferStockTx(ctx, TransferStockTxParams{
				TransationID: stockTransation.TransationID,
				CreateUser:   stockTransation.CrUser,
			})
			tc.checkTransferTx(t, &res, err)
			if tc.checkFund != nil {
				f, err := testQueries.GetFund(ctx, fund.FundID)
				require.NoError(t, err)
				require.NotEmpty(t, f)
				tc.checkFund(t, &f)
			}

			if tc.checkUserStock != nil {
				us, err := testQueries.GetUserStock(ctx, userStock.UserStockID)
				require.NoError(t, err)
				require.NotEmpty(t, us)
				tc.checkUserStock(t, &us)
			}
		})
	}

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
