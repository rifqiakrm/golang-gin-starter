// Code generated by MockGen. DO NOT EDIT.
// Source: ./modules/cms/repository/user.repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"

	entity "gin-starter/entity"
)

// MockUserRepositoryUseCase is a mock of UserRepositoryUseCase interface.
type MockUserRepositoryUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryUseCaseMockRecorder
}

// MockUserRepositoryUseCaseMockRecorder is the mock recorder for MockUserRepositoryUseCase.
type MockUserRepositoryUseCaseMockRecorder struct {
	mock *MockUserRepositoryUseCase
}

// NewMockUserRepositoryUseCase creates a new mock instance.
func NewMockUserRepositoryUseCase(ctrl *gomock.Controller) *MockUserRepositoryUseCase {
	mock := &MockUserRepositoryUseCase{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryUseCase) EXPECT() *MockUserRepositoryUseCaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepositoryUseCase) CreateUser(ctx context.Context, user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryUseCaseMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepositoryUseCase)(nil).CreateUser), ctx, user)
}

// DeleteAdmin mocks base method.
func (m *MockUserRepositoryUseCase) DeleteAdmin(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdmin", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAdmin indicates an expected call of DeleteAdmin.
func (mr *MockUserRepositoryUseCaseMockRecorder) DeleteAdmin(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdmin", reflect.TypeOf((*MockUserRepositoryUseCase)(nil).DeleteAdmin), ctx, id)
}

// GetAdminUsers mocks base method.
func (m *MockUserRepositoryUseCase) GetAdminUsers(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.User, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdminUsers", ctx, query, sort, order, limit, offset)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAdminUsers indicates an expected call of GetAdminUsers.
func (mr *MockUserRepositoryUseCaseMockRecorder) GetAdminUsers(ctx, query, sort, order, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdminUsers", reflect.TypeOf((*MockUserRepositoryUseCase)(nil).GetAdminUsers), ctx, query, sort, order, limit, offset)
}

// GetUserByID mocks base method.
func (m *MockUserRepositoryUseCase) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepositoryUseCaseMockRecorder) GetUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepositoryUseCase)(nil).GetUserByID), ctx, id)
}

// GetUsers mocks base method.
func (m *MockUserRepositoryUseCase) GetUsers(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.User, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, query, sort, order, limit, offset)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserRepositoryUseCaseMockRecorder) GetUsers(ctx, query, sort, order, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserRepositoryUseCase)(nil).GetUsers), ctx, query, sort, order, limit, offset)
}

// UpdateUser mocks base method.
func (m *MockUserRepositoryUseCase) UpdateUser(ctx context.Context, user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryUseCaseMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepositoryUseCase)(nil).UpdateUser), ctx, user)
}

// UpdateUserStatus mocks base method.
func (m *MockUserRepositoryUseCase) UpdateUserStatus(ctx context.Context, id uuid.UUID, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserStatus", ctx, id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserStatus indicates an expected call of UpdateUserStatus.
func (mr *MockUserRepositoryUseCaseMockRecorder) UpdateUserStatus(ctx, id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserStatus", reflect.TypeOf((*MockUserRepositoryUseCase)(nil).UpdateUserStatus), ctx, id, status)
}
