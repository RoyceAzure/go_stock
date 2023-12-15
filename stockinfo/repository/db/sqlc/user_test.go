package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"github.com/stretchr/testify/require"
)

// go get github.com/stretchr/testify  使用require檢查錯誤
func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func CreateRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(10))
	require.NoError(t, err)
	arg := CreateUserParams{
		UserName:       util.RandomString(6),
		Email:          util.RandomString(10),
		HashedPassword: hashedPassword,
		SsoIdentifer:   util.StringToSqlNiStr(util.RandomString(2)),
		CrUser:         "royce",
	}

	user, err := testQueries.CreateUser(context.TODO(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserName, user.UserName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.SsoIdentifer.String, user.SsoIdentifer.String)
	require.Equal(t, arg.UserName, user.UserName)
	require.True(t, user.PasswordChangedAt.UTC().IsZero())

	//檢查db自動產生  不能為0值
	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CrDate)
	return user
}

func CreateRandomUserNoTest() User {
	arg := CreateUserParams{
		UserName:       util.RandomString(6),
		Email:          util.RandomString(10),
		HashedPassword: util.RandomString(5),
		SsoIdentifer:   util.StringToSqlNiStr(util.RandomString(2)),
		CrUser:         "royce",
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
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.Equal(t, user.SsoIdentifer.String, user2.SsoIdentifer.String)

	require.WithinDuration(t, user.CrDate, user2.CrDate, time.Second)
}

func TestUpdateOnlyUserName(t *testing.T) {
	user := CreateRandomUser(t)

	arg := UpdateUserParams{
		UserID:   user.UserID,
		UserName: util.StringToSqlNiStr(util.RandomString(10)),
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.UserName.String, user2.UserName)
	require.NotEqual(t, user2.UserName, user.UserName)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.Equal(t, user.SsoIdentifer.String, user2.SsoIdentifer.String)

	require.WithinDuration(t, user.CrDate, user2.CrDate, time.Second)
}

func TestUpdateOnlyUserSSO(t *testing.T) {
	user := CreateRandomUser(t)
	newSSO := util.StringToSqlNiStr(util.RandomSSOTypeStr())
	arg := UpdateUserParams{
		UserID:       user.UserID,
		SsoIdentifer: newSSO,
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.SsoIdentifer, user2.SsoIdentifer)
	require.NotEqual(t, user.SsoIdentifer, arg.SsoIdentifer)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.Equal(t, newSSO.String, user2.SsoIdentifer.String)

	require.WithinDuration(t, user.CrDate, user2.CrDate, time.Second)
}

func TestUpdateOnlyUserPassword(t *testing.T) {
	user := CreateRandomUser(t)
	var update_time time.Time = time.Now().UTC()
	arg := UpdateUserParams{
		UserID:            user.UserID,
		HashedPassword:    util.StringToSqlNiStr(util.RandomString(10)),
		PasswordChangedAt: util.TimeToSqlNiTime(update_time),
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.HashedPassword.String, user2.HashedPassword)
	require.NotEqual(t, user2.HashedPassword, user.HashedPassword)
	require.Equal(t, user.Email, user2.Email)

	require.WithinDuration(t, user.CrDate, user2.CrDate, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user := CreateRandomUser(t)
	var update_time time.Time = time.Now().UTC()
	arg := UpdateUserParams{
		UserID:            user.UserID,
		UserName:          util.StringToSqlNiStr(util.RandomString(10)),
		SsoIdentifer:      util.StringToSqlNiStr(util.RandomSSOTypeStr()),
		HashedPassword:    util.StringToSqlNiStr(util.RandomString(10)),
		PasswordChangedAt: util.TimeToSqlNiTime(update_time),
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.HashedPassword.String, user2.HashedPassword)
	require.NotEqual(t, user2.HashedPassword, user.HashedPassword)
	require.NotEqual(t, user2.UserName, user.UserName)
	require.NotEqual(t, user2.SsoIdentifer, user.SsoIdentifer)
	require.Equal(t, user.Email, user2.Email)

	require.WithinDuration(t, user.CrDate, user2.CrDate, time.Second)
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
	require.GreaterOrEqual(t, len(users), 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
