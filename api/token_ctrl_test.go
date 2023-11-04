package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-api/token"
	mockdb "github.com/RoyceAzure/go-stockinfo-project/db/mock"
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRenewAccessToken(t *testing.T) {
	user := randomUser()
	tokenMaker, err := token.NewPasetoMaker("12345678123456781234567812345678")
	accessToken, accessPayload, err := tokenMaker.CreateToken(user.Email, user.UserID, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, accessToken)
	require.NotEmpty(t, accessPayload)

	refreshToken, refreshPayload, err := tokenMaker.CreateToken(user.Email, user.UserID, time.Hour)
	require.NoError(t, err)
	require.NotEmpty(t, refreshToken)
	require.NotEmpty(t, refreshPayload)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	session := db.Session{
		ID:           refreshPayload.ID,
		UserID:       user.UserID,
		RefreshToken: refreshToken,
		UserAgent:    "test",
		ClientIp:     "test",
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
		CrDate:       time.Now().UTC(),
	}

	store.EXPECT().
		GetSession(gomock.Any(), gomock.Any()).
		Times(1).Return(session, nil)
	server := newTestServer(t, store)
	server.tokenMaker = tokenMaker
	recoder := httptest.NewRecorder()
	url := "/token/renew_access"
	reqBody := renewAccessTokenRequest{
		RefreshToken: refreshToken,
	}
	data, err := json.Marshal(reqBody)
	require.NoError(t, err)
	require.NotEmpty(t, data)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	server.router.ServeHTTP(recoder, request)
	require.Equal(t, http.StatusOK, recoder.Code)
}
