package stockinfo_dao

import (
	"context"
	"fmt"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/constants"
	"google.golang.org/grpc/metadata"
)

func (stockinfoDao *StockInfoDao) GetUserStock(ctx context.Context, req *pb.GetUserStockRequest, accessToken string) (*pb.GetUserStockResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := stockinfoDao.client.GetUserStock(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) GetUserStockById(ctx context.Context, req *pb.GetUserStockByIdRequest, accessToken string) (*pb.GetUserStockBuIdResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := stockinfoDao.client.GetUserStockById(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
