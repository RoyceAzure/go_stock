package service

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/util"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/constants"
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
	SyncFund(ctx context.Context) (int64, error)
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

func (service *SchdulerService) SyncFund(ctx context.Context) (int64, error) {
	return 0, errors.New("not implement")
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
