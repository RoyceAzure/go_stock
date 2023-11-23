package repository

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/random"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func TestCreateUserRegister(t *testing.T) {
	CreateUserRegister(t)
}

func CreateUserRegister(t *testing.T) UserRegister {
	config, err := config.LoadConfig("../../../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := pgxpool.New(ctx, config.DBSource)
	require.NoError(t, err)
	require.NotEmpty(t, conn)

	dao := NewSQLDistributorDao(conn)
	require.NotEmpty(t, dao)

	StockCode := random.RandomStrInt(10000)
	// CreatedAt := time.Now().UTC()
	res, err := dao.CreateUserRegister(ctx, CreateUserRegisterParams{
		UserID:    testUserId,
		StockCode: StockCode,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, testUserId, res.UserID)
	require.Equal(t, StockCode, res.StockCode)
	return res
}

func TestDeleteUserRegister(t *testing.T) {
	config, err := config.LoadConfig("../../../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := pgxpool.New(ctx, config.DBSource)
	require.NoError(t, err)
	require.NotEmpty(t, conn)

	dao := NewSQLDistributorDao(conn)
	require.NotEmpty(t, dao)

	// CreatedAt := time.Now().UTC()
	ur := CreateUserRegister(t)
	require.NotEmpty(t, ur)

	arg := DeleteUserRegisterParams{
		UserID:    ur.UserID,
		StockCode: ur.StockCode,
	}

	err = dao.DeleteUserRegister(ctx, arg)
	require.NoError(t, err)
}

func TestGetUserRegisterByUserID(t *testing.T) {
	config, err := config.LoadConfig("../../../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := pgxpool.New(ctx, config.DBSource)
	require.NoError(t, err)
	require.NotEmpty(t, conn)

	dao := NewSQLDistributorDao(conn)
	require.NotEmpty(t, dao)

	// CreatedAt := time.Now().UTC()
	ur := CreateUserRegister(t)
	require.NotEmpty(t, ur)

	users, err := dao.GetUserRegisterByUserID(ctx, ur.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.Equal(t, ur.UserID, user.UserID)
	}
}

func TestGetUserRegisters(t *testing.T) {
	config, err := config.LoadConfig("../../../")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := pgxpool.New(ctx, config.DBSource)
	require.NoError(t, err)
	require.NotEmpty(t, conn)

	dao := NewSQLDistributorDao(conn)
	require.NotEmpty(t, dao)

	// CreatedAt := time.Now().UTC()
	for i := 0; i < 5; i++ {
		CreateUserRegister(t)
	}

	arg := GetUserRegistersParams{
		Limit:  5,
		Offset: 0,
	}

	users, err := dao.GetUserRegisters(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	require.Len(t, users, 5)
}
