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
	GetSprCache(ctx context.Context) *SprCache
	SetSprCache(ctx context.Context, value []*pb.StockPriceRealTime, resultTimeStr string)
	GetPreSprTime(ctx context.Context) (string, error)
	SetPreSprTime(ctx context.Context, resultTimeStr string)
	GetStockPriceRealTime(ctx context.Context) (*pb.StockPriceRealTimeResponse, error)
}

/*
parms:

	resultTimeStr :
		最新資料時間，是用5分鐘為區間的文字表示時間
*/
type SprCache struct {
	ResultTimeStr string
	Result        []*pb.StockPriceRealTime
}

/*
有狀態  跟spr_service為一對一關係  紀錄service取過那些資料
*/
type JSchdulerInfoDao struct {
	client     pb.StockInfoSchdulerClient
	sprCache   SprCache
	mutex      sync.RWMutex
	preSprTime string
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

func (dao *JSchdulerInfoDao) GetSprCache(ctx context.Context) *SprCache {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()
	result := &dao.sprCache
	return result
}

func (dao *JSchdulerInfoDao) SetSprCache(ctx context.Context, value []*pb.StockPriceRealTime, resultTimeStr string) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	if dao.sprCache.ResultTimeStr == resultTimeStr {
		return
	}
	dao.sprCache.ResultTimeStr = resultTimeStr
	dao.sprCache.Result = value
}

func (dao *JSchdulerInfoDao) GetPreSprTime(ctx context.Context) (string, error) {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()
	sprTimedao := dao.preSprTime
	if sprTimedao == "" {
		return "", fmt.Errorf("spr pre time is empty")
	}
	return dao.preSprTime, nil
}

func (dao *JSchdulerInfoDao) SetPreSprTime(ctx context.Context, preSprTime string) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	dao.preSprTime = preSprTime
}
