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
	CreateClientRegister(ctx context.Context, in *pb.CreateClientRegisterRequest, accessToken string) (*pb.CreateClientRegisterResponse, error)
	DeleteClientRegister(ctx context.Context, in *pb.DeleteClientRegisterRequest, accessToken string) (*pb.DeleteClientRegisterResponse, error)
	GetClientRegisterByClientUID(ctx context.Context, in *pb.GetClientRegisterByClientUIDRequest, accessToken string) (*pb.GetClientRegisterResponse, error)
	CreateFrontendClient(ctx context.Context, in *pb.CreateFrontendClientRequest, accessToken string) (*pb.CreateFrontendClientResponse, error)
	DeleteFrontendClient(ctx context.Context, in *pb.DeleteFrontendClientRequest, accessToken string) (*pb.DeleteFrontendClientResponse, error)
	GetFrontendClientByIP(ctx context.Context, in *pb.GetFrontendClientByIPRequest, accessToken string) (*pb.GetFrontendClientByIPResponse, error)
}

type DistributorDao struct {
	client pb.StockInfoDistributorClient
	Conn   *grpc.ClientConn
}

func NewDistributorDao(address string) (IDistributorDao, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("can't connect distributor grpc server")
	}

	client := pb.NewStockInfoDistributorClient(conn)
	return &DistributorDao{
		client: client,
		Conn:   conn,
	}, nil
}
