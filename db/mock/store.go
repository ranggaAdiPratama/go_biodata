// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ranggaAdiPratama/go_biodata/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/ranggaAdiPratama/go_biodata/db/sqlc"
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

// CheckHobbyWithPage mocks base method.
func (m *MockStore) CheckHobbyWithPage(arg0 context.Context, arg1 db.CheckHobbyWithPageParams) ([]db.Hobby, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHobbyWithPage", arg0, arg1)
	ret0, _ := ret[0].([]db.Hobby)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckHobbyWithPage indicates an expected call of CheckHobbyWithPage.
func (mr *MockStoreMockRecorder) CheckHobbyWithPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHobbyWithPage", reflect.TypeOf((*MockStore)(nil).CheckHobbyWithPage), arg0, arg1)
}

// CreateHobby mocks base method.
func (m *MockStore) CreateHobby(arg0 context.Context, arg1 db.CreateHobbyParams) (db.Hobby, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHobby", arg0, arg1)
	ret0, _ := ret[0].(db.Hobby)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHobby indicates an expected call of CreateHobby.
func (mr *MockStoreMockRecorder) CreateHobby(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHobby", reflect.TypeOf((*MockStore)(nil).CreateHobby), arg0, arg1)
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

// DeleteHobby mocks base method.
func (m *MockStore) DeleteHobby(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHobby", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHobby indicates an expected call of DeleteHobby.
func (mr *MockStoreMockRecorder) DeleteHobby(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHobby", reflect.TypeOf((*MockStore)(nil).DeleteHobby), arg0, arg1)
}

// GetAllUser mocks base method.
func (m *MockStore) GetAllUser(arg0 context.Context) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUser", arg0)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUser indicates an expected call of GetAllUser.
func (mr *MockStoreMockRecorder) GetAllUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUser", reflect.TypeOf((*MockStore)(nil).GetAllUser), arg0)
}

// GetHobby mocks base method.
func (m *MockStore) GetHobby(arg0 context.Context) ([]db.Hobby, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHobby", arg0)
	ret0, _ := ret[0].([]db.Hobby)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHobby indicates an expected call of GetHobby.
func (mr *MockStoreMockRecorder) GetHobby(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHobby", reflect.TypeOf((*MockStore)(nil).GetHobby), arg0)
}

// GetHobbyByUserId mocks base method.
func (m *MockStore) GetHobbyByUserId(arg0 context.Context, arg1 int64) ([]db.Hobby, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHobbyByUserId", arg0, arg1)
	ret0, _ := ret[0].([]db.Hobby)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHobbyByUserId indicates an expected call of GetHobbyByUserId.
func (mr *MockStoreMockRecorder) GetHobbyByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHobbyByUserId", reflect.TypeOf((*MockStore)(nil).GetHobbyByUserId), arg0, arg1)
}

// GetHobbyForUpdate mocks base method.
func (m *MockStore) GetHobbyForUpdate(arg0 context.Context, arg1 int64) (db.Hobby, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHobbyForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Hobby)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHobbyForUpdate indicates an expected call of GetHobbyForUpdate.
func (mr *MockStoreMockRecorder) GetHobbyForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHobbyForUpdate", reflect.TypeOf((*MockStore)(nil).GetHobbyForUpdate), arg0, arg1)
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

// GetUserByUsername mocks base method.
func (m *MockStore) GetUserByUsername(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStoreMockRecorder) GetUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), arg0, arg1)
}

// GetUserForUpdate mocks base method.
func (m *MockStore) GetUserForUpdate(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserForUpdate indicates an expected call of GetUserForUpdate.
func (mr *MockStoreMockRecorder) GetUserForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserForUpdate", reflect.TypeOf((*MockStore)(nil).GetUserForUpdate), arg0, arg1)
}

// UpdateHobby mocks base method.
func (m *MockStore) UpdateHobby(arg0 context.Context, arg1 db.UpdateHobbyParams) (db.Hobby, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHobby", arg0, arg1)
	ret0, _ := ret[0].(db.Hobby)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateHobby indicates an expected call of UpdateHobby.
func (mr *MockStoreMockRecorder) UpdateHobby(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHobby", reflect.TypeOf((*MockStore)(nil).UpdateHobby), arg0, arg1)
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

// UpdateUserPassword mocks base method.
func (m *MockStore) UpdateUserPassword(arg0 context.Context, arg1 db.UpdateUserPasswordParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockStoreMockRecorder) UpdateUserPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockStore)(nil).UpdateUserPassword), arg0, arg1)
}
