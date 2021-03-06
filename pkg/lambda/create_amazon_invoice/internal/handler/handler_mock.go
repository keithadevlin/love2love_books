// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package handler is a generated GoMock package.
package handler

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	mws "github.com/keithadevlin/love2love_books/pkg/shared/mws"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateInvoice mocks base method
func (m *MockService) CreateInvoice(ctx context.Context, orderReportItem mws.OrderReportItem, invoiceNumber int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvoice", ctx, orderReportItem, invoiceNumber)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInvoice indicates an expected call of CreateInvoice
func (mr *MockServiceMockRecorder) CreateInvoice(ctx, orderReportItem, invoiceNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvoice", reflect.TypeOf((*MockService)(nil).CreateInvoice), ctx, orderReportItem, invoiceNumber)
}
