package db

import (
	"context"
	"testing"
	"time"

	util "github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUserStock(t *testing.T) {
	CreateRandomUserStock(t)
}

func CreateRandomUserStock(t *testing.T) (User, Fund, Stock, UserStock) {
	user, fund := CreateRandomFund(t)
	stock := CreateRandomStock(t)
	quantity := int32(util.RandomInt(1, 10))
	userStock, err := testQueries.CreateUserStock(context.Background(), CreateUserStockParams{
		UserID:                user.UserID,
		StockID:               stock.StockID,
		Quantity:              quantity,
		PurchasePricePerShare: stock.CurrentPrice,
		PurchasedDate:         time.Now().UTC(),
		CrUser:                "royce",
	})

	require.NoError(t, err)
	require.NotEmpty(t, userStock)

	require.Equal(t, user.UserID, userStock.UserID)
	require.Equal(t, stock.StockID, userStock.StockID)
	require.Equal(t, quantity, userStock.Quantity)
	require.Equal(t, stock.CurrentPrice, userStock.PurchasePricePerShare)

	return user, fund, stock, userStock
}

func CreateRandomUserStockNoTest(min_fund int64, max_fund int64, min_quanty int64, max_quanty int64) (user User, fund Fund, stock Stock, userStock *UserStock) {
	user = CreateRandomUserNoTest()
	_, fund = CreateRandomFundNoTest(user, min_fund, max_fund)
	stock = CreateRandomStockNoTest()
	quantity := util.RandomInt(min_quanty, max_quanty)
	if quantity == 0 {
		userStock = nil
	} else {
		userStocktemp, _ := testQueries.CreateUserStock(context.Background(), CreateUserStockParams{
			UserID:                user.UserID,
			StockID:               stock.StockID,
			Quantity:              int32(quantity),
			PurchasePricePerShare: stock.CurrentPrice,
			PurchasedDate:         time.Now().UTC(),
			CrUser:                "royce",
		})
		userStock = &userStocktemp
	}

	return user, fund, stock, userStock
}
