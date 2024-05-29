// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/base-org/pessimism/internal/client (interfaces: SNSClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	client "github.com/base-org/pessimism/internal/client"
	gomock "github.com/golang/mock/gomock"
)

// MockSNSClient is a mock of SNSClient interface.
type MockSNSClient struct {
	ctrl     *gomock.Controller
	recorder *MockSNSClientMockRecorder
}

// MockSNSClientMockRecorder is the mock recorder for MockSNSClient.
type MockSNSClientMockRecorder struct {
	mock *MockSNSClient
}

// NewMockSNSClient creates a new mock instance.
func NewMockSNSClient(ctrl *gomock.Controller) *MockSNSClient {
	mock := &MockSNSClient{ctrl: ctrl}
	mock.recorder = &MockSNSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSNSClient) EXPECT() *MockSNSClientMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockSNSClient) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockSNSClientMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockSNSClient)(nil).GetName))
}

// PostEvent mocks base method.
func (m *MockSNSClient) PostEvent(arg0 context.Context, arg1 *client.AlertEventTrigger) (*client.AlertAPIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostEvent", arg0, arg1)
	ret0, _ := ret[0].(*client.AlertAPIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostEvent indicates an expected call of PostEvent.
func (mr *MockSNSClientMockRecorder) PostEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostEvent", reflect.TypeOf((*MockSNSClient)(nil).PostEvent), arg0, arg1)
}