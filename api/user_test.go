package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	mockdb "github.com/RoyceAzure/go-stockinfo-project/db/mock"
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

// password為明碼
type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e *eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := utility.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}
	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e *eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v, password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return &eqCreateUserParamsMatcher{arg, password}
}

// 隨便創建一個user
// 建立mock ctrl
// 建立mock store
// 設定要測試的func stub  接收user id
// 建立server
// 建立response
// 建立req  發送固定request 並且帶入userid
// 目的在檢查req發送過程，資料接收過程，以及API (controller層)是否有正確把stub給的user回傳
func TestGetUserApi(t *testing.T) {
	user := randomUser()

	testCase := []struct {
		name         string //子測試名稱
		userId       int64
		buildStub    func(store *mockdb.MockStore)
		checkReponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			userId: user.UserID,
			//把測試流程中的buildStub過程與checkReponse獨立出來成匿名function
			//那你要分離出來自定義的流程裡面，所需的物件都要變成參數傳入
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					//注意參數gomock.Any(), 雖然GetUser是mock, 但是API層有要負責處理store的參數
					//所以處理的store參數也必須接受驗證，不然萬一你API層處理參數有問題，會抓不出error
					GetUser(gomock.Any(), gomock.Eq(user.UserID)).
					Times(1).Return(user, nil)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
				requireBodyMatchUser(t, recoder.Body, user)
			},
		},
		{
			name:   "NotFound",
			userId: user.UserID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.UserID)).
					Times(1).Return(db.User{}, sql.ErrNoRows)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
			},
		},
		{
			name:   "InternalServerError",
			userId: user.UserID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.UserID)).
					Times(1).Return(db.User{}, sql.ErrConnDone)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name:   "InvalidID",
			userId: 0,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					//使用any無法dected錯誤的參數
					GetUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
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
			//使用Gin  *gin.Engine建立server
			//new Server已經把所有的路由都設定好
			server := newTestServer(t, store)
			//Response Recoder是做甚麼的?
			//ResponseRecorder 是一个实现了 http.ResponseWriter 接口的类型
			recoder := httptest.NewRecorder()

			url := fmt.Sprintf("/user/%d", tc.userId)
			request, err := http.NewRequest(http.MethodGet, url, nil)

			require.NoError(t, err)

			//這裡的router實際上是 *gin.Engine
			//自己發送自己接收?
			server.router.ServeHTTP(recoder, request)
			tc.checkReponse(t, recoder)
		})
	}
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

func randomUserDTO() createUserRequest {
	return createUserRequest{
		UserName:     utility.RandomString(10),
		Password:     utility.RandomString(10),
		Email:        utility.RandomString(10) + "@gmail.com",
		SsoIdentifer: utility.RandomSSOTypeStr(),
	}
}

// 所以*bytes.Buffer也是io.Reader
// body 是個joson 格式的encoding 資料
// 注意  body將會是DTO  而這裡的user是model  所以有點問題  這裡應該要做DTO相等的比較
func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser UserResponseDTO
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, newUserResponse(user), gotUser)
}

func requireBodyMatchLoginUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var loginRes loginUserResponse
	err = json.Unmarshal(data, &loginRes)
	require.NoError(t, err)
	require.Equal(t, newUserResponse(user), loginRes.User)
}

func TestCreateUserApi(t *testing.T) {
	// 或許在測試API時  randomUser就應該產生DTO 而不是model
	user := randomUser()

	testCase := []struct {
		name         string //子測試名稱
		body         gin.H
		buildStub    func(store *mockdb.MockStore)
		checkReponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"user_name":     user.UserName,
				"email":         user.Email,
				"password":      user.HashedPassword,
				"sso_identifer": user.SsoIdentifer.String,
			},
			buildStub: func(store *mockdb.MockStore) {
				//這裡手動模擬API處理參數
				arg := db.CreateUserParams{
					UserName:       user.UserName,
					Email:          user.Email,
					HashedPassword: user.HashedPassword,
					SsoIdentifer:   user.SsoIdentifer,
					CrUser:         "SYSTEM",
				}
				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(arg, user.HashedPassword)).
					Times(1).Return(user, nil)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusAccepted, recoder.Code)
				requireBodyMatchUser(t, recoder.Body, user)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"user_name":     user.UserName,
				"email":         "123456",
				"password":      user.HashedPassword,
				"sso_identifer": "GOOGLE",
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "UniqueViolation",
			body: gin.H{
				"user_name":     user.UserName,
				"email":         user.Email,
				"password":      user.HashedPassword,
				"sso_identifer": "GOOGLE",
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).Return(db.User{}, &pq.Error{Code: "23505"})
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recoder.Code)
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
			//使用Gin  *gin.Engine建立server
			//new Server已經把所有的路由都設定好
			server := newTestServer(t, store)
			//Response Recoder是做甚麼的?
			//ResponseRecorder 是一个实现了 http.ResponseWriter 接口的类型
			recoder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/user"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))

			require.NoError(t, err)

			//這裡的router實際上是 *gin.Engine
			//自己發送自己接收?
			server.router.ServeHTTP(recoder, request)
			tc.checkReponse(t, recoder)
		})
	}
}

func TestUserLoginApi(t *testing.T) {
	// 或許在測試API時  randomUser就應該產生DTO 而不是model
	user := randomUser()
	hashed_password, err := utility.HashPassword(user.HashedPassword)
	dbuer := user
	dbuer.HashedPassword = hashed_password
	require.NoError(t, err)
	testCase := []struct {
		name         string //子測試名稱
		body         gin.H
		buildStub    func(store *mockdb.MockStore)
		checkReponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"user_name":     user.UserName,
				"email":         user.Email,
				"password":      user.HashedPassword,
				"sso_identifer": user.SsoIdentifer.String,
			},
			buildStub: func(store *mockdb.MockStore) {
				//這裡手動模擬API處理參數
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).Return(dbuer, nil)
				store.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
					Times(1)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
				requireBodyMatchLoginUser(t, recoder.Body, dbuer)
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
			//使用Gin  *gin.Engine建立server
			//new Server已經把所有的路由都設定好
			server := newTestServer(t, store)
			//Response Recoder是做甚麼的?
			//ResponseRecorder 是一个实现了 http.ResponseWriter 接口的类型
			recoder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/user/login"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))

			require.NoError(t, err)

			//這裡的router實際上是 *gin.Engine
			//自己發送自己接收?
			server.router.ServeHTTP(recoder, request)
			tc.checkReponse(t, recoder)
		})
	}
}
