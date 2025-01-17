// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-azure/client/services (interfaces: VirtualNetworksClient,SecurityGroupsClient,WatchersClient,PublicIPAddressesClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	gomock "github.com/golang/mock/gomock"
)

// MockVirtualNetworksClient is a mock of VirtualNetworksClient interface.
type MockVirtualNetworksClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualNetworksClientMockRecorder
}

// MockVirtualNetworksClientMockRecorder is the mock recorder for MockVirtualNetworksClient.
type MockVirtualNetworksClientMockRecorder struct {
	mock *MockVirtualNetworksClient
}

// NewMockVirtualNetworksClient creates a new mock instance.
func NewMockVirtualNetworksClient(ctrl *gomock.Controller) *MockVirtualNetworksClient {
	mock := &MockVirtualNetworksClient{ctrl: ctrl}
	mock.recorder = &MockVirtualNetworksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualNetworksClient) EXPECT() *MockVirtualNetworksClientMockRecorder {
	return m.recorder
}

// ListAll mocks base method.
func (m *MockVirtualNetworksClient) ListAll(arg0 context.Context) (network.VirtualNetworkListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.VirtualNetworkListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockVirtualNetworksClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockVirtualNetworksClient)(nil).ListAll), arg0)
}

// MockSecurityGroupsClient is a mock of SecurityGroupsClient interface.
type MockSecurityGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityGroupsClientMockRecorder
}

// MockSecurityGroupsClientMockRecorder is the mock recorder for MockSecurityGroupsClient.
type MockSecurityGroupsClientMockRecorder struct {
	mock *MockSecurityGroupsClient
}

// NewMockSecurityGroupsClient creates a new mock instance.
func NewMockSecurityGroupsClient(ctrl *gomock.Controller) *MockSecurityGroupsClient {
	mock := &MockSecurityGroupsClient{ctrl: ctrl}
	mock.recorder = &MockSecurityGroupsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecurityGroupsClient) EXPECT() *MockSecurityGroupsClientMockRecorder {
	return m.recorder
}

// ListAll mocks base method.
func (m *MockSecurityGroupsClient) ListAll(arg0 context.Context) (network.SecurityGroupListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.SecurityGroupListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockSecurityGroupsClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockSecurityGroupsClient)(nil).ListAll), arg0)
}

// MockWatchersClient is a mock of WatchersClient interface.
type MockWatchersClient struct {
	ctrl     *gomock.Controller
	recorder *MockWatchersClientMockRecorder
}

// MockWatchersClientMockRecorder is the mock recorder for MockWatchersClient.
type MockWatchersClientMockRecorder struct {
	mock *MockWatchersClient
}

// NewMockWatchersClient creates a new mock instance.
func NewMockWatchersClient(ctrl *gomock.Controller) *MockWatchersClient {
	mock := &MockWatchersClient{ctrl: ctrl}
	mock.recorder = &MockWatchersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatchersClient) EXPECT() *MockWatchersClientMockRecorder {
	return m.recorder
}

// GetFlowLogStatus mocks base method.
func (m *MockWatchersClient) GetFlowLogStatus(arg0 context.Context, arg1, arg2 string, arg3 network.FlowLogStatusParameters) (network.WatchersGetFlowLogStatusFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlowLogStatus", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.WatchersGetFlowLogStatusFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlowLogStatus indicates an expected call of GetFlowLogStatus.
func (mr *MockWatchersClientMockRecorder) GetFlowLogStatus(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlowLogStatus", reflect.TypeOf((*MockWatchersClient)(nil).GetFlowLogStatus), arg0, arg1, arg2, arg3)
}

// ListAll mocks base method.
func (m *MockWatchersClient) ListAll(arg0 context.Context) (network.WatcherListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.WatcherListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockWatchersClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockWatchersClient)(nil).ListAll), arg0)
}

// MockPublicIPAddressesClient is a mock of PublicIPAddressesClient interface.
type MockPublicIPAddressesClient struct {
	ctrl     *gomock.Controller
	recorder *MockPublicIPAddressesClientMockRecorder
}

// MockPublicIPAddressesClientMockRecorder is the mock recorder for MockPublicIPAddressesClient.
type MockPublicIPAddressesClientMockRecorder struct {
	mock *MockPublicIPAddressesClient
}

// NewMockPublicIPAddressesClient creates a new mock instance.
func NewMockPublicIPAddressesClient(ctrl *gomock.Controller) *MockPublicIPAddressesClient {
	mock := &MockPublicIPAddressesClient{ctrl: ctrl}
	mock.recorder = &MockPublicIPAddressesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPublicIPAddressesClient) EXPECT() *MockPublicIPAddressesClientMockRecorder {
	return m.recorder
}

// ListAll mocks base method.
func (m *MockPublicIPAddressesClient) ListAll(arg0 context.Context) (network.PublicIPAddressListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.PublicIPAddressListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockPublicIPAddressesClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockPublicIPAddressesClient)(nil).ListAll), arg0)
}
