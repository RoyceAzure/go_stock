package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/RoyceAzure/go-stockinfo-scheduler/util"
	"github.com/stretchr/testify/require"
)

func TestCreateRealizedProfitLoss(t *testing.T) {
	CreateRandomRealizedProfitLoss(t)
}

func CreateRandomRealizedProfitLoss(t *testing.T) RealizedProfitLoss {
	ctx := context.Background()
	testUser, testFund, testStock, testUserStock = CreateRandomUserStockNoTest(100, 200, 0, 1)
	transation := createRandomTransation(t)
	res, err := testQueries.CreateRealizedProfitLoss(ctx, CreateRealizedProfitLossParams{
		TransationID:    transation.TransationID,
		UserID:          testUser.UserID,
		ProductName:     fmt.Sprintf("%s %s", testStock.StockCode, testStock.StockName),
		CostPerPrice:    util.RandomFloatString(5, 2),
		CostTotalPrice:  util.RandomFloatString(5, 2),
		Realized:        util.RandomFloatString(7, 2),
		RealizedPrecent: util.RandomFloatString(1000, 2),
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)
	return res
}

func TestGetRealizedProfitLoss(t *testing.T) {
	data := CreateRandomRealizedProfitLoss(t)
	res, err := testQueries.GetRealizedProfitLoss(context.Background(), data.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.EqualValues(t, data, res)
}

// func TestGetRealizedProfitLosssByUserId(t *testing.T) {
// 	var dataList []RealizedProfitLoss
// 	for i := 0; i < 5; i++ {
// 		dataList = append(dataList, CreateRandomRealizedProfitLoss(t))
// 	}
// 	res, err := testQueries.GetRealizedProfitLosssByUserId(context.Background(), GetRealizedProfitLosssByUserIdParams{
// 		UserID: dataList[0].UserID,
// 		Limit:  5,
// 		Offset: 0,
// 	})
// 	require.NoError(t, err)
// 	require.NotEmpty(t, res)
// 	require.Len(t, res, 5)
// 	for _, data := range res {
// 		require.Equal(t, dataList[0].UserID, data.UserID)
// 	}
// }
