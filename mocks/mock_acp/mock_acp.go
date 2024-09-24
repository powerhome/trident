// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/acp (interfaces: TridentACP)
//
// Generated by this command:
//
//	mockgen -destination=../mocks/mock_acp/mock_acp.go github.com/netapp/trident/acp TridentACP
//

// Package mock_acp is a generated GoMock package.
package mock_acp

import (
	context "context"
	reflect "reflect"

	version "github.com/netapp/trident/utils/version"
	gomock "go.uber.org/mock/gomock"
)

// MockTridentACP is a mock of TridentACP interface.
type MockTridentACP struct {
	ctrl     *gomock.Controller
	recorder *MockTridentACPMockRecorder
}

// MockTridentACPMockRecorder is the mock recorder for MockTridentACP.
type MockTridentACPMockRecorder struct {
	mock *MockTridentACP
}

// NewMockTridentACP creates a new mock instance.
func NewMockTridentACP(ctrl *gomock.Controller) *MockTridentACP {
	mock := &MockTridentACP{ctrl: ctrl}
	mock.recorder = &MockTridentACPMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTridentACP) EXPECT() *MockTridentACPMockRecorder {
	return m.recorder
}

// Enabled mocks base method.
func (m *MockTridentACP) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockTridentACPMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockTridentACP)(nil).Enabled))
}

// GetVersion mocks base method.
func (m *MockTridentACP) GetVersion(arg0 context.Context) (*version.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion", arg0)
	ret0, _ := ret[0].(*version.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockTridentACPMockRecorder) GetVersion(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockTridentACP)(nil).GetVersion), arg0)
}

// GetVersionWithBackoff mocks base method.
func (m *MockTridentACP) GetVersionWithBackoff(arg0 context.Context) (*version.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersionWithBackoff", arg0)
	ret0, _ := ret[0].(*version.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersionWithBackoff indicates an expected call of GetVersionWithBackoff.
func (mr *MockTridentACPMockRecorder) GetVersionWithBackoff(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionWithBackoff", reflect.TypeOf((*MockTridentACP)(nil).GetVersionWithBackoff), arg0)
}

// IsFeatureEnabled mocks base method.
func (m *MockTridentACP) IsFeatureEnabled(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFeatureEnabled", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsFeatureEnabled indicates an expected call of IsFeatureEnabled.
func (mr *MockTridentACPMockRecorder) IsFeatureEnabled(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFeatureEnabled", reflect.TypeOf((*MockTridentACP)(nil).IsFeatureEnabled), arg0, arg1)
}
