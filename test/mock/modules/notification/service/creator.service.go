// Code generated by MockGen. DO NOT EDIT.
// Source: ./modules/notification/service/creator.service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNotificationCreatorUseCase is a mock of NotificationCreatorUseCase interface.
type MockNotificationCreatorUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationCreatorUseCaseMockRecorder
}

// MockNotificationCreatorUseCaseMockRecorder is the mock recorder for MockNotificationCreatorUseCase.
type MockNotificationCreatorUseCaseMockRecorder struct {
	mock *MockNotificationCreatorUseCase
}

// NewMockNotificationCreatorUseCase creates a new mock instance.
func NewMockNotificationCreatorUseCase(ctrl *gomock.Controller) *MockNotificationCreatorUseCase {
	mock := &MockNotificationCreatorUseCase{ctrl: ctrl}
	mock.recorder = &MockNotificationCreatorUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationCreatorUseCase) EXPECT() *MockNotificationCreatorUseCaseMockRecorder {
	return m.recorder
}

// InsertNotification mocks base method.
func (m *MockNotificationCreatorUseCase) InsertNotification(ctx context.Context, userID, title, message, notifType, extra string, isRead bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNotification", ctx, userID, title, message, notifType, extra, isRead)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertNotification indicates an expected call of InsertNotification.
func (mr *MockNotificationCreatorUseCaseMockRecorder) InsertNotification(ctx, userID, title, message, notifType, extra, isRead interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNotification", reflect.TypeOf((*MockNotificationCreatorUseCase)(nil).InsertNotification), ctx, userID, title, message, notifType, extra, isRead)
}
