// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/utils (interfaces: NVMeInterface)

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	utils "github.com/netapp/trident/utils"
	models "github.com/netapp/trident/utils/models"
)

// MockNVMeInterface is a mock of NVMeInterface interface.
type MockNVMeInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNVMeInterfaceMockRecorder
}

// MockNVMeInterfaceMockRecorder is the mock recorder for MockNVMeInterface.
type MockNVMeInterfaceMockRecorder struct {
	mock *MockNVMeInterface
}

// NewMockNVMeInterface creates a new mock instance.
func NewMockNVMeInterface(ctrl *gomock.Controller) *MockNVMeInterface {
	mock := &MockNVMeInterface{ctrl: ctrl}
	mock.recorder = &MockNVMeInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNVMeInterface) EXPECT() *MockNVMeInterfaceMockRecorder {
	return m.recorder
}

// AddPublishedNVMeSession mocks base method.
func (m *MockNVMeInterface) AddPublishedNVMeSession(arg0 *utils.NVMeSessions, arg1 *models.VolumePublishInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddPublishedNVMeSession", arg0, arg1)
}

// AddPublishedNVMeSession indicates an expected call of AddPublishedNVMeSession.
func (mr *MockNVMeInterfaceMockRecorder) AddPublishedNVMeSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPublishedNVMeSession", reflect.TypeOf((*MockNVMeInterface)(nil).AddPublishedNVMeSession), arg0, arg1)
}

// GetHostNqn mocks base method.
func (m *MockNVMeInterface) GetHostNqn(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostNqn", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostNqn indicates an expected call of GetHostNqn.
func (mr *MockNVMeInterfaceMockRecorder) GetHostNqn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostNqn", reflect.TypeOf((*MockNVMeInterface)(nil).GetHostNqn), arg0)
}

// InspectNVMeSessions mocks base method.
func (m *MockNVMeInterface) InspectNVMeSessions(arg0 context.Context, arg1, arg2 *utils.NVMeSessions) []utils.NVMeSubsystem {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectNVMeSessions", arg0, arg1, arg2)
	ret0, _ := ret[0].([]utils.NVMeSubsystem)
	return ret0
}

// InspectNVMeSessions indicates an expected call of InspectNVMeSessions.
func (mr *MockNVMeInterfaceMockRecorder) InspectNVMeSessions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectNVMeSessions", reflect.TypeOf((*MockNVMeInterface)(nil).InspectNVMeSessions), arg0, arg1, arg2)
}

// NVMeActiveOnHost mocks base method.
func (m *MockNVMeInterface) NVMeActiveOnHost(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NVMeActiveOnHost", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NVMeActiveOnHost indicates an expected call of NVMeActiveOnHost.
func (mr *MockNVMeInterfaceMockRecorder) NVMeActiveOnHost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NVMeActiveOnHost", reflect.TypeOf((*MockNVMeInterface)(nil).NVMeActiveOnHost), arg0)
}

// NewNVMeDevice mocks base method.
func (m *MockNVMeInterface) NewNVMeDevice(arg0 context.Context, arg1 string) (utils.NVMeDeviceInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNVMeDevice", arg0, arg1)
	ret0, _ := ret[0].(utils.NVMeDeviceInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewNVMeDevice indicates an expected call of NewNVMeDevice.
func (mr *MockNVMeInterfaceMockRecorder) NewNVMeDevice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNVMeDevice", reflect.TypeOf((*MockNVMeInterface)(nil).NewNVMeDevice), arg0, arg1)
}

// NewNVMeSubsystem mocks base method.
func (m *MockNVMeInterface) NewNVMeSubsystem(arg0 context.Context, arg1 string) utils.NVMeSubsystemInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNVMeSubsystem", arg0, arg1)
	ret0, _ := ret[0].(utils.NVMeSubsystemInterface)
	return ret0
}

// NewNVMeSubsystem indicates an expected call of NewNVMeSubsystem.
func (mr *MockNVMeInterfaceMockRecorder) NewNVMeSubsystem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNVMeSubsystem", reflect.TypeOf((*MockNVMeInterface)(nil).NewNVMeSubsystem), arg0, arg1)
}

// PopulateCurrentNVMeSessions mocks base method.
func (m *MockNVMeInterface) PopulateCurrentNVMeSessions(arg0 context.Context, arg1 *utils.NVMeSessions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopulateCurrentNVMeSessions", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PopulateCurrentNVMeSessions indicates an expected call of PopulateCurrentNVMeSessions.
func (mr *MockNVMeInterfaceMockRecorder) PopulateCurrentNVMeSessions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopulateCurrentNVMeSessions", reflect.TypeOf((*MockNVMeInterface)(nil).PopulateCurrentNVMeSessions), arg0, arg1)
}

// RectifyNVMeSession mocks base method.
func (m *MockNVMeInterface) RectifyNVMeSession(arg0 context.Context, arg1 utils.NVMeSubsystem, arg2 *utils.NVMeSessions) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RectifyNVMeSession", arg0, arg1, arg2)
}

// RectifyNVMeSession indicates an expected call of RectifyNVMeSession.
func (mr *MockNVMeInterfaceMockRecorder) RectifyNVMeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RectifyNVMeSession", reflect.TypeOf((*MockNVMeInterface)(nil).RectifyNVMeSession), arg0, arg1, arg2)
}

// RemovePublishedNVMeSession mocks base method.
func (m *MockNVMeInterface) RemovePublishedNVMeSession(arg0 *utils.NVMeSessions, arg1, arg2 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePublishedNVMeSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// RemovePublishedNVMeSession indicates an expected call of RemovePublishedNVMeSession.
func (mr *MockNVMeInterfaceMockRecorder) RemovePublishedNVMeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePublishedNVMeSession", reflect.TypeOf((*MockNVMeInterface)(nil).RemovePublishedNVMeSession), arg0, arg1, arg2)
}
