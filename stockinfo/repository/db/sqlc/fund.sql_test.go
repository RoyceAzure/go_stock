package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestCreateFund(t *testing.T) {
	CreateRandomFund(t)
}

func CreateRandomFund(t *testing.T) (User, Fund) {
	user := CreateRandomUser(t)
	balance := "100"
	arg := CreateFundParams{
		UserID:       user.UserID,
		Balance:      balance,
		CurrencyType: util.RandomCurrencyTypeStr(),
		CrUser:       "royce",
	}

	fund, err := testQueries.CreateFund(context.TODO(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fund)

	require.Equal(t, user.UserID, fund.UserID)
	require.Equal(t, arg.Balance, fund.Balance)
	require.Equal(t, arg.CurrencyType, fund.CurrencyType)

	//檢查db自動產生  不能為0值
	require.NotZero(t, fund.FundID)
	require.NotZero(t, fund.CrDate)
	return user, fund
}

func CreateRandomFundNoTest(user User, min int64, max int64) (User, Fund) {
	arg := CreateFundParams{
		UserID:       user.UserID,
		Balance:      decimal.NewFromInt(util.RandomInt(100000, 5000000)).String(),
		CurrencyType: string(constants.TW),
		CrUser:       "royce",
	}

	fund, _ := testQueries.CreateFund(context.TODO(), arg)
	return user, fund
}

func CreateFundSpecify(t *testing.T, userId int64, balance string, currencyType string) Fund {
	arg := CreateFundParams{
		UserID:       userId,
		Balance:      balance,
		CurrencyType: string(currencyType),
		CrUser:       "test",
	}

	fund, err := testQueries.CreateFund(context.TODO(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, fund)
	return fund
}

func TestGetFund(t *testing.T) {
	_, fund := CreateRandomFund(t)
	fund2, err := testQueries.GetFund(context.Background(), fund.FundID)

	require.NoError(t, err)
	require.NotEmpty(t, fund2)

	require.Equal(t, fund.UserID, fund2.UserID)
	require.Equal(t, fund.Balance, fund2.Balance)
	require.Equal(t, fund.CurrencyType, fund2.CurrencyType)
	require.Equal(t, fund.CrDate, fund2.CrDate)
}

// func TestUpdateUser(t *testing.T) {
// 	user := CreateRandomUser(t)

// 	arg := UpdateUserParams{
// 		UserID:   user.UserID,
// 		UserName: utility.RandomString(10),
// 	}
// 	user2, err := testQueries.UpdateUser(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, arg.UserName, user2.UserName)
// 	require.Equal(t, user.Email, user2.Email)
// 	require.Equal(t, user.Password.String, user2.Password.String)
// 	require.Equal(t, user.SsoIdentifer.String, user2.SsoIdentifer.String)

// 	require.WithinDuration(t, user.CrDate.Time, user2.CrDate.Time, time.Second)
// }

func TestDeleteFund(t *testing.T) {
	_, fund := CreateRandomFund(t)

	err := testQueries.DeleteFund(context.Background(), fund.FundID)

	require.NoError(t, err)

	fund2, err := testQueries.GetFund(context.Background(), fund.FundID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, fund2)
}

func TestGetFunds(t *testing.T) {
	var user User
	for i := 0; i < 5; i++ {
		user, _ = CreateRandomFund(t)
	}

	arg := GetFundByUserIdParams{
		UserID: user.UserID,
		Limit:  5,
		Offset: 0,
	}
	funds, err := testQueries.GetFundByUserId(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, funds)

	for _, fund := range funds {
		require.NotEmpty(t, fund)
		require.Equal(t, user.UserID, fund.UserID)
	}
}
