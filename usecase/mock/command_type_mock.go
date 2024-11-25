// Code generated by MockGen. DO NOT EDIT.
// Source: command_type.go
//
// Generated by this command:
//
//	mockgen -source=command_type.go -destination=mock/command_type_mock.go -package=mock
//
// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	usecase "github.com/sora-ichigo/github-project-automation/usecase"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectV2Setter is a mock of ProjectV2Setter interface.
type MockProjectV2Setter struct {
	ctrl     *gomock.Controller
	recorder *MockProjectV2SetterMockRecorder
}

// MockProjectV2SetterMockRecorder is the mock recorder for MockProjectV2Setter.
type MockProjectV2SetterMockRecorder struct {
	mock *MockProjectV2Setter
}

// NewMockProjectV2Setter creates a new mock instance.
func NewMockProjectV2Setter(ctrl *gomock.Controller) *MockProjectV2Setter {
	mock := &MockProjectV2Setter{ctrl: ctrl}
	mock.recorder = &MockProjectV2SetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectV2Setter) EXPECT() *MockProjectV2SetterMockRecorder {
	return m.recorder
}

// Set mocks base method.
func (m *MockProjectV2Setter) Set(categoryID, statusID string, projectItems []usecase.ProjectItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", categoryID, statusID, projectItems)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockProjectV2SetterMockRecorder) Set(categoryID, statusID, projectItems any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockProjectV2Setter)(nil).Set), categoryID, statusID, projectItems)
}
