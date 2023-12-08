package gapi

import (
	"context"
	"time"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
	worker "github.com/RoyceAzure/go-stockinfo/worker"
	"github.com/hibiken/asynq"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreteUserRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}
	hashed_password, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to to hash password : %s", err)
	}

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			UserName:       req.GetUserName(),
			Email:          req.GetEmail(),
			HashedPassword: hashed_password,
			SsoIdentifer:   util.StringToSqlNiStr(req.GetSsoIdentifer()),
			CrUser:         "SYSTEM",
		},
		AfterCreate: func(user db.User) error {
			//send and vertify email
			taskPayload := &worker.PayloadSendVerifyEmail{
				UserName: user.UserName,
				UserId:   user.UserID,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.MailQueue),
			}
			return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}

	//注意  這裡的ctx是由gin.Context提供，這就表示要不要中止process是由gin框架控制
	txResult, err := server.store.CreateUserTx(ctx, arg)
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
		User: convertUser(txResult.User),
	}
	return res, nil
}

func validateCreteUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := util.ValidateUsername(req.GetUserName()); err != nil {
		violations = append(violations, util.FieldViolation("username", err))
	}
	if err := util.ValidEmail(req.GetEmail()); err != nil {
		violations = append(violations, util.FieldViolation("email", err))
	}
	if err := util.ValidPassword(req.GetPassword()); err != nil {
		violations = append(violations, util.FieldViolation("password", err))
	}
	if err := util.ValidSSO(req.GetSsoIdentifer()); err != nil {
		violations = append(violations, util.FieldViolation("sso_identifer", err))
	}
	return violations
}
