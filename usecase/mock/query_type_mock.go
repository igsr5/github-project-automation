// Code generated by MockGen. DO NOT EDIT.
// Source: query_type.go
//
// Generated by this command:
//
//	mockgen -source=query_type.go -destination=mock/query_type_mock.go -package=mock
//
// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	usecase "github.com/igsr5/github-project-automation/usecase"
	gomock "go.uber.org/mock/gomock"
)

// MockIssueFetcher is a mock of IssueFetcher interface.
type MockIssueFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockIssueFetcherMockRecorder
}

// MockIssueFetcherMockRecorder is the mock recorder for MockIssueFetcher.
type MockIssueFetcherMockRecorder struct {
	mock *MockIssueFetcher
}

// NewMockIssueFetcher creates a new mock instance.
func NewMockIssueFetcher(ctrl *gomock.Controller) *MockIssueFetcher {
	mock := &MockIssueFetcher{ctrl: ctrl}
	mock.recorder = &MockIssueFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssueFetcher) EXPECT() *MockIssueFetcherMockRecorder {
	return m.recorder
}

// MyIssues mocks base method.
func (m *MockIssueFetcher) MyIssues() ([]usecase.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MyIssues")
	ret0, _ := ret[0].([]usecase.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MyIssues indicates an expected call of MyIssues.
func (mr *MockIssueFetcherMockRecorder) MyIssues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MyIssues", reflect.TypeOf((*MockIssueFetcher)(nil).MyIssues))
}

// MockPrFetcher is a mock of PrFetcher interface.
type MockPrFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockPrFetcherMockRecorder
}

// MockPrFetcherMockRecorder is the mock recorder for MockPrFetcher.
type MockPrFetcherMockRecorder struct {
	mock *MockPrFetcher
}

// NewMockPrFetcher creates a new mock instance.
func NewMockPrFetcher(ctrl *gomock.Controller) *MockPrFetcher {
	mock := &MockPrFetcher{ctrl: ctrl}
	mock.recorder = &MockPrFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrFetcher) EXPECT() *MockPrFetcherMockRecorder {
	return m.recorder
}

// ApprovedPrs mocks base method.
func (m *MockPrFetcher) ApprovedPrs() ([]usecase.PullRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApprovedPrs")
	ret0, _ := ret[0].([]usecase.PullRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApprovedPrs indicates an expected call of ApprovedPrs.
func (mr *MockPrFetcherMockRecorder) ApprovedPrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApprovedPrs", reflect.TypeOf((*MockPrFetcher)(nil).ApprovedPrs))
}

// ChangesRequestedPrs mocks base method.
func (m *MockPrFetcher) ChangesRequestedPrs() ([]usecase.PullRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangesRequestedPrs")
	ret0, _ := ret[0].([]usecase.PullRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangesRequestedPrs indicates an expected call of ChangesRequestedPrs.
func (mr *MockPrFetcherMockRecorder) ChangesRequestedPrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangesRequestedPrs", reflect.TypeOf((*MockPrFetcher)(nil).ChangesRequestedPrs))
}

// CommentedPrs mocks base method.
func (m *MockPrFetcher) CommentedPrs() ([]usecase.PullRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommentedPrs")
	ret0, _ := ret[0].([]usecase.PullRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommentedPrs indicates an expected call of CommentedPrs.
func (mr *MockPrFetcherMockRecorder) CommentedPrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommentedPrs", reflect.TypeOf((*MockPrFetcher)(nil).CommentedPrs))
}

// UnReviewedPrs mocks base method.
func (m *MockPrFetcher) UnReviewedPrs() ([]usecase.PullRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnReviewedPrs")
	ret0, _ := ret[0].([]usecase.PullRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnReviewedPrs indicates an expected call of UnReviewedPrs.
func (mr *MockPrFetcherMockRecorder) UnReviewedPrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnReviewedPrs", reflect.TypeOf((*MockPrFetcher)(nil).UnReviewedPrs))
}

// MockReviewPrFetcher is a mock of ReviewPrFetcher interface.
type MockReviewPrFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockReviewPrFetcherMockRecorder
}

// MockReviewPrFetcherMockRecorder is the mock recorder for MockReviewPrFetcher.
type MockReviewPrFetcherMockRecorder struct {
	mock *MockReviewPrFetcher
}

// NewMockReviewPrFetcher creates a new mock instance.
func NewMockReviewPrFetcher(ctrl *gomock.Controller) *MockReviewPrFetcher {
	mock := &MockReviewPrFetcher{ctrl: ctrl}
	mock.recorder = &MockReviewPrFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewPrFetcher) EXPECT() *MockReviewPrFetcherMockRecorder {
	return m.recorder
}

// ApprovedPrs mocks base method.
func (m *MockReviewPrFetcher) ApprovedPrs() ([]usecase.PullRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApprovedPrs")
	ret0, _ := ret[0].([]usecase.PullRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApprovedPrs indicates an expected call of ApprovedPrs.
func (mr *MockReviewPrFetcherMockRecorder) ApprovedPrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApprovedPrs", reflect.TypeOf((*MockReviewPrFetcher)(nil).ApprovedPrs))
}

// ChangesRequestedPrs mocks base method.
func (m *MockReviewPrFetcher) ChangesRequestedPrs() ([]usecase.PullRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangesRequestedPrs")
	ret0, _ := ret[0].([]usecase.PullRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangesRequestedPrs indicates an expected call of ChangesRequestedPrs.
func (mr *MockReviewPrFetcherMockRecorder) ChangesRequestedPrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangesRequestedPrs", reflect.TypeOf((*MockReviewPrFetcher)(nil).ChangesRequestedPrs))
}

// CommentedPrs mocks base method.
func (m *MockReviewPrFetcher) CommentedPrs() ([]usecase.PullRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommentedPrs")
	ret0, _ := ret[0].([]usecase.PullRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommentedPrs indicates an expected call of CommentedPrs.
func (mr *MockReviewPrFetcherMockRecorder) CommentedPrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommentedPrs", reflect.TypeOf((*MockReviewPrFetcher)(nil).CommentedPrs))
}

// UnReviewedPrs mocks base method.
func (m *MockReviewPrFetcher) UnReviewedPrs() ([]usecase.PullRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnReviewedPrs")
	ret0, _ := ret[0].([]usecase.PullRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnReviewedPrs indicates an expected call of UnReviewedPrs.
func (mr *MockReviewPrFetcherMockRecorder) UnReviewedPrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnReviewedPrs", reflect.TypeOf((*MockReviewPrFetcher)(nil).UnReviewedPrs))
}
