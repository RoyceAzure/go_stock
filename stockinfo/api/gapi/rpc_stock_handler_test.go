package gapi

import (
	"context"
	"errors"
	"testing"

	mock_gapi "github.com/RoyceAzure/go-stockinfo/api/gapi/mock"
	"github.com/RoyceAzure/go-stockinfo/api/token"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	mock_scheduler_client "github.com/RoyceAzure/go-stockinfo/shared/pb/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestInitStock(t *testing.T) {
	testCase := []struct {
		name         string //子測試名稱
		buildStub    func(*mock_gapi.MockSchedulerClientFactory, *mock_scheduler_client.MockStockInfoSchdulerClient)
		buildContext func(*testing.T, token.Maker) context.Context
		checkReponse func(*testing.T, *pb.InitStockResponse, error)
	}{
		{
			name: "connect failed",
			buildStub: func(factory *mock_gapi.MockSchedulerClientFactory, client *mock_scheduler_client.MockStockInfoSchdulerClient) {
				factory.EXPECT().NewClient().Return(nil, nil, errors.New(codes.Internal.String())).
					Times(1)
			},
			checkReponse: func(t *testing.T, res *pb.InitStockResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "get SDA failed",
			buildStub: func(factory *mock_gapi.MockSchedulerClientFactory, client *mock_scheduler_client.MockStockInfoSchdulerClient) {
				client.EXPECT().GetStockDayAvg(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("get SDA failed"))
				factory.EXPECT().NewClient().Return(&client, func() {}, nil).
					Times(1)
			},
			checkReponse: func(t *testing.T, res *pb.InitStockResponse, err error) {
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
			var factory *mock_gapi.MockSchedulerClientFactory
			var ctx context.Context
			//建立mockdb
			//mockStore 裡面包含了所有store行為的介面  且你可以對所有介面設定其stub行為
			if tc.buildStub != nil {
				factory = mock_gapi.NewMockSchedulerClientFactory(ctrl)
				client := mock_scheduler_client.NewMockStockInfoSchdulerClient(ctrl)
				tc.buildStub(factory, client)
			}

			//grpc server可以直接call func
			server := newTestServer(t, nil, nil, factory)

			if tc.buildContext != nil {
				ctx = tc.buildContext(t, server.tokenMaker)
			}

			res, err := server.InitStock(ctx, nil)
			tc.checkReponse(t, res, err)
		})
	}
}
