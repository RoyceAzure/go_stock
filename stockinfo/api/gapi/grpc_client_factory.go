package gapi

import (
	"github.com/RoyceAzure/go-stockinfo/api/pb"
	"github.com/RoyceAzure/go-stockinfo/shared/utility/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
Simple Factory Pattern
*/
type SchedulerClientFactory interface {
	NewClient() (pb.StockInfoSchdulerClient, func(), error)
}

type GRPCSchedulerClientFactory struct {
	config *config.Config
}

func NewGRPCSchedulerClientFactory(config *config.Config) SchedulerClientFactory {
	return &GRPCSchedulerClientFactory{
		config: config,
	}
}

func (f *GRPCSchedulerClientFactory) NewClient() (pb.StockInfoSchdulerClient, func(), error) {
	conn, err := grpc.Dial(
		f.config.GRPCSchedulerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, nil, err
	}
	return pb.NewStockInfoSchdulerClient(conn), func() { conn.Close() }, nil
}
