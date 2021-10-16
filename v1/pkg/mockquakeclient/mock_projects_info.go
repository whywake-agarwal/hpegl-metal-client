// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

// Code generated by MockGen. DO NOT EDIT.
// Source: interface_projects_info.go

// Package mockquakeclient is a generated GoMock package.
package mockquakeclient

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/hpe-hcss/quake-client/v1/pkg/client"
)

// MockProjectsInfoAPI is a mock of ProjectsInfoAPI interface.
type MockProjectsInfoAPI struct {
	ctrl     *gomock.Controller
	recorder *MockProjectsInfoAPIMockRecorder
}

// MockProjectsInfoAPIMockRecorder is the mock recorder for MockProjectsInfoAPI.
type MockProjectsInfoAPIMockRecorder struct {
	mock *MockProjectsInfoAPI
}

// NewMockProjectsInfoAPI creates a new mock instance.
func NewMockProjectsInfoAPI(ctrl *gomock.Controller) *MockProjectsInfoAPI {
	mock := &MockProjectsInfoAPI{ctrl: ctrl}
	mock.recorder = &MockProjectsInfoAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectsInfoAPI) EXPECT() *MockProjectsInfoAPIMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockProjectsInfoAPI) List(ctx context.Context) (client.ProjectsInfo, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].(client.ProjectsInfo)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockProjectsInfoAPIMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProjectsInfoAPI)(nil).List), ctx)
}
