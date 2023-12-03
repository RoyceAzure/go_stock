// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/RoyceAzure/go-stockinfo-project/db/sqlc (interfaces: Store)

// Package mock_sqlc is a generated GoMock package.
package mock_sqlc

import (
        context "context"
        reflect "reflect"

        db "github.com/RoyceAzure/go-stockinfo/project/db/sqlc"
        gomock "github.com/golang/mock/gomock"
        uuid "github.com/google/uuid"

)

// MockStore is a mock of Store interface.
type MockStore struct {
        ctrl     *gomock.Controller
        recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
        mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
        mock := &MockStore{ctrl: ctrl}
        mock.recorder = &MockStoreMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
        return m.recorder
}

// CreateFund mocks base method.
func (m *MockStore) CreateFund(arg0 context.Context, arg1 db.CreateFundParams) (db.Fund, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateFund", arg0, arg1)
        ret0, _ := ret[0].(db.Fund)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateFund indicates an expected call of CreateFund.
func (mr *MockStoreMockRecorder) CreateFund(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFund", reflect.TypeOf((*MockStore)(nil).CreateFund), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
        ret0, _ := ret[0].(db.Session)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateStock mocks base method.
func (m *MockStore) CreateStock(arg0 context.Context, arg1 db.CreateStockParams) (db.Stock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateStock", arg0, arg1)
        ret0, _ := ret[0].(db.Stock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateStock indicates an expected call of CreateStock.
func (mr *MockStoreMockRecorder) CreateStock(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStock", reflect.TypeOf((*MockStore)(nil).CreateStock), arg0, arg1)
}

// CreateStockTransaction mocks base method.
func (m *MockStore) CreateStockTransaction(arg0 context.Context, arg1 db.CreateStockTransactionParams) (db.StockTransaction, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateStockTransaction", arg0, arg1)
        ret0, _ := ret[0].(db.StockTransaction)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateStockTransaction indicates an expected call of CreateStockTransaction.
func (mr *MockStoreMockRecorder) CreateStockTransaction(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStockTransaction", reflect.TypeOf((*MockStore)(nil).CreateStockTransaction), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
        ret0, _ := ret[0].(db.User)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateUserStock mocks base method.
func (m *MockStore) CreateUserStock(arg0 context.Context, arg1 db.CreateUserStockParams) (db.UserStock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateUserStock", arg0, arg1)
        ret0, _ := ret[0].(db.UserStock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateUserStock indicates an expected call of CreateUserStock.
func (mr *MockStoreMockRecorder) CreateUserStock(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserStock", reflect.TypeOf((*MockStore)(nil).CreateUserStock), arg0, arg1)       
}

// CreateUserTx mocks base method.
func (m *MockStore) CreateUserTx(arg0 context.Context, arg1 db.CreateUserTxParams) (db.CreateUserTxResults, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateUserTx", arg0, arg1)
        ret0, _ := ret[0].(db.CreateUserTxResults)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateUserTx indicates an expected call of CreateUserTx.
func (mr *MockStoreMockRecorder) CreateUserTx(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserTx", reflect.TypeOf((*MockStore)(nil).CreateUserTx), arg0, arg1)
}

// CreateVerifyEmail mocks base method.
func (m *MockStore) CreateVerifyEmail(arg0 context.Context, arg1 db.CreateVerifyEmailParams) (db.VerifyEmail, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateVerifyEmail", arg0, arg1)
        ret0, _ := ret[0].(db.VerifyEmail)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateVerifyEmail indicates an expected call of CreateVerifyEmail.
func (mr *MockStoreMockRecorder) CreateVerifyEmail(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVerifyEmail", reflect.TypeOf((*MockStore)(nil).CreateVerifyEmail), arg0, arg1)   
}

// DeleteFund mocks base method.
func (m *MockStore) DeleteFund(arg0 context.Context, arg1 int64) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteFund", arg0, arg1)
        ret0, _ := ret[0].(error)
        return ret0
}

// DeleteFund indicates an expected call of DeleteFund.
func (mr *MockStoreMockRecorder) DeleteFund(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFund", reflect.TypeOf((*MockStore)(nil).DeleteFund), arg0, arg1)
}

// DeleteStock mocks base method.
func (m *MockStore) DeleteStock(arg0 context.Context, arg1 int64) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteStock", arg0, arg1)
        ret0, _ := ret[0].(error)
        return ret0
}

// DeleteStock indicates an expected call of DeleteStock.
func (mr *MockStoreMockRecorder) DeleteStock(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStock", reflect.TypeOf((*MockStore)(nil).DeleteStock), arg0, arg1)
}

// DeleteStockTransaction mocks base method.
func (m *MockStore) DeleteStockTransaction(arg0 context.Context, arg1 int64) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteStockTransaction", arg0, arg1)
        ret0, _ := ret[0].(error)
        return ret0
}

// DeleteStockTransaction indicates an expected call of DeleteStockTransaction.
func (mr *MockStoreMockRecorder) DeleteStockTransaction(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStockTransaction", reflect.TypeOf((*MockStore)(nil).DeleteStockTransaction), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int64) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
        ret0, _ := ret[0].(error)
        return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// DeleteUserStock mocks base method.
func (m *MockStore) DeleteUserStock(arg0 context.Context, arg1 int64) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteUserStock", arg0, arg1)
        ret0, _ := ret[0].(error)
        return ret0
}

// DeleteUserStock indicates an expected call of DeleteUserStock.
func (mr *MockStoreMockRecorder) DeleteUserStock(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserStock", reflect.TypeOf((*MockStore)(nil).DeleteUserStock), arg0, arg1)       
}

// GetFund mocks base method.
func (m *MockStore) GetFund(arg0 context.Context, arg1 int64) (db.Fund, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetFund", arg0, arg1)
        ret0, _ := ret[0].(db.Fund)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetFund indicates an expected call of GetFund.
func (mr *MockStoreMockRecorder) GetFund(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFund", reflect.TypeOf((*MockStore)(nil).GetFund), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
        ret0, _ := ret[0].(db.Session)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetStock mocks base method.
func (m *MockStore) GetStock(arg0 context.Context, arg1 int64) (db.Stock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetStock", arg0, arg1)
        ret0, _ := ret[0].(db.Stock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetStock indicates an expected call of GetStock.
func (mr *MockStoreMockRecorder) GetStock(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStock", reflect.TypeOf((*MockStore)(nil).GetStock), arg0, arg1)
}

// GetStockTransaction mocks base method.
func (m *MockStore) GetStockTransaction(arg0 context.Context, arg1 int64) (db.StockTransaction, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetStockTransaction", arg0, arg1)
        ret0, _ := ret[0].(db.StockTransaction)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetStockTransaction indicates an expected call of GetStockTransaction.
func (mr *MockStoreMockRecorder) GetStockTransaction(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStockTransaction", reflect.TypeOf((*MockStore)(nil).GetStockTransaction), arg0, arg1)
}

// GetStockTransactions mocks base method.
func (m *MockStore) GetStockTransactions(arg0 context.Context, arg1 db.GetStockTransactionsParams) ([]db.StockTransaction, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetStockTransactions", arg0, arg1)
        ret0, _ := ret[0].([]db.StockTransaction)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetStockTransactions indicates an expected call of GetStockTransactions.
func (mr *MockStoreMockRecorder) GetStockTransactions(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStockTransactions", reflect.TypeOf((*MockStore)(nil).GetStockTransactions), arg0, arg1)
}

// GetStockTransactionsByDate mocks base method.
func (m *MockStore) GetStockTransactionsByDate(arg0 context.Context, arg1 db.GetStockTransactionsByDateParams) ([]db.StockTransaction, error) {       
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetStockTransactionsByDate", arg0, arg1)
        ret0, _ := ret[0].([]db.StockTransaction)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetStockTransactionsByDate indicates an expected call of GetStockTransactionsByDate.
func (mr *MockStoreMockRecorder) GetStockTransactionsByDate(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStockTransactionsByDate", reflect.TypeOf((*MockStore)(nil).GetStockTransactionsByDate), arg0, arg1)
}

// GetStockTransactionsByStockId mocks base method.
func (m *MockStore) GetStockTransactionsByStockId(arg0 context.Context, arg1 db.GetStockTransactionsByStockIdParams) ([]db.StockTransaction, error) { 
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetStockTransactionsByStockId", arg0, arg1)
        ret0, _ := ret[0].([]db.StockTransaction)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetStockTransactionsByStockId indicates an expected call of GetStockTransactionsByStockId.
func (mr *MockStoreMockRecorder) GetStockTransactionsByStockId(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStockTransactionsByStockId", reflect.TypeOf((*MockStore)(nil).GetStockTransactionsByStockId), arg0, arg1)
}

// GetStockTransactionsByUserId mocks base method.
func (m *MockStore) GetStockTransactionsByUserId(arg0 context.Context, arg1 db.GetStockTransactionsByUserIdParams) ([]db.StockTransaction, error) {   
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetStockTransactionsByUserId", arg0, arg1)
        ret0, _ := ret[0].([]db.StockTransaction)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetStockTransactionsByUserId indicates an expected call of GetStockTransactionsByUserId.
func (mr *MockStoreMockRecorder) GetStockTransactionsByUserId(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStockTransactionsByUserId", reflect.TypeOf((*MockStore)(nil).GetStockTransactionsByUserId), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int64) (db.User, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
        ret0, _ := ret[0].(db.User)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockStore) GetUserByEmail(arg0 context.Context, arg1 string) (db.User, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
        ret0, _ := ret[0].(db.User)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserForUpdateNoKey mocks base method.
func (m *MockStore) GetUserForUpdateNoKey(arg0 context.Context, arg1 int64) (db.User, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetUserForUpdateNoKey", arg0, arg1)
        ret0, _ := ret[0].(db.User)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetUserForUpdateNoKey indicates an expected call of GetUserForUpdateNoKey.
func (mr *MockStoreMockRecorder) GetUserForUpdateNoKey(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserForUpdateNoKey", reflect.TypeOf((*MockStore)(nil).GetUserForUpdateNoKey), arg0, 
arg1)
}

// GetUserStock mocks base method.
func (m *MockStore) GetUserStock(arg0 context.Context, arg1 int64) (db.UserStock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetUserStock", arg0, arg1)
        ret0, _ := ret[0].(db.UserStock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetUserStock indicates an expected call of GetUserStock.
func (mr *MockStoreMockRecorder) GetUserStock(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserStock", reflect.TypeOf((*MockStore)(nil).GetUserStock), arg0, arg1)
}

// GetUserStocks mocks base method.
func (m *MockStore) GetUserStocks(arg0 context.Context, arg1 db.GetUserStocksParams) ([]db.UserStock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetUserStocks", arg0, arg1)
        ret0, _ := ret[0].([]db.UserStock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetUserStocks indicates an expected call of GetUserStocks.
func (mr *MockStoreMockRecorder) GetUserStocks(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserStocks", reflect.TypeOf((*MockStore)(nil).GetUserStocks), arg0, arg1)
}

// GetUserStocksByPDate mocks base method.
func (m *MockStore) GetUserStocksByPDate(arg0 context.Context, arg1 db.GetUserStocksByPDateParams) ([]db.UserStock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetUserStocksByPDate", arg0, arg1)
        ret0, _ := ret[0].([]db.UserStock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetUserStocksByPDate indicates an expected call of GetUserStocksByPDate.
func (mr *MockStoreMockRecorder) GetUserStocksByPDate(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserStocksByPDate", reflect.TypeOf((*MockStore)(nil).GetUserStocksByPDate), arg0, arg1)
}

// GetUserStocksByStockId mocks base method.
func (m *MockStore) GetUserStocksByStockId(arg0 context.Context, arg1 db.GetUserStocksByStockIdParams) ([]db.UserStock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetUserStocksByStockId", arg0, arg1)
        ret0, _ := ret[0].([]db.UserStock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetUserStocksByStockId indicates an expected call of GetUserStocksByStockId.
func (mr *MockStoreMockRecorder) GetUserStocksByStockId(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserStocksByStockId", reflect.TypeOf((*MockStore)(nil).GetUserStocksByStockId), arg0, arg1)
}

// GetUserStocksByUserAStock mocks base method.
func (m *MockStore) GetUserStocksByUserAStock(arg0 context.Context, arg1 db.GetUserStocksByUserAStockParams) ([]db.UserStock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetUserStocksByUserAStock", arg0, arg1)
        ret0, _ := ret[0].([]db.UserStock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetUserStocksByUserAStock indicates an expected call of GetUserStocksByUserAStock.
func (mr *MockStoreMockRecorder) GetUserStocksByUserAStock(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserStocksByUserAStock", reflect.TypeOf((*MockStore)(nil).GetUserStocksByUserAStock), arg0, arg1)
}

// GetUserStocksByUserId mocks base method.
func (m *MockStore) GetUserStocksByUserId(arg0 context.Context, arg1 db.GetUserStocksByUserIdParams) ([]db.UserStock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetUserStocksByUserId", arg0, arg1)
        ret0, _ := ret[0].([]db.UserStock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetUserStocksByUserId indicates an expected call of GetUserStocksByUserId.
func (mr *MockStoreMockRecorder) GetUserStocksByUserId(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserStocksByUserId", reflect.TypeOf((*MockStore)(nil).GetUserStocksByUserId), arg0, 
arg1)
}

// GetfundByUidandFid mocks base method.
func (m *MockStore) GetfundByUidandFid(arg0 context.Context, arg1 db.GetfundByUidandFidParams) (db.Fund, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetfundByUidandFid", arg0, arg1)
        ret0, _ := ret[0].(db.Fund)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetfundByUidandFid indicates an expected call of GetfundByUidandFid.
func (mr *MockStoreMockRecorder) GetfundByUidandFid(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetfundByUidandFid", reflect.TypeOf((*MockStore)(nil).GetfundByUidandFid), arg0, arg1) 
}

// GetfundByUidandFidForUpdateNoK mocks base method.
func (m *MockStore) GetfundByUidandFidForUpdateNoK(arg0 context.Context, arg1 db.GetfundByUidandFidForUpdateNoKParams) (db.Fund, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetfundByUidandFidForUpdateNoK", arg0, arg1)
        ret0, _ := ret[0].(db.Fund)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetfundByUidandFidForUpdateNoK indicates an expected call of GetfundByUidandFidForUpdateNoK.
func (mr *MockStoreMockRecorder) GetfundByUidandFidForUpdateNoK(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetfundByUidandFidForUpdateNoK", reflect.TypeOf((*MockStore)(nil).GetfundByUidandFidForUpdateNoK), arg0, arg1)
}

// GetfundByUserId mocks base method.
func (m *MockStore) GetfundByUserId(arg0 context.Context, arg1 db.GetfundByUserIdParams) ([]db.Fund, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetfundByUserId", arg0, arg1)
        ret0, _ := ret[0].([]db.Fund)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetfundByUserId indicates an expected call of GetfundByUserId.
func (mr *MockStoreMockRecorder) GetfundByUserId(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetfundByUserId", reflect.TypeOf((*MockStore)(nil).GetfundByUserId), arg0, arg1)       
}

// Getfunds mocks base method.
func (m *MockStore) Getfunds(arg0 context.Context, arg1 db.GetfundsParams) ([]db.Fund, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Getfunds", arg0, arg1)
        ret0, _ := ret[0].([]db.Fund)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Getfunds indicates an expected call of Getfunds.
func (mr *MockStoreMockRecorder) Getfunds(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getfunds", reflect.TypeOf((*MockStore)(nil).Getfunds), arg0, arg1)
}

// GetserStockByUidandSid mocks base method.
func (m *MockStore) GetserStockByUidandSid(arg0 context.Context, arg1 db.GetserStockByUidandSidParams) (db.UserStock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetserStockByUidandSid", arg0, arg1)
        ret0, _ := ret[0].(db.UserStock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetserStockByUidandSid indicates an expected call of GetserStockByUidandSid.
func (mr *MockStoreMockRecorder) GetserStockByUidandSid(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetserStockByUidandSid", reflect.TypeOf((*MockStore)(nil).GetserStockByUidandSid), arg0, arg1)
}

// GetserStockByUidandSidForUpdateNoK mocks base method.
func (m *MockStore) GetserStockByUidandSidForUpdateNoK(arg0 context.Context, arg1 db.GetserStockByUidandSidForUpdateNoKParams) (db.UserStock, error) {        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetserStockByUidandSidForUpdateNoK", arg0, arg1)
        ret0, _ := ret[0].(db.UserStock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetserStockByUidandSidForUpdateNoK indicates an expected call of GetserStockByUidandSidForUpdateNoK.
func (mr *MockStoreMockRecorder) GetserStockByUidandSidForUpdateNoK(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetserStockByUidandSidForUpdateNoK", reflect.TypeOf((*MockStore)(nil).GetserStockByUidandSidForUpdateNoK), arg0, arg1)
}

// GetstockByCN mocks base method.
func (m *MockStore) GetstockByCN(arg0 context.Context, arg1 db.GetstockByCNParams) ([]db.Stock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetstockByCN", arg0, arg1)
        ret0, _ := ret[0].([]db.Stock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetstockByCN indicates an expected call of GetstockByCN.
func (mr *MockStoreMockRecorder) GetstockByCN(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetstockByCN", reflect.TypeOf((*MockStore)(nil).GetstockByCN), arg0, arg1)
}

// GetstockByTS mocks base method.
func (m *MockStore) GetstockByTS(arg0 context.Context, arg1 db.GetstockByTSParams) ([]db.Stock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetstockByTS", arg0, arg1)
        ret0, _ := ret[0].([]db.Stock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetstockByTS indicates an expected call of GetstockByTS.
func (mr *MockStoreMockRecorder) GetstockByTS(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetstockByTS", reflect.TypeOf((*MockStore)(nil).GetstockByTS), arg0, arg1)
}

// Getstocks mocks base method.
func (m *MockStore) Getstocks(arg0 context.Context, arg1 db.GetstocksParams) ([]db.Stock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Getstocks", arg0, arg1)
        ret0, _ := ret[0].([]db.Stock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Getstocks indicates an expected call of Getstocks.
func (mr *MockStoreMockRecorder) Getstocks(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getstocks", reflect.TypeOf((*MockStore)(nil).Getstocks), arg0, arg1)
}

// Getusers mocks base method.
func (m *MockStore) Getusers(arg0 context.Context, arg1 db.GetusersParams) ([]db.User, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Getusers", arg0, arg1)
        ret0, _ := ret[0].([]db.User)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Getusers indicates an expected call of Getusers.
func (mr *MockStoreMockRecorder) Getusers(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getusers", reflect.TypeOf((*MockStore)(nil).Getusers), arg0, arg1)
}

// TransferStockTx mocks base method.
func (m *MockStore) TransferStockTx(arg0 context.Context, arg1 db.TransferStockTxParams) (db.TransferStockTxResults, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "TransferStockTx", arg0, arg1)
        ret0, _ := ret[0].(db.TransferStockTxResults)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// TransferStockTx indicates an expected call of TransferStockTx.
func (mr *MockStoreMockRecorder) TransferStockTx(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferStockTx", reflect.TypeOf((*MockStore)(nil).TransferStockTx), arg0, arg1)       
}

// UpdateFund mocks base method.
func (m *MockStore) UpdateFund(arg0 context.Context, arg1 db.UpdateFundParams) (db.Fund, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "UpdateFund", arg0, arg1)
        ret0, _ := ret[0].(db.Fund)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// UpdateFund indicates an expected call of UpdateFund.
func (mr *MockStoreMockRecorder) UpdateFund(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFund", reflect.TypeOf((*MockStore)(nil).UpdateFund), arg0, arg1)
}

// UpdateStock mocks base method.
func (m *MockStore) UpdateStock(arg0 context.Context, arg1 db.UpdateStockParams) (db.Stock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "UpdateStock", arg0, arg1)
        ret0, _ := ret[0].(db.Stock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// UpdateStock indicates an expected call of UpdateStock.
func (mr *MockStoreMockRecorder) UpdateStock(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStock", reflect.TypeOf((*MockStore)(nil).UpdateStock), arg0, arg1)
}

// UpdateStockCPByCode mocks base method.
func (m *MockStore) UpdateStockCPByCode(arg0 context.Context, arg1 db.UpdateStockCPByCodeParams) (db.Stock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "UpdateStockCPByCode", arg0, arg1)
        ret0, _ := ret[0].(db.Stock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// UpdateStockCPByCode indicates an expected call of UpdateStockCPByCode.
func (mr *MockStoreMockRecorder) UpdateStockCPByCode(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStockCPByCode", reflect.TypeOf((*MockStore)(nil).UpdateStockCPByCode), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
        ret0, _ := ret[0].(db.User)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

// UpdateUserStock mocks base method.
func (m *MockStore) UpdateUserStock(arg0 context.Context, arg1 db.UpdateUserStockParams) (db.UserStock, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "UpdateUserStock", arg0, arg1)
        ret0, _ := ret[0].(db.UserStock)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// UpdateUserStock indicates an expected call of UpdateUserStock.
func (mr *MockStoreMockRecorder) UpdateUserStock(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserStock", reflect.TypeOf((*MockStore)(nil).UpdateUserStock), arg0, arg1)       
}

// UpdateVerifyEmail mocks base method.
func (m *MockStore) UpdateVerifyEmail(arg0 context.Context, arg1 db.UpdateVerifyEmailParams) (db.VerifyEmail, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "UpdateVerifyEmail", arg0, arg1)
        ret0, _ := ret[0].(db.VerifyEmail)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// UpdateVerifyEmail indicates an expected call of UpdateVerifyEmail.
func (mr *MockStoreMockRecorder) UpdateVerifyEmail(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVerifyEmail", reflect.TypeOf((*MockStore)(nil).UpdateVerifyEmail), arg0, arg1)   
}

// VerifyEmailTx mocks base method.
func (m *MockStore) VerifyEmailTx(arg0 context.Context, arg1 db.VerifyEmailTxParams) (db.VerifyEmailTxResults, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "VerifyEmailTx", arg0, arg1)
        ret0, _ := ret[0].(db.VerifyEmailTxResults)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// VerifyEmailTx indicates an expected call of VerifyEmailTx.
func (mr *MockStoreMockRecorder) VerifyEmailTx(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyEmailTx", reflect.TypeOf((*MockStore)(nil).VerifyEmailTx), arg0, arg1)
}