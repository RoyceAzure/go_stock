package gapi

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
)

func (s *SchedulerServer) GetStockPriceRealTime(ctx context.Context, req *pb.StockPriceRealTimeRequest) (*pb.StockPriceRealTimeResponse, error) {
	_, _, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	return s.schedulerDao.GetStockPriceRealTime(ctx, req)
}

func (s *SchedulerServer) GetStockDayAvg(ctx context.Context, req *pb.StockDayAvgRequest) (*pb.StockDayAvgResponse, error) {
	_, _, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	return s.schedulerDao.GetStockDayAvg(ctx, req)
}
