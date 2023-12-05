package gapi

import (
	"github.com/RoyceAzure/go-stockinfo-broker/api"
	stockinfo_dao "github.com/RoyceAzure/go-stockinfo-broker/repository/remote_dao/stockinfoDao"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
)

type StockInfoServer struct {
	pb.UnimplementedStockInfoServer
	stockinfoDao stockinfo_dao.IStockInfoDao
	authorizer   api.IAuthorizer
}

func NewStockInfoServer(stockinfoDao stockinfo_dao.IStockInfoDao, authorizer api.IAuthorizer) (*StockInfoServer, error) {
	server := &StockInfoServer{
		stockinfoDao: stockinfoDao,
		authorizer:   authorizer,
	}
	return server, nil
}
