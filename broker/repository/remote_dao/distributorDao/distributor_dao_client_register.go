package distributor_dao

import (
	"context"
	"fmt"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/constants"
	"google.golang.org/grpc/metadata"
)

func (distributorDao *DistributorDao) CreateClientRegister(ctx context.Context, req *pb.CreateClientRegisterRequest, accessToken string) (*pb.CreateClientRegisterResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := distributorDao.client.CreateClientRegister(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (distributorDao *DistributorDao) DeleteClientRegister(ctx context.Context, req *pb.DeleteClientRegisterRequest, accessToken string) (*pb.DeleteClientRegisterResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := distributorDao.client.DeleteClientRegister(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (distributorDao *DistributorDao) GetClientRegisterByClientUID(ctx context.Context, req *pb.GetClientRegisterByClientUIDRequest, accessToken string) (*pb.GetClientRegisterResponse, error) {
	md := metadata.New(map[string]string{
		constants.AuthorizationHeaderKey: fmt.Sprintf("%s %s", constants.AuthorizationTypeBearer, accessToken),
	})
	newCtx := metadata.NewOutgoingContext(ctx, md)
	res, err := distributorDao.client.GetClientRegisterByClientUID(newCtx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
