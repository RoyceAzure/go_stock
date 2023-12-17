package distributor_dao

import (
	"context"
	"fmt"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/constants"
	"google.golang.org/grpc/metadata"
)

func (distributorDao *DistributorDao) CreateFrontendClient(ctx context.Context, req *pb.CreateFrontendClientRequest, accessToken string) (*pb.CreateFrontendClientResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := distributorDao.client.CreateFrontendClient(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (distributorDao *DistributorDao) DeleteFrontendClient(ctx context.Context, req *pb.DeleteFrontendClientRequest, accessToken string) (*pb.DeleteFrontendClientResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := distributorDao.client.DeleteFrontendClient(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (distributorDao *DistributorDao) GetFrontendClientByIP(ctx context.Context, req *pb.GetFrontendClientByIPRequest, accessToken string) (*pb.GetFrontendClientByIPResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := distributorDao.client.GetFrontendClientByIP(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
