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
	_, err := s.authorizer.AuthorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}
	return s.stockinfoDao.UpdateUser(ctx, req)
}

func (s *StockInfoServer) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	violations := validateLoginRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}
	return s.stockinfoDao.LoginUser(ctx, req)
}

func (s *StockInfoServer) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	violations := validateVerifyEmailRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}
	return s.stockinfoDao.VerifyEmail(ctx, req)
}

func (s *StockInfoServer) InitStock(ctx context.Context, req *pb.InitStockRequest) (*pb.InitStockResponse, error) {
	_, err := s.authorizer.AuthorizUser(ctx)
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