package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/RoyceAzure/go-stockinfo/api/token"
	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	violations := validateLoginRequest(req)
	if violations != nil {
		return nil, util.InvalidArgumentError(violations)
	}
	user, err := server.store.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, constants.ErrUserNotEsixts.Error())
		}
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "%s", fmt.Errorf("wrong password"))
	}

	var accessToken string
	var accessPayload *token.Payload
	var session db.Session
	session, err = server.store.GetSessionByUserId(ctx, user.UserID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, status.Errorf(codes.Internal, "%s", err)
		}
		session, err = server.createSession(ctx, user.Email, user.UserID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "%s", err)
		}
	} else if time.Now().After(session.ExpiredAt) {
		err = server.store.DeleteSession(ctx, session.ID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "%s", err)
		}
		session, err = server.createSession(ctx, user.Email, user.UserID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "%s", err)
		}
	}

	accessToken, accessPayload, err = server.tokenMaker.CreateToken(
		req.Email,
		user.UserID,
		server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	rsp := &pb.LoginUserResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		AccessTokenExpiredAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          session.RefreshToken,
		RefreshTokenExpiredAt: timestamppb.New(session.ExpiredAt),
	}
	return rsp, nil
}

func validateLoginRequest(req *pb.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := util.ValidEmail(req.GetEmail()); err != nil {
		violations = append(violations, util.FieldViolation("email", err))
	}
	if err := util.ValidPassword(req.GetPassword()); err != nil {
		violations = append(violations, util.FieldViolation("password", err))
	}
	return violations
}

func (server *Server) createSession(ctx context.Context, email string, userId int64) (session db.Session, err error) {
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		email,
		userId,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return session, status.Errorf(codes.Internal, "%s", err)
	}
	//session跟token綁定??
	mtda := util.ExtractMetaData(ctx)
	session, err = server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       userId,
		RefreshToken: refreshToken,
		UserAgent:    mtda.UserAgent, //TODO : fillit
		ClientIp:     mtda.ClientIP,  //TODO : fillit
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return session, err
	}
	return session, nil
}
