package gapi

import (
	"context"
	"testing"

	"github.com/RoyceAzure/go-stockinfo/api/token"
	mockdb "github.com/RoyceAzure/go-stockinfo/repository/db/mock"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetRealizedProfitLoss(t *testing.T) {
	// user := randomUser()
	testCase := []struct {
		name         string //子測試名稱
		arg          *pb.GetRealizedProfitLossRequest
		buildContext func(*testing.T, token.Maker) context.Context
		buildStub    func(store *mockdb.MockStore)
		checkReponse func(t *testing.T, res *pb.GetRealizedProfitLossResponse, err error)
	}{
		{
			name: "empty auth",
			arg:  &pb.GetRealizedProfitLossRequest{},
			buildStub: func(store *mockdb.MockStore) {
				//這裡手動模擬API處理參數
				store.EXPECT().
					GetRealizedProfitLosssByUserIdDetial(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(*testing.T, token.Maker) context.Context {
				return context.Background()
			},
			checkReponse: func(t *testing.T, res *pb.GetRealizedProfitLossResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
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

			if tc.buildStub != nil {
				tc.buildStub(store)
			}
			//grpc server可以直接call func
			server := newTestServer(t, store, nil, nil)
			ctx := tc.buildContext(t, server.tokenMaker)
			res, err := server.GetRealizedProfitLoss(ctx, tc.arg)
			tc.checkReponse(t, res, err)
		})
	}
}
