// Code generated by MockGen. DO NOT EDIT.
// Source: db/userClient.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "github.com/amoriartyCH/go-sample/models/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserClient is a mock of UserClient interface
type MockUserClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserClientMockRecorder
}

// MockUserClientMockRecorder is the mock recorder for MockUserClient
type MockUserClientMockRecorder struct {
	mock *MockUserClient
}

// NewMockUserClient creates a new mock instance
func NewMockUserClient(ctrl *gomock.Controller) *MockUserClient {
	mock := &MockUserClient{ctrl: ctrl}
	mock.recorder = &MockUserClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserClient) EXPECT() *MockUserClientMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockUserClient) CreateUser(entity *entity.UserDao) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserClientMockRecorder) CreateUser(entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserClient)(nil).CreateUser), entity)
}

// GetUser mocks base method
func (m *MockUserClient) GetUser(id string) (*entity.UserDao, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(*entity.UserDao)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserClientMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserClient)(nil).GetUser), id)
}

// GetAllUsers mocks base method
func (m *MockUserClient) GetAllUsers() (*[]*entity.UserDao, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].(*[]*entity.UserDao)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers
func (mr *MockUserClientMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserClient)(nil).GetAllUsers))
}

// Shutdown mocks base method
func (m *MockUserClient) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockUserClientMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockUserClient)(nil).Shutdown))
}