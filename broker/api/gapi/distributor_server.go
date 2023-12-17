package gapi

import (
	"github.com/RoyceAzure/go-stockinfo-broker/api"
	distributor_dao "github.com/RoyceAzure/go-stockinfo-broker/repository/remote_dao/distributorDao"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
)

type DistributorServer struct {
	pb.UnimplementedStockInfoDistributorServer
	distributorDao distributor_dao.IDistributorDao
	authorizer     api.IAuthorizer
}

func NewDistributorServer(distributorDao distributor_dao.IDistributorDao, authorizer api.IAuthorizer) (*DistributorServer, error) {
	server := &DistributorServer{
		distributorDao: distributorDao,
		authorizer:     authorizer,
	}
	return server, nil
}
