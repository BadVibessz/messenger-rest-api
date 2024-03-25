// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ew0s/ewos-to-go-hw/chat-server/internal/service/user (interfaces: UserRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/ew0s/ewos-to-go-hw/chat-server/internal/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserRepo) AddUser(arg0 context.Context, arg1 entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserRepoMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserRepo)(nil).AddUser), arg0, arg1)
}

// CheckUniqueConstraints mocks base method.
func (m *MockUserRepo) CheckUniqueConstraints(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUniqueConstraints", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUniqueConstraints indicates an expected call of CheckUniqueConstraints.
func (mr *MockUserRepoMockRecorder) CheckUniqueConstraints(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUniqueConstraints", reflect.TypeOf((*MockUserRepo)(nil).CheckUniqueConstraints), arg0, arg1, arg2)
}

// DeleteUser mocks base method.
func (m *MockUserRepo) DeleteUser(arg0 context.Context, arg1 int) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepoMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepo)(nil).DeleteUser), arg0, arg1)
}

// GetAllUsers mocks base method.
func (m *MockUserRepo) GetAllUsers(arg0 context.Context, arg1, arg2 int) []*entity.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*entity.User)
	return ret0
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserRepoMockRecorder) GetAllUsers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserRepo)(nil).GetAllUsers), arg0, arg1, arg2)
}

// GetUserByEmail mocks base method.
func (m *MockUserRepo) GetUserByEmail(arg0 context.Context, arg1 string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepoMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepo)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserByID mocks base method.
func (m *MockUserRepo) GetUserByID(arg0 context.Context, arg1 int) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepoMockRecorder) GetUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepo)(nil).GetUserByID), arg0, arg1)
}

// GetUserByUsername mocks base method.
func (m *MockUserRepo) GetUserByUsername(arg0 context.Context, arg1 string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockUserRepoMockRecorder) GetUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockUserRepo)(nil).GetUserByUsername), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockUserRepo) UpdateUser(arg0 context.Context, arg1 int, arg2 entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepoMockRecorder) UpdateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepo)(nil).UpdateUser), arg0, arg1, arg2)
}
