package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "github.com/RoyceAzure/go-stockinfo/shared/pb"
	"google.golang.org/grpc"
)

type SchdulerInfoDao interface {
	GetSprData(ctx context.Context) SprData
	SetSprData(ctx context.Context, value SprData)
	GetStockPriceRealTime(ctx context.Context) (SprData, error)
}

/*
外不需要使用這格結構  所以expose
使用複製上給呼叫者使用，所以成員都是大寫

parms:

	DataTime :
		最新資料與時間，時間是由scheduler設定
*/
type SprData struct {
	DataTime string
	Data     []*pb.StockPriceRealTime
}

/*
有狀態  跟spr_service為一對一關係  紀錄service取過那些資料

取得sprData 需要透過鎖，不能直接取得，所以sprData 沒有expose
parms:

	preSprTime: 上一次成功去schduler取得spr時間
*/
type JSchdulerInfoDao struct {
	client     pb.StockInfoSchdulerClient
	sprData    SprData
	mutex      sync.RWMutex
	preSprTime time.Time
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

func (dao *JSchdulerInfoDao) GetSprData(ctx context.Context) SprData {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()
	result := dao.sprData
	return result
}

/*
set sprCache.ResultTimeStr and sprCache.Result with mutex.Lock
*/
func (dao *JSchdulerInfoDao) SetSprData(ctx context.Context, value SprData) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	if dao.sprData.DataTime == value.DataTime {
		return
	}
	dao.sprData = value
}
