package gapi

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
)

func (s *SchedulerServer) GetStockPriceRealTime(ctx context.Context, req *pb.StockPriceRealTimeRequest) (*pb.StockPriceRealTimeResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	_, err = s.stockinfoDao.ValidateToken(ctx, &pb.ValidateTokenRequest{
		AccessToken: token,
	})
	if err != nil {
		return nil, err
	}
	return s.schedulerDao.GetStockPriceRealTime(ctx, req)
}

func (s *SchedulerServer) GetStockDayAvg(ctx context.Context, req *pb.StockDayAvgRequest) (*pb.StockDayAvgResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	_, err = s.stockinfoDao.ValidateToken(ctx, &pb.ValidateTokenRequest{
		AccessToken: token,
	})
	if err != nil {
		return nil, err
	}
	return s.schedulerDao.GetStockDayAvg(ctx, req)
}
