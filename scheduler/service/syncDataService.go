package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/util"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/constants"
	worker "github.com/RoyceAzure/go-stockinfo-schduler/worker"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

const (
	BATCH_SIZE = 2000
)

/*
download data STOCK_DAY_AVG_ALL
TODO: 身分驗證 , insert 失敗資料且入DB
*/

type SyncDataService interface {
	DownloadAndInsertDataSVAA(ctx context.Context) (int64, error)
	SyncStock(ctx context.Context) (int64, []error)
}

type STOCK_DAY_AVG_ALL_DTO struct {
	Code            string `json:"Code"`
	StockName       string `json:"Name"`
	ClosingPrice    string `json:"ClosingPrice"`
	MonthlyAVGPRice string `json:"MonthlyAveragePrice"`
}

func (service *SchdulerService) DownloadAndInsertDataSVAA(ctx context.Context) (int64, error) {
	var dtos []STOCK_DAY_AVG_ALL_DTO
	var entities []repository.BulkInsertDAVGALLParams
	byteData, err := util.SendRequest(constants.METHOD_GET,
		constants.URL_STOCK_DAY_AVG_ALL,
		nil)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(byteData, &dtos)
	if err != nil {
		return 0, err
	}

	var wg sync.WaitGroup

	unprocessed := make(chan []STOCK_DAY_AVG_ALL_DTO)
	processed := make(chan repository.BulkInsertDAVGALLParams)
	wg.Add(3)
	go util.TaskDistributor(unprocessed, BATCH_SIZE, dtos, &wg)
	go util.TaskWorker("worker 1", unprocessed, processed, convertSDAVGALLDTO2BulkEntity, func(err error) {
		log.Warn().Err(err).Msg("err to process data")
	}, &wg)
	go util.TaskWorker("worker 1", unprocessed, processed, convertSDAVGALLDTO2BulkEntity, func(err error) {
		log.Warn().Err(err).Msg("err to process data")
	}, &wg)

	go func() {
		wg.Wait()
		close(processed)
	}()

	insertDatas := int64(0)
	for processData := range processed {
		entities = append(entities, processData)
		if len(entities)%BATCH_SIZE == 0 {
			res, err := service.dao.BulkInsertDAVGALL(ctx, entities)
			if err != nil {
				log.Warn().Err(err).Msg("bulk insert DAVGALL get some error")
				entities = make([]repository.BulkInsertDAVGALLParams, 0)
				continue
			}
			insertDatas += res
			entities = make([]repository.BulkInsertDAVGALLParams, 0)
		}
	}
	if len(entities) > 0 {
		res, err := service.dao.BulkInsertDAVGALL(ctx, entities)
		if err != nil {
			log.Warn().Err(err).Msg("bulk insert DAVGALL get some error")
		} else {
			insertDatas += res
		}
	}
	return insertDatas, nil
}

/*
batchSize 1000
TODO : 若當日沒有資料，要有警示
*/
func (service *SchdulerService) SyncStock(ctx context.Context) (int64, []error) {
	startTime := time.Now().UTC()
	var errs []error
	cr_date_start := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())
	cr_date_end := cr_date_start.AddDate(0, 0, 1).Add(-time.Nanosecond)
	stock_day_alls, err := service.dao.GetSDAVGALLs(ctx, repository.GetSDAVGALLsParams{
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
		errs = append(errs, fmt.Errorf("failed to get SDAVGALL, %w", err))
		return 0, errs
	}
	if len(stock_day_alls) == 0 {
		errs = append(errs, fmt.Errorf("stock_day_alls has no data today"))
		return 0, errs
	}
	var wg sync.WaitGroup
	var payload []worker.BatchUpdateStockPayload

	wg.Add(3)
	unprocessed := make(chan []repository.StockDayAvgAll)
	processed := make(chan worker.BatchUpdateStockPayload)
	batchSize := 1000
	go util.TaskDistributor(unprocessed, batchSize, stock_day_alls, &wg)
	go util.TaskWorker("woker1", unprocessed, processed, cvSDAVGALLEntity2BatchPayload, nil, &wg)
	go util.TaskWorker("worker2", unprocessed, processed, cvSDAVGALLEntity2BatchPayload, nil, &wg)

	go func() {
		wg.Wait()
		close(processed)
	}()

	task_count := 0
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	for batch := range processed {
		payload = append(payload, batch)
		if len(payload)%batchSize == 0 {
			err := service.taskDistributor.DistributeTaskBatchUpdateStock(ctx, payload, opts...)
			if err != nil {
				errs = append(errs, err)
			} else {
				task_count += batchSize
			}
			payload = make([]worker.BatchUpdateStockPayload, 0)
		}
	}
	if len(payload) > 0 {
		err := service.taskDistributor.DistributeTaskBatchUpdateStock(ctx, payload, opts...)
		if err != nil {
			errs = append(errs, err)
		} else {
			task_count += batchSize
		}
	}

	return int64(task_count), errs
}

func convertSDAVGALLDTO2E(dto STOCK_DAY_AVG_ALL_DTO) (repository.CreateSDAVGALLParams, error) {
	var result repository.CreateSDAVGALLParams
	if dto.ClosingPrice == "" {
		dto.ClosingPrice = constants.STR_ZERO
	}
	if dto.MonthlyAVGPRice == "" {
		dto.MonthlyAVGPRice = constants.STR_ZERO
	}
	var cp, mp pgtype.Numeric
	err := cp.Scan(dto.ClosingPrice)
	if err != nil {
		return result, err
	}
	err = mp.Scan(dto.ClosingPrice)
	if err != nil {
		return result, err
	}
	return repository.CreateSDAVGALLParams{
		Code:            dto.Code,
		StockName:       dto.StockName,
		ClosePrice:      cp,
		MonthlyAvgPrice: mp,
	}, nil
}

func convertSDAVGALLDTO2BulkEntity(dto STOCK_DAY_AVG_ALL_DTO) (repository.BulkInsertDAVGALLParams, error) {
	var result repository.BulkInsertDAVGALLParams
	if dto.ClosingPrice == "" {
		dto.ClosingPrice = constants.STR_ZERO
	}
	if dto.MonthlyAVGPRice == "" {
		dto.MonthlyAVGPRice = constants.STR_ZERO
	}
	var cp, mp pgtype.Numeric
	err := cp.Scan(dto.ClosingPrice)
	if err != nil {
		return result, err
	}
	err = mp.Scan(dto.ClosingPrice)
	if err != nil {
		return result, err
	}
	return repository.BulkInsertDAVGALLParams{
		Code:            dto.Code,
		StockName:       dto.StockName,
		ClosePrice:      cp,
		MonthlyAvgPrice: mp,
	}, nil
}

/*
BatchUpdateStockPayload.MarketCap 目前寫死10000
*/
func cvSDAVGALLEntity2BatchPayload(entity repository.StockDayAvgAll) (worker.BatchUpdateStockPayload, error) {
	var result worker.BatchUpdateStockPayload
	close_price, err := entity.ClosePrice.Float64Value()
	if err != nil {
		return result, fmt.Errorf("error to convert stock_day_avg_all to batch_update_stoc_payload")
	}

	return worker.BatchUpdateStockPayload{
		StockCode:    entity.Code,
		StockName:    entity.StockName,
		CurrentPrice: strconv.FormatFloat(close_price.Float64, 'f', -1, 64),
		MarketCap:    10000,
	}, nil
}
