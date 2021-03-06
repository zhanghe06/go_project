// Code generated by MockGen. DO NOT EDIT.
// Source: ./user.go

// Package mock is a generated GoMock package.
package mock

import (
	model "go_project/infra/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepoInterface is a mock of UserRepoInterface interface.
type MockUserRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoInterfaceMockRecorder
}

// MockUserRepoInterfaceMockRecorder is the mock recorder for MockUserRepoInterface.
type MockUserRepoInterfaceMockRecorder struct {
	mock *MockUserRepoInterface
}

// NewMockUserRepoInterface creates a new mock instance.
func NewMockUserRepoInterface(ctrl *gomock.Controller) *MockUserRepoInterface {
	mock := &MockUserRepoInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepoInterface) EXPECT() *MockUserRepoInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepoInterface) Create(data *model.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", data)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserRepoInterfaceMockRecorder) Create(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepoInterface)(nil).Create), data)
}

// Delete mocks base method.
func (m *MockUserRepoInterface) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserRepoInterfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserRepoInterface)(nil).Delete), id)
}

// GetInfo mocks base method.
func (m *MockUserRepoInterface) GetInfo(id int) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockUserRepoInterfaceMockRecorder) GetInfo(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockUserRepoInterface)(nil).GetInfo), id)
}

// GetList mocks base method.
func (m *MockUserRepoInterface) GetList(filter map[string]interface{}) (int64, []*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", filter)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].([]*model.User)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetList indicates an expected call of GetList.
func (mr *MockUserRepoInterfaceMockRecorder) GetList(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockUserRepoInterface)(nil).GetList), filter)
}

// Update mocks base method.
func (m *MockUserRepoInterface) Update(id int, data map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepoInterfaceMockRecorder) Update(id, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepoInterface)(nil).Update), id, data)
}
