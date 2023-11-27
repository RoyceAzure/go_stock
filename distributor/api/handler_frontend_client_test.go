package api

import (
	"bytes"
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

func TestCreateFrontendClient(t *testing.T) {
	fc := randomFC()
	testhostIP := "127.0.0.1:"
	testIP := "127.0.0.1"
	arg := repository.CreateFrontendClientParams{
		Ip:     testIP,
		Region: fc.Region,
	}
	taseCase := []struct {
		name         string //子測試名稱
		body         gin.H
		setUpReqIP   func(req *http.Request)
		buildStub    func(dbDao *mock_repository.MockDistributorDao)
		checkReponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"region": fc.Region,
			},
			setUpReqIP: func(req *http.Request) {
				req.RemoteAddr = testhostIP
			},
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					CreateFrontendClient(gomock.Any(), arg).
					Times(1).
					Return(fc, nil)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusAccepted, recoder.Code)
			},
		},
		{
			name: "ip is empty",
			body: gin.H{
				"region": fc.Region,
			},
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					CreateFrontendClient(gomock.Any(), arg).
					Times(0)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "invalid ip format",
			body: gin.H{
				"region": fc.Region,
			},
			setUpReqIP: func(req *http.Request) {
				req.RemoteAddr = "123456789"
			},
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					CreateFrontendClient(gomock.Any(), arg).
					Times(0)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "ip is exists",
			body: gin.H{
				"region": fc.Region,
			},
			setUpReqIP: func(req *http.Request) {
				req.RemoteAddr = testhostIP
			},
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					CreateFrontendClient(gomock.Any(), arg).
					Times(1).
					Return(repository.FrontendClient{}, repository.ErrUniqueViolation)
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

			url := "/frontend_client"

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

func TestDeleteFrontendClient(t *testing.T) {
	fc := randomFC()
	taseCase := []struct {
		name         string //子測試名稱
		clientID     string
		buildStub    func(dbDao *mock_repository.MockDistributorDao)
		checkReponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:     "OK",
			clientID: fc.ClientUid.String(),
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					DeleteFrontendClient(gomock.Any(), fc.ClientUid).
					Times(1).
					Return(nil)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusAccepted, recoder.Code)
			},
		},
		{
			name:     "invalid uuid",
			clientID: "123456789",
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					DeleteFrontendClient(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
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

			url := fmt.Sprintf("/frontend_client/%s", tc.clientID)

			req, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recoder, req)
			tc.checkReponse(t, recoder)
		})
	}
}

func TestGetFrontendClientByIP(t *testing.T) {
	fc := randomFC()
	testhostIP := "127.0.0.1:"
	testIP := "127.0.0.1"
	invalidIp := "12345678"
	taseCase := []struct {
		name         string //子測試名稱
		setUpReqIP   func(req *http.Request)
		buildStub    func(dbDao *mock_repository.MockDistributorDao)
		checkReponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setUpReqIP: func(req *http.Request) {
				req.RemoteAddr = testhostIP
			},
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					GetFrontendClientByIP(gomock.Any(), testIP).
					Times(1).
					Return(fc, nil)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusAccepted, recoder.Code)
			},
		},
		{
			name: "invalid ip",
			setUpReqIP: func(req *http.Request) {
				req.RemoteAddr = invalidIp + ":"
			},
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					GetFrontendClientByIP(gomock.Any(), invalidIp).
					Times(0)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "empty ip",
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					GetFrontendClientByIP(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "not found",
			setUpReqIP: func(req *http.Request) {
				req.RemoteAddr = testhostIP
			},
			buildStub: func(dbDao *mock_repository.MockDistributorDao) {
				//這裡手動模擬API處理參數
				dbDao.EXPECT().
					GetFrontendClientByIP(gomock.Any(), testIP).
					Times(1).
					Return(repository.FrontendClient{}, repository.ErrRecordNotFound)
			},
			checkReponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recoder.Code)
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

			url := "/frontend_client"

			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			if tc.setUpReqIP != nil {
				tc.setUpReqIP(req)
			}
			server.router.ServeHTTP(recoder, req)
			tc.checkReponse(t, recoder)
		})
	}
}

func randomFC() repository.FrontendClient {
	return repository.FrontendClient{
		ClientUid: uuid.New(),
		Ip:        random.RandomString(5),
		Region:    random.RandomString(5),
		CreatedAt: time.Now().UTC(),
	}
}
