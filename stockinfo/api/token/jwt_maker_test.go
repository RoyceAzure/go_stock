package token

import (
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	username := util.RandomString(10)
	userID := util.RandomInt(100, 1000)
	duration := time.Minute

	issuedAt := time.Now().UTC()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username, userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err = maker.VertifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.UPN)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	username := util.RandomString(10)
	userID := util.RandomInt(100, 1000)
	duration := time.Minute

	//createToken沒有禁止使用負數的duration
	token, payload, err := maker.CreateToken(username, userID, -duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err = maker.VertifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	//自己產生jwt token
	//payload
	username := util.RandomString(10)
	userID := util.RandomInt(100, 1000)
	payload, err := NewPayload(username, userID, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	//選擇加密演算法製作claim
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	//指定key作加密
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	payload, err = maker.VertifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
