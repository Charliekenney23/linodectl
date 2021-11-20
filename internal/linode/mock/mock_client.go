// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Charliekenney23/linodectl/internal/linode (interfaces: Client)

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
