package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/RoyceAzure/go-stockinfo-api/token"
	mockdb "github.com/RoyceAzure/go-stockinfo-project/db/mock"
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func randomFund() (db.User, db.Fund) {
	user := randomUser() //假設db存在的user  並沒有寫入db
	fund := CreateRandomFund(user)
	return user, fund
}

func CreateRandomFund(user db.User) db.Fund {
	return db.Fund{
		FundID:       utility.RandomInt(1, 100),
		UserID:       user.UserID,
		Balance:      strconv.FormatInt(utility.RandomInt(10000, 1000000), 10),
		CurrencyType: utility.RandomCurrencyTypeStr(),
		CrDate:       time.Now().UTC(),
		CrUser:       "SYSTEM",
	}
}

func requireBodyMatchFund(t *testing.T, body *bytes.Buffer, fund db.Fund) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)
	var fundRes FundResponseDTO
	err = json.Unmarshal(data, &fundRes)
	require.NoError(t, err)
	require.Equal(t, fundRes, fundRes)
}
func TestGetFund(t *testing.T) {
	user, fund := randomFund()

	testCases := []struct {
		name         string
		fundId       int64
		setupAuth    func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub    func(store *mockdb.MockStore)
		checkReponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			fundId: fund.FundID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				AddAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			//把測試流程中的buildStub過程與checkReponse獨立出來成匿名function
			//那你要分離出來自定義的流程裡面，所需的物件都要變成參數傳入
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					//注意參數gomock.Any(), 雖然GetUser是mock, 但是API層有要負責處理store的參數
					//所以處理的store參數也必須接受驗證，不然萬一你API層處理參數有問題，會抓不出error
					GetFund(gomock.Any(), gomock.Eq(fund.FundID)).
					Times(1).Return(fund, nil)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
				requireBodyMatchFund(t, recoder.Body, fund)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			//gomock.Controller。這個 ctrl 或 Controller 負責控制模擬物件的生命週期以及其期望行為的驗證
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)

			tc.buildStub(store)

			server := newTestServer(t, store)
			recoder := httptest.NewRecorder()
			url := fmt.Sprintf("/fund/%d", tc.fundId)
			request, err := http.NewRequest(http.MethodGet, url, nil)

			tc.setupAuth(t, request, server.tokenMaker)

			require.NoError(t, err)

			server.router.ServeHTTP(recoder, request)
			tc.checkReponse(t, recoder)
		})
	}
}
