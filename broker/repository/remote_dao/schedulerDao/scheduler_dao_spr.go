package scheduler_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
)

func (schedulerDao *SchedulerDao) GetStockPriceRealTime(ctx context.Context, req *pb.StockPriceRealTimeRequest) (*pb.StockPriceRealTimeResponse, error) {
	res, err := schedulerDao.client.GetStockPriceRealTime(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
