// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/ec2instanceconnect/ec2instanceconnectiface (interfaces: EC2InstanceConnectAPI)

// Package aws is a generated GoMock package.
package aws

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	ec2instanceconnect "github.com/aws/aws-sdk-go/service/ec2instanceconnect"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEC2InstanceConnectAPI is a mock of EC2InstanceConnectAPI interface
type MockEC2InstanceConnectAPI struct {
	ctrl     *gomock.Controller
	recorder *MockEC2InstanceConnectAPIMockRecorder
}

// MockEC2InstanceConnectAPIMockRecorder is the mock recorder for MockEC2InstanceConnectAPI
type MockEC2InstanceConnectAPIMockRecorder struct {
	mock *MockEC2InstanceConnectAPI
}

// NewMockEC2InstanceConnectAPI creates a new mock instance
func NewMockEC2InstanceConnectAPI(ctrl *gomock.Controller) *MockEC2InstanceConnectAPI {
	mock := &MockEC2InstanceConnectAPI{ctrl: ctrl}
	mock.recorder = &MockEC2InstanceConnectAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEC2InstanceConnectAPI) EXPECT() *MockEC2InstanceConnectAPIMockRecorder {
	return m.recorder
}

// SendSSHPublicKey mocks base method
func (m *MockEC2InstanceConnectAPI) SendSSHPublicKey(arg0 *ec2instanceconnect.SendSSHPublicKeyInput) (*ec2instanceconnect.SendSSHPublicKeyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSSHPublicKey", arg0)
	ret0, _ := ret[0].(*ec2instanceconnect.SendSSHPublicKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendSSHPublicKey indicates an expected call of SendSSHPublicKey
func (mr *MockEC2InstanceConnectAPIMockRecorder) SendSSHPublicKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSSHPublicKey", reflect.TypeOf((*MockEC2InstanceConnectAPI)(nil).SendSSHPublicKey), arg0)
}

// SendSSHPublicKeyRequest mocks base method
func (m *MockEC2InstanceConnectAPI) SendSSHPublicKeyRequest(arg0 *ec2instanceconnect.SendSSHPublicKeyInput) (*request.Request, *ec2instanceconnect.SendSSHPublicKeyOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSSHPublicKeyRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*ec2instanceconnect.SendSSHPublicKeyOutput)
	return ret0, ret1
}

// SendSSHPublicKeyRequest indicates an expected call of SendSSHPublicKeyRequest
func (mr *MockEC2InstanceConnectAPIMockRecorder) SendSSHPublicKeyRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSSHPublicKeyRequest", reflect.TypeOf((*MockEC2InstanceConnectAPI)(nil).SendSSHPublicKeyRequest), arg0)
}

// SendSSHPublicKeyWithContext mocks base method
func (m *MockEC2InstanceConnectAPI) SendSSHPublicKeyWithContext(arg0 context.Context, arg1 *ec2instanceconnect.SendSSHPublicKeyInput, arg2 ...request.Option) (*ec2instanceconnect.SendSSHPublicKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendSSHPublicKeyWithContext", varargs...)
	ret0, _ := ret[0].(*ec2instanceconnect.SendSSHPublicKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendSSHPublicKeyWithContext indicates an expected call of SendSSHPublicKeyWithContext
func (mr *MockEC2InstanceConnectAPIMockRecorder) SendSSHPublicKeyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSSHPublicKeyWithContext", reflect.TypeOf((*MockEC2InstanceConnectAPI)(nil).SendSSHPublicKeyWithContext), varargs...)
}
