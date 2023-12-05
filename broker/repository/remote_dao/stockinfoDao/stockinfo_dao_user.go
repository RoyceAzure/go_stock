package stockinfo_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
)

func (stockinfoDao *StockInfoDao) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := stockinfoDao.client.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (stockinfoDao *StockInfoDao) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	res, err := stockinfoDao.client.UpdateUser(ctx, req)
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
