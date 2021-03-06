// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/internal/controller/controllers (interfaces: CRDEventsHandler)

// Package controllers is a generated GoMock package.
package controllers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	event "sigs.k8s.io/controller-runtime/pkg/event"
)

// MockCRDEventsHandler is a mock of CRDEventsHandler interface
type MockCRDEventsHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCRDEventsHandlerMockRecorder
}

// MockCRDEventsHandlerMockRecorder is the mock recorder for MockCRDEventsHandler
type MockCRDEventsHandlerMockRecorder struct {
	mock *MockCRDEventsHandler
}

// NewMockCRDEventsHandler creates a new mock instance
func NewMockCRDEventsHandler(ctrl *gomock.Controller) *MockCRDEventsHandler {
	mock := &MockCRDEventsHandler{ctrl: ctrl}
	mock.recorder = &MockCRDEventsHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCRDEventsHandler) EXPECT() *MockCRDEventsHandlerMockRecorder {
	return m.recorder
}

// GetAgentUpdates mocks base method
func (m *MockCRDEventsHandler) GetAgentUpdates() chan event.GenericEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentUpdates")
	ret0, _ := ret[0].(chan event.GenericEvent)
	return ret0
}

// GetAgentUpdates indicates an expected call of GetAgentUpdates
func (mr *MockCRDEventsHandlerMockRecorder) GetAgentUpdates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentUpdates", reflect.TypeOf((*MockCRDEventsHandler)(nil).GetAgentUpdates))
}

// GetClusterDeploymentUpdates mocks base method
func (m *MockCRDEventsHandler) GetClusterDeploymentUpdates() chan event.GenericEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterDeploymentUpdates")
	ret0, _ := ret[0].(chan event.GenericEvent)
	return ret0
}

// GetClusterDeploymentUpdates indicates an expected call of GetClusterDeploymentUpdates
func (mr *MockCRDEventsHandlerMockRecorder) GetClusterDeploymentUpdates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterDeploymentUpdates", reflect.TypeOf((*MockCRDEventsHandler)(nil).GetClusterDeploymentUpdates))
}

// GetInstallEnvUpdates mocks base method
func (m *MockCRDEventsHandler) GetInstallEnvUpdates() chan event.GenericEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallEnvUpdates")
	ret0, _ := ret[0].(chan event.GenericEvent)
	return ret0
}

// GetInstallEnvUpdates indicates an expected call of GetInstallEnvUpdates
func (mr *MockCRDEventsHandlerMockRecorder) GetInstallEnvUpdates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallEnvUpdates", reflect.TypeOf((*MockCRDEventsHandler)(nil).GetInstallEnvUpdates))
}

// NotifyAgentUpdates mocks base method
func (m *MockCRDEventsHandler) NotifyAgentUpdates(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyAgentUpdates", arg0, arg1)
}

// NotifyAgentUpdates indicates an expected call of NotifyAgentUpdates
func (mr *MockCRDEventsHandlerMockRecorder) NotifyAgentUpdates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyAgentUpdates", reflect.TypeOf((*MockCRDEventsHandler)(nil).NotifyAgentUpdates), arg0, arg1)
}

// NotifyClusterDeploymentUpdates mocks base method
func (m *MockCRDEventsHandler) NotifyClusterDeploymentUpdates(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyClusterDeploymentUpdates", arg0, arg1)
}

// NotifyClusterDeploymentUpdates indicates an expected call of NotifyClusterDeploymentUpdates
func (mr *MockCRDEventsHandlerMockRecorder) NotifyClusterDeploymentUpdates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyClusterDeploymentUpdates", reflect.TypeOf((*MockCRDEventsHandler)(nil).NotifyClusterDeploymentUpdates), arg0, arg1)
}

// NotifyInstallEnvUpdates mocks base method
func (m *MockCRDEventsHandler) NotifyInstallEnvUpdates(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyInstallEnvUpdates", arg0, arg1)
}

// NotifyInstallEnvUpdates indicates an expected call of NotifyInstallEnvUpdates
func (mr *MockCRDEventsHandlerMockRecorder) NotifyInstallEnvUpdates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyInstallEnvUpdates", reflect.TypeOf((*MockCRDEventsHandler)(nil).NotifyInstallEnvUpdates), arg0, arg1)
}
