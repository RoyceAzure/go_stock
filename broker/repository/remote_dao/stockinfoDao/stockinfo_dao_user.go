package stockinfo_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
)

func (stockinfoDao *StockInfoDao) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := stockinfoDao.client.CreateUser(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) GetUser(ctx context.Context, req *pb.GetUserRequest, accessToken string) (*pb.GetUserResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, accessToken)
	res, err := stockinfoDao.client.GetUser(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest, accessToken string) (*pb.UpdateUserResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, accessToken)
	res, err := stockinfoDao.client.UpdateUser(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (stockinfoDao *StockInfoDao) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := stockinfoDao.client.LoginUser(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (stockinfoDao *StockInfoDao) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := stockinfoDao.client.VerifyEmail(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
