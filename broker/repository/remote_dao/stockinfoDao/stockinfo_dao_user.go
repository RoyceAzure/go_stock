package stockinfo_dao

import (
	"context"
	"fmt"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/constants"
	"google.golang.org/grpc/metadata"
)

func (stockinfoDao *StockInfoDao) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := stockinfoDao.client.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) GetUser(ctx context.Context, req *pb.GetUserRequest, accessToken string) (*pb.GetUserResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := stockinfoDao.client.GetUser(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest, accessToken string) (*pb.UpdateUserResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := stockinfoDao.client.UpdateUser(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (stockinfoDao *StockInfoDao) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	res, err := stockinfoDao.client.LoginUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (stockinfoDao *StockInfoDao) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	res, err := stockinfoDao.client.VerifyEmail(ctx, req)
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
