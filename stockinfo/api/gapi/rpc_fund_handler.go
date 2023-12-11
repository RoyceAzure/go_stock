package gapi

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	utility "github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
讀取使用者所有funds
*/
func (server *Server) GetFund(ctx context.Context, req *pb.GetFundRequest) (*pb.GetFundResponse, error) {
	payload, err := server.authorizUser(ctx)
	if err != nil {
		return nil, utility.UnauthticatedError(err)
	}
	funds, err := server.store.GetFundByUserId(ctx, db.GetFundByUserIdParams{
		UserID: payload.UserId,
		Limit:  DEFAULT_PAGE_SIZE,
		Offset: (DEFAULT_PAGE - 1) * DEFAULT_PAGE_SIZE,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).Msg("fund is empty")
			return nil, status.Errorf(codes.NotFound, "fund not found")
		}
		log.Error().Err(err).Msg("get fund by id failed")
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	var res pb.GetFundResponse

	for _, fund := range funds {
		if err != nil {
			log.Error().Err(err).Msg("failed to convert fund balance")
			return &res, status.Error(codes.Internal, "failed to convert fund balance")
		}
		res.Data = append(res.Data, &pb.Fund{
			UserId:       fund.UserID,
			Balance:      fund.Balance,
			CurrencyType: fund.CurrencyType,
		})
	}
	return &res, nil
}

/*
為了方便  儲值使用者自己金額?  或者發送email給管理員?
TODO要使用transation  ok
*/
func (server *Server) AddFund(ctx context.Context, req *pb.AddFundRequest) (*pb.AddFundResponse, error) {
	payload, err := server.authorizUser(ctx)
	if err != nil {
		return nil, utility.UnauthticatedError(err)
	}

	violations := validateAddTWFundRequest(req)
	if violations != nil {
		return nil, utility.InvalidArgumentError(violations)
	}
	userId := payload.UserId
	txRes, err := server.store.UpdateFundTx(ctx, db.UpdateFundTxParams{
		UserID: userId,
		UPN:    payload.UPN,
		Amount: req.Amount,
	})
	if err != nil {
		return nil, utility.InternalError(err)
	}

	return &pb.AddFundResponse{
		Data: &pb.Fund{
			UserId:       txRes.Fund.UserID,
			Balance:      txRes.Fund.Balance,
			CurrencyType: txRes.Fund.CurrencyType,
		},
		Result: "successed",
	}, nil
}

/*
validate fund for type is TW
default use tw type this version, so no need to check CurrencyType
*/
func validateAddTWFundRequest(req *pb.AddFundRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	value, err := utility.ValidateStringToDecimal(req.GetAmount())
	if err != nil {
		violations = append(violations, utility.FieldViolation("amount", err))
	} else if value.IsZero() {
		violations = append(violations, utility.FieldViolation("amount", fmt.Errorf("amount must not be zero")))
	} else if value.IsNegative() {
		violations = append(violations, utility.FieldViolation("amount", fmt.Errorf("amount must be positive")))
	} else if !value.IsInteger() {
		violations = append(violations, utility.FieldViolation("amount", fmt.Errorf("amount must be integer")))
	}

	// if ok := utility.IsSupportedCurrencyType(req.GetCurrencyType()); !ok {
	// 	violations = append(violations, utility.FieldViolation("currrency_type", fmt.Errorf("%s currency type not supported", req.GetCurrencyType())))
	// }

	return violations
}
