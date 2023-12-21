package stockinfo_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"

)

func (stockinfoDao *StockInfoDao) GetRealizedProfitLoss(ctx context.Context, req *pb.GetRealizedProfitLossRequest, accessToken string) (*pb.GetRealizedProfitLossResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, accessToken)
	res, err := stockinfoDao.client.GetRealizedProfitLoss(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) GetUnRealizedProfitLoss(ctx context.Context, req *pb.GetUnRealizedProfitLossRequest, accessToken string) (*pb.GetUnRealizedProfitLossResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, accessToken)
	res, err := stockinfoDao.client.GetUnRealizedProfitLoss(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
