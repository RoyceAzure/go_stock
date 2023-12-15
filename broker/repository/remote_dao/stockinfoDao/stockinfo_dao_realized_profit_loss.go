package stockinfo_dao

import (
	"context"
	"fmt"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/constants"
	"google.golang.org/grpc/metadata"
)

func (stockinfoDao *StockInfoDao) GetRealizedProfitLoss(ctx context.Context, req *pb.GetRealizedProfitLossRequest, accessToken string) (*pb.GetRealizedProfitLossResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := stockinfoDao.client.GetRealizedProfitLoss(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) GetUnRealizedProfitLoss(ctx context.Context, req *pb.GetUnRealizedProfitLossRequest, accessToken string) (*pb.GetUnRealizedProfitLossResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := stockinfoDao.client.GetUnRealizedProfitLoss(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
