package gapi

import (
	"context"
	"strconv"
	"time"

	"github.com/RoyceAzure/go-stockinfo-scheduler/api/pb"
	repository "github.com/RoyceAzure/go-stockinfo-scheduler/repository/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/*
取得今天的SDA
*/
func (server *Server) GetStockDayAvg(ctx context.Context, req *pb.StockDayAvgRequest) (*pb.StockDayAvgResponse, error) {
	startTime := time.Now().UTC()
	crDateStart := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())
	crDateEnd := crDateStart.AddDate(0, 0, 1).Add(-time.Nanosecond)
	limit := 65535
	page := 1
	offset := (page - 1) * limit
	entities, err := server.dao.GetSDAVGALLs(ctx, repository.GetSDAVGALLsParams{
		CrDateStart: pgtype.Timestamptz{
			Time:  crDateStart,
			Valid: true,
		},
		CrDateEnd: pgtype.Timestamptz{
			Time:  crDateEnd,
			Valid: true,
		},
		Limits:  int32(limit),
		Offsets: int32(offset),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get stock day avg : %s", err)
	}
	if len(entities) == 0 {
		return nil, status.Error(codes.Internal, "has no stack day avg data")
	}
	var resultList []*pb.StockDayAvg
	for _, entity := range entities {
		resultList = append(resultList, cvSDAVGALL2GrpcRes(entity))
	}

	return &pb.StockDayAvgResponse{
		Result: resultList,
	}, nil
}

func cvSDAVGALL2GrpcRes(entity repository.StockDayAvgAll) *pb.StockDayAvg {
	var close_price string
	var monthly_avg string

	close_price_temp, err := entity.ClosePrice.Float64Value()
	if err != nil {
		close_price = "0.00"
	} else {
		close_price = strconv.FormatFloat(close_price_temp.Float64, 'f', -1, 64)
	}

	monthly_avg_temp, err := entity.ClosePrice.Float64Value()
	if err != nil {
		monthly_avg = "0.00"
	} else {
		monthly_avg = strconv.FormatFloat(monthly_avg_temp.Float64, 'f', -1, 64)
	}

	return &pb.StockDayAvg{
		StockCode:  entity.Code,
		StockName:  entity.StockName,
		ClosePrice: close_price,
		MonthlyAvg: monthly_avg,
		UpDate:     timestamppb.New(entity.CrDate),
	}
}
