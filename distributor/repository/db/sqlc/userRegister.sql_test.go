package repository

import (
	"context"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func TestCreateUserRegister(t *testing.T) {
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

	UserID := int64(10)
	StockCode := "testCode"
	// CreatedAt := time.Now().UTC()
	res, err := dao.CreateUserRegister(ctx, CreateUserRegisterParams{
		UserID:    UserID,
		StockCode: StockCode,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, UserID, res.UserID)
	require.Equal(t, StockCode, res.StockCode)
}
