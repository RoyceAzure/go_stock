package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	dto "github.com/RoyceAzure/go-stockinfo-distributor/shared/model/dto"
	pb "github.com/RoyceAzure/go-stockinfo-distributor/shared/pb/stock_info_scheduler"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

var SprAlreadyRetrives error = errors.New("spr already retirves")

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
	sprCache := s.schdulerDao.GetSprCache(ctx)
	if sprCache == nil {
		return nil, fmt.Errorf("sprCache is empty")
	}

	var res []dto.StockPriceRealTimeDTO

	batchSize := 2000

	upprodessed := make(chan []*pb.StockPriceRealTime)
	processed := make(chan dto.StockPriceRealTimeDTO)

	var wg sync.WaitGroup
	wg.Add(3)
	go util.SliceBatchIterator(upprodessed, batchSize, sprCache.Result, []func([]*pb.StockPriceRealTime) []*pb.StockPriceRealTime{
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

func (s *DistributorService) GetPreSprtime(ctx context.Context) string {
	s.sprLock.RLock()
	defer s.sprLock.RUnlock()
	return s.preSprtime
}

func (s *DistributorService) SetPreSprtime(ctx context.Context, preSprtime string) {
	s.sprLock.Lock()
	defer s.sprLock.Unlock()
	s.preSprtime = preSprtime
}

/*
必須是有狀態的  因為要避免塞入重複資料造成效能損耗

會對應多個client  取出所有註冊的distinct stock code，撈取對應資料  放入kafka
*/
func (s *DistributorService) GetFilterSPRByIPAndSendToKa(ctx context.Context, ip string) error {
	log.Info().Msg("start get filter spr by ip")

	preDataTime := s.GetPreSprtime(ctx)

	sprCache := s.schdulerDao.GetSprCache(ctx)
	if sprCache == nil {
		return fmt.Errorf("sprCache is empty")
	}

	if preDataTime == sprCache.ResultTimeStr {
		return SprAlreadyRetrives
	}

	//get frontend register stock
	regisStockCodes, err := s.dbDao.GetDistinctStockCode(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	//make map to filter
	var stockCodeMap map[string]bool
	for _, item := range regisStockCodes {
		stockCodeMap[item] = true
	}

	//get and filter spr

	var res []kafka.Message

	batchSize := 2000

	upprodessed := make(chan []*pb.StockPriceRealTime)
	processed := make(chan kafka.Message)

	var wg sync.WaitGroup
	wg.Add(3)
	go util.SliceBatchIterator(upprodessed, batchSize, sprCache.Result, []func([]*pb.StockPriceRealTime) []*pb.StockPriceRealTime{
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

	go util.TaskWorker("worker1", upprodessed, processed, cvSpr2KafkaMsg, errfunc, &wg, sprCache.ResultTimeStr)
	go util.TaskWorker("worker2", upprodessed, processed, cvSpr2KafkaMsg, errfunc, &wg, sprCache.ResultTimeStr)

	go func() {
		wg.Wait()
		close(processed)
	}()

	for item := range processed {
		res = append(res, item)
	}

	err = s.jkafkaWrite.WriteMessages(ctx, res)
	if err != nil {
		return err
	}

	s.schdulerDao.SetPreSprTime(ctx, sprCache.ResultTimeStr)

	log.Info().Msg("end get filter spr by ip")
	return nil
}

func cvSpr2DTO(value *pb.StockPriceRealTime, parms ...any) (dto.StockPriceRealTimeDTO, error) {
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

/*
parms must pass spr time, and only spr time

	重每個dto裡面取出code  分組, topic就是組名稱 key是code  組合成[]kafka messafe
	秘等操作  過濾已經有的資料時間
*/
func cvSpr2KafkaMsg(value *pb.StockPriceRealTime, parms ...any) (kafka.Message, error) {
	var res kafka.Message

	if len(parms) == 0 {
		return res, fmt.Errorf("spr time is empty")
	}

	dto := dto.StockPriceRealTimeDTO{
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
	}

	msgValue, err := json.Marshal(dto)
	if err != nil {
		return res, fmt.Errorf("marsh spr 2 kafka msg get err : %w", err)
	}
	sprTime, ok := parms[0].(string)
	if !ok {
		return res, fmt.Errorf("spr time must be string")
	}

	key := value.StockCode

	keyValue, err := json.Marshal(key)
	if err != nil {
		return res, fmt.Errorf("marsh spr 2 kafka msg get err : %w", err)
	}

	topic := sprTime + string([]rune(key)[:2])

	res.Topic = topic
	res.Key = keyValue
	res.Value = msgValue
	return res, nil
}
