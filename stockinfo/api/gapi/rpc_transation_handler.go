package gapi

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	"github.com/RoyceAzure/go-stockinfo/shared/util"
	utility "github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
	worker "github.com/RoyceAzure/go-stockinfo/worker"
	"github.com/hibiken/asynq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) TransationStock(ctx context.Context, req *pb.TransationRequest) (*pb.TransationResponse, error) {
	violations := validateTransationStockRequest(req)
	if violations != nil {
		return nil, utility.InvalidArgumentError(violations)
	}
	payload, err := server.authorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}

	err = server.taskDistributor.DistributeTaskStockTransation(ctx, &worker.PayloadTransation{
		UserID:    payload.UserId,
		StockCode: req.StockCode,
		TransType: req.TransationType,
		Amt:       int32(req.TransAmt),
		Operator:  payload.UPN,
	}, asynq.ProcessIn(time.Millisecond), asynq.MaxRetry(3))
	if err != nil {
		return nil, util.InternalError(err)
	}
	return nil, nil
}

/*
get user sotcks by user id, default page 1 size 10
*/
func (server *Server) GetAllTransations(ctx context.Context, req *pb.GetAllStockTransationRequest) (*pb.StockTransatsionResponse, error) {
	// tokenPayload, err := server.authorizUser(ctx)
	// if err != nil {
	// 	return nil, utility.UnauthticatedError(err)
	// }

	violations := validateGetAllTransationsRequest(req)
	if violations != nil {
		return nil, utility.InvalidArgumentError(violations)
	}

	var page, page_size int32

	if req.Page == 0 {
		page = constants.DEFAULT_PAGE
	} else {
		page = req.Page
	}

	if req.PageSize == 0 {
		page_size = constants.DEFAULT_PAGE_SIZE
	} else {
		page_size = req.PageSize
	}

	var transactionType sql.NullString
	if req.TransationType != "" {
		transactionType.Valid = false
	} else {
		transactionType.Valid = true
		transactionType.String = req.TransationType
	}

	transations, err := server.store.GetStockTransactionsFilter(ctx, db.GetStockTransactionsFilterParams{
		UserID:          req.UserId,
		StockID:         req.StockId,
		TransactionType: transactionType,
		Limits:          page_size,
		Offsets:         (page - 1) * page_size,
	})

	if err != nil {
		return nil, utility.InternalError(err)
	}

	var data []*pb.StockTransation

	for _, transation := range transations {
		data = append(data, &pb.StockTransation{
			UserId:              transation.UserID,
			StockId:             transation.StockID,
			TransationType:      transation.TransactionType,
			TransAmt:            int64(transation.TransationAmt),
			TransPricesPerShare: transation.TransationPricePerShare,
		})
	}

	res := pb.StockTransatsionResponse{
		Data: data,
	}
	return &res, nil
}

func validateGetAllTransationsRequest(req *pb.GetAllStockTransationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.TransationType != "" {
		if !utility.IsSupportedTransationType(req.TransationType) {
			violations = append(violations, utility.FieldViolation("transation_type", errors.New("transation_type not supported")))
		}
	}
	return violations
}

func validateTransationStockRequest(req *pb.TransationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.TransationType != "" {
		if !utility.IsSupportedTransationType(req.TransationType) {
			violations = append(violations, utility.FieldViolation("transation_type", errors.New("transation_type not supported")))
		}
	}
	if req.StockCode == "" {
		violations = append(violations, utility.FieldViolation("stock_code", errors.New("stock_code must not empty")))
	}
	if req.TransAmt == 0 {
		violations = append(violations, utility.FieldViolation("trans_amt", errors.New("trans_amt must not empty")))
	}
	return violations
}
