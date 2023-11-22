package redisService

import (
	"context"
	"fmt"
	"sync"
	"time"

	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/util"
	decimal "github.com/jackc/pgx-shopspring-decimal"
)

type StockPriceRealtimeDTO struct {
	Code         string          `json:"code"`
	StockName    string          `json:"stock_name"`
	TradeVolume  decimal.Decimal `json:"trade_volume"`
	TradeValue   decimal.Decimal `json:"trade_value"`
	OpeningPrice decimal.Decimal `json:"opening_price"`
	HighestPrice decimal.Decimal `json:"highest_price"`
	LowestPrice  decimal.Decimal `json:"lowest_price"`
	ClosingPrice decimal.Decimal `json:"closing_price"`
	Change       decimal.Decimal `json:"change"`
	Transaction  decimal.Decimal `json:"transaction"`
	TransTime    time.Time       `json:"trans_time"`
}

type RedisServiceSPR interface {
	GetLatestSPR(ctx context.Context) ([]StockPriceRealtimeDTO, error)
}

func (j *JRedisService) GetLatestSPR(ctx context.Context) ([]StockPriceRealtimeDTO, error) {
	latestKey := j.redisDao.GetSPRLatestKey()
	if latestKey == "" {
		return nil, fmt.Errorf("there is no latest data in redis")
	}

	sprs, err := j.redisDao.FindSPRByID(ctx, latestKey)
	if err != nil {
		return nil, fmt.Errorf("get latest spr data in redis return err : %w", err)
	}

	var result []StockPriceRealtimeDTO

	var wg sync.WaitGroup
	batchSize := 2000
	unporcessed := make(chan []repository.StockPriceRealtime)
	processed := make(chan *StockPriceRealtimeDTO, batchSize/2)
	wg.Add(4)
	go util.TaskDistributor(unporcessed, batchSize, sprs, &wg)
	go util.TaskWorker("woker1", unporcessed, processed, cvSprEntity2DTO, nil, &wg)
	go util.TaskWorker("woker2", unporcessed, processed, cvSprEntity2DTO, nil, &wg)
	go util.TaskWorker("woker3", unporcessed, processed, cvSprEntity2DTO, nil, &wg)
	go func() {
		wg.Wait()
		close(processed)
	}()
	for data := range processed {
		result = append(result, *data)
	}

	return result, nil
}

func cvSprEntity2DTO(value repository.StockPriceRealtime) (*StockPriceRealtimeDTO, error) {
	var tradeVol, tradeVal, openPri, highPri, LowPri, Clopri, change, tra decimal.Decimal
	err := tradeVol.ScanNumeric(value.TradeVolume)
	if err != nil {
		return nil, fmt.Errorf("convert spr to dto dget error : %w", err)
	}
	err = tradeVal.ScanNumeric(value.TradeValue)
	if err != nil {
		return nil, fmt.Errorf("convert spr to dto dget error : %w", err)
	}

	err = openPri.ScanNumeric(value.OpeningPrice)
	if err != nil {
		return nil, fmt.Errorf("convert spr to dto dget error : %w", err)
	}

	err = highPri.ScanNumeric(value.HighestPrice)
	if err != nil {
		return nil, fmt.Errorf("convert spr to dto dget error : %w", err)
	}

	err = LowPri.ScanNumeric(value.LowestPrice)
	if err != nil {
		return nil, fmt.Errorf("convert spr to dto dget error : %w", err)
	}

	err = change.ScanNumeric(value.Change)
	if err != nil {
		return nil, fmt.Errorf("convert spr to dto dget error : %w", err)
	}

	err = tra.ScanNumeric(value.Transaction)
	if err != nil {
		return nil, fmt.Errorf("convert spr to dto dget error : %w", err)
	}

	return &StockPriceRealtimeDTO{
		Code:         value.Code,
		StockName:    value.StockName,
		TradeVolume:  tradeVol,
		TradeValue:   tradeVal,
		OpeningPrice: openPri,
		HighestPrice: highPri,
		LowestPrice:  LowPri,
		ClosingPrice: Clopri,
		Change:       change,
		Transaction:  tra,
		TransTime:    value.TransTime,
	}, nil
}
