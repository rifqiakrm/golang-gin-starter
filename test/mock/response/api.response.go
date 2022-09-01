// Code generated by MockGen. DO NOT EDIT.
// Source: ./response/api.response.go

// Package mock_response is a generated GoMock package.
package mock_response

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockApiResponseList is a mock of ApiResponseList interface.
type MockApiResponseList struct {
	ctrl     *gomock.Controller
	recorder *MockApiResponseListMockRecorder
}

// MockApiResponseListMockRecorder is the mock recorder for MockApiResponseList.
type MockApiResponseListMockRecorder struct {
	mock *MockApiResponseList
}

// NewMockApiResponseList creates a new mock instance.
func NewMockApiResponseList(ctrl *gomock.Controller) *MockApiResponseList {
	mock := &MockApiResponseList{ctrl: ctrl}
	mock.recorder = &MockApiResponseListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApiResponseList) EXPECT() *MockApiResponseListMockRecorder {
	return m.recorder
}

// GetCode mocks base method.
func (m *MockApiResponseList) GetCode() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCode")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetCode indicates an expected call of GetCode.
func (mr *MockApiResponseListMockRecorder) GetCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCode", reflect.TypeOf((*MockApiResponseList)(nil).GetCode))
}

// GetData mocks base method.
func (m *MockApiResponseList) GetData() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetData indicates an expected call of GetData.
func (mr *MockApiResponseListMockRecorder) GetData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockApiResponseList)(nil).GetData))
}

// GetMessage mocks base method.
func (m *MockApiResponseList) GetMessage() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMessage indicates an expected call of GetMessage.
func (mr *MockApiResponseListMockRecorder) GetMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockApiResponseList)(nil).GetMessage))
}
