package scheduler_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
)

func (schedulerDao *SchedulerDao) GetStockPriceRealTime(ctx context.Context, req *pb.StockPriceRealTimeRequest) (*pb.StockPriceRealTimeResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := schedulerDao.client.GetStockPriceRealTime(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
