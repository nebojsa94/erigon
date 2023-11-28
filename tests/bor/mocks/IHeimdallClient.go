// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nebojsa94/erigon/consensus/bor (interfaces: IHeimdallClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clerk "github.com/nebojsa94/erigon/consensus/bor/clerk"
	checkpoint "github.com/nebojsa94/erigon/consensus/bor/heimdall/checkpoint"
	span "github.com/nebojsa94/erigon/consensus/bor/heimdall/span"
)

// MockIHeimdallClient is a mock of IHeimdallClient interface.
type MockIHeimdallClient struct {
	ctrl     *gomock.Controller
	recorder *MockIHeimdallClientMockRecorder
}

// MockIHeimdallClientMockRecorder is the mock recorder for MockIHeimdallClient.
type MockIHeimdallClientMockRecorder struct {
	mock *MockIHeimdallClient
}

// NewMockIHeimdallClient creates a new mock instance.
func NewMockIHeimdallClient(ctrl *gomock.Controller) *MockIHeimdallClient {
	mock := &MockIHeimdallClient{ctrl: ctrl}
	mock.recorder = &MockIHeimdallClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHeimdallClient) EXPECT() *MockIHeimdallClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIHeimdallClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockIHeimdallClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIHeimdallClient)(nil).Close))
}

// FetchCheckpoint mocks base method.
func (m *MockIHeimdallClient) FetchCheckpoint(arg0 context.Context, arg1 int64) (*checkpoint.Checkpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCheckpoint", arg0, arg1)
	ret0, _ := ret[0].(*checkpoint.Checkpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCheckpoint indicates an expected call of FetchCheckpoint.
func (mr *MockIHeimdallClientMockRecorder) FetchCheckpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCheckpoint", reflect.TypeOf((*MockIHeimdallClient)(nil).FetchCheckpoint), arg0, arg1)
}

// FetchCheckpointCount mocks base method.
func (m *MockIHeimdallClient) FetchCheckpointCount(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCheckpointCount", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCheckpointCount indicates an expected call of FetchCheckpointCount.
func (mr *MockIHeimdallClientMockRecorder) FetchCheckpointCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCheckpointCount", reflect.TypeOf((*MockIHeimdallClient)(nil).FetchCheckpointCount), arg0)
}

// Span mocks base method.
func (m *MockIHeimdallClient) Span(arg0 context.Context, arg1 uint64) (*span.HeimdallSpan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Span", arg0, arg1)
	ret0, _ := ret[0].(*span.HeimdallSpan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Span indicates an expected call of Span.
func (mr *MockIHeimdallClientMockRecorder) Span(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Span", reflect.TypeOf((*MockIHeimdallClient)(nil).Span), arg0, arg1)
}

// StateSyncEvents mocks base method.
func (m *MockIHeimdallClient) StateSyncEvents(arg0 context.Context, arg1 uint64, arg2 int64) ([]*clerk.EventRecordWithTime, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSyncEvents", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*clerk.EventRecordWithTime)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSyncEvents indicates an expected call of StateSyncEvents.
func (mr *MockIHeimdallClientMockRecorder) StateSyncEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSyncEvents", reflect.TypeOf((*MockIHeimdallClient)(nil).StateSyncEvents), arg0, arg1, arg2)
}
