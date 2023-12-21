package distributor_dao

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
)

func (distributorDao *DistributorDao) CreateFrontendClient(ctx context.Context, req *pb.CreateFrontendClientRequest) (*pb.CreateFrontendClientResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := distributorDao.client.CreateFrontendClient(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (distributorDao *DistributorDao) DeleteFrontendClient(ctx context.Context, req *pb.DeleteFrontendClientRequest) (*pb.DeleteFrontendClientResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := distributorDao.client.DeleteFrontendClient(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (distributorDao *DistributorDao) GetFrontendClientByIP(ctx context.Context, req *pb.GetFrontendClientByIPRequest) (*pb.GetFrontendClientByIPResponse, error) {
	newCtx := util.NewOutGoingMetaData(ctx, "")
	res, err := distributorDao.client.GetFrontendClientByIP(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
