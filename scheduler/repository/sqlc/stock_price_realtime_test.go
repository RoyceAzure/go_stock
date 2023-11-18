package repository

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-schduler/util"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

// go get github.com/stretchr/testify  使用require檢查錯誤
func TestCreateSPR(t *testing.T) {
	CreateRandomSPR(t)
}

func CreateRandomSPR(t *testing.T) StockPriceRealtime {
	trade_vol, _ := util.RandomNumeric(6, 2)
	trade_val, _ := util.RandomNumeric(6, 2)
	open_price, _ := util.RandomNumeric(6, 2)
	heightest_price, _ := util.RandomNumeric(6, 2)
	lowest_price, _ := util.RandomNumeric(6, 2)
	closing_price, _ := util.RandomNumeric(6, 2)
	change, _ := util.RandomNumeric(6, 2)
	trasation, _ := util.RandomNumeric(6, 2)
	arg := CreateSPRParams{
		Code:         utility.RandomString(6),
		StockName:    utility.RandomString(10),
		TradeVolume:  trade_vol,
		TradeValue:   trade_val,
		OpeningPrice: open_price,
		HighestPrice: heightest_price,
		LowestPrice:  lowest_price,
		ClosingPrice: closing_price,
		Change:       change,
		Transaction:  trasation,
		TransTime:    time.Now().UTC(),
	}
	spr, err := testDao.CreateSPR(context.TODO(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, spr)

	require.Equal(t, arg.Code, spr.Code)
	require.Equal(t, arg.StockName, spr.StockName)
	require.Equal(t, arg.TradeVolume, spr.TradeVolume)
	require.Equal(t, arg.TradeValue, spr.TradeValue)
	require.Equal(t, arg.OpeningPrice, spr.OpeningPrice)
	require.Equal(t, arg.LowestPrice, spr.LowestPrice)
	require.Equal(t, arg.ClosingPrice, spr.ClosingPrice)
	require.Equal(t, arg.Change, spr.Change)
	require.Equal(t, arg.Transaction, spr.Transaction)
	require.WithinDuration(t, time.Now().UTC(), arg.TransTime, time.Second)
	return spr
}

func TestGetSPR(t *testing.T) {
	spr := CreateRandomSPR(t)

	testCase := []struct {
		name      string
		argFunc   func() GetSPRsParams
		checkFunc func(t *testing.T, data []StockPriceRealtime)
	}{
		{
			name: "Get all",
			argFunc: func() GetSPRsParams {
				limit := 65535
				page := 1
				offset := (page - 1) * limit
				return GetSPRsParams{
					Limits:  int32(limit),
					Offsets: int32(offset),
				}
			},
			checkFunc: func(t *testing.T, data []StockPriceRealtime) {
				require.Greater(t, len(data), 0)
			},
		},
		{
			name: "Get by stock name",
			argFunc: func() GetSPRsParams {
				limit := 1
				page := 1
				offset := (page - 1) * limit
				return GetSPRsParams{
					StockName: pgtype.Text{
						String: spr.StockName,
						Valid:  true,
					},
					Limits:  int32(limit),
					Offsets: int32(offset),
				}
			},
			checkFunc: func(t *testing.T, data []StockPriceRealtime) {
				require.Equal(t, spr.StockName, data[0].StockName)
				require.Len(t, data, 1)
			},
		},
		{
			name: "Get by Code",
			argFunc: func() GetSPRsParams {
				limit := 1
				page := 1
				offset := (page - 1) * limit
				return GetSPRsParams{
					Code: pgtype.Text{
						String: spr.Code,
						Valid:  true,
					},
					Limits:  int32(limit),
					Offsets: int32(offset),
				}
			},
			checkFunc: func(t *testing.T, data []StockPriceRealtime) {
				require.Equal(t, spr.Code, data[0].Code)
				require.Len(t, data, 1)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			arg := tc.argFunc()
			data, err := testDao.GetSPRs(context.Background(), arg)
			require.NoError(t, err)
			tc.checkFunc(t, data)
		})
	}
}
