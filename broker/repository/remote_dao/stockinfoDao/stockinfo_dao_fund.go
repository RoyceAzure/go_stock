package stockinfo_dao

import (
	"context"
	"fmt"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/constants"
	"google.golang.org/grpc/metadata"
)

func (stockinfoDao *StockInfoDao) GetFund(ctx context.Context, req *pb.GetFundRequest, accessToken string) (*pb.GetFundResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := stockinfoDao.client.GetFund(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) AddFund(ctx context.Context, req *pb.AddFundRequest, accessToken string) (*pb.AddFundResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := stockinfoDao.client.AddFund(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
