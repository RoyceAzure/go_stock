package gapi

import (
	"context"
	"errors"
	"testing"

	"github.com/RoyceAzure/go-stockinfo-scheduler/api/pb"
	mock_repo "github.com/RoyceAzure/go-stockinfo-scheduler/repository/mock"
	mock_redis_dao "github.com/RoyceAzure/go-stockinfo-scheduler/repository/redis/mock"
	mock_redis_service "github.com/RoyceAzure/go-stockinfo-scheduler/service/redisService/mock"
	"github.com/RoyceAzure/go-stockinfo-scheduler/util/config"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

func TestGetStockPriceRealTime(t *testing.T) {
	testCase := []struct {
		name         string //子測試名稱
		buildStub    func(redisService *mock_redis_service.MockRedisService, dao *mock_redis_dao.MockJRedisDao)
		checkReponse func(*testing.T, *pb.StockPriceRealTimeResponse, error)
	}{
		{
			name: "get spr err",
			buildStub: func(redisService *mock_redis_service.MockRedisService, dao *mock_redis_dao.MockJRedisDao) {
				redisService.EXPECT().GetLatestSPR(gomock.Any()).
					Return(nil, "", errors.New("error")).
					Times(1)
			},
			checkReponse: func(t *testing.T, res *pb.StockPriceRealTimeResponse, err error) {
				require.Error(t, err)
				s, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, s.Code())
				require.Nil(t, res)
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
			defer ctrl.Finish()
			config := config.Config{}

			mockRedisDao := mock_redis_dao.NewMockJRedisDao(ctrl)
			mockRedisSesrevice := mock_redis_service.NewMockRedisService(ctrl)
			//grpc server可以直接call func

			tc.buildStub(mockRedisSesrevice, mockRedisDao)
			server := newTestServer(t, config, dao, nil, mockRedisSesrevice)
			res, err := server.GetStockPriceRealTime(context.Background(), &pb.StockPriceRealTimeRequest{})
			tc.checkReponse(t, res, err)
		})
	}
}
