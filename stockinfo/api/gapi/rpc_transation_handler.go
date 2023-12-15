package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	"github.com/RoyceAzure/go-stockinfo/shared/util"
	utility "github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
	worker "github.com/RoyceAzure/go-stockinfo/worker"
	"github.com/hibiken/asynq"
	"github.com/shopspring/decimal"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

/*
從 scheduler撈全部  一 一比對是不是目標stock
*/
func (server *Server) TransationStock(ctx context.Context, req *pb.TransationRequest) (*pb.TransationResponse, error) {
	violations := validateTransationStockRequest(req)
	if violations != nil {
		return nil, utility.InvalidArgumentError(violations)
	}
	payload, err := server.authorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}

	fund, err := server.store.GetFundByUidandCurForUpdateNoK(ctx, db.GetFundByUidandCurForUpdateNoKParams{
		UserID:       payload.UserId,
		CurrencyType: string(constants.TW),
	})
	if err != nil {
		return nil, utility.InValidateOperation(err)
	}

	stock, err := server.store.GetStockByCode(ctx, req.StockCode)
	if err != nil {
		return nil, utility.InValidateOperation(err)
	}

	schedulerClient, cancel, err := server.clientFactory.NewClient()
	defer cancel()
	if err != nil {
		return nil, utility.InternalError(err)
	}
	sprCache, err := schedulerClient.GetStockPriceRealTime(ctx, &pb.StockPriceRealTimeRequest{})
	if err != nil {
		return nil, utility.InternalError(err)
	}

	var targetSPR *pb.StockPriceRealTime

	for _, spr := range sprCache.Result {
		if spr.StockCode == stock.StockCode {
			targetSPR = spr
			break
		}
	}

	if targetSPR == nil {
		err := fmt.Errorf("%w : can't fetch %s current price", constants.ErrInternal, stock.StockName)
		return nil, utility.InValidateOperation(err)
	}

	//不需要查看是否失敗，重試情況下再建立一個新的
	//成功要寫到以實現損益  成交回報   stockTrans要記錄交易成功與失敗

	_, err = decimal.NewFromString(targetSPR.OpenPrice)
	if err != nil {
		err := fmt.Errorf("%w : error convert price", constants.ErrInternal)
		return nil, utility.InValidateOperation(err)
	}

	stockTrans, err := server.store.CreateStockTransaction(ctx, db.CreateStockTransactionParams{
		UserID:                  payload.UserId,
		StockID:                 stock.StockID,
		FundID:                  fund.FundID,
		TransactionType:         req.TransationType,
		TransactionDate:         time.Now().UTC(),
		TransationAmt:           req.TransAmt,
		TransationPricePerShare: targetSPR.OpenPrice,
	})
	if err != nil {
		err := fmt.Errorf("%w : create transation failed", constants.ErrInternal)
		return nil, utility.InternalError(err)
	}

	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.MailQueue),
	}

	err = server.taskDistributor.DistributeTaskStockTransation(ctx, &worker.PayloadTransation{
		TransationID: stockTrans.TransationID,
		Operator:     payload.UPN,
	}, opts...)

	if err != nil {
		return nil, util.InternalError(err)
	}
	return &pb.TransationResponse{
		Result: fmt.Sprintf("commit transation Successed, id %d", stockTrans.TransationID),
	}, nil
}

/*
get user sotcks by user id, default page 1 size 10
目前先固定搜索登入者的資料
*/
func (server *Server) GetAllTransations(ctx context.Context, req *pb.GetAllStockTransationRequest) (*pb.StockTransatsionResponse, error) {
	tokenPayload, err := server.authorizUser(ctx)
	if err != nil {
		return nil, utility.UnauthticatedError(err)
	}

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

	var userId, stockId sql.NullInt64
	if req.UserId == 0 {
		userId.Valid = false
	} else {
		userId.Valid = true
		userId.Int64 = req.UserId
	}

	if req.StockId == 0 {
		stockId.Valid = false
	} else {
		stockId.Valid = true
		stockId.Int64 = req.StockId
	}

	var transactionType sql.NullString
	if req.TransationType == "" {
		transactionType.Valid = false
	} else {
		transactionType.Valid = true
		transactionType.String = req.TransationType
	}

	transations, err := server.store.GetStockTransactionsFilter(ctx, db.GetStockTransactionsFilterParams{
		UserID: sql.NullInt64{
			Int64: tokenPayload.UserId,
			Valid: true,
		},
		StockID:         stockId,
		TransactionType: transactionType,
		Limits:          page_size,
		Offsets:         (page - 1) * page_size,
	})

	if err != nil {
		return nil, utility.InternalError(err)
	}

	var data []*pb.GetransationsResponse

	for _, transation := range transations {
		data = append(data, &pb.GetransationsResponse{
			TransationId:        transation.TransationID,
			UserId:              transation.UserID,
			StockId:             transation.StockID,
			StockCode:           transation.StockCode.String,
			StockName:           transation.StockName.String,
			TransationType:      transation.TransactionType,
			TransAmt:            int64(transation.TransationAmt),
			TransPricesPerShare: transation.TransationPricePerShare,
			Result:              string(transation.Result),
			Msg:                 transation.Msg.String,
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
