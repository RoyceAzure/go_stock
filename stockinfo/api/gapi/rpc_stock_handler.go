package gapi

import (
	"context"
	"database/sql"
	"time"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
grpc call scheduler get SDA, tansform to stock data
*/
func (server *Server) InitStock(ctx context.Context, req *pb.InitStockRequest) (*pb.InitStockResponse, error) {
	startTime := time.Now().UTC()
	log.Info().Msg("initSyncStock start")
	client, closeConn, err := server.clientFactory.NewClient()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't connect grpc server")
	}
	defer closeConn()

	res, err := client.GetStockDayAvg(ctx, &pb.StockDayAvgRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	successedCount := 0
	failedCount := 0

	err = server.store.TruncateStocks(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, dto := range res.Result {
		insertEnty := cvStockDayAvg2CreateParm(dto)
		_, err := server.store.CreateStock(ctx, insertEnty)
		if err != nil {
			log.Warn().Msg("create stock get err")
			failedCount++
		} else {
			successedCount++
		}
	}

	response := pb.InitStockResponse{
		SuccessCount: int64(successedCount),
		FailedCount:  int64(failedCount),
	}
	endTime := time.Now().UTC()
	duration := endTime.Sub(startTime)
	log.Info().Float64("elpase second", duration.Seconds()).Msg("initSyncStock end")
	return &response, nil
}

func cvStockDayAvg2CreateParm(value *pb.StockDayAvg) db.CreateStockParams {
	return db.CreateStockParams{
		StockCode:    value.StockCode,
		StockName:    value.StockName,
		CurrentPrice: value.ClosePrice,
		MarketCap:    int64(100000),
		CrUser:       "SYSTEM",
	}
}

/*
取得單一stock ???
*/
func (server *Server) GetStock(ctx context.Context, req *pb.GetStockRequest) (*pb.GetStockResponse, error) {
	return nil, nil
}

/*
TODO : 資料應該要從redis 來  而不是DB
*/
func (server *Server) GetStocks(ctx context.Context, req *pb.GetStocksRequest) (*pb.GetStocksResponse, error) {
	var pageSize, page int32

	if req.PageSize != 0 {
		pageSize = req.PageSize
	} else {
		pageSize = DEFAULT_PAGE_SIZE
	}

	if req.Page != 0 {
		page = req.Page
	} else {
		page = DEFAULT_PAGE
	}

	stocks, err := server.store.GetStocks(ctx, db.GetStocksParams{
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "stock not found : %s", err)
		}
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	var res pb.GetStocksResponse

	for _, stock := range stocks {
		res.Data = append(res.Data, &pb.Stock{
			StockCode:    stock.StockCode,
			StockName:    stock.StockName,
			CurrentPrice: stock.CurrentPrice,
		})
	}

	return &res, nil
}
