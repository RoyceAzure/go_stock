package fake

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	logger "github.com/RoyceAzure/go-stockinfo-scheduler/repository/logger_distributor"
	repository "github.com/RoyceAzure/go-stockinfo-scheduler/repository/sqlc"
	dto "github.com/RoyceAzure/go-stockinfo-scheduler/shared/model/DTO"
	"github.com/RoyceAzure/go-stockinfo-scheduler/util"
	"github.com/RoyceAzure/go-stockinfo-scheduler/util/constants"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
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

func NewFakeSPRDataService(dao repository.Dao) (*FakeSPRDataService, error) {
	fakeDataService := &FakeSPRDataService{
		dao: dao,
	}

	prototype, err := fakeDataService.GeneratePrototype(true)
	if err != nil {
		return nil, err
	}
	fakeDataService.SetPrototype(prototype)

	return fakeDataService, nil
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
if presesrve, new fakedata will save as prototype
*/
func (fakeService *FakeSPRDataService) GenerateFakeData(presesrve bool) ([]repository.StockPriceRealtime, error) {

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
		logger.Logger.Warn().Err(err).Msg("err to process data")
	}, &wg)
	go util.TaskWorker("worker 2", unprocessed, processed, createSCRFromPrototype, func(err error) {
		logger.Logger.Warn().Err(err).Msg("err to process data")
	}, &wg)
	go util.TaskWorker("worker 3", unprocessed, processed, createSCRFromPrototype, func(err error) {
		logger.Logger.Warn().Err(err).Msg("err to process data")
	}, &wg)

	go func() {
		wg.Wait()
		close(processed)
	}()

	for process := range processed {
		result = append(result, process)
	}

	if presesrve {
		fakeService.SetPrototype(result)
	}
	return result, nil
}

/*
generate prototype from STOCK_DAY_ALL.json
if refresh, will download STOCK_DAY_ALL and save as json file
*/
func (fakeService *FakeSPRDataService) GeneratePrototype(refresh bool) ([]repository.StockPriceRealtime, error) {
	logger.Logger.Trace().
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
	logger.Logger.Trace().Int64("elpase_time_(ms)", int64(elapsed/time.Millisecond)).Msg("end of generate prototype")
	return result, nil
}

/*
最大震幅3%
除了OpeningPrice  其餘欄位都是隨機產生 沒有參考價值
*/
func createSCRFromPrototype(prototype repository.StockPriceRealtime) (repository.StockPriceRealtime, error) {
	var result repository.StockPriceRealtime
	var openPrice, closePrice pgxdecimal.Decimal
	err := openPrice.ScanNumeric(prototype.OpeningPrice)
	if err != nil {
		logger.Logger.Warn().Err(err).Msg("stock open price can't convert to float64")
		return result, err
	}
	err = closePrice.ScanNumeric(prototype.ClosingPrice)
	if err != nil {
		logger.Logger.Warn().Err(err).Msg("stock open price can't convert to float64")
		return result, err
	}

	var finalOpenPrice, finalClosePrice *decimal.Decimal

	randomFloatInRange := decimal.NewFromFloat32(rand.Float32() * 3 / 100)
	randomDir := rand.Int31n(1)

	finalOpenPrice, err = processPrice(openPrice, randomDir, randomFloatInRange)
	if err != nil {
		logger.Logger.Warn().Err(err).Msg("stock open price can't convert to decimal")
		return result, err
	}
	finalClosePrice, err = processPrice(openPrice, randomDir, randomFloatInRange)
	if err != nil {
		logger.Logger.Warn().Err(err).Msg("stock close price can't convert to decimal")
		return result, err
	}

	var OpeningPrice, ClosingPrice pgtype.Numeric
	err = OpeningPrice.Scan(finalOpenPrice.String())
	if err != nil {
		logger.Logger.Warn().Err(err).Msg("stock open price can't convert to numeric")
		return result, err
	}
	err = ClosingPrice.Scan(finalClosePrice.String())
	if err != nil {
		logger.Logger.Warn().Err(err).Msg("stock close price can't convert to numeric")
		return result, err
	}

	trade_vol, _ := util.RandomNumeric(3, 0)
	trade_val, _ := util.RandomNumeric(6, 2)
	heightest_price, _ := util.RandomNumeric(6, 2)
	lowest_price, _ := util.RandomNumeric(6, 2)
	change, _ := util.RandomNumeric(6, 2)
	trasation, _ := util.RandomNumeric(6, 2)

	return repository.StockPriceRealtime{
		Code:         prototype.Code,
		StockName:    prototype.StockName,
		TradeVolume:  trade_vol,
		TradeValue:   trade_val,
		OpeningPrice: OpeningPrice,
		HighestPrice: heightest_price,
		LowestPrice:  lowest_price,
		ClosingPrice: ClosingPrice,
		Change:       change,
		Transaction:  trasation,
		TransTime:    time.Now().UTC(),
	}, nil
}

/*
for generate prototype
use SDA.MonthlyAvgPrice as prototype OpeningPrice
use SDA.ClosePrice as prototype ClosingPrice
*/
func Sdavg2StockPriceRealTime(value dto.StockDayAvgAllDTO) (*repository.StockPriceRealtime, error) {
	var openPrice pgtype.Numeric
	err := openPrice.Scan(value.MonthlyAvgPrice)
	if err != nil {
		return nil, err
	}
	//string to numeric
	var closePrice pgtype.Numeric
	err = closePrice.Scan(value.ClosePrice)
	if err != nil {
		return nil, err
	}

	return &repository.StockPriceRealtime{
		Code:         value.Code,
		StockName:    value.StockName,
		TradeVolume:  pgtype.Numeric{Valid: true},
		TradeValue:   pgtype.Numeric{Valid: true},
		OpeningPrice: openPrice,
		HighestPrice: pgtype.Numeric{Valid: true},
		LowestPrice:  pgtype.Numeric{Valid: true},
		ClosingPrice: closePrice,
		Change:       pgtype.Numeric{Valid: true},
		Transaction:  pgtype.Numeric{Valid: true},
		TransTime:    time.Now().UTC(),
	}, nil
}

/*
des:

	將pgxdecimal.Decimal 隨機變換百分比趴樹的值  回傳*decimal.Decimal

parm:

	price (pgxdecimal.Decimal) : value to be mutiply by randomFloatInRange
	randomDir (int32) : change direction, 0 down, 1 up
	randomFloatInRange (decimal.Decimal) : precent of 震幅

return:

	*decimal.Decimal
*/
func processPrice(price pgxdecimal.Decimal, randomDir int32, randomFloatInRange decimal.Decimal) (*decimal.Decimal, error) {
	var finalPrice decimal.Decimal
	decimalOne := decimal.NewFromInt(1)

	priceStr := decimal.Decimal(price).String()

	// 將 pgxdecimal.Decimal 轉換為 decimal.Decimal
	decPrice, err := decimal.NewFromString(priceStr)
	if err != nil {
		return nil, err
	}

	// 進行計算
	if randomDir == 0 {
		finalPrice = decPrice.Mul(decimalOne.Sub(randomFloatInRange))
	} else {
		finalPrice = decPrice.Mul(decimalOne.Add(randomFloatInRange))
	}

	return &finalPrice, nil
}
