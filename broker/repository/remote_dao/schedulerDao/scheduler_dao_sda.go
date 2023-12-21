package scheduler_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
)

func (schedulerDao *SchedulerDao) GetStockDayAvg(ctx context.Context, req *pb.StockDayAvgRequest) (*pb.StockDayAvgResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := schedulerDao.client.GetStockDayAvg(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
