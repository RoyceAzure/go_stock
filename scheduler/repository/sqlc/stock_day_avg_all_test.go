package repository

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-schduler/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func RadomCreateSDAVGALL() StockDayAvgAll {
	cs, _ := util.RandomNumeric(4, 2)
	ms, _ := util.RandomNumeric(4, 2)
	return StockDayAvgAll{
		ID:              util.RandomInt64(10, 1000),
		Code:            "Test" + util.RandomString(5),
		StockName:       "Test" + util.RandomString(10),
		ClosePrice:      cs,
		MonthlyAvgPrice: ms,
		CrDate:          time.Now().UTC(),
		CrUser:          util.RandomString(5),
	}
}

func RadomCreateSDAVGALLParams() BulkInsertDAVGALLParams {
	cs, _ := util.RandomNumeric(4, 2)
	ms, _ := util.RandomNumeric(4, 2)
	return BulkInsertDAVGALLParams{
		Code:            "BulkTest" + util.RandomString(5),
		StockName:       "BulkTest" + util.RandomString(10),
		ClosePrice:      cs,
		MonthlyAvgPrice: ms,
	}
}

func CreateRandomSDAVGALL(t *testing.T) StockDayAvgAll {
	startTime := time.Now()
	insertData := RadomCreateSDAVGALL()
	newData, err := testDao.CreateSDAVGALL(context.Background(), CreateSDAVGALLParams{
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

func TestBulkCreateSDAVGALL(t *testing.T) {
	var insertList []BulkInsertDAVGALLParams
	for i := 0; i < 50000; i++ {
		insertList = append(insertList, RadomCreateSDAVGALLParams())
	}
	res, err := testDao.BulkInsertDAVGALL(context.Background(), insertList)
	require.NoError(t, err)
	require.NotZero(t, res)
}

func TestGetSDAVGALLs(t *testing.T) {
	startTime := time.Now()

	limit := 20
	page := 1500
	offset := (page - 1) * limit
	data, err := testDao.GetSDAVGALLs(context.Background(),
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
	datas, err := testDao.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			Code: pgtype.Text{
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
	datas, err := testDao.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			StockName: pgtype.Text{
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
	startTime := time.Now()
	limit := 10
	page := 1
	offset := (page - 1) * limit
	cp_upper := util.RandomFloatString(6, 2)
	cp_lower := util.RandomFloatString(6, 2)
	if cp_lower < cp_upper {
		temp := cp_lower
		cp_lower = cp_upper
		cp_lower = temp
	}
	cp_num_upper, err := util.StringToNumeric(cp_upper)
	require.NoError(t, err)
	cp_num_lower, err := util.StringToNumeric(cp_lower)
	require.NoError(t, err)
	_, err = testDao.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			CpUpper: cp_num_upper,
			CpLower: cp_num_lower,
			Limits:  int32(limit),
			Offsets: int32(offset),
		})
	require.NoError(t, err)

	require.WithinDuration(t, time.Now(), startTime, time.Second)
}

func TestGetSDAVGALLsByMapInterval(t *testing.T) {
	startTime := time.Now()
	limit := 10
	page := 1
	offset := (page - 1) * limit
	mp_upper := util.RandomFloatString(6, 2)
	mp_lower := util.RandomFloatString(6, 2)
	if mp_lower < mp_upper {
		temp := mp_lower
		mp_lower = mp_upper
		mp_lower = temp
	}
	cp_num_upper, err := util.StringToNumeric(mp_upper)
	require.NoError(t, err)
	cp_num_lower, err := util.StringToNumeric(mp_lower)
	require.NoError(t, err)
	_, err = testDao.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			MapUpper: cp_num_upper,
			MapLower: cp_num_lower,
			Limits:   int32(limit),
			Offsets:  int32(offset),
		})
	require.NoError(t, err)
	require.WithinDuration(t, time.Now(), startTime, time.Second)
}
func TestGetSDAVGALLsByCrDateInterval(t *testing.T) {
	CreateRandomSDAVGALL(t)
	startTime := time.Now().UTC()
	crDateStart := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())
	crDateEnd := crDateStart.AddDate(0, 0, 1).Add(-time.Nanosecond)
	limit := 10
	page := 1
	offset := (page - 1) * limit
	datas, err := testDao.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			CrDateStart: pgtype.Timestamptz{
				Valid: true,
				Time:  crDateStart,
			},
			CrDateEnd: pgtype.Timestamptz{
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
	datas, err := testDao.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			CrDateStart: pgtype.Timestamptz{
				Valid: true,
				Time:  timetemp,
			},
			CrDateEnd: pgtype.Timestamptz{
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
	datas, err := testDao.GetSDAVGALLs(context.Background(),
		GetSDAVGALLsParams{
			CrDateStart: pgtype.Timestamptz{
				Valid: true,
				Time:  timetemp,
			},
			CrDateEnd: pgtype.Timestamptz{
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

func TestDeleteSDAVGALLCodePrexForTest(t *testing.T) {
	err := testDao.DeleteSDAVGALLCodePrexForTest(context.Background(), DeleteSDAVGALLCodePrexForTestParams{
		Len:        9,
		CodePrefix: "BulkTest",
	})
	require.NoError(t, err)
}
