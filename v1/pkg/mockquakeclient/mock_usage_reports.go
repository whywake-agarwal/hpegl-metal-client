// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

// Code generated by MockGen. DO NOT EDIT.
// Source: interface_usage_reports.go

// Package mockquakeclient is a generated GoMock package.
package mockquakeclient

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/hewlettpackard/hpegl-metal-client/v1/pkg/client"
)

// MockUsageReportsAPI is a mock of UsageReportsAPI interface.
type MockUsageReportsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockUsageReportsAPIMockRecorder
}

// MockUsageReportsAPIMockRecorder is the mock recorder for MockUsageReportsAPI.
type MockUsageReportsAPIMockRecorder struct {
	mock *MockUsageReportsAPI
}

// NewMockUsageReportsAPI creates a new mock instance.
func NewMockUsageReportsAPI(ctrl *gomock.Controller) *MockUsageReportsAPI {
	mock := &MockUsageReportsAPI{ctrl: ctrl}
	mock.recorder = &MockUsageReportsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsageReportsAPI) EXPECT() *MockUsageReportsAPIMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockUsageReportsAPI) Get(ctx context.Context, start string, localVarOptionals *client.UsageReportsApiGetOpts) (client.UsageReport, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, start, localVarOptionals)
	ret0, _ := ret[0].(client.UsageReport)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockUsageReportsAPIMockRecorder) Get(ctx, start, localVarOptionals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsageReportsAPI)(nil).Get), ctx, start, localVarOptionals)
}
