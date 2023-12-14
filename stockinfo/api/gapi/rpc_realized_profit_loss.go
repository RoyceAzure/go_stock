package gapi

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetRealizedProfitLoss(ctx context.Context, req *pb.GetRealizedProfitLossRequest) (*pb.GetRealizedProfitLossResponse, error) {
	var rsp pb.GetRealizedProfitLossResponse
	payload, err := server.authorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	entities, err := server.store.GetRealizedProfitLosssByUserIdDetial(ctx, db.GetRealizedProfitLosssByUserIdDetialParams{
		UserID: payload.UserId,
		Limit:  constants.DEFAULT_PAGE_SIZE,
		Offset: (constants.DEFAULT_PAGE - 1) * constants.DEFAULT_PAGE_SIZE,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, status.Errorf(codes.Internal, "%s", err)
	}
	var datas []*pb.RealizedProfitLoss

	for _, entity := range entities {
		realized, err := decimal.NewFromString(entity.Realized)
		if err != nil {
			return nil, util.InternalError(err)
		}

		realizedPrecent, err := decimal.NewFromString(entity.RealizedPrecent)
		if err != nil {
			return nil, util.InternalError(err)
		}

		datas = append(datas, &pb.RealizedProfitLoss{
			UserId:          entity.UserID,
			ProductName:     entity.ProductName,
			CostPerPrice:    entity.CostPerPrice,
			CostTotalPrice:  entity.CostPerPrice,
			Amt:             fmt.Sprintf("%d", entity.TransationAmt.Int32),
			DealPerPrice:    entity.TransationPricePerShare.String,
			Realized:        realized.Round(2).String(),
			RealizedPrecent: fmt.Sprintf("%s%%", realizedPrecent.Round(2).String()),
			TransAt:         timestamppb.New(entity.TransAt.Time),
		})
	}

	rsp.Data = datas
	return &rsp, nil
}

/*
 */
func (server *Server) GetUnRealizedProfitLoss(ctx context.Context, req *pb.GetUnRealizedProfitLossRequest) (*pb.GetUnRealizedProfitLossResponse, error) {
	payload, err := server.authorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}

	userStocks, err := server.store.GetUserStocksByUserId(ctx, db.GetUserStocksByUserIdParams{
		UserID: sql.NullInt64{
			Int64: payload.UserId,
			Valid: true,
		},
		Limits:  constants.DEFAULT_PAGE_SIZE,
		Offsets: (constants.DEFAULT_PAGE - 1) * constants.DEFAULT_PAGE_SIZE,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	stockCodeMap := make(map[string]*db.GetUserStocksByUserIdRow)

	for _, userStock := range userStocks {
		if !userStock.StockName.Valid {
			return nil, status.Errorf(codes.Internal, "%s", fmt.Errorf("stock id mismatch"))
		}
		stockCodeMap[userStock.StockCode.String] = &userStock
	}

	schedulerClient, cancel, err := server.clientFactory.NewClient()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", fmt.Errorf("can't connect to scheduler grpc server"))
	}
	defer cancel()

	sprs, err := schedulerClient.GetStockPriceRealTime(ctx, &pb.StockPriceRealTimeRequest{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", fmt.Errorf("can't fetch spr datas from scheduler"))
	}

	res := pb.GetUnRealizedProfitLossResponse{}
	// var data []*pb.RealizedProfitLoss

	for _, spr := range sprs.Result {
		if val, exists := stockCodeMap[spr.StockCode]; exists {
			costPerPrice, err := decimal.NewFromString(val.PurchasePricePerShare)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "%s", fmt.Errorf("convert purchase price failed"))
			}
			quantity := decimal.NewFromInt32(val.Quantity)
			costTotalPrice := costPerPrice.Mul(quantity)
			currentPrice, err := decimal.NewFromString(spr.OpenPrice)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "%s", fmt.Errorf("convert purchase price failed"))
			}
			currentTotalPrice := currentPrice.Mul(quantity)
			realized := currentTotalPrice.Sub(costTotalPrice)

			var realizedPrecent decimal.Decimal
			if realized.Equal(decimal.NewFromInt32(0)) {
				realizedPrecent = realized.Div(decimal.NewFromInt32(1))
			} else {
				realizedPrecent = realized.Div(costTotalPrice).Mul(decimal.NewFromInt(100))
			}
			res.Data = append(res.Data, &pb.RealizedProfitLoss{
				ProductName:     fmt.Sprintf("%s %s", spr.StockCode, spr.StockName),
				CostPerPrice:    val.PurchasePricePerShare,
				CostTotalPrice:  costTotalPrice.Round(2).String(),
				CurrentPrice:    currentPrice.String(),
				Amt:             fmt.Sprintf("%d", val.Quantity),
				Realized:        realized.Round(2).String(),
				RealizedPrecent: fmt.Sprintf("%s%%", realizedPrecent.Round(2).String()),
			})
		}
	}
	// res.Data = data
	return &res, nil
}
