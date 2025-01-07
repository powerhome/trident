// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/operator/clients (interfaces: OperatorCRDClientInterface,TridentCRDClientInterface,SnapshotCRDClientInterface)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/mock_operator/mock_clients/mock_api.go github.com/netapp/trident/operator/clients OperatorCRDClientInterface,TridentCRDClientInterface,SnapshotCRDClientInterface
//

// Package mock_clients is a generated GoMock package.
package mock_clients

import (
	reflect "reflect"

	v1 "github.com/kubernetes-csi/external-snapshotter/client/v8/apis/volumesnapshot/v1"
	v10 "github.com/netapp/trident/operator/crd/apis/netapp/v1"
	v11 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	gomock "go.uber.org/mock/gomock"
	types "k8s.io/apimachinery/pkg/types"
)

// MockOperatorCRDClientInterface is a mock of OperatorCRDClientInterface interface.
type MockOperatorCRDClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorCRDClientInterfaceMockRecorder
}

// MockOperatorCRDClientInterfaceMockRecorder is the mock recorder for MockOperatorCRDClientInterface.
type MockOperatorCRDClientInterfaceMockRecorder struct {
	mock *MockOperatorCRDClientInterface
}

// NewMockOperatorCRDClientInterface creates a new mock instance.
func NewMockOperatorCRDClientInterface(ctrl *gomock.Controller) *MockOperatorCRDClientInterface {
	mock := &MockOperatorCRDClientInterface{ctrl: ctrl}
	mock.recorder = &MockOperatorCRDClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorCRDClientInterface) EXPECT() *MockOperatorCRDClientInterfaceMockRecorder {
	return m.recorder
}

// GetControllingTorcCR mocks base method.
func (m *MockOperatorCRDClientInterface) GetControllingTorcCR() (*v10.TridentOrchestrator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetControllingTorcCR")
	ret0, _ := ret[0].(*v10.TridentOrchestrator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetControllingTorcCR indicates an expected call of GetControllingTorcCR.
func (mr *MockOperatorCRDClientInterfaceMockRecorder) GetControllingTorcCR() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControllingTorcCR", reflect.TypeOf((*MockOperatorCRDClientInterface)(nil).GetControllingTorcCR))
}

// GetTconfCR mocks base method.
func (m *MockOperatorCRDClientInterface) GetTconfCR(arg0 string) (*v10.TridentConfigurator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTconfCR", arg0)
	ret0, _ := ret[0].(*v10.TridentConfigurator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTconfCR indicates an expected call of GetTconfCR.
func (mr *MockOperatorCRDClientInterfaceMockRecorder) GetTconfCR(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTconfCR", reflect.TypeOf((*MockOperatorCRDClientInterface)(nil).GetTconfCR), arg0)
}

// GetTconfCRList mocks base method.
func (m *MockOperatorCRDClientInterface) GetTconfCRList() (*v10.TridentConfiguratorList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTconfCRList")
	ret0, _ := ret[0].(*v10.TridentConfiguratorList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTconfCRList indicates an expected call of GetTconfCRList.
func (mr *MockOperatorCRDClientInterfaceMockRecorder) GetTconfCRList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTconfCRList", reflect.TypeOf((*MockOperatorCRDClientInterface)(nil).GetTconfCRList))
}

// GetTorcCRList mocks base method.
func (m *MockOperatorCRDClientInterface) GetTorcCRList() (*v10.TridentOrchestratorList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTorcCRList")
	ret0, _ := ret[0].(*v10.TridentOrchestratorList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTorcCRList indicates an expected call of GetTorcCRList.
func (mr *MockOperatorCRDClientInterfaceMockRecorder) GetTorcCRList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTorcCRList", reflect.TypeOf((*MockOperatorCRDClientInterface)(nil).GetTorcCRList))
}

// UpdateTridentConfiguratorStatus mocks base method.
func (m *MockOperatorCRDClientInterface) UpdateTridentConfiguratorStatus(arg0 *v10.TridentConfigurator, arg1 v10.TridentConfiguratorStatus) (*v10.TridentConfigurator, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTridentConfiguratorStatus", arg0, arg1)
	ret0, _ := ret[0].(*v10.TridentConfigurator)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateTridentConfiguratorStatus indicates an expected call of UpdateTridentConfiguratorStatus.
func (mr *MockOperatorCRDClientInterfaceMockRecorder) UpdateTridentConfiguratorStatus(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTridentConfiguratorStatus", reflect.TypeOf((*MockOperatorCRDClientInterface)(nil).UpdateTridentConfiguratorStatus), arg0, arg1)
}

// MockTridentCRDClientInterface is a mock of TridentCRDClientInterface interface.
type MockTridentCRDClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTridentCRDClientInterfaceMockRecorder
}

// MockTridentCRDClientInterfaceMockRecorder is the mock recorder for MockTridentCRDClientInterface.
type MockTridentCRDClientInterfaceMockRecorder struct {
	mock *MockTridentCRDClientInterface
}

// NewMockTridentCRDClientInterface creates a new mock instance.
func NewMockTridentCRDClientInterface(ctrl *gomock.Controller) *MockTridentCRDClientInterface {
	mock := &MockTridentCRDClientInterface{ctrl: ctrl}
	mock.recorder = &MockTridentCRDClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTridentCRDClientInterface) EXPECT() *MockTridentCRDClientInterfaceMockRecorder {
	return m.recorder
}

// CheckTridentBackendConfigExists mocks base method.
func (m *MockTridentCRDClientInterface) CheckTridentBackendConfigExists(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTridentBackendConfigExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckTridentBackendConfigExists indicates an expected call of CheckTridentBackendConfigExists.
func (mr *MockTridentCRDClientInterfaceMockRecorder) CheckTridentBackendConfigExists(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTridentBackendConfigExists", reflect.TypeOf((*MockTridentCRDClientInterface)(nil).CheckTridentBackendConfigExists), arg0, arg1)
}

// DeleteTridentBackendConfig mocks base method.
func (m *MockTridentCRDClientInterface) DeleteTridentBackendConfig(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTridentBackendConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTridentBackendConfig indicates an expected call of DeleteTridentBackendConfig.
func (mr *MockTridentCRDClientInterfaceMockRecorder) DeleteTridentBackendConfig(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTridentBackendConfig", reflect.TypeOf((*MockTridentCRDClientInterface)(nil).DeleteTridentBackendConfig), arg0, arg1)
}

// GetTridentBackendConfig mocks base method.
func (m *MockTridentCRDClientInterface) GetTridentBackendConfig(arg0, arg1 string) (*v11.TridentBackendConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTridentBackendConfig", arg0, arg1)
	ret0, _ := ret[0].(*v11.TridentBackendConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTridentBackendConfig indicates an expected call of GetTridentBackendConfig.
func (mr *MockTridentCRDClientInterfaceMockRecorder) GetTridentBackendConfig(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTridentBackendConfig", reflect.TypeOf((*MockTridentCRDClientInterface)(nil).GetTridentBackendConfig), arg0, arg1)
}

// ListTridentBackend mocks base method.
func (m *MockTridentCRDClientInterface) ListTridentBackend(arg0 string) (*v11.TridentBackendList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTridentBackend", arg0)
	ret0, _ := ret[0].(*v11.TridentBackendList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTridentBackend indicates an expected call of ListTridentBackend.
func (mr *MockTridentCRDClientInterfaceMockRecorder) ListTridentBackend(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTridentBackend", reflect.TypeOf((*MockTridentCRDClientInterface)(nil).ListTridentBackend), arg0)
}

// PatchTridentBackendConfig mocks base method.
func (m *MockTridentCRDClientInterface) PatchTridentBackendConfig(arg0, arg1 string, arg2 []byte, arg3 types.PatchType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchTridentBackendConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchTridentBackendConfig indicates an expected call of PatchTridentBackendConfig.
func (mr *MockTridentCRDClientInterfaceMockRecorder) PatchTridentBackendConfig(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchTridentBackendConfig", reflect.TypeOf((*MockTridentCRDClientInterface)(nil).PatchTridentBackendConfig), arg0, arg1, arg2, arg3)
}

// MockSnapshotCRDClientInterface is a mock of SnapshotCRDClientInterface interface.
type MockSnapshotCRDClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotCRDClientInterfaceMockRecorder
}

// MockSnapshotCRDClientInterfaceMockRecorder is the mock recorder for MockSnapshotCRDClientInterface.
type MockSnapshotCRDClientInterfaceMockRecorder struct {
	mock *MockSnapshotCRDClientInterface
}

// NewMockSnapshotCRDClientInterface creates a new mock instance.
func NewMockSnapshotCRDClientInterface(ctrl *gomock.Controller) *MockSnapshotCRDClientInterface {
	mock := &MockSnapshotCRDClientInterface{ctrl: ctrl}
	mock.recorder = &MockSnapshotCRDClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnapshotCRDClientInterface) EXPECT() *MockSnapshotCRDClientInterfaceMockRecorder {
	return m.recorder
}

// CheckVolumeSnapshotClassExists mocks base method.
func (m *MockSnapshotCRDClientInterface) CheckVolumeSnapshotClassExists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckVolumeSnapshotClassExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckVolumeSnapshotClassExists indicates an expected call of CheckVolumeSnapshotClassExists.
func (mr *MockSnapshotCRDClientInterfaceMockRecorder) CheckVolumeSnapshotClassExists(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckVolumeSnapshotClassExists", reflect.TypeOf((*MockSnapshotCRDClientInterface)(nil).CheckVolumeSnapshotClassExists), arg0)
}

// DeleteVolumeSnapshotClass mocks base method.
func (m *MockSnapshotCRDClientInterface) DeleteVolumeSnapshotClass(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolumeSnapshotClass", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolumeSnapshotClass indicates an expected call of DeleteVolumeSnapshotClass.
func (mr *MockSnapshotCRDClientInterfaceMockRecorder) DeleteVolumeSnapshotClass(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolumeSnapshotClass", reflect.TypeOf((*MockSnapshotCRDClientInterface)(nil).DeleteVolumeSnapshotClass), arg0)
}

// GetVolumeSnapshotClass mocks base method.
func (m *MockSnapshotCRDClientInterface) GetVolumeSnapshotClass(arg0 string) (*v1.VolumeSnapshotClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeSnapshotClass", arg0)
	ret0, _ := ret[0].(*v1.VolumeSnapshotClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeSnapshotClass indicates an expected call of GetVolumeSnapshotClass.
func (mr *MockSnapshotCRDClientInterfaceMockRecorder) GetVolumeSnapshotClass(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeSnapshotClass", reflect.TypeOf((*MockSnapshotCRDClientInterface)(nil).GetVolumeSnapshotClass), arg0)
}

// PatchVolumeSnapshotClass mocks base method.
func (m *MockSnapshotCRDClientInterface) PatchVolumeSnapshotClass(arg0 string, arg1 []byte, arg2 types.PatchType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchVolumeSnapshotClass", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchVolumeSnapshotClass indicates an expected call of PatchVolumeSnapshotClass.
func (mr *MockSnapshotCRDClientInterfaceMockRecorder) PatchVolumeSnapshotClass(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchVolumeSnapshotClass", reflect.TypeOf((*MockSnapshotCRDClientInterface)(nil).PatchVolumeSnapshotClass), arg0, arg1, arg2)
}
