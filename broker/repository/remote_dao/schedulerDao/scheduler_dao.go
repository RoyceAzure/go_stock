package scheduler_dao

import (
	"context"
	"fmt"
	"time"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

)

type ISchedulerDao interface {
	GetStockDayAvg(ctx context.Context, in *pb.StockDayAvgRequest) (*pb.StockDayAvgResponse, error)
	GetStockPriceRealTime(ctx context.Context, in *pb.StockPriceRealTimeRequest) (*pb.StockPriceRealTimeResponse, error)
}

type SchedulerDao struct {
	client pb.StockInfoSchdulerClient
}

func NewSchedulerDao(address string) (ISchedulerDao, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("can't connect grpc server")
	}

	client := pb.NewStockInfoSchdulerClient(conn)
	return &SchedulerDao{
		client: client,
	}, nil
}
