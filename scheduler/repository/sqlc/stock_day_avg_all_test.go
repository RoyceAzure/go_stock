package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-schduler/util"
	"github.com/stretchr/testify/require"
)

func RadomCreateSDAVGALL() StockDayAvgAll {
	return StockDayAvgAll{
		ID:              util.RandomInt64(10, 1000),
		Code:            util.RandomString(5),
		StockName:       util.RandomString(10),
		ClosePrice:      util.RandomFloatString(4, 2),
		MonthlyAvgPrice: util.RandomFloatString(4, 2),
		CrDate:          time.Now().UTC(),
		CrUser:          util.RandomString(5),
	}
}

func CreateRandomSDAVGALL(t *testing.T) StockDayAvgAll {
	startTime := time.Now()
	insertData := RadomCreateSDAVGALL()
	newData, err := testQueries.CreateSDAVGALL(context.Background(), CreateSDAVGALLParams{
		Code:            insertData.Code,
		StockName:       insertData.StockName,
		ClosePrice:      insertData.ClosePrice,
		MonthlyAvgPrice: insertData.MonthlyAvgPrice,
	})
	require.NoError(t, err)
	require.NotEmpty(t, newData)

	require.Equal(t, insertData.Code, newData.Code)
	require.Equal(t, insertData.StockName, newData.StockName)
	require.Equal(t, insertData.ClosePrice, newData.ClosePrice)
	require.Equal(t, insertData.MonthlyAvgPrice, newData.MonthlyAvgPrice)

	require.WithinDuration(t, time.Now(), startTime, time.Second)
	return newData
}

func TestCreateSDAVGALL(t *testing.T) {
	CreateRandomSDAVGALL(t)
}

func TestGetSDAVGALLs(t *testing.T) {
	startTime := time.Now()

	limit := 20
	page := 1500
	offset := (page - 1) * limit
	data, err := testQueries.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			Limits:  500,
			Offsets: int32(offset),
		})
	require.NoError(t, err)
	require.NotEmpty(t, data)
	require.Len(t, data, 500)
	require.WithinDuration(t, time.Now(), startTime, time.Second)
}

func TestGetSDAVGALLsByCode(t *testing.T) {
	newData := CreateRandomSDAVGALL(t)
	startTime := time.Now()
	limit := 2
	page := 1
	offset := (page - 1) * limit
	datas, err := testQueries.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			Code: sql.NullString{
				Valid:  true,
				String: newData.Code,
			},
			Limits:  500,
			Offsets: int32(offset),
		})
	require.NoError(t, err)
	require.NotEmpty(t, datas)
	for _, data := range datas {
		require.Equal(t, newData.Code, data.Code)
	}

	require.WithinDuration(t, time.Now(), startTime, time.Second)
}

func TestGetSDAVGALLsByStockName(t *testing.T) {
	newData := CreateRandomSDAVGALL(t)
	startTime := time.Now()
	limit := 2
	page := 1
	offset := (page - 1) * limit
	datas, err := testQueries.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			StockName: sql.NullString{
				Valid:  true,
				String: newData.StockName,
			},
			Limits:  500,
			Offsets: int32(offset),
		})
	require.NoError(t, err)
	require.NotEmpty(t, datas)
	for _, data := range datas {
		require.Equal(t, newData.StockName, data.StockName)
	}

	require.WithinDuration(t, time.Now(), startTime, time.Second)
}

func TestGetSDAVGALLsByCPInterval(t *testing.T) {
	newData := CreateRandomSDAVGALL(t)
	startTime := time.Now()
	limit := 10
	page := 1
	offset := (page - 1) * limit
	datas, err := testQueries.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			CpUpper: sql.NullString{
				Valid:  true,
				String: newData.ClosePrice,
			},
			CpLower: sql.NullString{
				Valid:  true,
				String: newData.ClosePrice,
			},
			Limits:  int32(limit),
			Offsets: int32(offset),
		})
	require.NoError(t, err)
	require.NotEmpty(t, datas)
	require.Len(t, datas, 1)
	for _, data := range datas {
		require.Equal(t, newData.Code, data.Code)
		require.Equal(t, newData.StockName, data.StockName)
		require.Equal(t, newData.ClosePrice, data.ClosePrice)
		require.Equal(t, newData.MonthlyAvgPrice, data.MonthlyAvgPrice)
	}
	require.WithinDuration(t, time.Now(), startTime, time.Second)
}

func TestGetSDAVGALLsByMapInterval(t *testing.T) {
	newData := CreateRandomSDAVGALL(t)
	startTime := time.Now()
	limit := 10
	page := 1
	offset := (page - 1) * limit
	datas, err := testQueries.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			MapUpper: sql.NullString{
				Valid:  true,
				String: newData.MonthlyAvgPrice,
			},
			MapLower: sql.NullString{
				Valid:  true,
				String: newData.MonthlyAvgPrice,
			},
			Limits:  int32(limit),
			Offsets: int32(offset),
		})
	require.NoError(t, err)
	require.NotEmpty(t, datas)
	require.Len(t, datas, 1)
	for _, data := range datas {
		require.Equal(t, newData.Code, data.Code)
		require.Equal(t, newData.StockName, data.StockName)
		require.Equal(t, newData.ClosePrice, data.ClosePrice)
		require.Equal(t, newData.MonthlyAvgPrice, data.MonthlyAvgPrice)
	}
	require.WithinDuration(t, time.Now(), startTime, time.Second)
}
func TestGetSDAVGALLsByCrDateInterval(t *testing.T) {
	CreateRandomSDAVGALL(t)
	startTime := time.Now()
	crDateStart := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())
	crDateEnd := crDateStart.AddDate(0, 0, 1).Add(-time.Nanosecond)
	limit := 10
	page := 1
	offset := (page - 1) * limit
	datas, err := testQueries.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			CrDateStart: sql.NullTime{
				Valid: true,
				Time:  crDateStart,
			},
			CrDateEnd: sql.NullTime{
				Valid: true,
				Time:  crDateEnd,
			},
			Limits:  int32(limit),
			Offsets: int32(offset),
		})
	require.NoError(t, err)
	require.NotEmpty(t, datas)
	for _, data := range datas {
		require.True(t, data.CrDate.After(crDateStart) && data.CrDate.Before(crDateEnd))
	}
	require.WithinDuration(t, time.Now(), startTime, time.Second)
}

func TestGetSDAVGALLsBySpecialCrDate(t *testing.T) {
	startTime := time.Now()
	//2023-11-07 13:34:06.506946+00
	timetemp := time.Date(2023, time.November, 7, 13, 34, 6, 506946000, time.UTC)
	limit := 10
	page := 1
	offset := (page - 1) * limit
	datas, err := testQueries.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			CrDateStart: sql.NullTime{
				Valid: true,
				Time:  timetemp,
			},
			CrDateEnd: sql.NullTime{
				Valid: true,
				Time:  timetemp,
			},
			Limits:  int32(limit),
			Offsets: int32(offset),
		})
	require.NoError(t, err)
	require.NotEmpty(t, datas)
	for _, data := range datas {
		require.True(t, data.CrDate.Equal(timetemp))
	}
	require.WithinDuration(t, time.Now(), startTime, time.Second)
}

func TestGetSDAVGALLsBySpecialCrDatetTz(t *testing.T) {
	startTime := time.Now()
	//2023-11-07 13:34:06.506946+00
	timetemp := time.Date(2023, time.November, 7, 21, 34, 6, 506946000, time.Local)
	limit := 10
	page := 1
	offset := (page - 1) * limit
	datas, err := testQueries.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			CrDateStart: sql.NullTime{
				Valid: true,
				Time:  timetemp,
			},
			CrDateEnd: sql.NullTime{
				Valid: true,
				Time:  timetemp,
			},
			Limits:  int32(limit),
			Offsets: int32(offset),
		})
	require.NoError(t, err)
	require.NotEmpty(t, datas)
	for _, data := range datas {
		require.True(t, data.CrDate.Equal(timetemp))
	}
	require.WithinDuration(t, time.Now(), startTime, time.Second)
}
