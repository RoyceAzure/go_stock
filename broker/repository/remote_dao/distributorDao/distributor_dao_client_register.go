package distributor_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
)

func (distributorDao *DistributorDao) CreateClientRegister(ctx context.Context, req *pb.CreateClientRegisterRequest) (*pb.CreateClientRegisterResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := distributorDao.client.CreateClientRegister(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (distributorDao *DistributorDao) DeleteClientRegister(ctx context.Context, req *pb.DeleteClientRegisterRequest) (*pb.DeleteClientRegisterResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := distributorDao.client.DeleteClientRegister(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (distributorDao *DistributorDao) GetClientRegisterByClientUID(ctx context.Context, req *pb.GetClientRegisterByClientUIDRequest) (*pb.GetClientRegisterResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := distributorDao.client.GetClientRegisterByClientUID(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
