package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiredAt time.Time `json:"access_token_expired_at"`
}

// TODO: session管理應該跟 refreshToken分開  兩者生命週期不一樣
func (server *Server) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	//valid refreshToken
	refreshPayload, err := server.tokenMaker.VertifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if session.IsBlocked {
		err = fmt.Errorf("blocked session")
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	if session.UserID != refreshPayload.UserId {
		err = fmt.Errorf("incorrect session user")
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	if session.RefreshToken != req.RefreshToken {
		err = fmt.Errorf("mismatch session token")
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	if time.Now().After(session.ExpiredAt) {
		err = fmt.Errorf("expired session")
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(refreshPayload.UPN, session.UserID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	res := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiredAt: accessPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, res)
}
