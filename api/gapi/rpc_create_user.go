package gapi

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo-api/pb"
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/RoyceAzure/go-stockinfo-shared/utility/constants"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreteUserRequest(req)
	if violations != nil {
		return nil, utility.InvalidArgumentError(violations)
	}
	hashed_password, err := utility.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to to hash password : %s", err)
	}

	arg := db.CreateUserParams{
		UserName:       req.GetUserName(),
		Email:          req.GetEmail(),
		HashedPassword: hashed_password,
		SsoIdentifer:   utility.StringToSqlNiStr(req.GetSsoIdentifer()),
		CrUser:         "SYSTEM",
	}
	//send and vertify email

	//注意  這裡的ctx是由gin.Context提供，這就表示要不要中止process是由gin框架控制
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			switch pgErr.Code.Name() {
			case constants.UniqueViolation:
				return nil, status.Errorf(codes.AlreadyExists, "user name already exists : %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user : %s", err)
	}

	res := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return res, nil
}

func validateCreteUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utility.ValidateUsername(req.GetUserName()); err != nil {
		violations = append(violations, utility.FieldViolation("username", err))
	}
	if err := utility.ValidEmail(req.GetEmail()); err != nil {
		violations = append(violations, utility.FieldViolation("email", err))
	}
	if err := utility.ValidPassword(req.GetPassword()); err != nil {
		violations = append(violations, utility.FieldViolation("password", err))
	}
	if err := utility.ValidSSO(req.GetSsoIdentifer()); err != nil {
		violations = append(violations, utility.FieldViolation("sso_identifer", err))
	}
	return violations
}
