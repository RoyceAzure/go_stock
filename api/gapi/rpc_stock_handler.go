package gapi

import (
	"context"
	"time"

	"github.com/RoyceAzure/go-stockinfo-api/pb"
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func (server *Server) InitStock(ctx context.Context, req *pb.InitStockRequest) (*pb.InitStockResponse, error) {
	startTime := time.Now().UTC()
	log.Info().Msg("initSyncStock start")
	schduler_host_url := "localhost:9091"
	conn, err := grpc.Dial(schduler_host_url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't connect grpc server")
	}
	defer conn.Close()

	client := pb.NewStockInfoSchdulerClient(conn)

	res, err := client.GetStockDayAvg(ctx, &pb.StockDayAvgRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "call grpc server get error")
	}
	successedCount := 0
	failedCount := 0
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
