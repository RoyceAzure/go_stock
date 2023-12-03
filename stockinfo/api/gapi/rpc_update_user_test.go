package gapi

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo/api/pb"
	"github.com/RoyceAzure/go-stockinfo/api/token"
	mockdb "github.com/RoyceAzure/go-stockinfo/project/db/mock"
	db "github.com/RoyceAzure/go-stockinfo/project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/utility"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUpdateUserApi(t *testing.T) {
	// 或許在測試API時  randomUser就應該產生DTO 而不是model
	user := randomUser()
	newName := utility.RandomString(10)
	invalidName := "fds"
	invalidSSO := "aa"
	// newPas := utility.RandomString(32)
	newSSo := utility.RandomSSOTypeStr()
	up_date := time.Now().UTC()
	testCase := []struct {
		name         string //子測試名稱
		body         *pb.UpdateUserRequest
		buildStub    func(store *mockdb.MockStore)
		buildContext func(*testing.T, token.Maker) context.Context
		checkReponse func(*testing.T, *pb.UpdateUserResponse, error)
	}{
		{
			name: "OK",
			body: &pb.UpdateUserRequest{
				UserId:       user.UserID,
				UserName:     &newName,
				SsoIdentifer: &newSSo,
			},
			buildStub: func(store *mockdb.MockStore) {
				//這裡手動模擬API處理參數
				arg := db.UpdateUserParams{
					UserID:       user.UserID,
					UserName:     utility.StringToSqlNiStr(newName),
					SsoIdentifer: utility.StringToSqlNiStr(newSSo),
				}
				updateUser := db.User{
					UserID:            user.UserID,
					UserName:          newName,
					Email:             user.Email,
					HashedPassword:    user.HashedPassword,
					PasswordChangedAt: user.PasswordChangedAt,
					SsoIdentifer:      utility.StringToSqlNiStr(newSSo),
					IsEmailVerified:   user.IsEmailVerified,
					CrDate:            user.CrDate,
					UpDate:            utility.TimeToSqlNiTime(up_date),
				}
				store.EXPECT().
					UpdateUser(gomock.Any(), gomock.Eq(arg)).
					Times(1).Return(updateUser, nil)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Email, user.UserID, time.Minute)
			},
			checkReponse: func(t *testing.T, res *pb.UpdateUserResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				updateUser := res.GetUser()
				require.Equal(t, updateUser.GetEmail(), user.Email)
				require.Equal(t, updateUser.GetUserName(), newName)
				require.Equal(t, updateUser.GetSsoIdentifer(), newSSo)

			},
		},
		{
			name: "User not found",
			body: &pb.UpdateUserRequest{
				UserId:       user.UserID,
				UserName:     &newName,
				SsoIdentifer: &newSSo,
			},
			buildStub: func(store *mockdb.MockStore) {
				//這裡手動模擬API處理參數
				arg := db.UpdateUserParams{
					UserID:       user.UserID,
					UserName:     utility.StringToSqlNiStr(newName),
					SsoIdentifer: utility.StringToSqlNiStr(newSSo),
				}
				store.EXPECT().
					UpdateUser(gomock.Any(), gomock.Eq(arg)).
					Times(1).Return(db.User{}, sql.ErrNoRows)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Email, user.UserID, time.Minute)
			},
			checkReponse: func(t *testing.T, res *pb.UpdateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.NotFound, st.Code())
			},
		},
		{
			name: "Expired Token",
			body: &pb.UpdateUserRequest{
				UserId:       user.UserID,
				UserName:     &newName,
				SsoIdentifer: &newSSo,
			},
			buildStub: func(store *mockdb.MockStore) {
				//這裡手動模擬API處理參數
				store.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Email, user.UserID, -time.Minute)
			},
			checkReponse: func(t *testing.T, res *pb.UpdateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "No Auth",
			body: &pb.UpdateUserRequest{
				UserId:       user.UserID,
				UserName:     &user.UserName,
				SsoIdentifer: &user.SsoIdentifer.String,
			},
			buildStub: func(store *mockdb.MockStore) {
				//這裡手動模擬API處理參數
				store.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return context.Background()
			},
			checkReponse: func(t *testing.T, res *pb.UpdateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "Invalid Arg",
			body: &pb.UpdateUserRequest{
				UserId:       user.UserID,
				UserName:     &invalidName,
				SsoIdentifer: &invalidSSO,
			},
			buildStub: func(store *mockdb.MockStore) {
				//這裡手動模擬API處理參數
				store.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)

			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, user.Email, user.UserID, time.Minute)
			},
			checkReponse: func(t *testing.T, res *pb.UpdateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			//建立mockdb
			//mockStore 裡面包含了所有store行為的介面  且你可以對所有介面設定其stub行為
			store := mockdb.NewMockStore(ctrl)

			tc.buildStub(store)
			//grpc server可以直接call func
			server := newTestServer(t, store, nil)

			ctx := tc.buildContext(t, server.tokenMaker)

			res, err := server.UpdateUser(ctx, tc.body)
			tc.checkReponse(t, res, err)
		})
	}
}
