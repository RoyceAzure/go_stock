package stockinfo_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
)

func (stockinfoDao *StockInfoDao) GetStock(ctx context.Context, req *pb.GetStockRequest) (*pb.GetStockResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := stockinfoDao.client.GetStock(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) GetStocks(ctx context.Context, req *pb.GetStocksRequest, accessToken string) (*pb.GetStocksResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, accessToken)
	res, err := stockinfoDao.client.GetStocks(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) InitStock(ctx context.Context, req *pb.InitStockRequest) (*pb.InitStockResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := stockinfoDao.client.InitStock(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
