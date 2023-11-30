package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mock_repository "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/mock"
	repository "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-distributor/shared/util/random"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateClientRegister(t *testing.T) {
	cr := randomCR()
	testhostIP := "127.0.0.1:"
	testIP := "127.0.0.1"
	taseCase := []struct {
		name         string //子測試名稱
		body         gin.H
		setUpReqIP   func(req *http.Request)
		buildStub    func(dbDao *mock_repository.MockDistributorDao)
		checkReponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "with client id, client not exists",
			body: gin.H{
				"client_uid": cr.ClientUid,
				"stock_code": cr.StockCode,
			},
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					GetFrontendClientByIP(gomock.Any(), gomock.Any()).
					Times(0)
				dbDao.EXPECT().
					GetFrontendClientByID(gomock.Any(), gomock.Eq(cr.ClientUid)).
					Times(1).
					Return(repository.FrontendClient{}, sql.ErrNoRows)
				dbDao.EXPECT().
					CreateClientRegister(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "with ip, client not exists",
			body: gin.H{
				"stock_code": cr.StockCode,
			},
			setUpReqIP: func(req *http.Request) {
				req.RemoteAddr = testhostIP
			},
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				dbDao.EXPECT().
					GetFrontendClientByIP(gomock.Any(), testIP).
					Times(1).
					Return(repository.FrontendClient{}, sql.ErrNoRows)
				dbDao.EXPECT().
					GetFrontendClientByID(gomock.Any(), gomock.Any()).
					Times(0)
				dbDao.EXPECT().
					CreateClientRegister(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "unique violation",
			body: gin.H{
				"stock_code": cr.StockCode,
			},
			setUpReqIP: func(req *http.Request) {
				req.RemoteAddr = testhostIP
			},
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				dbDao.EXPECT().
					GetFrontendClientByIP(gomock.Any(), testIP).
					Times(1).
					Return(repository.FrontendClient{}, nil)
				dbDao.EXPECT().
					GetFrontendClientByID(gomock.Any(), gomock.Any()).
					Times(0)
				dbDao.EXPECT().
					CreateClientRegister(gomock.Any(), gomock.Any()).
					Times(1).
					Return(repository.ClientRegister{}, repository.ErrUniqueViolation)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recoder.Code)
			},
		},
	}

	for _, tc := range taseCase {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockDao := mock_repository.NewMockDistributorDao(ctrl)
			tc.buildStub(mockDao)

			server := NewServer(mockDao, nil)

			recoder := httptest.NewRecorder()
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/client_register"

			req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			if tc.setUpReqIP != nil {
				tc.setUpReqIP(req)
			}
			server.router.ServeHTTP(recoder, req)
			tc.checkReponse(t, recoder)
		})
	}
}

func TestGetClientRegisterByClientUID(t *testing.T) {
	cr := randomCR()

	testRes := createSpcitCRList(cr.ClientUid, "0050", "0051", "0052")
	taseCase := []struct {
		name         string //子測試名稱
		clientId     string
		buildStub    func(dbDao *mock_repository.MockDistributorDao, res []repository.ClientRegister)
		checkReponse func(t *testing.T, recoder *httptest.ResponseRecorder, exceped []repository.ClientRegister)
	}{
		{
			name:     "ok",
			clientId: "/" + cr.ClientUid.String(),
			buildStub: func(dbDao *mock_repository.MockDistributorDao, res []repository.ClientRegister) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					GetClientRegisterByClientUID(gomock.Any(), gomock.Any()).
					Times(1).Return(res, nil)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder, exceped []repository.ClientRegister) {
				require.Equal(t, http.StatusAccepted, recoder.Code)
				var res GetClientRegisterByClientUIDResponse
				err := json.Unmarshal(recoder.Body.Bytes(), &res)
				require.NoError(t, err)

				require.Len(t, exceped, len(res.Result))
				for i := range exceped {
					require.Equal(t, exceped[i].ClientUid, res.Result[i].ClientUid)
				}
			},
		},
		{
			name:     "with invalid client id",
			clientId: "/12345889",
			buildStub: func(dbDao *mock_repository.MockDistributorDao, res []repository.ClientRegister) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					GetClientRegisterByClientUID(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder, exceped []repository.ClientRegister) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for _, tc := range taseCase {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockDao := mock_repository.NewMockDistributorDao(ctrl)
			tc.buildStub(mockDao, testRes)

			server := NewServer(mockDao, nil)

			recoder := httptest.NewRecorder()

			url := "/client_register" + fmt.Sprintf("%s", tc.clientId)

			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recoder, req)
			tc.checkReponse(t, recoder, testRes)
		})
	}
}

func randomCR() repository.ClientRegister {
	return repository.ClientRegister{
		ClientUid: uuid.New(),
		StockCode: random.RandomString(5),
		CreatedAt: time.Now().UTC(),
	}
}

func createSpcitCR(uid uuid.UUID, stockCode string) repository.ClientRegister {
	return repository.ClientRegister{
		ClientUid: uid,
		StockCode: stockCode,
		CreatedAt: time.Now().UTC(),
	}
}

func createSpcitCRList(uid uuid.UUID, stockCode ...string) []repository.ClientRegister {
	var res []repository.ClientRegister
	for _, code := range stockCode {
		res = append(res, repository.ClientRegister{
			ClientUid: uid,
			StockCode: code,
			CreatedAt: time.Now().UTC(),
		})
	}
	return res
}
