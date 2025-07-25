// Code generated by MockGen. DO NOT EDIT.
// Source: user_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	entity "MuchUp/backend/internal/domain/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserUsecase) CreateUser(user *entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}


// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserUsecaseMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserUsecase)(nil).CreateUser), user)
}

// DeleteUser mocks base method.
func (m *MockUserUsecase) DeleteUser(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserUsecaseMockRecorder) DeleteUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserUsecase)(nil).DeleteUser), id)
}

// GetUserByID mocks base method.
func (m *MockUserUsecase) GetUserByID(id string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", id)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserUsecaseMockRecorder) GetUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserUsecase)(nil).GetUserByID), id)
}

// GetUsers mocks base method.
func (m *MockUserUsecase) GetUsers(limit, offset int) ([]*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", limit, offset)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserUsecaseMockRecorder) GetUsers(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserUsecase)(nil).GetUsers), limit, offset)
}

// GetUsersByGroup mocks base method.
func (m *MockUserUsecase) GetUsersByGroup(groupID string) ([]*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByGroup", groupID)
	ret0, _ := ret[0].([]*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByGroup indicates an expected call of GetUsersByGroup.
func (mr *MockUserUsecaseMockRecorder) GetUsersByGroup(groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByGroup", reflect.TypeOf((*MockUserUsecase)(nil).GetUsersByGroup), groupID)
}

// JoinGroup mocks base method.
func (m *MockUserUsecase) JoinGroup(userID, groupID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JoinGroup", userID, groupID)
	ret0, _ := ret[0].(error)
	return ret0
}

// JoinGroup indicates an expected call of JoinGroup.
func (mr *MockUserUsecaseMockRecorder) JoinGroup(userID, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JoinGroup", reflect.TypeOf((*MockUserUsecase)(nil).JoinGroup), userID, groupID)
}

// LeaveGroup mocks base method.
func (m *MockUserUsecase) LeaveGroup(userID, groupID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaveGroup", userID, groupID)
	ret0, _ := ret[0].(error)
	return ret0
}

// LeaveGroup indicates an expected call of LeaveGroup.
func (mr *MockUserUsecaseMockRecorder) LeaveGroup(userID, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveGroup", reflect.TypeOf((*MockUserUsecase)(nil).LeaveGroup), userID, groupID)
}

// Login mocks base method.
func (m *MockUserUsecase) Login(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserUsecaseMockRecorder) Login(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserUsecase)(nil).Login), email, password)
}

// UpdateUser mocks base method.
func (m *MockUserUsecase) UpdateUser(user *entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserUsecaseMockRecorder) UpdateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserUsecase)(nil).UpdateUser), user)
}
