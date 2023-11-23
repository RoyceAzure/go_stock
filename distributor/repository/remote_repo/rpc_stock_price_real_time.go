package repository

import (
	"context"
	"fmt"

	pb "github.com/RoyceAzure/go-stockinfo-distributor/shared/pb/stock_info_scheduler"
	"github.com/rs/zerolog/log"
)

func (dao *JSchdulerInfoDao) GetStockPriceRealTime(ctx context.Context) (*pb.StockPriceRealTimeResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err() // ctx.Err() 将是 context.DeadlineExceeded 如果超时
	default:
		// 正常处理请求
		log.Info().Msg("start get stock price realtime")
		req := &pb.StockPriceRealTimeRequest{}

		res, err := dao.client.GetStockPriceRealTime(ctx, req)
		if err != nil {
			cache := dao.GetSprCache()
			if cache == nil {
				return nil, fmt.Errorf("get stock price real time get err : %w", err)
			}
			log.Warn().Msg("get stock price realtime from cache")
			res.Result = cache
			return res, nil
		}
		dao.SetSprCache(res.Result)
		log.Info().Msg("end get stock price realtime")
		return res, nil
	}
}
