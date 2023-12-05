// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/RoyceAzure/go-stockinfo/api/pb (interfaces: StockInfoSchdulerClient)

// Package mock_scheduler_client is a generated GoMock package.
package mock_scheduler_client

import (
	context "context"
	reflect "reflect"

	pb "github.com/RoyceAzure/go-stockinfo/api/pb"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockStockInfoSchdulerClient is a mock of StockInfoSchdulerClient interface.
type MockStockInfoSchdulerClient struct {
	ctrl     *gomock.Controller
	recorder *MockStockInfoSchdulerClientMockRecorder
}

// MockStockInfoSchdulerClientMockRecorder is the mock recorder for MockStockInfoSchdulerClient.
type MockStockInfoSchdulerClientMockRecorder struct {
	mock *MockStockInfoSchdulerClient
}

// NewMockStockInfoSchdulerClient creates a new mock instance.
func NewMockStockInfoSchdulerClient(ctrl *gomock.Controller) *MockStockInfoSchdulerClient {
	mock := &MockStockInfoSchdulerClient{ctrl: ctrl}
	mock.recorder = &MockStockInfoSchdulerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStockInfoSchdulerClient) EXPECT() *MockStockInfoSchdulerClientMockRecorder {
	return m.recorder
}

// GetStockDayAvg mocks base method.
func (m *MockStockInfoSchdulerClient) GetStockDayAvg(arg0 context.Context, arg1 *pb.StockDayAvgRequest, arg2 ...grpc.CallOption) (*pb.StockDayAvgResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStockDayAvg", varargs...)
	ret0, _ := ret[0].(*pb.StockDayAvgResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStockDayAvg indicates an expected call of GetStockDayAvg.
func (mr *MockStockInfoSchdulerClientMockRecorder) GetStockDayAvg(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStockDayAvg", reflect.TypeOf((*MockStockInfoSchdulerClient)(nil).GetStockDayAvg), varargs...)
}
