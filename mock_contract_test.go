// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package logger_lib is a generated GoMock package.
package logger_lib

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLoggerInterface is a mock of LoggerInterface interface.
type MockLoggerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerInterfaceMockRecorder
}

// MockLoggerInterfaceMockRecorder is the mock recorder for MockLoggerInterface.
type MockLoggerInterfaceMockRecorder struct {
	mock *MockLoggerInterface
}

// NewMockLoggerInterface creates a new mock instance.
func NewMockLoggerInterface(ctrl *gomock.Controller) *MockLoggerInterface {
	mock := &MockLoggerInterface{ctrl: ctrl}
	mock.recorder = &MockLoggerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoggerInterface) EXPECT() *MockLoggerInterfaceMockRecorder {
	return m.recorder
}

// AddFuncName mocks base method.
func (m *MockLoggerInterface) AddFuncName(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddFuncName", name)
}

// AddFuncName indicates an expected call of AddFuncName.
func (mr *MockLoggerInterfaceMockRecorder) AddFuncName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFuncName", reflect.TypeOf((*MockLoggerInterface)(nil).AddFuncName), name)
}

// Error mocks base method.
func (m *MockLoggerInterface) Error(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", msg)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerInterfaceMockRecorder) Error(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLoggerInterface)(nil).Error), msg)
}

// Info mocks base method.
func (m *MockLoggerInterface) Info(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Info", msg)
}

// Info indicates an expected call of Info.
func (mr *MockLoggerInterfaceMockRecorder) Info(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLoggerInterface)(nil).Info), msg)
}

// Warn mocks base method.
func (m *MockLoggerInterface) Warn(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Warn", msg)
}

// Warn indicates an expected call of Warn.
func (mr *MockLoggerInterfaceMockRecorder) Warn(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLoggerInterface)(nil).Warn), msg)
}
