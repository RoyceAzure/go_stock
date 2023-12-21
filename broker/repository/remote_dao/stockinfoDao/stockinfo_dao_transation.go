package stockinfo_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
)

func (stockinfoDao *StockInfoDao) TransationStock(ctx context.Context, req *pb.TransationRequest, accessToken string) (*pb.TransationResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, accessToken)
	res, err := stockinfoDao.client.TransationStock(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) GetAllTransations(ctx context.Context, req *pb.GetAllStockTransationRequest, accessToken string) (*pb.StockTransatsionResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, accessToken)
	res, err := stockinfoDao.client.GetAllTransations(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
