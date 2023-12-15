package stockinfo_dao

import (
	"context"
	"fmt"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/constants"
	"google.golang.org/grpc/metadata"
)

func (stockinfoDao *StockInfoDao) GetStock(ctx context.Context, req *pb.GetStockRequest) (*pb.GetStockResponse, error) {
	res, err := stockinfoDao.client.GetStock(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) GetStocks(ctx context.Context, req *pb.GetStocksRequest, accessToken string) (*pb.GetStocksResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := stockinfoDao.client.GetStocks(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) InitStock(ctx context.Context, req *pb.InitStockRequest) (*pb.InitStockResponse, error) {
	res, err := stockinfoDao.client.InitStock(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
