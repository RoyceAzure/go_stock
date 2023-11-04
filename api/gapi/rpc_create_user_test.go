package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-api/pb"
	mockdb "github.com/RoyceAzure/go-stockinfo-project/db/mock"
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	worker "github.com/RoyceAzure/go-stockinfo-worker"
	mockwk "github.com/RoyceAzure/go-stockinfo-worker/mock"
	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// password為明碼
type eqCreateUserParamsTxMatcher struct {
	arg      db.CreateUserTxParams
	password string
	user     db.User
}

/*
時做gomock.Matcher  來給gomock使用  目的是比較hashed password
*/
func (excepted *eqCreateUserParamsTxMatcher) Matches(x interface{}) bool {
	actuelArg, ok := x.(db.CreateUserTxParams)
	if !ok {
		return false
	}

	err := utility.CheckPassword(excepted.password, actuelArg.HashedPassword)
	if err != nil {
		return false
	}
	excepted.arg.HashedPassword = actuelArg.HashedPassword
	err = actuelArg.AfterCreate(excepted.user)
	if err != nil {
		return false
	}
	return reflect.DeepEqual(excepted.arg.CreateUserParams, actuelArg.CreateUserParams)
}

func (excepted *eqCreateUserParamsTxMatcher) String() string {
	return fmt.Sprintf("matches arg %v, password %v", excepted.arg, excepted.password)
}

func EqCreateUserTxParams(arg db.CreateUserTxParams, password string, user db.User) gomock.Matcher {
	return &eqCreateUserParamsTxMatcher{arg, password, user}
}

func randomUser() db.User {
	return db.User{
		UserID:         utility.RandomInt(1, 100),
		UserName:       utility.RandomString(10),
		HashedPassword: utility.RandomString(10),
		Email:          utility.RandomString(10) + "@gmail.com",
		SsoIdentifer:   utility.StringToSqlNiStr(utility.RandomSSOTypeStr()),
		CrDate:         time.Now().UTC(),
		CrUser:         "royce",
	}
}

func TestCreateUserApi(t *testing.T) {
	// 或許在測試API時  randomUser就應該產生DTO 而不是model
	user := randomUser()

	testCase := []struct {
		name         string //子測試名稱
		body         *pb.CreateUserRequest
		buildStub    func(store *mockdb.MockStore, taskDistributor *mockwk.MockTaskDistributor)
		checkReponse func(*testing.T, *pb.CreateUserResponse, error)
	}{
		{
			name: "OK",
			body: &pb.CreateUserRequest{
				UserName:     user.UserName,
				Email:        user.Email,
				Password:     user.HashedPassword,
				SsoIdentifer: user.SsoIdentifer.String,
			},
			buildStub: func(store *mockdb.MockStore, taskDistributor *mockwk.MockTaskDistributor) {
				//這裡手動模擬API處理參數
				arg := db.CreateUserTxParams{
					CreateUserParams: db.CreateUserParams{
						UserName:       user.UserName,
						Email:          user.Email,
						HashedPassword: user.HashedPassword,
						SsoIdentifer:   user.SsoIdentifer,
						CrUser:         "SYSTEM",
					},
				}
				store.EXPECT().
					CreateUserTx(gomock.Any(), EqCreateUserTxParams(arg, user.HashedPassword, user)).
					Times(1).Return(db.CreateUserTxResults{
					User: user,
				}, nil)

				taskPayload := &worker.PayloadSendVerifyEmail{
					UserName: user.UserName,
					UserId:   user.UserID,
				}

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), taskPayload, gomock.Any()).
					Times(1).
					Return(nil)
			},
			checkReponse: func(t *testing.T, res *pb.CreateUserResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				createUser := res.GetUser()
				require.Equal(t, createUser.GetEmail(), user.Email)
				require.Equal(t, createUser.GetUserName(), user.UserName)
				require.Equal(t, createUser.GetSsoIdentifer(), user.SsoIdentifer.String)

			},
		},
		{
			name: "Internal Server err",
			body: &pb.CreateUserRequest{
				UserName:     user.UserName,
				Email:        user.Email,
				Password:     user.HashedPassword,
				SsoIdentifer: user.SsoIdentifer.String,
			},
			buildStub: func(store *mockdb.MockStore, taskDistributor *mockwk.MockTaskDistributor) {
				//這裡手動模擬API Layer 處理repo layer 參數
				store.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(1).Return(db.CreateUserTxResults{}, sql.ErrConnDone)

				taskPayload := &worker.PayloadSendVerifyEmail{
					UserName: user.UserName,
					UserId:   user.UserID,
				}

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), taskPayload, gomock.Any()).
					Times(0)
			},
			checkReponse: func(t *testing.T, res *pb.CreateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())

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

			taskCtrl := gomock.NewController(t)

			defer ctrl.Finish()

			taskDistributor := mockwk.NewMockTaskDistributor(taskCtrl)

			tc.buildStub(store, taskDistributor)

			//grpc server可以直接call func
			server := newTestServer(t, store, taskDistributor)
			res, err := server.CreateUser(context.Background(), tc.body)
			tc.checkReponse(t, res, err)
		})
	}
}
