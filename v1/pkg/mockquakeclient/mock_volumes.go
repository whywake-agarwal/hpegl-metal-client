// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

// Code generated by MockGen. DO NOT EDIT.
// Source: interface_volumes.go

// Package mockquakeclient is a generated GoMock package.
package mockquakeclient

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/hewlettpackard/hpegl-metal-client/v1/pkg/client"
)

// MockVolumesAPI is a mock of VolumesAPI interface.
type MockVolumesAPI struct {
	ctrl     *gomock.Controller
	recorder *MockVolumesAPIMockRecorder
}

// MockVolumesAPIMockRecorder is the mock recorder for MockVolumesAPI.
type MockVolumesAPIMockRecorder struct {
	mock *MockVolumesAPI
}

// NewMockVolumesAPI creates a new mock instance.
func NewMockVolumesAPI(ctrl *gomock.Controller) *MockVolumesAPI {
	mock := &MockVolumesAPI{ctrl: ctrl}
	mock.recorder = &MockVolumesAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVolumesAPI) EXPECT() *MockVolumesAPIMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockVolumesAPI) Add(ctx context.Context, newVolume client.NewVolume) (client.Volume, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, newVolume)
	ret0, _ := ret[0].(client.Volume)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Add indicates an expected call of Add.
func (mr *MockVolumesAPIMockRecorder) Add(ctx, newVolume interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockVolumesAPI)(nil).Add), ctx, newVolume)
}

// Attach mocks base method.
func (m *MockVolumesAPI) Attach(ctx context.Context, volumeId string, volumeAttachHostUuid client.VolumeAttachHostUuid) (client.VolumeAttachment, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attach", ctx, volumeId, volumeAttachHostUuid)
	ret0, _ := ret[0].(client.VolumeAttachment)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Attach indicates an expected call of Attach.
func (mr *MockVolumesAPIMockRecorder) Attach(ctx, volumeId, volumeAttachHostUuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attach", reflect.TypeOf((*MockVolumesAPI)(nil).Attach), ctx, volumeId, volumeAttachHostUuid)
}

// Delete mocks base method.
func (m *MockVolumesAPI) Delete(ctx context.Context, volumeId string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, volumeId)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockVolumesAPIMockRecorder) Delete(ctx, volumeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVolumesAPI)(nil).Delete), ctx, volumeId)
}

// Detach mocks base method.
func (m *MockVolumesAPI) Detach(ctx context.Context, volumeId string, volumeAttachHostUuid client.VolumeAttachHostUuid) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Detach", ctx, volumeId, volumeAttachHostUuid)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Detach indicates an expected call of Detach.
func (mr *MockVolumesAPIMockRecorder) Detach(ctx, volumeId, volumeAttachHostUuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detach", reflect.TypeOf((*MockVolumesAPI)(nil).Detach), ctx, volumeId, volumeAttachHostUuid)
}

// GetByID mocks base method.
func (m *MockVolumesAPI) GetByID(ctx context.Context, volumeId string) (client.Volume, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, volumeId)
	ret0, _ := ret[0].(client.Volume)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByID indicates an expected call of GetByID.
func (mr *MockVolumesAPIMockRecorder) GetByID(ctx, volumeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockVolumesAPI)(nil).GetByID), ctx, volumeId)
}

// List mocks base method.
func (m *MockVolumesAPI) List(ctx context.Context) ([]client.Volume, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]client.Volume)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockVolumesAPIMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVolumesAPI)(nil).List), ctx)
}

// Update mocks base method.
func (m *MockVolumesAPI) Update(ctx context.Context, volumeId string, volume client.Volume) (client.Volume, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, volumeId, volume)
	ret0, _ := ret[0].(client.Volume)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockVolumesAPIMockRecorder) Update(ctx, volumeId, volume interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVolumesAPI)(nil).Update), ctx, volumeId, volume)
}
