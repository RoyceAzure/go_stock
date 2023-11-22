package fake

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	dto "github.com/RoyceAzure/go-stockinfo-schduler/shared/model/DTO"
	"github.com/RoyceAzure/go-stockinfo-schduler/util"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/constants"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

var BATCH_SIZE = 2000

type FakeDataService[T any] interface {
	GenerateFakeData() []T
	SetPrototype(prototype []T)
	GetPrototype() []T
	GeneratePrototype(bool) ([]T, error)
}

type FakeSPRDataService struct {
	dao       repository.Dao
	prototype []repository.StockPriceRealtime
	mutex     sync.RWMutex
}

func NewFakeSPRDataService(dao repository.Dao) *FakeSPRDataService {
	return &FakeSPRDataService{
		dao: dao,
	}
}

func (fakeService *FakeSPRDataService) SetPrototype(prototype []repository.StockPriceRealtime) {
	fakeService.mutex.Lock()
	fakeService.prototype = prototype
	fakeService.mutex.Unlock()
}

func (fakeService *FakeSPRDataService) GetPrototype() []repository.StockPriceRealtime {
	fakeService.mutex.RLock()
	prototype := fakeService.prototype
	fakeService.mutex.RUnlock()
	return prototype
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
	go util.TaskDistributor(unprocessed, BATCH_SIZE, fakeService.GetPrototype(), &wg)
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

func (fakeService *FakeSPRDataService) GeneratePrototype(refresh bool) ([]repository.StockPriceRealtime, error) {
	log.Info().
		Msg("start of generate prototype")
	var result []repository.StockPriceRealtime
	var stock_day_alls []dto.StockDayAvgAllDTO
	var filePath string = "./doc/fake_source/STOCK_DAY_ALL.json"
	var byteData []byte
	var err error
	startTime := time.Now().UTC()

	if refresh {
		byteData, err = util.SendRequest(constants.METHOD_GET,
			constants.URL_STOCK_DAY_AVG_ALL,
			nil)
		if err != nil {
			return nil, fmt.Errorf("generate prototype get err : %w", err)
		}

		err = util.WriteJsonFile(filePath, byteData)
		if err != nil {
			return nil, fmt.Errorf("generate prototype write file get err : %w", err)
		}
	}

	byteData, err = util.ReadJsonFile(filePath)
	if err != nil {
		return result, fmt.Errorf("generate prototype get err : %w", err)
	}
	err = json.Unmarshal(byteData, &stock_day_alls)
	if err != nil {
		return result, fmt.Errorf("generate prototype , unmarshal get err : %w", err)
	}

	if len(stock_day_alls) == 0 {
		return result, fmt.Errorf("generate prototype get err failed, no source data")
	}

	//製作prototype
	var wg sync.WaitGroup

	wg.Add(5)
	unprocessed := make(chan []dto.StockDayAvgAllDTO)
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

	fakeService.SetPrototype(result)

	elapsed := time.Since(startTime)
	log.Info().Int64("elpase time (ms)", int64(elapsed/time.Millisecond)).Msg("end of generate prototype")
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
func Sdavg2StockPriceRealTime(value dto.StockDayAvgAllDTO) (*repository.StockPriceRealtime, error) {
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
		TransTime:    time.Now().UTC(),
	}, nil
}
