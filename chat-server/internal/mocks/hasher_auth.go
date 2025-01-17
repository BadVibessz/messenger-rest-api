// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ew0s/ewos-to-go-hw/http5/homework/chat-server/internal/service/auth (interfaces: Hasher)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthHasher is a mock of Hasher interface.
type MockAuthHasher struct {
	ctrl     *gomock.Controller
	recorder *MockAuthHasherMockRecorder
}

// MockAuthHasherMockRecorder is the mock recorder for MockAuthHasher.
type MockAuthHasherMockRecorder struct {
	mock *MockAuthHasher
}

// NewMockAuthHasher creates a new mock instance.
func NewMockAuthHasher(ctrl *gomock.Controller) *MockAuthHasher {
	mock := &MockAuthHasher{ctrl: ctrl}
	mock.recorder = &MockAuthHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthHasher) EXPECT() *MockAuthHasherMockRecorder {
	return m.recorder
}

// CompareHashAndPassword mocks base method.
func (m *MockAuthHasher) CompareHashAndPassword(arg0, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareHashAndPassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompareHashAndPassword indicates an expected call of CompareHashAndPassword.
func (mr *MockAuthHasherMockRecorder) CompareHashAndPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareHashAndPassword", reflect.TypeOf((*MockAuthHasher)(nil).CompareHashAndPassword), arg0, arg1)
}
