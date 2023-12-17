package distributor_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"

)

func (distributorDao *DistributorDao) CreateFrontendClient(ctx context.Context, req *pb.CreateFrontendClientRequest) (*pb.CreateFrontendClientResponse, error) {
	res, err := distributorDao.client.CreateFrontendClient(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (distributorDao *DistributorDao) DeleteFrontendClient(ctx context.Context, req *pb.DeleteFrontendClientRequest) (*pb.DeleteFrontendClientResponse, error) {
	res, err := distributorDao.client.DeleteFrontendClient(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (distributorDao *DistributorDao) GetFrontendClientByIP(ctx context.Context, req *pb.GetFrontendClientByIPRequest) (*pb.GetFrontendClientByIPResponse, error) {
	res, err := distributorDao.client.GetFrontendClientByIP(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
