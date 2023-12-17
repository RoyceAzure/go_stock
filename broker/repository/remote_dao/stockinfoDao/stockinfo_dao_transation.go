package stockinfo_dao

import (
	"context"
	"fmt"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/constants"
	"google.golang.org/grpc/metadata"
)

func (stockinfoDao *StockInfoDao) TransationStock(ctx context.Context, req *pb.TransationRequest, accessToken string) (*pb.TransationResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := stockinfoDao.client.TransationStock(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) GetAllTransations(ctx context.Context, req *pb.GetAllStockTransationRequest, accessToken string) (*pb.StockTransatsionResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := stockinfoDao.client.GetAllTransations(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
