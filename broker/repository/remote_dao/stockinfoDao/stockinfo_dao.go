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
	GetFund(ctx context.Context, req *pb.GetFundRequest, accessToken string) (*pb.GetFundResponse, error)
	AddFund(ctx context.Context, req *pb.AddFundRequest, accessToken string) (*pb.AddFundResponse, error)
	GetRealizedProfitLoss(ctx context.Context, req *pb.GetRealizedProfitLossRequest, accessToken string) (*pb.GetRealizedProfitLossResponse, error)
	GetUnRealizedProfitLoss(ctx context.Context, req *pb.GetUnRealizedProfitLossRequest, accessToken string) (*pb.GetUnRealizedProfitLossResponse, error)
	GetStock(ctx context.Context, req *pb.GetStockRequest) (*pb.GetStockResponse, error)
	GetStocks(ctx context.Context, req *pb.GetStocksRequest, accessToken string) (*pb.GetStocksResponse, error)
	TransationStock(ctx context.Context, req *pb.TransationRequest, accessToken string) (*pb.TransationResponse, error)
	GetAllTransations(ctx context.Context, req *pb.GetAllStockTransationRequest, accessToken string) (*pb.StockTransatsionResponse, error)
	GetUserStock(ctx context.Context, req *pb.GetUserStockRequest, accessToken string) (*pb.GetUserStockResponse, error)
	GetUserStockById(ctx context.Context, req *pb.GetUserStockByIdRequest, accessToken string) (*pb.GetUserStockBuIdResponse, error)
}

type StockInfoDao struct {
	client pb.StockInfoClient
}

func NewStockInfoDao(address string) (IStockInfoDao, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, nil, fmt.Errorf("can't connect grpc server")
	}

	client := pb.NewStockInfoClient(conn)
	return &StockInfoDao{
		client: client,
	}, func() { conn.Close() }, nil
}
