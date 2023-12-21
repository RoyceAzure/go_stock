package stockinfo_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
)

func (stockinfoDao *StockInfoDao) GetFund(ctx context.Context, req *pb.GetFundRequest, accessToken string) (*pb.GetFundResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, accessToken)
	res, err := stockinfoDao.client.GetFund(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) AddFund(ctx context.Context, req *pb.AddFundRequest, accessToken string) (*pb.AddFundResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, accessToken)
	res, err := stockinfoDao.client.AddFund(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
