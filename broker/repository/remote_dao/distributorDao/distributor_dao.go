package distributor_dao

import (
	"context"
	"fmt"
	"time"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IDistributorDao interface {
	CreateClientRegister(ctx context.Context, in *pb.CreateClientRegisterRequest) (*pb.CreateClientRegisterResponse, error)
	DeleteClientRegister(ctx context.Context, in *pb.DeleteClientRegisterRequest) (*pb.DeleteClientRegisterResponse, error)
	GetClientRegisterByClientUID(ctx context.Context, in *pb.GetClientRegisterByClientUIDRequest) (*pb.GetClientRegisterResponse, error)
	CreateFrontendClient(ctx context.Context, in *pb.CreateFrontendClientRequest) (*pb.CreateFrontendClientResponse, error)
	DeleteFrontendClient(ctx context.Context, in *pb.DeleteFrontendClientRequest) (*pb.DeleteFrontendClientResponse, error)
	GetFrontendClientByIP(ctx context.Context, in *pb.GetFrontendClientByIPRequest) (*pb.GetFrontendClientByIPResponse, error)
}

type DistributorDao struct {
	client pb.StockInfoDistributorClient
}

func NewDistributorDao(address string) (IDistributorDao, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, nil, fmt.Errorf("can't connect distributor grpc server")
	}

	client := pb.NewStockInfoDistributorClient(conn)
	return &DistributorDao{
		client: client,
	}, func() { conn.Close() }, nil
}
