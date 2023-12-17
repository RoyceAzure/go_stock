package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	logger "github.com/RoyceAzure/go-stockinfo-distributor/repository/logger_distributor"
	dto "github.com/RoyceAzure/go-stockinfo-distributor/shared/model/dto"
	pb "github.com/RoyceAzure/go-stockinfo-distributor/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

var ErrAlreadyRetrieved error = errors.New("spr already retirves")

type SPRKafkaVO struct {
	DataTime string
	Data     dto.StockPriceRealTimeDTO
}

/*
Get frontend register spr data
Params:

	ip - frontend client ip

Returns:

	spr的dto slice
*/
func (s *DistributorService) GetFilterSPRByIP(ctx context.Context, ip string) ([]dto.StockPriceRealTimeDTO, error) {
	logger.Logger.Info().Msg("start get filter spr by ip")

	//get frontend register stock
	client, err := s.dbDao.GetFrontendClientByIP(ctx, ip)
	if err != nil {
		return nil, fmt.Errorf("client ip is not exists, %w", err)
	}

	crs, err := s.dbDao.GetClientRegisterByClientUID(ctx, client.ClientUid)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	//make map to filter
	stockCodeMap := make(map[string]struct{})
	for _, cr := range crs {
		stockCodeMap[cr.StockCode] = struct{}{}
	}
	//get and filter spr
	sprCache := s.schdulerDao.GetSprData(ctx)
	if sprCache.DataTime == "" {
		return nil, fmt.Errorf("sprCache is empty")
	}

	var res []dto.StockPriceRealTimeDTO

	batchSize := 2000

	upprodessed := make(chan []*pb.StockPriceRealTime)
	processed := make(chan dto.StockPriceRealTimeDTO)

	var wg sync.WaitGroup
	wg.Add(3)
	go util.TaskDistributor(upprodessed, batchSize, sprCache.Data, []func([]*pb.StockPriceRealTime) []*pb.StockPriceRealTime{
		func(sprcahce []*pb.StockPriceRealTime) []*pb.StockPriceRealTime {
			var result []*pb.StockPriceRealTime
			length := len(sprcahce)
			for i := 0; i < length; i++ {
				if _, exists := stockCodeMap[sprcahce[i].StockCode]; exists {
					result = append(result, sprcahce[i])
				}
			}
			return result
		},
	}, &wg)

	errfunc := func(err error) {
		logger.Logger.Warn().Err(err)
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

	logger.Logger.Info().Msg("end get filter spr by ip")
	return res, nil
}

func (s *DistributorService) GetPreSuccessedSprtime(ctx context.Context) string {
	s.sprLock.RLock()
	defer s.sprLock.RUnlock()
	return s.preSprtime
}

func (s *DistributorService) SetPreSuccessedSprtime(ctx context.Context, preSprtime string) {
	s.sprLock.Lock()
	defer s.sprLock.Unlock()
	s.preSprtime = preSprtime
}

/*
透過schdulerDao取資料，無法得知資料是最新從sdchduler回傳，還是cache裡的資料

必須是有狀態的  因為要避免塞入重複資料造成效能損耗

會對應多個client  取出所有註冊的distinct stock code，撈取對應資料  放入kafka
*/
func (s *DistributorService) GetAllRegisStockAndSendToKa(ctx context.Context) error {
	logger.Logger.Info().Msg("start get filter spr by ip and send to ka")
	var sprDatas []*pb.StockPriceRealTime
	var sprTime string

	sprRes, err := s.schdulerDao.GetStockPriceRealTime(ctx)
	if err != nil {
		return fmt.Errorf("get spr from schduler failed%w", err)
	}
	preDataTime := s.GetPreSuccessedSprtime(ctx)

	if preDataTime == sprRes.DataTime {
		log.Info().Msg("spr already retrives!!")
		return nil
	}
	sprDatas = sprRes.Data
	sprTime = sprRes.DataTime

	//get frontend register stock
	regisStockCodes, err := s.dbDao.GetDistinctStockCode(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	//make map to filter
	stockCodeMap := make(map[string]struct{})
	for _, item := range regisStockCodes {
		stockCodeMap[item] = struct{}{}
	}

	//get and filter spr

	var res []kafka.Message

	batchSize := 5000

	upprodessed := make(chan []*pb.StockPriceRealTime, 10)
	processed := make(chan kafka.Message, 10)

	var wg sync.WaitGroup
	wg.Add(3)
	go util.TaskDistributor(upprodessed, batchSize, sprDatas, []func([]*pb.StockPriceRealTime) []*pb.StockPriceRealTime{
		func(sprcahce []*pb.StockPriceRealTime) []*pb.StockPriceRealTime {
			var result []*pb.StockPriceRealTime
			length := len(sprcahce)
			for i := 0; i < length; i++ {
				if _, exists := stockCodeMap[sprcahce[i].StockCode]; exists {
					result = append(result, sprcahce[i])
				}
			}
			return result
		},
	}, &wg)

	errfunc := func(err error) {
		logger.Logger.Warn().Err(err)
	}

	go util.TaskWorker("worker1", upprodessed, processed, cvSpr2KafkaMsg, errfunc, &wg, sprTime)
	go util.TaskWorker("worker2", upprodessed, processed, cvSpr2KafkaMsg, errfunc, &wg, sprTime)

	go func() {
		wg.Wait()
		close(processed)
	}()

	for item := range processed {
		res = append(res, item)
	}

	if len(res) > 0 {
		err = s.jkafkaWrite.WriteMessages(ctx, res)
		if err != nil {
			return err
		}

		s.SetPreSuccessedSprtime(ctx, sprTime)
	}

	logger.Logger.Info().Msg("end get filter spr by ip and send to ka")
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
		return res, fmt.Errorf("spr data time is empty")
	}
	sprTime := parms[0].(string)
	dataTime := strings.Replace(strings.Split(sprTime, "_")[1], ":", "", -1)
	vo := SPRKafkaVO{
		DataTime: dataTime,
		Data: dto.StockPriceRealTimeDTO{
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
		},
	}

	msgValue, err := json.Marshal(vo)
	if err != nil {
		return res, fmt.Errorf("marsh spr 2 kafka msg get err : %w", err)
	}

	key := value.StockCode

	keyValue, err := json.Marshal(key)
	if err != nil {
		return res, fmt.Errorf("marsh spr 2 kafka msg get err : %w", err)
	}

	topic := string([]rune(key)[:2])

	// time := strings.Replace(strings.Split(sprTime, "_")[1], ":", "", -1)

	res.Topic = topic
	res.Key = keyValue
	res.Value = msgValue
	return res, nil
}
