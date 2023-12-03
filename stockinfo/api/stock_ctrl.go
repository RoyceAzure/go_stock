package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RoyceAzure/go-stockinfo/api/pb"
	db "github.com/RoyceAzure/go-stockinfo/project/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
config
需要trans
trunctable??
*/
func (server *Server) initSyncStock(ctx *gin.Context) {
	startTime := time.Now().UTC()
	log.Info().Msg("initSyncStock start")
	schduler_host_url := "localhost:9091"
	conn, err := grpc.Dial(schduler_host_url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("connect grpc server get error : %w", err))
		return
	}
	defer conn.Close()

	client := pb.NewStockInfoSchdulerClient(conn)

	res, err := client.GetStockDayAvg(ctx, &pb.StockDayAvgRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("call grpc server get error : %w", err))
		return
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

	response := struct {
		SucessCount int `josn:"success_rows"`
		FailedCount int `josn:"failed_rows"`
	}{
		SucessCount: successedCount,
		FailedCount: failedCount,
	}
	endTime := time.Now().UTC()
	duration := endTime.Sub(startTime)
	log.Info().Float64("elpase second", duration.Seconds()).Msg("initSyncStock end")
	ctx.JSON(http.StatusAccepted, response)
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
