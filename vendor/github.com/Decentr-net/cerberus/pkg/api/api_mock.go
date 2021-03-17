// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package api is a generated GoMock package.
package api

import (
	context "context"
	schema "github.com/Decentr-net/cerberus/pkg/schema"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCerberus is a mock of Cerberus interface
type MockCerberus struct {
	ctrl     *gomock.Controller
	recorder *MockCerberusMockRecorder
}

// MockCerberusMockRecorder is the mock recorder for MockCerberus
type MockCerberusMockRecorder struct {
	mock *MockCerberus
}

// NewMockCerberus creates a new mock instance
func NewMockCerberus(ctrl *gomock.Controller) *MockCerberus {
	mock := &MockCerberus{ctrl: ctrl}
	mock.recorder = &MockCerberusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCerberus) EXPECT() *MockCerberusMockRecorder {
	return m.recorder
}

// SavePDV mocks base method
func (m *MockCerberus) SavePDV(ctx context.Context, p schema.PDV) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePDV", ctx, p)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SavePDV indicates an expected call of SavePDV
func (mr *MockCerberusMockRecorder) SavePDV(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePDV", reflect.TypeOf((*MockCerberus)(nil).SavePDV), ctx, p)
}

// ListPDV mocks base method
func (m *MockCerberus) ListPDV(ctx context.Context, owner string, from uint64, limit uint16) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPDV", ctx, owner, from, limit)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPDV indicates an expected call of ListPDV
func (mr *MockCerberusMockRecorder) ListPDV(ctx, owner, from, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPDV", reflect.TypeOf((*MockCerberus)(nil).ListPDV), ctx, owner, from, limit)
}

// ReceivePDV mocks base method
func (m *MockCerberus) ReceivePDV(ctx context.Context, owner string, id uint64) (schema.PDV, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceivePDV", ctx, owner, id)
	ret0, _ := ret[0].(schema.PDV)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceivePDV indicates an expected call of ReceivePDV
func (mr *MockCerberusMockRecorder) ReceivePDV(ctx, owner, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivePDV", reflect.TypeOf((*MockCerberus)(nil).ReceivePDV), ctx, owner, id)
}

// GetPDVMeta mocks base method
func (m *MockCerberus) GetPDVMeta(ctx context.Context, owner string, id uint64) (PDVMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPDVMeta", ctx, owner, id)
	ret0, _ := ret[0].(PDVMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPDVMeta indicates an expected call of GetPDVMeta
func (mr *MockCerberusMockRecorder) GetPDVMeta(ctx, owner, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPDVMeta", reflect.TypeOf((*MockCerberus)(nil).GetPDVMeta), ctx, owner, id)
}
