// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/danmx/sigil/pkg/aws (interfaces: Cloud,CloudInstances,CloudSessions,CloudSSH)

// Package session is a generated GoMock package.
package session

import (
	aws "github.com/danmx/sigil/pkg/aws"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCloud is a mock of Cloud interface
type MockCloud struct {
	ctrl     *gomock.Controller
	recorder *MockCloudMockRecorder
}

// MockCloudMockRecorder is the mock recorder for MockCloud
type MockCloudMockRecorder struct {
	mock *MockCloud
}

// NewMockCloud creates a new mock instance
func NewMockCloud(ctrl *gomock.Controller) *MockCloud {
	mock := &MockCloud{ctrl: ctrl}
	mock.recorder = &MockCloudMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloud) EXPECT() *MockCloudMockRecorder {
	return m.recorder
}

// NewWithConfig mocks base method
func (m *MockCloud) NewWithConfig(arg0 *aws.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWithConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewWithConfig indicates an expected call of NewWithConfig
func (mr *MockCloudMockRecorder) NewWithConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWithConfig", reflect.TypeOf((*MockCloud)(nil).NewWithConfig), arg0)
}

// MockCloudInstances is a mock of CloudInstances interface
type MockCloudInstances struct {
	ctrl     *gomock.Controller
	recorder *MockCloudInstancesMockRecorder
}

// MockCloudInstancesMockRecorder is the mock recorder for MockCloudInstances
type MockCloudInstancesMockRecorder struct {
	mock *MockCloudInstances
}

// NewMockCloudInstances creates a new mock instance
func NewMockCloudInstances(ctrl *gomock.Controller) *MockCloudInstances {
	mock := &MockCloudInstances{ctrl: ctrl}
	mock.recorder = &MockCloudInstancesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudInstances) EXPECT() *MockCloudInstancesMockRecorder {
	return m.recorder
}

// ListInstances mocks base method
func (m *MockCloudInstances) ListInstances() ([]*aws.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInstances")
	ret0, _ := ret[0].([]*aws.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstances indicates an expected call of ListInstances
func (mr *MockCloudInstancesMockRecorder) ListInstances() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstances", reflect.TypeOf((*MockCloudInstances)(nil).ListInstances))
}

// NewWithConfig mocks base method
func (m *MockCloudInstances) NewWithConfig(arg0 *aws.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWithConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewWithConfig indicates an expected call of NewWithConfig
func (mr *MockCloudInstancesMockRecorder) NewWithConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWithConfig", reflect.TypeOf((*MockCloudInstances)(nil).NewWithConfig), arg0)
}

// StartSession mocks base method
func (m *MockCloudInstances) StartSession(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartSession indicates an expected call of StartSession
func (mr *MockCloudInstancesMockRecorder) StartSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSession", reflect.TypeOf((*MockCloudInstances)(nil).StartSession), arg0, arg1)
}

// MockCloudSessions is a mock of CloudSessions interface
type MockCloudSessions struct {
	ctrl     *gomock.Controller
	recorder *MockCloudSessionsMockRecorder
}

// MockCloudSessionsMockRecorder is the mock recorder for MockCloudSessions
type MockCloudSessionsMockRecorder struct {
	mock *MockCloudSessions
}

// NewMockCloudSessions creates a new mock instance
func NewMockCloudSessions(ctrl *gomock.Controller) *MockCloudSessions {
	mock := &MockCloudSessions{ctrl: ctrl}
	mock.recorder = &MockCloudSessionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudSessions) EXPECT() *MockCloudSessionsMockRecorder {
	return m.recorder
}

// ListSessions mocks base method
func (m *MockCloudSessions) ListSessions() ([]*aws.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSessions")
	ret0, _ := ret[0].([]*aws.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSessions indicates an expected call of ListSessions
func (mr *MockCloudSessionsMockRecorder) ListSessions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSessions", reflect.TypeOf((*MockCloudSessions)(nil).ListSessions))
}

// NewWithConfig mocks base method
func (m *MockCloudSessions) NewWithConfig(arg0 *aws.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWithConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewWithConfig indicates an expected call of NewWithConfig
func (mr *MockCloudSessionsMockRecorder) NewWithConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWithConfig", reflect.TypeOf((*MockCloudSessions)(nil).NewWithConfig), arg0)
}

// TerminateSession mocks base method
func (m *MockCloudSessions) TerminateSession(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TerminateSession indicates an expected call of TerminateSession
func (mr *MockCloudSessionsMockRecorder) TerminateSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateSession", reflect.TypeOf((*MockCloudSessions)(nil).TerminateSession), arg0)
}

// MockCloudSSH is a mock of CloudSSH interface
type MockCloudSSH struct {
	ctrl     *gomock.Controller
	recorder *MockCloudSSHMockRecorder
}

// MockCloudSSHMockRecorder is the mock recorder for MockCloudSSH
type MockCloudSSHMockRecorder struct {
	mock *MockCloudSSH
}

// NewMockCloudSSH creates a new mock instance
func NewMockCloudSSH(ctrl *gomock.Controller) *MockCloudSSH {
	mock := &MockCloudSSH{ctrl: ctrl}
	mock.recorder = &MockCloudSSHMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudSSH) EXPECT() *MockCloudSSHMockRecorder {
	return m.recorder
}

// NewWithConfig mocks base method
func (m *MockCloudSSH) NewWithConfig(arg0 *aws.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWithConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewWithConfig indicates an expected call of NewWithConfig
func (mr *MockCloudSSHMockRecorder) NewWithConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWithConfig", reflect.TypeOf((*MockCloudSSH)(nil).NewWithConfig), arg0)
}

// StartSSH mocks base method
func (m *MockCloudSSH) StartSSH(arg0, arg1, arg2 string, arg3 uint64, arg4 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartSSH", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartSSH indicates an expected call of StartSSH
func (mr *MockCloudSSHMockRecorder) StartSSH(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartSSH", reflect.TypeOf((*MockCloudSSH)(nil).StartSSH), arg0, arg1, arg2, arg3, arg4)
}
