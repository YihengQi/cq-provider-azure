// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-azure/client/services (interfaces: DisksClient,VirtualMachinesClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	gomock "github.com/golang/mock/gomock"
)

// MockDisksClient is a mock of DisksClient interface.
type MockDisksClient struct {
	ctrl     *gomock.Controller
	recorder *MockDisksClientMockRecorder
}

// MockDisksClientMockRecorder is the mock recorder for MockDisksClient.
type MockDisksClientMockRecorder struct {
	mock *MockDisksClient
}

// NewMockDisksClient creates a new mock instance.
func NewMockDisksClient(ctrl *gomock.Controller) *MockDisksClient {
	mock := &MockDisksClient{ctrl: ctrl}
	mock.recorder = &MockDisksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDisksClient) EXPECT() *MockDisksClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockDisksClient) List(arg0 context.Context) (compute.DiskListPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(compute.DiskListPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockDisksClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDisksClient)(nil).List), arg0)
}

// MockVirtualMachinesClient is a mock of VirtualMachinesClient interface.
type MockVirtualMachinesClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachinesClientMockRecorder
}

// MockVirtualMachinesClientMockRecorder is the mock recorder for MockVirtualMachinesClient.
type MockVirtualMachinesClientMockRecorder struct {
	mock *MockVirtualMachinesClient
}

// NewMockVirtualMachinesClient creates a new mock instance.
func NewMockVirtualMachinesClient(ctrl *gomock.Controller) *MockVirtualMachinesClient {
	mock := &MockVirtualMachinesClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachinesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMachinesClient) EXPECT() *MockVirtualMachinesClientMockRecorder {
	return m.recorder
}

// ListAll mocks base method.
func (m *MockVirtualMachinesClient) ListAll(arg0 context.Context, arg1 string) (compute.VirtualMachineListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0, arg1)
	ret0, _ := ret[0].(compute.VirtualMachineListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockVirtualMachinesClientMockRecorder) ListAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockVirtualMachinesClient)(nil).ListAll), arg0, arg1)
}