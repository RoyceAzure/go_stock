package stockinfo_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
)

func (stockinfoDao *StockInfoDao) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	res, err := stockinfoDao.client.ValidateToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stockinfoDao *StockInfoDao) RenewToken(ctx context.Context, req *pb.RenewTokenRequest) (*pb.RenewTokenResponse, error) {
	res, err := stockinfoDao.client.RenewToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
