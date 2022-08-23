// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/0xch4z/linodectl/internal/linode (interfaces: Client)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	linodego "github.com/linode/linodego"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateInstance mocks base method.
func (m *MockClient) CreateInstance(arg0 context.Context, arg1 linodego.InstanceCreateOptions) (*linodego.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstance", arg0, arg1)
	ret0, _ := ret[0].(*linodego.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInstance indicates an expected call of CreateInstance.
func (mr *MockClientMockRecorder) CreateInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstance", reflect.TypeOf((*MockClient)(nil).CreateInstance), arg0, arg1)
}

// CreateLKECluster mocks base method.
func (m *MockClient) CreateLKECluster(arg0 context.Context, arg1 linodego.LKEClusterCreateOptions) (*linodego.LKECluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLKECluster", arg0, arg1)
	ret0, _ := ret[0].(*linodego.LKECluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLKECluster indicates an expected call of CreateLKECluster.
func (mr *MockClientMockRecorder) CreateLKECluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLKECluster", reflect.TypeOf((*MockClient)(nil).CreateLKECluster), arg0, arg1)
}

// CreateObjectStorageBucket mocks base method.
func (m *MockClient) CreateObjectStorageBucket(arg0 context.Context, arg1 linodego.ObjectStorageBucketCreateOptions) (*linodego.ObjectStorageBucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateObjectStorageBucket", arg0, arg1)
	ret0, _ := ret[0].(*linodego.ObjectStorageBucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateObjectStorageBucket indicates an expected call of CreateObjectStorageBucket.
func (mr *MockClientMockRecorder) CreateObjectStorageBucket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateObjectStorageBucket", reflect.TypeOf((*MockClient)(nil).CreateObjectStorageBucket), arg0, arg1)
}

// CreateObjectStorageKey mocks base method.
func (m *MockClient) CreateObjectStorageKey(arg0 context.Context, arg1 linodego.ObjectStorageKeyCreateOptions) (*linodego.ObjectStorageKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateObjectStorageKey", arg0, arg1)
	ret0, _ := ret[0].(*linodego.ObjectStorageKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateObjectStorageKey indicates an expected call of CreateObjectStorageKey.
func (mr *MockClientMockRecorder) CreateObjectStorageKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateObjectStorageKey", reflect.TypeOf((*MockClient)(nil).CreateObjectStorageKey), arg0, arg1)
}

// DeleteInstance mocks base method.
func (m *MockClient) DeleteInstance(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInstance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInstance indicates an expected call of DeleteInstance.
func (mr *MockClientMockRecorder) DeleteInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstance", reflect.TypeOf((*MockClient)(nil).DeleteInstance), arg0, arg1)
}

// DeleteLKECluster mocks base method.
func (m *MockClient) DeleteLKECluster(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLKECluster", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLKECluster indicates an expected call of DeleteLKECluster.
func (mr *MockClientMockRecorder) DeleteLKECluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLKECluster", reflect.TypeOf((*MockClient)(nil).DeleteLKECluster), arg0, arg1)
}

// DeleteObjectStorageBucket mocks base method.
func (m *MockClient) DeleteObjectStorageBucket(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjectStorageBucket", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObjectStorageBucket indicates an expected call of DeleteObjectStorageBucket.
func (mr *MockClientMockRecorder) DeleteObjectStorageBucket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectStorageBucket", reflect.TypeOf((*MockClient)(nil).DeleteObjectStorageBucket), arg0, arg1, arg2)
}

// DeleteObjectStorageKey mocks base method.
func (m *MockClient) DeleteObjectStorageKey(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjectStorageKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObjectStorageKey indicates an expected call of DeleteObjectStorageKey.
func (mr *MockClientMockRecorder) DeleteObjectStorageKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectStorageKey", reflect.TypeOf((*MockClient)(nil).DeleteObjectStorageKey), arg0, arg1)
}

// DeleteStackscript mocks base method.
func (m *MockClient) DeleteStackscript(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStackscript", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStackscript indicates an expected call of DeleteStackscript.
func (mr *MockClientMockRecorder) DeleteStackscript(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStackscript", reflect.TypeOf((*MockClient)(nil).DeleteStackscript), arg0, arg1)
}

// GetInstance mocks base method.
func (m *MockClient) GetInstance(arg0 context.Context, arg1 int) (*linodego.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstance", arg0, arg1)
	ret0, _ := ret[0].(*linodego.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstance indicates an expected call of GetInstance.
func (mr *MockClientMockRecorder) GetInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstance", reflect.TypeOf((*MockClient)(nil).GetInstance), arg0, arg1)
}

// GetLKEClusterKubeconfig mocks base method.
func (m *MockClient) GetLKEClusterKubeconfig(arg0 context.Context, arg1 int) (*linodego.LKEClusterKubeconfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLKEClusterKubeconfig", arg0, arg1)
	ret0, _ := ret[0].(*linodego.LKEClusterKubeconfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLKEClusterKubeconfig indicates an expected call of GetLKEClusterKubeconfig.
func (mr *MockClientMockRecorder) GetLKEClusterKubeconfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLKEClusterKubeconfig", reflect.TypeOf((*MockClient)(nil).GetLKEClusterKubeconfig), arg0, arg1)
}

// GetProfile mocks base method.
func (m *MockClient) GetProfile(arg0 context.Context) (*linodego.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", arg0)
	ret0, _ := ret[0].(*linodego.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockClientMockRecorder) GetProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockClient)(nil).GetProfile), arg0)
}

// ListInstances mocks base method.
func (m *MockClient) ListInstances(arg0 context.Context, arg1 *linodego.ListOptions) ([]linodego.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInstances", arg0, arg1)
	ret0, _ := ret[0].([]linodego.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstances indicates an expected call of ListInstances.
func (mr *MockClientMockRecorder) ListInstances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstances", reflect.TypeOf((*MockClient)(nil).ListInstances), arg0, arg1)
}

// ListLKEClusterPools mocks base method.
func (m *MockClient) ListLKEClusterPools(arg0 context.Context, arg1 int, arg2 *linodego.ListOptions) ([]linodego.LKEClusterPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLKEClusterPools", arg0, arg1, arg2)
	ret0, _ := ret[0].([]linodego.LKEClusterPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLKEClusterPools indicates an expected call of ListLKEClusterPools.
func (mr *MockClientMockRecorder) ListLKEClusterPools(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLKEClusterPools", reflect.TypeOf((*MockClient)(nil).ListLKEClusterPools), arg0, arg1, arg2)
}

// ListLKEClusters mocks base method.
func (m *MockClient) ListLKEClusters(arg0 context.Context, arg1 *linodego.ListOptions) ([]linodego.LKECluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLKEClusters", arg0, arg1)
	ret0, _ := ret[0].([]linodego.LKECluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLKEClusters indicates an expected call of ListLKEClusters.
func (mr *MockClientMockRecorder) ListLKEClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLKEClusters", reflect.TypeOf((*MockClient)(nil).ListLKEClusters), arg0, arg1)
}

// ListObjectStorageBuckets mocks base method.
func (m *MockClient) ListObjectStorageBuckets(arg0 context.Context, arg1 *linodego.ListOptions) ([]linodego.ObjectStorageBucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectStorageBuckets", arg0, arg1)
	ret0, _ := ret[0].([]linodego.ObjectStorageBucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjectStorageBuckets indicates an expected call of ListObjectStorageBuckets.
func (mr *MockClientMockRecorder) ListObjectStorageBuckets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectStorageBuckets", reflect.TypeOf((*MockClient)(nil).ListObjectStorageBuckets), arg0, arg1)
}

// ListStackscripts mocks base method.
func (m *MockClient) ListStackscripts(arg0 context.Context, arg1 *linodego.ListOptions) ([]linodego.Stackscript, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStackscripts", arg0, arg1)
	ret0, _ := ret[0].([]linodego.Stackscript)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStackscripts indicates an expected call of ListStackscripts.
func (mr *MockClientMockRecorder) ListStackscripts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStackscripts", reflect.TypeOf((*MockClient)(nil).ListStackscripts), arg0, arg1)
}

// UpdateInstance mocks base method.
func (m *MockClient) UpdateInstance(arg0 context.Context, arg1 int, arg2 linodego.InstanceUpdateOptions) (*linodego.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInstance", arg0, arg1, arg2)
	ret0, _ := ret[0].(*linodego.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInstance indicates an expected call of UpdateInstance.
func (mr *MockClientMockRecorder) UpdateInstance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstance", reflect.TypeOf((*MockClient)(nil).UpdateInstance), arg0, arg1, arg2)
}

// UpdateLKECluster mocks base method.
func (m *MockClient) UpdateLKECluster(arg0 context.Context, arg1 int, arg2 linodego.LKEClusterUpdateOptions) (*linodego.LKECluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLKECluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(*linodego.LKECluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLKECluster indicates an expected call of UpdateLKECluster.
func (mr *MockClientMockRecorder) UpdateLKECluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLKECluster", reflect.TypeOf((*MockClient)(nil).UpdateLKECluster), arg0, arg1, arg2)
}

// UpdateLKEClusterPool mocks base method.
func (m *MockClient) UpdateLKEClusterPool(arg0 context.Context, arg1, arg2 int, arg3 linodego.LKEClusterPoolUpdateOptions) (*linodego.LKEClusterPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLKEClusterPool", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*linodego.LKEClusterPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLKEClusterPool indicates an expected call of UpdateLKEClusterPool.
func (mr *MockClientMockRecorder) UpdateLKEClusterPool(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLKEClusterPool", reflect.TypeOf((*MockClient)(nil).UpdateLKEClusterPool), arg0, arg1, arg2, arg3)
}
