// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

// Code generated by MockGen. DO NOT EDIT.
// Source: interface_ippools.go

// Package mockquakeclient is a generated GoMock package.
package mockquakeclient

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/hewlettpackard/hpegl-metal-client/v1/pkg/client"
)

// MockIPPoolsAPI is a mock of IPPoolsAPI interface.
type MockIPPoolsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIPPoolsAPIMockRecorder
}

// MockIPPoolsAPIMockRecorder is the mock recorder for MockIPPoolsAPI.
type MockIPPoolsAPIMockRecorder struct {
	mock *MockIPPoolsAPI
}

// NewMockIPPoolsAPI creates a new mock instance.
func NewMockIPPoolsAPI(ctrl *gomock.Controller) *MockIPPoolsAPI {
	mock := &MockIPPoolsAPI{ctrl: ctrl}
	mock.recorder = &MockIPPoolsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPPoolsAPI) EXPECT() *MockIPPoolsAPIMockRecorder {
	return m.recorder
}

// AllocateIPs mocks base method.
func (m *MockIPPoolsAPI) AllocateIPs(ctx context.Context, ippoolId string, iPAllocation []client.IpAllocation) (client.IpPool, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllocateIPs", ctx, ippoolId, iPAllocation)
	ret0, _ := ret[0].(client.IpPool)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AllocateIPs indicates an expected call of AllocateIPs.
func (mr *MockIPPoolsAPIMockRecorder) AllocateIPs(ctx, ippoolId, iPAllocation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocateIPs", reflect.TypeOf((*MockIPPoolsAPI)(nil).AllocateIPs), ctx, ippoolId, iPAllocation)
}

// GetByID mocks base method.
func (m *MockIPPoolsAPI) GetByID(ctx context.Context, ippoolId string) (client.IpPool, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, ippoolId)
	ret0, _ := ret[0].(client.IpPool)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIPPoolsAPIMockRecorder) GetByID(ctx, ippoolId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIPPoolsAPI)(nil).GetByID), ctx, ippoolId)
}

// List mocks base method.
func (m *MockIPPoolsAPI) List(ctx context.Context) ([]client.IpPool, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]client.IpPool)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockIPPoolsAPIMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIPPoolsAPI)(nil).List), ctx)
}

// ReturnIPs mocks base method.
func (m *MockIPPoolsAPI) ReturnIPs(ctx context.Context, ippoolId string, requestBody []string) (client.IpPool, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturnIPs", ctx, ippoolId, requestBody)
	ret0, _ := ret[0].(client.IpPool)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReturnIPs indicates an expected call of ReturnIPs.
func (mr *MockIPPoolsAPIMockRecorder) ReturnIPs(ctx, ippoolId, requestBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturnIPs", reflect.TypeOf((*MockIPPoolsAPI)(nil).ReturnIPs), ctx, ippoolId, requestBody)
}

// Update mocks base method.
func (m *MockIPPoolsAPI) Update(ctx context.Context, ippoolId string, ipPool client.IpPool) (client.IpPool, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, ippoolId, ipPool)
	ret0, _ := ret[0].(client.IpPool)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockIPPoolsAPIMockRecorder) Update(ctx, ippoolId, ipPool interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIPPoolsAPI)(nil).Update), ctx, ippoolId, ipPool)
}
