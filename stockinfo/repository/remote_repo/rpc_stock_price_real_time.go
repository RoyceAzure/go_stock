package repository

import (
	"context"
	"fmt"
	"time"

	logger "github.com/RoyceAzure/go-stockinfo/repository/logger_distributor"
	pb "github.com/RoyceAzure/go-stockinfo/shared/pb"
)

const SPR_SAVED_DUR = time.Second * 5

/*
從schduler 取spr資料，若取不到，將會從cache拿備份回傳
*/
func (dao *JSchdulerInfoDao) GetStockPriceRealTime(ctx context.Context) (SprData, error) {
	select {
	case <-ctx.Done():
		return SprData{}, ctx.Err() // ctx.Err() 将是 context.DeadlineExceeded 如果超时
	default:
		// 正常处理请求
		logger.Logger.Info().Msg("start get stock price realtime")
		now := time.Now().UTC()
		if now.Sub(dao.preSprTime) >= SPR_SAVED_DUR {
			res, err := dao.client.GetStockPriceRealTime(ctx, &pb.StockPriceRealTimeRequest{})
			if err == nil {
				sprData := cvSprRes2SprData(res)
				dao.SetSprData(ctx, sprData)
				dao.preSprTime = now
				logger.Logger.Info().Msg("end get stock price realtime")
				return sprData, nil
			}
		}
		sprData := dao.GetSprData(ctx)
		if sprData.DataTime == "" {
			return SprData{}, fmt.Errorf("get stock price real time get err")
		}
		logger.Logger.Info().Msg("end get stock price realtime")
		return sprData, nil
	}
}

func cvSprRes2SprData(value *pb.StockPriceRealTimeResponse) SprData {
	return SprData{
		DataTime: value.GetKeyTime(),
		Data:     value.GetResult(),
	}
}
