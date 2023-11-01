package gapi

import (
	"context"
	"database/sql"

	"github.com/RoyceAzure/go-stockinfo-api/pb"
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	violations := validateLoginRequest(req)
	if violations != nil {
		return nil, utility.InvalidArgumentError(violations)
	}
	user, err := server.store.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found : %s", err)
		}
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	err = utility.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "%s", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(req.Email,
		user.UserID,
		server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		req.Email,
		user.UserID,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}
	//session跟token綁定??
	mtda := server.extractMetaData(ctx)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       user.UserID,
		RefreshToken: refreshToken,
		UserAgent:    mtda.UserAgent, //TODO : fillit
		ClientIp:     mtda.ClientIP,  //TODO : fillit
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	rsp := &pb.LoginUserResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}

func validateLoginRequest(req *pb.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utility.ValidEmail(req.GetEmail()); err != nil {
		violations = append(violations, utility.FieldViolation("email", err))
	}
	if err := utility.ValidPassword(req.GetPassword()); err != nil {
		violations = append(violations, utility.FieldViolation("password", err))
	}
	return violations
}
