package gapi

import (
	"context"
	"database/sql"
	"time"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 沒有權限的users直接回傳錯誤，並不會有violations 錯誤訊息
// 因為此API提供給grpc gatway and http req, 所以驗證auth header邏輯直接寫在此handler裡面，不使用intecpter
func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	payload, err := server.authorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}
	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}

	if payload.UserId != req.GetUserId() {
		return nil, status.Errorf(codes.PermissionDenied, "can't update other user's info")
	}

	arg := db.UpdateUserParams{
		UserID:       req.UserId,
		UserName:     util.StringToSqlNiStr(req.GetUserName()),
		SsoIdentifer: util.StringToSqlNiStr(req.GetSsoIdentifer()),
	}
	if req.Password != nil {
		hashed_password, err := util.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to to hash password : %s", err)
		}
		arg.HashedPassword = util.StringToSqlNiStr(hashed_password)
		arg.PasswordChangedAt = util.TimeToSqlNiTime(time.Now().UTC())
	}

	//注意  這裡的ctx是由gin.Context提供，這就表示要不要中止process是由gin框架控制
	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		// if pgErr, ok := err.(*pq.Error); ok {
		// 	switch pgErr.Code.Name() {
		// 	case constants.UniqueViolation:
		// 		return nil, status.Errorf(codes.AlreadyExists, "user name already exists : %s", err)
		// 	}
		// }
		return nil, status.Errorf(codes.Internal, "failed to Update user : %s", err)
	}

	res := &pb.UpdateUserResponse{
		User: convertUser(user),
	}
	return res, nil
}

// allowed user_name password sso_identifer to be nil
func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := util.ValidateMustNotZeroInt64(req.UserId); err != nil {
		violations = append(violations, util.FieldViolation("userID", err))
	}
	if req.UserName != nil {
		if err := util.ValidateUsername(req.GetUserName()); err != nil {
			violations = append(violations, util.FieldViolation("username", err))
		}
	}
	if req.Password != nil {
		if err := util.ValidPassword(req.GetPassword()); err != nil {
			violations = append(violations, util.FieldViolation("password", err))
		}
	}
	if req.SsoIdentifer != nil {
		if err := util.ValidSSO(req.GetSsoIdentifer()); err != nil {
			violations = append(violations, util.FieldViolation("sso_identifer", err))
		}
	}
	return violations
}
