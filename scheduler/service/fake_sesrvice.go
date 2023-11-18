package service

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type FakeDataService[T any] interface {
	GenerateFakeData() []T
	SetPrototype(prototype []T)
	GeneratePrototype() ([]T, error)
}

type FakeSPRDataService struct {
	dao       repository.Dao
	prototype []repository.StockPriceRealtime
}

func NewFakeSPRDataService(dao repository.Dao) *FakeSPRDataService {
	return &FakeSPRDataService{
		dao: dao,
	}
}

func (fakeService *FakeSPRDataService) SetPrototype(prototype []repository.StockPriceRealtime) {
	// sourceData, err := fakeService.dao.GetSDAVGALLs(context.Background(), repository.GetSDAVGALLsParams{
	// 	CrDateStart: pgtype.Timestamptz{
	// 		Time:  startTime,
	// 		Valid: true,
	// 	},
	// 	CrDateEnd: pgtype.Timestamptz{
	// 		Time:  endTime,
	// 		Valid: true,
	// 	},
	// })
	// if err != nil {
	// 	return fmt.Errorf("get ssdavg has err: %w", err)
	// }
	fakeService.prototype = prototype
}

/*
根據prototype產生 fake data
*/
func (fakeService *FakeSPRDataService) GenerateFakeData() ([]repository.StockPriceRealtime, error) {

	var result []repository.StockPriceRealtime

	if fakeService == nil {
		return nil, fmt.Errorf("fake service is not init")
	}

	if len(fakeService.prototype) == 0 {
		return nil, fmt.Errorf("prototype is eempty")
	}

	var wg sync.WaitGroup
	unprocessed := make(chan []repository.StockPriceRealtime)
	processed := make(chan repository.StockPriceRealtime)
	wg.Add(4)
	go util.TaskDistributor(unprocessed, BATCH_SIZE, fakeService.prototype, &wg)
	go util.TaskWorker("worker 1", unprocessed, processed, createSCRFromPrototype, func(err error) {
		log.Warn().Err(err).Msg("err to process data")
	}, &wg)
	go util.TaskWorker("worker 2", unprocessed, processed, createSCRFromPrototype, func(err error) {
		log.Warn().Err(err).Msg("err to process data")
	}, &wg)
	go util.TaskWorker("worker 3", unprocessed, processed, createSCRFromPrototype, func(err error) {
		log.Warn().Err(err).Msg("err to process data")
	}, &wg)

	go func() {
		wg.Wait()
		close(processed)
	}()

	for process := range processed {
		result = append(result, process)
	}
	return result, nil
}

func (fakeService *FakeSPRDataService) GeneratePrototype() ([]repository.StockPriceRealtime, error) {
	var result []repository.StockPriceRealtime

	startTime := time.Now().UTC()
	cr_date_start := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())
	cr_date_end := cr_date_start.AddDate(0, 0, 1).Add(-time.Nanosecond)
	ctx := context.Background()
	stock_day_alls, err := fakeService.dao.GetSDAVGALLs(ctx, repository.GetSDAVGALLsParams{
		CrDateStart: pgtype.Timestamptz{
			Time:  cr_date_start,
			Valid: true,
		},
		CrDateEnd: pgtype.Timestamptz{
			Time:  cr_date_end,
			Valid: true,
		},
		Limits:  100000,
		Offsets: 0,
	})

	if err != nil {
		return result, fmt.Errorf("generate prottype get err : %w", err)
	}

	if len(stock_day_alls) == 0 {
		return result, fmt.Errorf("generate protoype get err failed, no source data")
	}

	//製作prototype
	var wg sync.WaitGroup

	wg.Add(5)
	unprocessed := make(chan []repository.StockDayAvgAll)
	processed := make(chan *repository.StockPriceRealtime)
	batchSize := 2000
	go util.TaskDistributor(unprocessed, batchSize, stock_day_alls, &wg)
	go util.TaskWorker("worker1", unprocessed, processed, Sdavg2StockPriceRealTime, nil, &wg)
	go util.TaskWorker("worker2", unprocessed, processed, Sdavg2StockPriceRealTime, nil, &wg)
	go util.TaskWorker("worker3", unprocessed, processed, Sdavg2StockPriceRealTime, nil, &wg)
	go util.TaskWorker("worker4", unprocessed, processed, Sdavg2StockPriceRealTime, nil, &wg)
	go func() {
		wg.Wait()
		close(processed)
	}()

	for batch := range processed {
		result = append(result, *batch)
	}

	return result, nil
}

/*
最大震幅3%
除了OpeningPrice  其餘欄位都是隨機產生 沒有參考價值
*/
func createSCRFromPrototype(prototype repository.StockPriceRealtime) (repository.StockPriceRealtime, error) {
	var result repository.StockPriceRealtime
	var open_priceOri float64
	open_priceOri, err := util.NumericToFloat64(prototype.OpeningPrice)
	if err != nil {
		log.Warn().Err(err).Msg("stock open price can't convert to float64")
		return result, err
	}

	randomFloatInRange := rand.Float64() * 3 / 100
	var finalPrice float64
	randomDir := rand.Int31n(1)
	if randomDir == 0 {
		finalPrice = open_priceOri * (1 - randomFloatInRange)
	} else {
		finalPrice = open_priceOri * (1 + randomFloatInRange)
	}
	open_price_numeric, err := util.StringToNumeric(util.Float64ToString(finalPrice))
	if err != nil {
		return result, err
	}

	trade_vol, _ := util.RandomNumeric(3, 0)
	trade_val, _ := util.RandomNumeric(6, 2)
	heightest_price, _ := util.RandomNumeric(6, 2)
	lowest_price, _ := util.RandomNumeric(6, 2)
	closing_price, _ := util.RandomNumeric(6, 2)
	change, _ := util.RandomNumeric(6, 2)
	trasation, _ := util.RandomNumeric(6, 2)

	return repository.StockPriceRealtime{
		Code:         prototype.Code,
		StockName:    prototype.StockName,
		TradeVolume:  trade_vol,
		TradeValue:   trade_val,
		OpeningPrice: open_price_numeric,
		HighestPrice: heightest_price,
		LowestPrice:  lowest_price,
		ClosingPrice: closing_price,
		Change:       change,
		Transaction:  trasation,
		TransTime:    time.Now().UTC(),
	}, nil
}

/*
for generate prototype
*/
func Sdavg2StockPriceRealTime(value repository.StockDayAvgAll) (*repository.StockPriceRealtime, error) {
	return &repository.StockPriceRealtime{
		Code:         value.Code,
		StockName:    value.StockName,
		TradeVolume:  pgtype.Numeric{Valid: true},
		TradeValue:   pgtype.Numeric{Valid: true},
		OpeningPrice: pgtype.Numeric{Valid: true},
		HighestPrice: pgtype.Numeric{Valid: true},
		LowestPrice:  pgtype.Numeric{Valid: true},
		ClosingPrice: pgtype.Numeric{Valid: true},
		Change:       pgtype.Numeric{Valid: true},
		Transaction:  pgtype.Numeric{Valid: true},
		TransTime:    value.CrDate,
	}, nil
}
