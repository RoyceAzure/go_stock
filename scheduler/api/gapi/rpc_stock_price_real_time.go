package gapi

import (
	"context"
	"sync"
	"time"

	"github.com/RoyceAzure/go-stockinfo-scheduler/api/pb"
	redisService "github.com/RoyceAzure/go-stockinfo-scheduler/service/redisService"
	"github.com/RoyceAzure/go-stockinfo-scheduler/util"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

)

/*
TODO : 驗證
*/
func (server *Server) GetStockPriceRealTime(ctx context.Context, req *pb.StockPriceRealTimeRequest) (*pb.StockPriceRealTimeResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			// 记录错误信息，可以使用日志库
			log.Error().Err(err.(error)).
				Msg("Recovered from error")
			// 返回一个通用的错误响应给客户端
		}
	}()
	log.Info().Msg("get stock price realtime start")
	startTime := time.Now().UTC()

	sprs, keyTime, err := server.redisService.GetLatestSPR(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get stock price real time : %s", err)
	}
	if len(sprs) == 0 {
		return nil, status.Error(codes.Internal, "there is no stock price real time data")
	}

	var resultList []*pb.StockPriceRealTime

	var wg sync.WaitGroup
	batchSize := 2000
	unporcessed := make(chan []redisService.StockPriceRealtimeDTO)
	processed := make(chan *pb.StockPriceRealTime, batchSize/2)
	wg.Add(4)
	go util.TaskDistributor(unporcessed, batchSize, sprs, &wg)
	go util.TaskWorker("woker1", unporcessed, processed, cvSprDTO2res, nil, &wg)
	go util.TaskWorker("woker2", unporcessed, processed, cvSprDTO2res, nil, &wg)
	go util.TaskWorker("woker3", unporcessed, processed, cvSprDTO2res, nil, &wg)
	go func() {
		wg.Wait()
		close(processed)
	}()
	for data := range processed {
		resultList = append(resultList, data)
	}

	log.Info().Dur("elapse time", time.Duration(time.Now().UTC().Sub(startTime))).Msg("get stock price realtime end")

	return &pb.StockPriceRealTimeResponse{
		KeyTime: keyTime,
		Result:  resultList,
	}, nil
}

func cvSprDTO2res(value redisService.StockPriceRealtimeDTO) (*pb.StockPriceRealTime, error) {
	trade_vol := decimal.Decimal(value.TradeVolume)
	trade_val := decimal.Decimal(value.TradeValue)
	open_price := decimal.Decimal(value.OpeningPrice)
	highest_price := decimal.Decimal(value.HighestPrice)
	lowest_price := decimal.Decimal(value.LowestPrice)
	close_price := decimal.Decimal(value.ClosingPrice)
	change := decimal.Decimal(value.Change)
	transation := decimal.Decimal(value.Transaction)
	return &pb.StockPriceRealTime{
		StockCode:    value.Code,
		StockName:    value.StockName,
		TradeVolume:  trade_vol.String(),
		TradeValue:   trade_val.String(),
		OpenPrice:    open_price.String(),
		HighestPrice: highest_price.String(),
		LowestPrice:  lowest_price.String(),
		ClosePrice:   close_price.String(),
		Change:       change.String(),
		Transaction:  transation.String(),
		TransTime:    timestamppb.New(value.TransTime),
	}, nil
}
