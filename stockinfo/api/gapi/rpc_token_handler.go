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
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TODO: session管理應該跟 refreshToken分開  兩者生命週期不一樣
func (server *Server) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	//valid refreshToken
	tokenPayload, err := server.tokenMaker.VertifyToken(req.AccessToken)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}

	session, err := server.store.GetSessionByUserId(ctx, tokenPayload.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, util.UnauthticatedError(fmt.Errorf("please login first"))
		}
		return nil, util.InternalError(err)
	}

	err = server.validateTokenSession(ctx, tokenPayload, &session)
	if err != nil {
		err = server.store.DeleteSession(ctx, session.ID)
		if err != nil {
			err = fmt.Errorf("please re login")
			return nil, util.InternalError(err)
		}
		return nil, util.UnauthticatedError(fmt.Errorf("please re login"))
	}

	// accessToken, accessPayload, err := server.tokenMaker.CreateToken(refreshPayload.UPN, session.UserID, server.config.AccessTokenDuration)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	// res := renewAccessTokenResponse{
	// 	AccessToken:          accessToken,
	// 	AccessTokenExpiredAt: accessPayload.ExpiredAt,
	// }
	return &pb.ValidateTokenResponse{
		Result: true,
	}, nil
}

/*
token 只能是refresh token 不能是 access token
*/
func (server *Server) RenewToken(ctx context.Context, req *pb.RenewTokenRequest) (*pb.RenewTokenResponse, error) {

	//valid refreshToken
	refreshPayload, err := server.tokenMaker.VertifyToken(req.RefreshToken)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}

	oriSession, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, util.UnauthticatedError(fmt.Errorf("please login first"))
		}
		return nil, util.InternalError(err)
	}

	err = server.validateTokenSession(ctx, refreshPayload, &oriSession)
	if err != nil {
		return nil, err
	}

	if oriSession.RefreshToken != req.RefreshToken {
		err = fmt.Errorf("mismatch session token")
		return nil, util.UnauthticatedError(err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(refreshPayload.UPN,
		oriSession.UserID,
		server.config.AccessTokenDuration)
	if err != nil {
		return nil, util.InternalError(err)
	}
	return &pb.RenewTokenResponse{
		AccessToken: accessToken,
		ExpiredAt:   timestamppb.New(accessPayload.ExpiredAt),
	}, nil
}

/*
for grpc
return error with grpc code
*/
func (server *Server) validateTokenSession(ctx context.Context, token *token.Payload, session *db.Session) error {
	var err error

	if session.IsBlocked {
		err = fmt.Errorf("blocked session")
		return util.UnauthticatedError(err)
	}

	if session.UserID != token.UserId {
		err = fmt.Errorf("incorrect session user")
		return util.UnauthticatedError(err)
	}

	if time.Now().After(session.ExpiredAt) {
		err = fmt.Errorf("expired session")
		return util.UnauthticatedError(err)
	}
	return nil
}
