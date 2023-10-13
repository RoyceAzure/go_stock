package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/RoyceAzure/go-stockinfo-project/db/mock"
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

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
					GetUser(gomock.Any(), gomock.Eq(user.UserID)).
					Times(1).Return(user, nil)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
				requireBodyNatchUser(t, recoder.Body, user)
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
			server := NewServer(store)
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
		Email:          utility.RandomString(10),
		SsoIdentifer:   utility.StringToSqlNiStr(utility.RandomSSOTypeStr()),
		CrDate:         time.Now().UTC(),
		CrUser:         "royce",
	}
}

// 所以*bytes.Buffer也是io.Reader
// body 是個joson 格式的encoding 資料
func requireBodyNatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user, gotUser)
}
