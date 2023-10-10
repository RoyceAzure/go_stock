package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/stretchr/testify/require"
)

// go get github.com/stretchr/testify  使用require檢查錯誤
func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		UserName:     utility.RandomString(6),
		Email:        utility.RandomString(10),
		Password:     utility.StringToSqlNiStr(utility.RandomString(5)),
		SsoIdentifer: utility.StringToSqlNiStr(utility.RandomString(2)),
		CrUser:       "royce",
	}

	user, err := testQueries.CreateUser(context.TODO(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserName, user.UserName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password.String, user.Password.String)
	require.Equal(t, arg.SsoIdentifer.String, user.SsoIdentifer.String)
	require.Equal(t, arg.UserName, user.UserName)

	//檢查db自動產生  不能為0值
	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CrDate)
	return user
}

func CreateRandomUserNoTest() User {
	arg := CreateUserParams{
		UserName:     utility.RandomString(6),
		Email:        utility.RandomString(10),
		Password:     utility.StringToSqlNiStr(utility.RandomString(5)),
		SsoIdentifer: utility.StringToSqlNiStr(utility.RandomString(2)),
		CrUser:       "royce",
	}

	user, _ := testQueries.CreateUser(context.TODO(), arg)
	return user
}

func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user.UserID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.UserName, user2.UserName)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.Password.String, user2.Password.String)
	require.Equal(t, user.SsoIdentifer.String, user2.SsoIdentifer.String)

	require.WithinDuration(t, user.CrDate.Time, user2.CrDate.Time, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user := CreateRandomUser(t)

	arg := UpdateUserParams{
		UserID:   user.UserID,
		UserName: utility.RandomString(10),
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.UserName, user2.UserName)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.Password.String, user2.Password.String)
	require.Equal(t, user.SsoIdentifer.String, user2.SsoIdentifer.String)

	require.WithinDuration(t, user.CrDate.Time, user2.CrDate.Time, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user := CreateRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user.UserID)

	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user.UserID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestGetUsers(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomUser(t)
	}

	arg := GetusersParams{
		Limit:  5,
		Offset: 5,
	}
	users, err := testQueries.Getusers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
