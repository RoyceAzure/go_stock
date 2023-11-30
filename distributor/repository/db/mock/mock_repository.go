package mock_repository

import (
	context "context"
	reflect "reflect"

	repository "github.com/RoyceAzure/go-stockinfo-distributor/repository/db/sqlc"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockDistributorDao is a mock of DistributorDao interface.
type MockDistributorDao struct {
	ctrl     *gomock.Controller
	recorder *MockDistributorDaoMockRecorder
}

// MockDistributorDaoMockRecorder is the mock recorder for MockDistributorDao.
type MockDistributorDaoMockRecorder struct {
	mock *MockDistributorDao
}

// NewMockDistributorDao creates a new mock instance.
func NewMockDistributorDao(ctrl *gomock.Controller) *MockDistributorDao {
	mock := &MockDistributorDao{ctrl: ctrl}
	mock.recorder = &MockDistributorDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDistributorDao) EXPECT() *MockDistributorDaoMockRecorder {
	return m.recorder
}

// CreateClientRegister mocks base method.
func (m *MockDistributorDao) CreateClientRegister(arg0 context.Context, arg1 repository.CreateClientRegisterParams) (repository.ClientRegister, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClientRegister", arg0, arg1)
	ret0, _ := ret[0].(repository.ClientRegister)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClientRegister indicates an expected call of CreateClientRegister.
func (mr *MockDistributorDaoMockRecorder) CreateClientRegister(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClientRegister", reflect.TypeOf((*MockDistributorDao)(nil).CreateClientRegister), arg0, arg1)
}

// CreateFrontendClient mocks base method.
func (m *MockDistributorDao) CreateFrontendClient(arg0 context.Context, arg1 repository.CreateFrontendClientParams) (repository.FrontendClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFrontendClient", arg0, arg1)
	ret0, _ := ret[0].(repository.FrontendClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFrontendClient indicates an expected call of CreateFrontendClient.
func (mr *MockDistributorDaoMockRecorder) CreateFrontendClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFrontendClient", reflect.TypeOf((*MockDistributorDao)(nil).CreateFrontendClient), arg0, arg1)
}

// CreateUserRegister mocks base method.
func (m *MockDistributorDao) CreateUserRegister(arg0 context.Context, arg1 repository.CreateUserRegisterParams) (repository.UserRegister, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserRegister", arg0, arg1)
	ret0, _ := ret[0].(repository.UserRegister)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserRegister indicates an expected call of CreateUserRegister.
func (mr *MockDistributorDaoMockRecorder) CreateUserRegister(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserRegister", reflect.TypeOf((*MockDistributorDao)(nil).CreateUserRegister), arg0, arg1)
}

// DeleteClientRegister mocks base method.
func (m *MockDistributorDao) DeleteClientRegister(arg0 context.Context, arg1 repository.DeleteClientRegisterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClientRegister", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClientRegister indicates an expected call of DeleteClientRegister.
func (mr *MockDistributorDaoMockRecorder) DeleteClientRegister(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClientRegister", reflect.TypeOf((*MockDistributorDao)(nil).DeleteClientRegister), arg0, arg1)
}

// DeleteFrontendClient mocks base method.
func (m *MockDistributorDao) DeleteFrontendClient(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFrontendClient", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFrontendClient indicates an expected call of DeleteFrontendClient.
func (mr *MockDistributorDaoMockRecorder) DeleteFrontendClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFrontendClient", reflect.TypeOf((*MockDistributorDao)(nil).DeleteFrontendClient), arg0, arg1)
}

// DeleteUserRegister mocks base method.
func (m *MockDistributorDao) DeleteUserRegister(arg0 context.Context, arg1 repository.DeleteUserRegisterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserRegister", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserRegister indicates an expected call of DeleteUserRegister.
func (mr *MockDistributorDaoMockRecorder) DeleteUserRegister(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserRegister", reflect.TypeOf((*MockDistributorDao)(nil).DeleteUserRegister), arg0, arg1)
}

// GetClientRegisterByClientUID mocks base method.
func (m *MockDistributorDao) GetClientRegisterByClientUID(arg0 context.Context, arg1 uuid.UUID) ([]repository.ClientRegister, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientRegisterByClientUID", arg0, arg1)
	ret0, _ := ret[0].([]repository.ClientRegister)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientRegisterByClientUID indicates an expected call of GetClientRegisterByClientUID.
func (mr *MockDistributorDaoMockRecorder) GetClientRegisterByClientUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientRegisterByClientUID", reflect.TypeOf((*MockDistributorDao)(nil).GetClientRegisterByClientUID), arg0, arg1)
}

// GetClientRegisters mocks base method.
func (m *MockDistributorDao) GetClientRegisters(arg0 context.Context, arg1 repository.GetClientRegistersParams) ([]repository.ClientRegister, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientRegisters", arg0, arg1)
	ret0, _ := ret[0].([]repository.ClientRegister)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientRegisters indicates an expected call of GetClientRegisters.
func (mr *MockDistributorDaoMockRecorder) GetClientRegisters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientRegisters", reflect.TypeOf((*MockDistributorDao)(nil).GetClientRegisters), arg0, arg1)
}

// GetDistinctStockCode mocks base method.
func (m *MockDistributorDao) GetDistinctStockCode(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDistinctStockCode", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDistinctStockCode indicates an expected call of GetDistinctStockCode.
func (mr *MockDistributorDaoMockRecorder) GetDistinctStockCode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDistinctStockCode", reflect.TypeOf((*MockDistributorDao)(nil).GetDistinctStockCode), arg0)
}

// GetFrontendClientByID mocks base method.
func (m *MockDistributorDao) GetFrontendClientByID(arg0 context.Context, arg1 uuid.UUID) (repository.FrontendClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFrontendClientByID", arg0, arg1)
	ret0, _ := ret[0].(repository.FrontendClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFrontendClientByID indicates an expected call of GetFrontendClientByID.
func (mr *MockDistributorDaoMockRecorder) GetFrontendClientByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFrontendClientByID", reflect.TypeOf((*MockDistributorDao)(nil).GetFrontendClientByID), arg0, arg1)
}

// GetFrontendClientByIP mocks base method.
func (m *MockDistributorDao) GetFrontendClientByIP(arg0 context.Context, arg1 string) (repository.FrontendClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFrontendClientByIP", arg0, arg1)
	ret0, _ := ret[0].(repository.FrontendClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFrontendClientByIP indicates an expected call of GetFrontendClientByIP.
func (mr *MockDistributorDaoMockRecorder) GetFrontendClientByIP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFrontendClientByIP", reflect.TypeOf((*MockDistributorDao)(nil).GetFrontendClientByIP), arg0, arg1)
}

// GetFrontendClients mocks base method.
func (m *MockDistributorDao) GetFrontendClients(arg0 context.Context, arg1 repository.GetFrontendClientsParams) ([]repository.FrontendClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFrontendClients", arg0, arg1)
	ret0, _ := ret[0].([]repository.FrontendClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFrontendClients indicates an expected call of GetFrontendClients.
func (mr *MockDistributorDaoMockRecorder) GetFrontendClients(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFrontendClients", reflect.TypeOf((*MockDistributorDao)(nil).GetFrontendClients), arg0, arg1)
}

// GetUserRegisterByUserID mocks base method.
func (m *MockDistributorDao) GetUserRegisterByUserID(arg0 context.Context, arg1 int64) ([]repository.UserRegister, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRegisterByUserID", arg0, arg1)
	ret0, _ := ret[0].([]repository.UserRegister)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRegisterByUserID indicates an expected call of GetUserRegisterByUserID.
func (mr *MockDistributorDaoMockRecorder) GetUserRegisterByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRegisterByUserID", reflect.TypeOf((*MockDistributorDao)(nil).GetUserRegisterByUserID), arg0, arg1)
}

// GetUserRegisters mocks base method.
func (m *MockDistributorDao) GetUserRegisters(arg0 context.Context, arg1 repository.GetUserRegistersParams) ([]repository.UserRegister, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRegisters", arg0, arg1)
	ret0, _ := ret[0].([]repository.UserRegister)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRegisters indicates an expected call of GetUserRegisters.
func (mr *MockDistributorDaoMockRecorder) GetUserRegisters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRegisters", reflect.TypeOf((*MockDistributorDao)(nil).GetUserRegisters), arg0, arg1)
}
