// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

// Code generated by MockGen. DO NOT EDIT.
// Source: interface_services.go

// Package mockquakeclient is a generated GoMock package.
package mockquakeclient

import (
	context "context"
	http "net/http"
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/hewlettpackard/hpegl-metal-client/v1/pkg/client"
)

// MockServicesAPI is a mock of ServicesAPI interface.
type MockServicesAPI struct {
	ctrl     *gomock.Controller
	recorder *MockServicesAPIMockRecorder
}

// MockServicesAPIMockRecorder is the mock recorder for MockServicesAPI.
type MockServicesAPIMockRecorder struct {
	mock *MockServicesAPI
}

// NewMockServicesAPI creates a new mock instance.
func NewMockServicesAPI(ctrl *gomock.Controller) *MockServicesAPI {
	mock := &MockServicesAPI{ctrl: ctrl}
	mock.recorder = &MockServicesAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicesAPI) EXPECT() *MockServicesAPIMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockServicesAPI) Add(ctx context.Context, fileName *os.File, localVarOptionals *client.ServicesApiAddOpts) (client.OsServiceImage, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, fileName, localVarOptionals)
	ret0, _ := ret[0].(client.OsServiceImage)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Add indicates an expected call of Add.
func (mr *MockServicesAPIMockRecorder) Add(ctx, fileName, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockServicesAPI)(nil).Add), ctx, fileName, localVarOptionals)
}

// Delete mocks base method.
func (m *MockServicesAPI) Delete(ctx context.Context, serviceId string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, serviceId)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServicesAPIMockRecorder) Delete(ctx, serviceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServicesAPI)(nil).Delete), ctx, serviceId)
}

// GetByID mocks base method.
func (m *MockServicesAPI) GetByID(ctx context.Context, serviceId string) (client.OsServiceImage, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, serviceId)
	ret0, _ := ret[0].(client.OsServiceImage)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockServicesAPIMockRecorder) GetByID(ctx, serviceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockServicesAPI)(nil).GetByID), ctx, serviceId)
}

// List mocks base method.
func (m *MockServicesAPI) List(ctx context.Context, localVarOptionals *client.ServicesApiListOpts) ([]client.OsServiceImage, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, localVarOptionals)
	ret0, _ := ret[0].([]client.OsServiceImage)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockServicesAPIMockRecorder) List(ctx, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServicesAPI)(nil).List), ctx, localVarOptionals)
}

// Update mocks base method.
func (m *MockServicesAPI) Update(ctx context.Context, serviceId string, fileName *os.File) (client.OsServiceImage, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, serviceId, fileName)
	ret0, _ := ret[0].(client.OsServiceImage)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockServicesAPIMockRecorder) Update(ctx, serviceId, fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServicesAPI)(nil).Update), ctx, serviceId, fileName)
}