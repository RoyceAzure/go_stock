package gapi

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/pb"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util"
	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (s *StockInfoServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreteUserRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}
	return s.stockinfoDao.CreateUser(ctx, req)
}

func (s *StockInfoServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}
	return s.stockinfoDao.UpdateUser(ctx, req, token)
}

func (s *StockInfoServer) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	violations := validateLoginRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}
	res, err := s.stockinfoDao.LoginUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StockInfoServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.stockinfoDao.GetUser(ctx, req, token)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StockInfoServer) GetFund(ctx context.Context, req *pb.GetFundRequest) (*pb.GetFundResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.stockinfoDao.GetFund(ctx, req, token)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StockInfoServer) AddFund(ctx context.Context, req *pb.AddFundRequest) (*pb.AddFundResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.stockinfoDao.AddFund(ctx, req, token)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StockInfoServer) GetRealizedProfitLoss(ctx context.Context, req *pb.GetRealizedProfitLossRequest) (*pb.GetRealizedProfitLossResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.stockinfoDao.GetRealizedProfitLoss(ctx, req, token)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StockInfoServer) GetUnRealizedProfitLoss(ctx context.Context, req *pb.GetUnRealizedProfitLossRequest) (*pb.GetUnRealizedProfitLossResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.stockinfoDao.GetUnRealizedProfitLoss(ctx, req, token)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StockInfoServer) GetStock(ctx context.Context, req *pb.GetStockRequest) (*pb.GetStockResponse, error) {

	res, err := s.stockinfoDao.GetStock(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StockInfoServer) GetStocks(ctx context.Context, req *pb.GetStocksRequest) (*pb.GetStocksResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.stockinfoDao.GetStocks(ctx, req, token)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StockInfoServer) TransationStock(ctx context.Context, req *pb.TransationRequest) (*pb.TransationResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.stockinfoDao.TransationStock(ctx, req, token)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StockInfoServer) GetAllTransations(ctx context.Context, req *pb.GetAllStockTransationRequest) (*pb.StockTransatsionResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.stockinfoDao.GetAllTransations(ctx, req, token)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StockInfoServer) GetUserStock(ctx context.Context, req *pb.GetUserStockRequest) (*pb.GetUserStockResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.stockinfoDao.GetUserStock(ctx, req, token)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StockInfoServer) GetUserStockById(ctx context.Context, req *pb.GetUserStockByIdRequest) (*pb.GetUserStockBuIdResponse, error) {
	_, token, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.stockinfoDao.GetUserStockById(ctx, req, token)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StockInfoServer) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	violations := validateVerifyEmailRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}
	return s.stockinfoDao.VerifyEmail(ctx, req)
}

func (s *StockInfoServer) InitStock(ctx context.Context, req *pb.InitStockRequest) (*pb.InitStockResponse, error) {
	_, _, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	return s.stockinfoDao.InitStock(ctx, req)
}

func validateCreteUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateUsername(req.GetUserName()); err != nil {
		violations = append(violations, util.FieldViolation("username", err))
	}
	if err := validate.ValidEmail(req.GetEmail()); err != nil {
		violations = append(violations, util.FieldViolation("email", err))
	}
	if err := validate.ValidPassword(req.GetPassword()); err != nil {
		violations = append(violations, util.FieldViolation("password", err))
	}
	if err := validate.ValidSSO(req.GetSsoIdentifer()); err != nil {
		violations = append(violations, util.FieldViolation("sso_identifer", err))
	}
	return violations
}

func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateMustNotZeroInt64(req.UserId); err != nil {
		violations = append(violations, util.FieldViolation("userID", err))
	}
	if req.UserName != nil {
		if err := validate.ValidateUsername(req.GetUserName()); err != nil {
			violations = append(violations, util.FieldViolation("username", err))
		}
	}
	if req.Password != nil {
		if err := validate.ValidPassword(req.GetPassword()); err != nil {
			violations = append(violations, util.FieldViolation("password", err))
		}
	}
	if req.SsoIdentifer != nil {
		if err := validate.ValidSSO(req.GetSsoIdentifer()); err != nil {
			violations = append(violations, util.FieldViolation("sso_identifer", err))
		}
	}
	return violations
}

func validateLoginRequest(req *pb.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidEmail(req.GetEmail()); err != nil {
		violations = append(violations, util.FieldViolation("email", err))
	}
	if err := validate.ValidPassword(req.GetPassword()); err != nil {
		violations = append(violations, util.FieldViolation("password", err))
	}
	return violations
}

func validateVerifyEmailRequest(req *pb.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidEmailID(req.GetEmailId()); err != nil {
		violations = append(violations, util.FieldViolation("email_id", err))
	}
	if err := validate.ValidSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, util.FieldViolation("secret_code", err))
	}
	return violations
}
