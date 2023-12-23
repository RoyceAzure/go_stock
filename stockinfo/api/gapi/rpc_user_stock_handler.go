package gapi

import (
	"context"
	"database/sql"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
)

/*
目前限定只能取得自己的資料，所以不接受userId參數
*/
func (server *Server) GetUserStock(ctx context.Context, req *pb.GetUserStockRequest) (*pb.GetUserStockResponse, error) {
	tokenPayload, err := server.authorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	var page_size, page int32
	if req.PageSize == 0 {
		page_size = constants.DEFAULT_PAGE_SIZE
	} else {
		page_size = req.PageSize
	}

	if req.Page == 0 {
		page = constants.DEFAULT_PAGE
	} else {
		page = req.Page
	}

	userStocks, err := server.store.GetUserStocksByUserId(ctx, db.GetUserStocksByUserIdParams{
		UserID: sql.NullInt64{
			Int64: tokenPayload.UserId,
			Valid: true,
		},
		Limits:  page_size,
		Offsets: (page - 1) * page_size,
	})

	if err != nil {
		return nil, util.InternalError(err)
	}

	var data []*pb.GetUserStockResponseDTO

	for _, usesrSotck := range userStocks {
		data = append(data, &pb.GetUserStockResponseDTO{
			UserId:                usesrSotck.UserID,
			StockId:               usesrSotck.StockID,
			StockCode:             usesrSotck.StockCode.String,
			StockName:             usesrSotck.StockName.String,
			Quantity:              usesrSotck.Quantity,
			PurchasePricePerShare: usesrSotck.PurchasePricePerShare,
		})
	}

	res := pb.GetUserStockResponse{
		Data: data,
	}

	return &res, nil
}

/*
get user sotcks by user id, default page 1 size 10
*/
func (server *Server) GetUserStockById(ctx context.Context, req *pb.GetUserStockByIdRequest) (*pb.GetUserStockBuIdResponse, error) {
	tokenPayload, err := server.authorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}

	userStocks, err := server.store.GetUserStocksByUserId(ctx, db.GetUserStocksByUserIdParams{
		UserID: sql.NullInt64{
			Int64: tokenPayload.UserId,
			Valid: true,
		},
		Limits:  constants.DEFAULT_PAGE_SIZE,
		Offsets: (constants.DEFAULT_PAGE - 1) * constants.DEFAULT_PAGE_SIZE,
	})

	if err != nil {
		return nil, util.InternalError(err)
	}

	var data []*pb.UserStock

	for _, usesrSotck := range userStocks {
		data = append(data, &pb.UserStock{
			UserId:                usesrSotck.UserID,
			StockId:               usesrSotck.StockID,
			Quantity:              usesrSotck.Quantity,
			PurchasePricePerShare: usesrSotck.PurchasePricePerShare,
		})
	}

	res := pb.GetUserStockBuIdResponse{
		Data: data,
	}
	return &res, nil
}
