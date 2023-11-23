package service

import (
	"context"
	"fmt"
	"sync"

	dto "github.com/RoyceAzure/go-stockinfo-distributor/shared/model/dto"
	pb "github.com/RoyceAzure/go-stockinfo-distributor/shared/pb/stock_info_scheduler"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util"
	"github.com/rs/zerolog/log"
)

/*
Get frontend register spr data
Params:

	ip - frontend client ip

Returns:

	spr的dto slice
*/
func (s *DistributorService) GetFilterSPRByIP(ctx context.Context, ip string) ([]dto.StockPriceRealTimeDTO, error) {
	log.Info().Msg("start get filter spr by ip")

	//get frontend register stock
	client, err := s.dbDao.GetFrontendClientByIP(ctx, ip)
	if err != nil {
		return nil, fmt.Errorf("client ip is not exists, %w", err)
	}

	clientRegister, err := s.dbDao.GetClientRegisterByClientUID(ctx, client.ClientUid)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	//make map to filter
	var stockCodeMap map[string]bool
	for _, item := range clientRegister {
		stockCodeMap[item.StockCode] = true
	}

	//get and filter spr
	sprCache := s.schdulerDao.GetSprCache()
	if sprCache == nil {
		return nil, fmt.Errorf("sprCache is empty")
	}

	var res []dto.StockPriceRealTimeDTO

	batchSize := 2000

	upprodessed := make(chan []*pb.StockPriceRealTime)
	processed := make(chan dto.StockPriceRealTimeDTO)

	var wg sync.WaitGroup
	wg.Add(3)
	go util.SliceBatchIterator(upprodessed, batchSize, sprCache, []func([]*pb.StockPriceRealTime) []*pb.StockPriceRealTime{
		func(sprcahce []*pb.StockPriceRealTime) []*pb.StockPriceRealTime {
			var result []*pb.StockPriceRealTime
			length := len(sprcahce)
			for i := 0; i < length; i++ {
				if stockCodeMap[sprcahce[i].StockCode] {
					result = append(result, sprcahce[i])
				}
			}
			return result
		},
	})

	errfunc := func(err error) {
		log.Warn().Err(err)
	}

	go util.TaskWorker("worker1", upprodessed, processed, cvSpr2DTO, errfunc, &wg)
	go util.TaskWorker("worker2", upprodessed, processed, cvSpr2DTO, errfunc, &wg)

	go func() {
		wg.Wait()
		close(processed)
	}()

	for item := range processed {
		res = append(res, item)
	}

	log.Info().Msg("end get filter spr by ip")
	return res, nil
}

func cvSpr2DTO(value *pb.StockPriceRealTime) (dto.StockPriceRealTimeDTO, error) {
	return dto.StockPriceRealTimeDTO{
		StockCode:    value.StockCode,
		StockName:    value.StockName,
		TradeVolume:  value.TradeVolume,
		TradeValue:   value.TradeValue,
		OpenPrice:    value.OpenPrice,
		HighestPrice: value.HighestPrice,
		LowestPrice:  value.LowestPrice,
		ClosePrice:   value.ClosePrice,
		Change:       value.Change,
		Transaction:  value.Transaction,
		TransTime:    value.TransTime.AsTime(),
	}, nil
}
