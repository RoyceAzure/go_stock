package distributor_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
)

func (distributorDao *DistributorDao) CreateClientRegister(ctx context.Context, req *pb.CreateClientRegisterRequest) (*pb.CreateClientRegisterResponse, error) {
	res, err := distributorDao.client.CreateClientRegister(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (distributorDao *DistributorDao) DeleteClientRegister(ctx context.Context, req *pb.DeleteClientRegisterRequest) (*pb.DeleteClientRegisterResponse, error) {
	res, err := distributorDao.client.DeleteClientRegister(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (distributorDao *DistributorDao) GetClientRegisterByClientUID(ctx context.Context, req *pb.GetClientRegisterByClientUIDRequest) (*pb.GetClientRegisterResponse, error) {
	res, err := distributorDao.client.GetClientRegisterByClientUID(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
