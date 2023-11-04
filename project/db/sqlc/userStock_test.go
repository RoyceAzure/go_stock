package db

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/stretchr/testify/require"
)

func TestCreateUserStock(t *testing.T) {
	CreateRandomUserStock(t)
}

func CreateRandomUserStock(t *testing.T) (User, Fund, Stock, UserStock) {
	user, fund := CreateRandomFund(t)
	stock := CreateRandomStock(t)
	quantity := int32(utility.RandomInt(1, 10))
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

func CreateRandomUserStockNoTest() (User, Fund, Stock, UserStock) {
	user := CreateRandomUserNoTest()
	_, fund := CreateRandomFundNoTest(user)
	stock := CreateRandomStockNoTest()
	quantity := utility.RandomInt(1, 10)
	userStock, _ := testQueries.CreateUserStock(context.Background(), CreateUserStockParams{
		UserID:                user.UserID,
		StockID:               stock.StockID,
		Quantity:              int32(quantity),
		PurchasePricePerShare: stock.CurrentPrice,
		PurchasedDate:         time.Now().UTC(),
		CrUser:                "royce",
	})
	return user, fund, stock, userStock
}
