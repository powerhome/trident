// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/storage_drivers/azure/api (interfaces: Azure)
//
// Generated by this command:
//
//	mockgen -destination=../../../mocks/mock_storage_drivers/mock_azure/mock_api.go github.com/netapp/trident/storage_drivers/azure/api Azure
//

// Package mock_api is a generated GoMock package.
package mock_api

import (
	context "context"
	reflect "reflect"
	time "time"

	storage "github.com/netapp/trident/storage"
	api "github.com/netapp/trident/storage_drivers/azure/api"
	gomock "go.uber.org/mock/gomock"
)

// MockAzure is a mock of Azure interface.
type MockAzure struct {
	ctrl     *gomock.Controller
	recorder *MockAzureMockRecorder
}

// MockAzureMockRecorder is the mock recorder for MockAzure.
type MockAzureMockRecorder struct {
	mock *MockAzure
}

// NewMockAzure creates a new mock instance.
func NewMockAzure(ctrl *gomock.Controller) *MockAzure {
	mock := &MockAzure{ctrl: ctrl}
	mock.recorder = &MockAzureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzure) EXPECT() *MockAzureMockRecorder {
	return m.recorder
}

// AvailabilityZones mocks base method.
func (m *MockAzure) AvailabilityZones(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilityZones", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AvailabilityZones indicates an expected call of AvailabilityZones.
func (mr *MockAzureMockRecorder) AvailabilityZones(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilityZones", reflect.TypeOf((*MockAzure)(nil).AvailabilityZones), arg0)
}

// CapacityPools mocks base method.
func (m *MockAzure) CapacityPools() *[]*api.CapacityPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CapacityPools")
	ret0, _ := ret[0].(*[]*api.CapacityPool)
	return ret0
}

// CapacityPools indicates an expected call of CapacityPools.
func (mr *MockAzureMockRecorder) CapacityPools() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CapacityPools", reflect.TypeOf((*MockAzure)(nil).CapacityPools))
}

// CapacityPoolsForStoragePool mocks base method.
func (m *MockAzure) CapacityPoolsForStoragePool(arg0 context.Context, arg1 storage.Pool, arg2 string) []*api.CapacityPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CapacityPoolsForStoragePool", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*api.CapacityPool)
	return ret0
}

// CapacityPoolsForStoragePool indicates an expected call of CapacityPoolsForStoragePool.
func (mr *MockAzureMockRecorder) CapacityPoolsForStoragePool(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CapacityPoolsForStoragePool", reflect.TypeOf((*MockAzure)(nil).CapacityPoolsForStoragePool), arg0, arg1, arg2)
}

// CapacityPoolsForStoragePools mocks base method.
func (m *MockAzure) CapacityPoolsForStoragePools(arg0 context.Context) []*api.CapacityPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CapacityPoolsForStoragePools", arg0)
	ret0, _ := ret[0].([]*api.CapacityPool)
	return ret0
}

// CapacityPoolsForStoragePools indicates an expected call of CapacityPoolsForStoragePools.
func (mr *MockAzureMockRecorder) CapacityPoolsForStoragePools(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CapacityPoolsForStoragePools", reflect.TypeOf((*MockAzure)(nil).CapacityPoolsForStoragePools), arg0)
}

// CreateSnapshot mocks base method.
func (m *MockAzure) CreateSnapshot(arg0 context.Context, arg1 *api.FileSystem, arg2 string) (*api.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockAzureMockRecorder) CreateSnapshot(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockAzure)(nil).CreateSnapshot), arg0, arg1, arg2)
}

// CreateVolume mocks base method.
func (m *MockAzure) CreateVolume(arg0 context.Context, arg1 *api.FilesystemCreateRequest) (*api.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", arg0, arg1)
	ret0, _ := ret[0].(*api.FileSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolume indicates an expected call of CreateVolume.
func (mr *MockAzureMockRecorder) CreateVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockAzure)(nil).CreateVolume), arg0, arg1)
}

// DeleteSnapshot mocks base method.
func (m *MockAzure) DeleteSnapshot(arg0 context.Context, arg1 *api.FileSystem, arg2 *api.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockAzureMockRecorder) DeleteSnapshot(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockAzure)(nil).DeleteSnapshot), arg0, arg1, arg2)
}

// DeleteVolume mocks base method.
func (m *MockAzure) DeleteVolume(arg0 context.Context, arg1 *api.FileSystem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolume indicates an expected call of DeleteVolume.
func (mr *MockAzureMockRecorder) DeleteVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockAzure)(nil).DeleteVolume), arg0, arg1)
}

// DiscoverAzureResources mocks base method.
func (m *MockAzure) DiscoverAzureResources(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoverAzureResources", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DiscoverAzureResources indicates an expected call of DiscoverAzureResources.
func (mr *MockAzureMockRecorder) DiscoverAzureResources(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverAzureResources", reflect.TypeOf((*MockAzure)(nil).DiscoverAzureResources), arg0)
}

// EnableAzureFeatures mocks base method.
func (m *MockAzure) EnableAzureFeatures(arg0 context.Context, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableAzureFeatures", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableAzureFeatures indicates an expected call of EnableAzureFeatures.
func (mr *MockAzureMockRecorder) EnableAzureFeatures(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableAzureFeatures", reflect.TypeOf((*MockAzure)(nil).EnableAzureFeatures), varargs...)
}

// EnsureVolumeInValidCapacityPool mocks base method.
func (m *MockAzure) EnsureVolumeInValidCapacityPool(arg0 context.Context, arg1 *api.FileSystem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureVolumeInValidCapacityPool", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureVolumeInValidCapacityPool indicates an expected call of EnsureVolumeInValidCapacityPool.
func (mr *MockAzureMockRecorder) EnsureVolumeInValidCapacityPool(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureVolumeInValidCapacityPool", reflect.TypeOf((*MockAzure)(nil).EnsureVolumeInValidCapacityPool), arg0, arg1)
}

// Features mocks base method.
func (m *MockAzure) Features() map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Features")
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// Features indicates an expected call of Features.
func (mr *MockAzureMockRecorder) Features() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Features", reflect.TypeOf((*MockAzure)(nil).Features))
}

// FilteredCapacityPoolMap mocks base method.
func (m *MockAzure) FilteredCapacityPoolMap(arg0 context.Context, arg1, arg2, arg3 []string) map[string]*api.CapacityPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilteredCapacityPoolMap", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(map[string]*api.CapacityPool)
	return ret0
}

// FilteredCapacityPoolMap indicates an expected call of FilteredCapacityPoolMap.
func (mr *MockAzureMockRecorder) FilteredCapacityPoolMap(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilteredCapacityPoolMap", reflect.TypeOf((*MockAzure)(nil).FilteredCapacityPoolMap), arg0, arg1, arg2, arg3)
}

// FilteredSubnetMap mocks base method.
func (m *MockAzure) FilteredSubnetMap(arg0 context.Context, arg1 []string, arg2, arg3 string) map[string]*api.Subnet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilteredSubnetMap", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(map[string]*api.Subnet)
	return ret0
}

// FilteredSubnetMap indicates an expected call of FilteredSubnetMap.
func (mr *MockAzureMockRecorder) FilteredSubnetMap(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilteredSubnetMap", reflect.TypeOf((*MockAzure)(nil).FilteredSubnetMap), arg0, arg1, arg2, arg3)
}

// HasFeature mocks base method.
func (m *MockAzure) HasFeature(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasFeature", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasFeature indicates an expected call of HasFeature.
func (mr *MockAzureMockRecorder) HasFeature(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasFeature", reflect.TypeOf((*MockAzure)(nil).HasFeature), arg0)
}

// Init mocks base method.
func (m *MockAzure) Init(arg0 context.Context, arg1 map[string]storage.Pool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockAzureMockRecorder) Init(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockAzure)(nil).Init), arg0, arg1)
}

// ModifyVolume mocks base method.
func (m *MockAzure) ModifyVolume(arg0 context.Context, arg1 *api.FileSystem, arg2 map[string]string, arg3 *string, arg4 *bool, arg5 *api.ExportRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyVolume", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyVolume indicates an expected call of ModifyVolume.
func (mr *MockAzureMockRecorder) ModifyVolume(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyVolume", reflect.TypeOf((*MockAzure)(nil).ModifyVolume), arg0, arg1, arg2, arg3, arg4, arg5)
}

// RandomSubnetForStoragePool mocks base method.
func (m *MockAzure) RandomSubnetForStoragePool(arg0 context.Context, arg1 storage.Pool) *api.Subnet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomSubnetForStoragePool", arg0, arg1)
	ret0, _ := ret[0].(*api.Subnet)
	return ret0
}

// RandomSubnetForStoragePool indicates an expected call of RandomSubnetForStoragePool.
func (mr *MockAzureMockRecorder) RandomSubnetForStoragePool(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomSubnetForStoragePool", reflect.TypeOf((*MockAzure)(nil).RandomSubnetForStoragePool), arg0, arg1)
}

// RefreshAzureResources mocks base method.
func (m *MockAzure) RefreshAzureResources(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshAzureResources", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshAzureResources indicates an expected call of RefreshAzureResources.
func (mr *MockAzureMockRecorder) RefreshAzureResources(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshAzureResources", reflect.TypeOf((*MockAzure)(nil).RefreshAzureResources), arg0)
}

// ResizeVolume mocks base method.
func (m *MockAzure) ResizeVolume(arg0 context.Context, arg1 *api.FileSystem, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResizeVolume indicates an expected call of ResizeVolume.
func (mr *MockAzureMockRecorder) ResizeVolume(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizeVolume", reflect.TypeOf((*MockAzure)(nil).ResizeVolume), arg0, arg1, arg2)
}

// RestoreSnapshot mocks base method.
func (m *MockAzure) RestoreSnapshot(arg0 context.Context, arg1 *api.FileSystem, arg2 *api.Snapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreSnapshot indicates an expected call of RestoreSnapshot.
func (mr *MockAzureMockRecorder) RestoreSnapshot(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreSnapshot", reflect.TypeOf((*MockAzure)(nil).RestoreSnapshot), arg0, arg1, arg2)
}

// SnapshotForVolume mocks base method.
func (m *MockAzure) SnapshotForVolume(arg0 context.Context, arg1 *api.FileSystem, arg2 string) (*api.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotForVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotForVolume indicates an expected call of SnapshotForVolume.
func (mr *MockAzureMockRecorder) SnapshotForVolume(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotForVolume", reflect.TypeOf((*MockAzure)(nil).SnapshotForVolume), arg0, arg1, arg2)
}

// SnapshotsForVolume mocks base method.
func (m *MockAzure) SnapshotsForVolume(arg0 context.Context, arg1 *api.FileSystem) (*[]*api.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotsForVolume", arg0, arg1)
	ret0, _ := ret[0].(*[]*api.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotsForVolume indicates an expected call of SnapshotsForVolume.
func (mr *MockAzureMockRecorder) SnapshotsForVolume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotsForVolume", reflect.TypeOf((*MockAzure)(nil).SnapshotsForVolume), arg0, arg1)
}

// SubnetsForStoragePool mocks base method.
func (m *MockAzure) SubnetsForStoragePool(arg0 context.Context, arg1 storage.Pool) []*api.Subnet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubnetsForStoragePool", arg0, arg1)
	ret0, _ := ret[0].([]*api.Subnet)
	return ret0
}

// SubnetsForStoragePool indicates an expected call of SubnetsForStoragePool.
func (mr *MockAzureMockRecorder) SubnetsForStoragePool(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubnetsForStoragePool", reflect.TypeOf((*MockAzure)(nil).SubnetsForStoragePool), arg0, arg1)
}

// Volume mocks base method.
func (m *MockAzure) Volume(arg0 context.Context, arg1 *storage.VolumeConfig) (*api.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Volume", arg0, arg1)
	ret0, _ := ret[0].(*api.FileSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Volume indicates an expected call of Volume.
func (mr *MockAzureMockRecorder) Volume(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Volume", reflect.TypeOf((*MockAzure)(nil).Volume), arg0, arg1)
}

// VolumeByCreationToken mocks base method.
func (m *MockAzure) VolumeByCreationToken(arg0 context.Context, arg1 string) (*api.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeByCreationToken", arg0, arg1)
	ret0, _ := ret[0].(*api.FileSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeByCreationToken indicates an expected call of VolumeByCreationToken.
func (mr *MockAzureMockRecorder) VolumeByCreationToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeByCreationToken", reflect.TypeOf((*MockAzure)(nil).VolumeByCreationToken), arg0, arg1)
}

// VolumeByID mocks base method.
func (m *MockAzure) VolumeByID(arg0 context.Context, arg1 string) (*api.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeByID", arg0, arg1)
	ret0, _ := ret[0].(*api.FileSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeByID indicates an expected call of VolumeByID.
func (mr *MockAzureMockRecorder) VolumeByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeByID", reflect.TypeOf((*MockAzure)(nil).VolumeByID), arg0, arg1)
}

// VolumeExists mocks base method.
func (m *MockAzure) VolumeExists(arg0 context.Context, arg1 *storage.VolumeConfig) (bool, *api.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*api.FileSystem)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VolumeExists indicates an expected call of VolumeExists.
func (mr *MockAzureMockRecorder) VolumeExists(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeExists", reflect.TypeOf((*MockAzure)(nil).VolumeExists), arg0, arg1)
}

// VolumeExistsByCreationToken mocks base method.
func (m *MockAzure) VolumeExistsByCreationToken(arg0 context.Context, arg1 string) (bool, *api.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeExistsByCreationToken", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*api.FileSystem)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VolumeExistsByCreationToken indicates an expected call of VolumeExistsByCreationToken.
func (mr *MockAzureMockRecorder) VolumeExistsByCreationToken(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeExistsByCreationToken", reflect.TypeOf((*MockAzure)(nil).VolumeExistsByCreationToken), arg0, arg1)
}

// VolumeExistsByID mocks base method.
func (m *MockAzure) VolumeExistsByID(arg0 context.Context, arg1 string) (bool, *api.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeExistsByID", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*api.FileSystem)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VolumeExistsByID indicates an expected call of VolumeExistsByID.
func (mr *MockAzureMockRecorder) VolumeExistsByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeExistsByID", reflect.TypeOf((*MockAzure)(nil).VolumeExistsByID), arg0, arg1)
}

// Volumes mocks base method.
func (m *MockAzure) Volumes(arg0 context.Context) (*[]*api.FileSystem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Volumes", arg0)
	ret0, _ := ret[0].(*[]*api.FileSystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Volumes indicates an expected call of Volumes.
func (mr *MockAzureMockRecorder) Volumes(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Volumes", reflect.TypeOf((*MockAzure)(nil).Volumes), arg0)
}

// WaitForSnapshotState mocks base method.
func (m *MockAzure) WaitForSnapshotState(arg0 context.Context, arg1 *api.Snapshot, arg2 *api.FileSystem, arg3 string, arg4 []string, arg5 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForSnapshotState", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForSnapshotState indicates an expected call of WaitForSnapshotState.
func (mr *MockAzureMockRecorder) WaitForSnapshotState(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForSnapshotState", reflect.TypeOf((*MockAzure)(nil).WaitForSnapshotState), arg0, arg1, arg2, arg3, arg4, arg5)
}

// WaitForVolumeState mocks base method.
func (m *MockAzure) WaitForVolumeState(arg0 context.Context, arg1 *api.FileSystem, arg2 string, arg3 []string, arg4 time.Duration, arg5 api.Operation) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForVolumeState", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForVolumeState indicates an expected call of WaitForVolumeState.
func (mr *MockAzureMockRecorder) WaitForVolumeState(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForVolumeState", reflect.TypeOf((*MockAzure)(nil).WaitForVolumeState), arg0, arg1, arg2, arg3, arg4, arg5)
}
