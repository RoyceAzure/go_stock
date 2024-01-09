package db

import (
	"context"
	"testing"

	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/stretchr/testify/require"
)

// go get github.com/stretchr/testify  使用require檢查錯誤
func TestCreateStock(t *testing.T) {
	CreateRandomStock(t)
}

func CreateRandomStock(t *testing.T) Stock {
	arg := CreateStockParams{
		StockCode:    util.RandomString(6),
		StockName:    util.RandomString(10),
		CurrentPrice: "600000",
		MarketCap:    99999999,
		CrUser:       "royce",
	}

	stock, err := testQueries.CreateStock(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, stock)

	require.Equal(t, arg.StockCode, stock.StockCode)
	require.Equal(t, arg.StockName, stock.StockName)
	require.Equal(t, arg.CurrentPrice, stock.CurrentPrice)
	require.Equal(t, arg.MarketCap, stock.MarketCap)
	require.Equal(t, arg.CrUser, stock.CrUser)

	//檢查db自動產生  不能為0值
	require.NotZero(t, stock.StockID)
	require.NotZero(t, stock.CrDate)
	return stock
}

func CreateRandomStockNoTest() Stock {
	arg := CreateStockParams{
		StockCode:    util.RandomString(6),
		StockName:    util.RandomString(10),
		CurrentPrice: "600000",
		MarketCap:    99999999,
		CrUser:       "royce",
	}

	stock, _ := testQueries.CreateStock(context.Background(), arg)
	return stock
}

func CreateStockSpecific(t *testing.T, stockCode string, stockName string, currentPrice string, marketCap int64) Stock {
	arg := CreateStockParams{
		StockCode:    stockCode,
		StockName:    stockName,
		CurrentPrice: currentPrice,
		MarketCap:    marketCap,
		CrUser:       "test",
	}

	stock, err := testQueries.CreateStock(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, stock)
	return stock
}