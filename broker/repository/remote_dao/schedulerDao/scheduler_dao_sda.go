package scheduler_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
)

func (schedulerDao *SchedulerDao) GetStockDayAvg(ctx context.Context, req *pb.StockDayAvgRequest) (*pb.StockDayAvgResponse, error) {
	res, err := schedulerDao.client.GetStockDayAvg(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
