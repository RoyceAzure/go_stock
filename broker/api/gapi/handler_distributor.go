package gapi

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
)

func (s *DistributorServer) CreateClientRegister(ctx context.Context, req *pb.CreateClientRegisterRequest) (*pb.CreateClientRegisterResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	_, err = s.stockinfoDao.ValidateToken(ctx, &pb.ValidateTokenRequest{
		AccessToken: token,
	})
	if err != nil {
		return nil, err
	}
	return s.distributorDao.CreateClientRegister(ctx, req)
}
func (s *DistributorServer) DeleteClientRegister(ctx context.Context, req *pb.DeleteClientRegisterRequest) (*pb.DeleteClientRegisterResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	_, err = s.stockinfoDao.ValidateToken(ctx, &pb.ValidateTokenRequest{
		AccessToken: token,
	})
	if err != nil {
		return nil, err
	}
	return s.distributorDao.DeleteClientRegister(ctx, req)
}
func (s *DistributorServer) GetClientRegisterByClientUID(ctx context.Context, req *pb.GetClientRegisterByClientUIDRequest) (*pb.GetClientRegisterResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	_, err = s.stockinfoDao.ValidateToken(ctx, &pb.ValidateTokenRequest{
		AccessToken: token,
	})
	if err != nil {
		return nil, err
	}
	return s.distributorDao.GetClientRegisterByClientUID(ctx, req)
}
func (s *DistributorServer) CreateFrontendClient(ctx context.Context, req *pb.CreateFrontendClientRequest) (*pb.CreateFrontendClientResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	_, err = s.stockinfoDao.ValidateToken(ctx, &pb.ValidateTokenRequest{
		AccessToken: token,
	})
	if err != nil {
		return nil, err
	}
	return s.distributorDao.CreateFrontendClient(ctx, req)
}
func (s *DistributorServer) DeleteFrontendClient(ctx context.Context, req *pb.DeleteFrontendClientRequest) (*pb.DeleteFrontendClientResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	_, err = s.stockinfoDao.ValidateToken(ctx, &pb.ValidateTokenRequest{
		AccessToken: token,
	})
	if err != nil {
		return nil, err
	}
	return s.distributorDao.DeleteFrontendClient(ctx, req)
}
func (s *DistributorServer) GetFrontendClientByIP(ctx context.Context, req *pb.GetFrontendClientByIPRequest) (*pb.GetFrontendClientByIPResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	_, err = s.stockinfoDao.ValidateToken(ctx, &pb.ValidateTokenRequest{
		AccessToken: token,
	})
	if err != nil {
		return nil, err
	}
	return s.distributorDao.GetFrontendClientByIP(ctx, req)
}
