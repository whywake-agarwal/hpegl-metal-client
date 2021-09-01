// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

// Code generated by MockGen. DO NOT EDIT.
// Source: interface_hosts.go

// Package mockquakeclient is a generated GoMock package.
package mockquakeclient

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	client "github.com/hpe-hcss/quake-client/v1/pkg/client"
	http "net/http"
	reflect "reflect"
)

// MockHostsAPI is a mock of HostsAPI interface
type MockHostsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockHostsAPIMockRecorder
}

// MockHostsAPIMockRecorder is the mock recorder for MockHostsAPI
type MockHostsAPIMockRecorder struct {
	mock *MockHostsAPI
}

// NewMockHostsAPI creates a new mock instance
func NewMockHostsAPI(ctrl *gomock.Controller) *MockHostsAPI {
	mock := &MockHostsAPI{ctrl: ctrl}
	mock.recorder = &MockHostsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHostsAPI) EXPECT() *MockHostsAPIMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockHostsAPI) Add(ctx context.Context, newHost client.NewHost) (client.Host, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, newHost)
	ret0, _ := ret[0].(client.Host)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Add indicates an expected call of Add
func (mr *MockHostsAPIMockRecorder) Add(ctx, newHost interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockHostsAPI)(nil).Add), ctx, newHost)
}

// Delete mocks base method
func (m *MockHostsAPI) Delete(ctx context.Context, hostId string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, hostId)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockHostsAPIMockRecorder) Delete(ctx, hostId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockHostsAPI)(nil).Delete), ctx, hostId)
}

// GetByID mocks base method
func (m *MockHostsAPI) GetByID(ctx context.Context, hostId string) (client.Host, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, hostId)
	ret0, _ := ret[0].(client.Host)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID
func (mr *MockHostsAPIMockRecorder) GetByID(ctx, hostId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockHostsAPI)(nil).GetByID), ctx, hostId)
}

// List mocks base method
func (m *MockHostsAPI) List(ctx context.Context) ([]client.Host, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]client.Host)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockHostsAPIMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockHostsAPI)(nil).List), ctx)
}

// PowerOff mocks base method
func (m *MockHostsAPI) PowerOff(ctx context.Context, hostId string) (client.Host, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOff", ctx, hostId)
	ret0, _ := ret[0].(client.Host)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PowerOff indicates an expected call of PowerOff
func (mr *MockHostsAPIMockRecorder) PowerOff(ctx, hostId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOff", reflect.TypeOf((*MockHostsAPI)(nil).PowerOff), ctx, hostId)
}

// PowerOn mocks base method
func (m *MockHostsAPI) PowerOn(ctx context.Context, hostId string) (client.Host, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOn", ctx, hostId)
	ret0, _ := ret[0].(client.Host)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PowerOn indicates an expected call of PowerOn
func (mr *MockHostsAPIMockRecorder) PowerOn(ctx, hostId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOn", reflect.TypeOf((*MockHostsAPI)(nil).PowerOn), ctx, hostId)
}

// Update mocks base method
func (m *MockHostsAPI) Update(ctx context.Context, hostId string, host client.Host) (client.Host, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, hostId, host)
	ret0, _ := ret[0].(client.Host)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update
func (mr *MockHostsAPIMockRecorder) Update(ctx, hostId, host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockHostsAPI)(nil).Update), ctx, hostId, host)
}
