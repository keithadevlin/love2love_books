// Code generated by MockGen. DO NOT EDIT.
// Source: persistfile.go

// Package persistfile is a generated GoMock package.
package persistfile

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFilePersistor is a mock of FilePersistor interface
type MockFilePersistor struct {
	ctrl     *gomock.Controller
	recorder *MockFilePersistorMockRecorder
}

// MockFilePersistorMockRecorder is the mock recorder for MockFilePersistor
type MockFilePersistorMockRecorder struct {
	mock *MockFilePersistor
}

// NewMockFilePersistor creates a new mock instance
func NewMockFilePersistor(ctrl *gomock.Controller) *MockFilePersistor {
	mock := &MockFilePersistor{ctrl: ctrl}
	mock.recorder = &MockFilePersistorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFilePersistor) EXPECT() *MockFilePersistorMockRecorder {
	return m.recorder
}

// PersistFilePayload mocks base method
func (m *MockFilePersistor) PersistFilePayload(ctx context.Context, fileName, payload string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersistFilePayload", ctx, fileName, payload)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PersistFilePayload indicates an expected call of PersistFilePayload
func (mr *MockFilePersistorMockRecorder) PersistFilePayload(ctx, fileName, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersistFilePayload", reflect.TypeOf((*MockFilePersistor)(nil).PersistFilePayload), ctx, fileName, payload)
}
