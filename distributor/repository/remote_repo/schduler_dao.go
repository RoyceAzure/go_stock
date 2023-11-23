package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "github.com/RoyceAzure/go-stockinfo-distributor/shared/pb/stock_info_scheduler"
	"google.golang.org/grpc"
)

type SchdulerInfoDao interface {
	GetSprCache() []*pb.StockPriceRealTime
	SetSprCache(value []*pb.StockPriceRealTime)
	GetStockPriceRealTime(ctx context.Context) (*pb.StockPriceRealTimeResponse, error)
}

type JSchdulerInfoDao struct {
	client   pb.StockInfoSchdulerClient
	sprCache []*pb.StockPriceRealTime
	mutex    sync.RWMutex
}

func NewJSchdulerInfoDao(address string) (SchdulerInfoDao, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("can't connect grpc server")
	}

	client := pb.NewStockInfoSchdulerClient(conn)
	return &JSchdulerInfoDao{
		client: client,
	}, nil
}

func (dao *JSchdulerInfoDao) GetSprCache() []*pb.StockPriceRealTime {
	dao.mutex.RLock()
	result := dao.sprCache
	dao.mutex.RUnlock()
	return result
}

func (dao *JSchdulerInfoDao) SetSprCache(value []*pb.StockPriceRealTime) {
	dao.mutex.Lock()
	dao.sprCache = value
	dao.mutex.Unlock()
}
