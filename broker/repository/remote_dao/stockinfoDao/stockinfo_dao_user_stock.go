package stockinfo_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
)

func (stockinfoDao *StockInfoDao) GetUserStock(ctx context.Context, req *pb.GetUserStockRequest, accessToken string) (*pb.GetUserStockResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, accessToken)
	res, err := stockinfoDao.client.GetUserStock(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) GetUserStockById(ctx context.Context, req *pb.GetUserStockByIdRequest, accessToken string) (*pb.GetUserStockBuIdResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, accessToken)
	res, err := stockinfoDao.client.GetUserStockById(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
