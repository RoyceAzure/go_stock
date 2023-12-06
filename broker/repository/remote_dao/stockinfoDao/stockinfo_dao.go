package stockinfo_dao

import (
	"context"
	"fmt"
	"time"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IStockInfoDao interface {
	CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	GetUser(ctx context.Context, in *pb.GetUserRequest, accessToken string) (*pb.GetUserResponse, error)
	UpdateUser(ctx context.Context, in *pb.UpdateUserRequest, accessToken string) (*pb.UpdateUserResponse, error)
	LoginUser(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error)
	VerifyEmail(ctx context.Context, in *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error)
	InitStock(ctx context.Context, in *pb.InitStockRequest) (*pb.InitStockResponse, error)
}

type StockInfoDao struct {
	client pb.StockInfoClient
}

func NewStockInfoDao(address string) (IStockInfoDao, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("can't connect grpc server")
	}

	client := pb.NewStockInfoClient(conn)
	return &StockInfoDao{
		client: client,
	}, nil
}
