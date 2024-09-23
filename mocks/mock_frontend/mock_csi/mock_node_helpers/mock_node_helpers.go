// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/frontend/csi/node_helpers (interfaces: NodeHelper,VolumePublishManager)

// Package mock_node_helpers is a generated GoMock package.
package mock_node_helpers

import (
	context "context"
	fs "io/fs"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	models "github.com/netapp/trident/utils/models"
)

// MockNodeHelper is a mock of NodeHelper interface.
type MockNodeHelper struct {
	ctrl     *gomock.Controller
	recorder *MockNodeHelperMockRecorder
}

// MockNodeHelperMockRecorder is the mock recorder for MockNodeHelper.
type MockNodeHelperMockRecorder struct {
	mock *MockNodeHelper
}

// NewMockNodeHelper creates a new mock instance.
func NewMockNodeHelper(ctrl *gomock.Controller) *MockNodeHelper {
	mock := &MockNodeHelper{ctrl: ctrl}
	mock.recorder = &MockNodeHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeHelper) EXPECT() *MockNodeHelperMockRecorder {
	return m.recorder
}

// AddPublishedPath mocks base method.
func (m *MockNodeHelper) AddPublishedPath(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPublishedPath", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPublishedPath indicates an expected call of AddPublishedPath.
func (mr *MockNodeHelperMockRecorder) AddPublishedPath(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPublishedPath", reflect.TypeOf((*MockNodeHelper)(nil).AddPublishedPath), arg0, arg1, arg2)
}

// DeleteFailedUpgradeTrackingFile mocks base method.
func (m *MockNodeHelper) DeleteFailedUpgradeTrackingFile(arg0 context.Context, arg1 fs.FileInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteFailedUpgradeTrackingFile", arg0, arg1)
}

// DeleteFailedUpgradeTrackingFile indicates an expected call of DeleteFailedUpgradeTrackingFile.
func (mr *MockNodeHelperMockRecorder) DeleteFailedUpgradeTrackingFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFailedUpgradeTrackingFile", reflect.TypeOf((*MockNodeHelper)(nil).DeleteFailedUpgradeTrackingFile), arg0, arg1)
}

// DeleteTrackingInfo mocks base method.
func (m *MockNodeHelper) DeleteTrackingInfo(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrackingInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrackingInfo indicates an expected call of DeleteTrackingInfo.
func (mr *MockNodeHelperMockRecorder) DeleteTrackingInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrackingInfo", reflect.TypeOf((*MockNodeHelper)(nil).DeleteTrackingInfo), arg0, arg1)
}

// GetVolumeTrackingFiles mocks base method.
func (m *MockNodeHelper) GetVolumeTrackingFiles() ([]fs.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeTrackingFiles")
	ret0, _ := ret[0].([]fs.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeTrackingFiles indicates an expected call of GetVolumeTrackingFiles.
func (mr *MockNodeHelperMockRecorder) GetVolumeTrackingFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeTrackingFiles", reflect.TypeOf((*MockNodeHelper)(nil).GetVolumeTrackingFiles))
}

// ListVolumeTrackingInfo mocks base method.
func (m *MockNodeHelper) ListVolumeTrackingInfo(arg0 context.Context) (map[string]*models.VolumeTrackingInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumeTrackingInfo", arg0)
	ret0, _ := ret[0].(map[string]*models.VolumeTrackingInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVolumeTrackingInfo indicates an expected call of ListVolumeTrackingInfo.
func (mr *MockNodeHelperMockRecorder) ListVolumeTrackingInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumeTrackingInfo", reflect.TypeOf((*MockNodeHelper)(nil).ListVolumeTrackingInfo), arg0)
}

// ReadTrackingInfo mocks base method.
func (m *MockNodeHelper) ReadTrackingInfo(arg0 context.Context, arg1 string) (*models.VolumeTrackingInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadTrackingInfo", arg0, arg1)
	ret0, _ := ret[0].(*models.VolumeTrackingInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadTrackingInfo indicates an expected call of ReadTrackingInfo.
func (mr *MockNodeHelperMockRecorder) ReadTrackingInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTrackingInfo", reflect.TypeOf((*MockNodeHelper)(nil).ReadTrackingInfo), arg0, arg1)
}

// RemovePublishedPath mocks base method.
func (m *MockNodeHelper) RemovePublishedPath(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePublishedPath", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePublishedPath indicates an expected call of RemovePublishedPath.
func (mr *MockNodeHelperMockRecorder) RemovePublishedPath(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePublishedPath", reflect.TypeOf((*MockNodeHelper)(nil).RemovePublishedPath), arg0, arg1, arg2)
}

// UpgradeVolumeTrackingFile mocks base method.
func (m *MockNodeHelper) UpgradeVolumeTrackingFile(arg0 context.Context, arg1 string, arg2 map[string]struct{}, arg3 map[string]string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeVolumeTrackingFile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeVolumeTrackingFile indicates an expected call of UpgradeVolumeTrackingFile.
func (mr *MockNodeHelperMockRecorder) UpgradeVolumeTrackingFile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeVolumeTrackingFile", reflect.TypeOf((*MockNodeHelper)(nil).UpgradeVolumeTrackingFile), arg0, arg1, arg2, arg3)
}

// ValidateTrackingFile mocks base method.
func (m *MockNodeHelper) ValidateTrackingFile(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTrackingFile", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateTrackingFile indicates an expected call of ValidateTrackingFile.
func (mr *MockNodeHelperMockRecorder) ValidateTrackingFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTrackingFile", reflect.TypeOf((*MockNodeHelper)(nil).ValidateTrackingFile), arg0, arg1)
}

// WriteTrackingInfo mocks base method.
func (m *MockNodeHelper) WriteTrackingInfo(arg0 context.Context, arg1 string, arg2 *models.VolumeTrackingInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTrackingInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTrackingInfo indicates an expected call of WriteTrackingInfo.
func (mr *MockNodeHelperMockRecorder) WriteTrackingInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTrackingInfo", reflect.TypeOf((*MockNodeHelper)(nil).WriteTrackingInfo), arg0, arg1, arg2)
}

// MockVolumePublishManager is a mock of VolumePublishManager interface.
type MockVolumePublishManager struct {
	ctrl     *gomock.Controller
	recorder *MockVolumePublishManagerMockRecorder
}

// MockVolumePublishManagerMockRecorder is the mock recorder for MockVolumePublishManager.
type MockVolumePublishManagerMockRecorder struct {
	mock *MockVolumePublishManager
}

// NewMockVolumePublishManager creates a new mock instance.
func NewMockVolumePublishManager(ctrl *gomock.Controller) *MockVolumePublishManager {
	mock := &MockVolumePublishManager{ctrl: ctrl}
	mock.recorder = &MockVolumePublishManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVolumePublishManager) EXPECT() *MockVolumePublishManagerMockRecorder {
	return m.recorder
}

// DeleteFailedUpgradeTrackingFile mocks base method.
func (m *MockVolumePublishManager) DeleteFailedUpgradeTrackingFile(arg0 context.Context, arg1 fs.FileInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteFailedUpgradeTrackingFile", arg0, arg1)
}

// DeleteFailedUpgradeTrackingFile indicates an expected call of DeleteFailedUpgradeTrackingFile.
func (mr *MockVolumePublishManagerMockRecorder) DeleteFailedUpgradeTrackingFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFailedUpgradeTrackingFile", reflect.TypeOf((*MockVolumePublishManager)(nil).DeleteFailedUpgradeTrackingFile), arg0, arg1)
}

// DeleteTrackingInfo mocks base method.
func (m *MockVolumePublishManager) DeleteTrackingInfo(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrackingInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrackingInfo indicates an expected call of DeleteTrackingInfo.
func (mr *MockVolumePublishManagerMockRecorder) DeleteTrackingInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrackingInfo", reflect.TypeOf((*MockVolumePublishManager)(nil).DeleteTrackingInfo), arg0, arg1)
}

// GetVolumeTrackingFiles mocks base method.
func (m *MockVolumePublishManager) GetVolumeTrackingFiles() ([]fs.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeTrackingFiles")
	ret0, _ := ret[0].([]fs.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeTrackingFiles indicates an expected call of GetVolumeTrackingFiles.
func (mr *MockVolumePublishManagerMockRecorder) GetVolumeTrackingFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeTrackingFiles", reflect.TypeOf((*MockVolumePublishManager)(nil).GetVolumeTrackingFiles))
}

// ListVolumeTrackingInfo mocks base method.
func (m *MockVolumePublishManager) ListVolumeTrackingInfo(arg0 context.Context) (map[string]*models.VolumeTrackingInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumeTrackingInfo", arg0)
	ret0, _ := ret[0].(map[string]*models.VolumeTrackingInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVolumeTrackingInfo indicates an expected call of ListVolumeTrackingInfo.
func (mr *MockVolumePublishManagerMockRecorder) ListVolumeTrackingInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumeTrackingInfo", reflect.TypeOf((*MockVolumePublishManager)(nil).ListVolumeTrackingInfo), arg0)
}

// ReadTrackingInfo mocks base method.
func (m *MockVolumePublishManager) ReadTrackingInfo(arg0 context.Context, arg1 string) (*models.VolumeTrackingInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadTrackingInfo", arg0, arg1)
	ret0, _ := ret[0].(*models.VolumeTrackingInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadTrackingInfo indicates an expected call of ReadTrackingInfo.
func (mr *MockVolumePublishManagerMockRecorder) ReadTrackingInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTrackingInfo", reflect.TypeOf((*MockVolumePublishManager)(nil).ReadTrackingInfo), arg0, arg1)
}

// UpgradeVolumeTrackingFile mocks base method.
func (m *MockVolumePublishManager) UpgradeVolumeTrackingFile(arg0 context.Context, arg1 string, arg2 map[string]struct{}, arg3 map[string]string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeVolumeTrackingFile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeVolumeTrackingFile indicates an expected call of UpgradeVolumeTrackingFile.
func (mr *MockVolumePublishManagerMockRecorder) UpgradeVolumeTrackingFile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeVolumeTrackingFile", reflect.TypeOf((*MockVolumePublishManager)(nil).UpgradeVolumeTrackingFile), arg0, arg1, arg2, arg3)
}

// ValidateTrackingFile mocks base method.
func (m *MockVolumePublishManager) ValidateTrackingFile(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTrackingFile", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateTrackingFile indicates an expected call of ValidateTrackingFile.
func (mr *MockVolumePublishManagerMockRecorder) ValidateTrackingFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTrackingFile", reflect.TypeOf((*MockVolumePublishManager)(nil).ValidateTrackingFile), arg0, arg1)
}

// WriteTrackingInfo mocks base method.
func (m *MockVolumePublishManager) WriteTrackingInfo(arg0 context.Context, arg1 string, arg2 *models.VolumeTrackingInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTrackingInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTrackingInfo indicates an expected call of WriteTrackingInfo.
func (mr *MockVolumePublishManagerMockRecorder) WriteTrackingInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTrackingInfo", reflect.TypeOf((*MockVolumePublishManager)(nil).WriteTrackingInfo), arg0, arg1, arg2)
}
