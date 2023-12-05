package gapi

import (
	"context"
	"errors"
	"testing"

	"github.com/RoyceAzure/go-stockinfo-scheduler/api/pb"
	mock_repo "github.com/RoyceAzure/go-stockinfo-scheduler/repository/mock"
	repository "github.com/RoyceAzure/go-stockinfo-scheduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-scheduler/util/config"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
交互測試
TODO : 使用mock
*/
func TestGetStockDayAvg(t *testing.T) {

	testCase := []struct {
		name         string //子測試名稱
		buildStub    func(mockDao *mock_repo.MockDao)
		checkReponse func(*testing.T, *pb.StockDayAvgResponse, error)
	}{
		{
			name: "get SDA failed",
			buildStub: func(mockDao *mock_repo.MockDao) {
				var res []repository.StockDayAvgAll
				mockDao.EXPECT().GetSDAVGALLs(gomock.Any(), gomock.Any()).
					Return(res, errors.New("get SDA failed")).
					Times(1)
			},
			checkReponse: func(t *testing.T, res *pb.StockDayAvgResponse, err error) {
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, s.Code())
			},
		},
		{
			name: "no SDA",
			buildStub: func(mockDao *mock_repo.MockDao) {
				var res []repository.StockDayAvgAll
				mockDao.EXPECT().GetSDAVGALLs(gomock.Any(), gomock.Any()).
					Return(res, nil).
					Times(1)
			},
			checkReponse: func(t *testing.T, res *pb.StockDayAvgResponse, err error) {
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, s.Code())

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
			dao := mock_repo.NewMockDao(ctrl)

			tc.buildStub(dao)
			config := config.Config{}

			//grpc server可以直接call func
			// jr := jredis.NewJredis(config)

			// jservice := redisService.NewJRedisService(jr)
			server := newTestServer(t, config, dao, nil, nil)

			res, err := server.GetStockDayAvg(context.Background(), &pb.StockDayAvgRequest{})
			tc.checkReponse(t, res, err)
		})
	}
}
