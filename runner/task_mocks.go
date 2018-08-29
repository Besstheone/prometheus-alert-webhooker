// Code generated by MockGen. DO NOT EDIT.
// Source: task.go

// Package runner is a generated GoMock package.
package runner

import (
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
	time "time"
)

// MockTask is a mock of Task interface
type MockTask struct {
	ctrl     *gomock.Controller
	recorder *MockTaskMockRecorder
}

// MockTaskMockRecorder is the mock recorder for MockTask
type MockTaskMockRecorder struct {
	mock *MockTask
}

// NewMockTask creates a new mock instance
func NewMockTask(ctrl *gomock.Controller) *MockTask {
	mock := &MockTask{ctrl: ctrl}
	mock.recorder = &MockTaskMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTask) EXPECT() *MockTaskMockRecorder {
	return m.recorder
}

// EventID mocks base method
func (m *MockTask) EventID() string {
	ret := m.ctrl.Call(m, "EventID")
	ret0, _ := ret[0].(string)
	return ret0
}

// EventID indicates an expected call of EventID
func (mr *MockTaskMockRecorder) EventID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventID", reflect.TypeOf((*MockTask)(nil).EventID))
}

// Rule mocks base method
func (m *MockTask) Rule() string {
	ret := m.ctrl.Call(m, "Rule")
	ret0, _ := ret[0].(string)
	return ret0
}

// Rule indicates an expected call of Rule
func (mr *MockTaskMockRecorder) Rule() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rule", reflect.TypeOf((*MockTask)(nil).Rule))
}

// Alert mocks base method
func (m *MockTask) Alert() string {
	ret := m.ctrl.Call(m, "Alert")
	ret0, _ := ret[0].(string)
	return ret0
}

// Alert indicates an expected call of Alert
func (mr *MockTaskMockRecorder) Alert() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alert", reflect.TypeOf((*MockTask)(nil).Alert))
}

// ExecutorName mocks base method
func (m *MockTask) ExecutorName() string {
	ret := m.ctrl.Call(m, "ExecutorName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ExecutorName indicates an expected call of ExecutorName
func (mr *MockTaskMockRecorder) ExecutorName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecutorName", reflect.TypeOf((*MockTask)(nil).ExecutorName))
}

// ExecutorDetails mocks base method
func (m *MockTask) ExecutorDetails() interface{} {
	ret := m.ctrl.Call(m, "ExecutorDetails")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// ExecutorDetails indicates an expected call of ExecutorDetails
func (mr *MockTaskMockRecorder) ExecutorDetails() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecutorDetails", reflect.TypeOf((*MockTask)(nil).ExecutorDetails))
}

// Fingerprint mocks base method
func (m *MockTask) Fingerprint() string {
	ret := m.ctrl.Call(m, "Fingerprint")
	ret0, _ := ret[0].(string)
	return ret0
}

// Fingerprint indicates an expected call of Fingerprint
func (mr *MockTaskMockRecorder) Fingerprint() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fingerprint", reflect.TypeOf((*MockTask)(nil).Fingerprint))
}

// Exec mocks base method
func (m *MockTask) Exec(logger *logrus.Logger) error {
	ret := m.ctrl.Call(m, "Exec", logger)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec
func (mr *MockTaskMockRecorder) Exec(logger interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockTask)(nil).Exec), logger)
}

// BlockTTL mocks base method
func (m *MockTask) BlockTTL() time.Duration {
	ret := m.ctrl.Call(m, "BlockTTL")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// BlockTTL indicates an expected call of BlockTTL
func (mr *MockTaskMockRecorder) BlockTTL() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockTTL", reflect.TypeOf((*MockTask)(nil).BlockTTL))
}
