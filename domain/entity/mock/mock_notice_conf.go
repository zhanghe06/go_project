// Code generated by MockGen. DO NOT EDIT.
// Source: ./notice_conf.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	vo "sap_cert_mgt/domain/vo"

	gomock "github.com/golang/mock/gomock"
)

// MockNoticeConfEntityInterface is a mock of NoticeConfEntityInterface interface.
type MockNoticeConfEntityInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNoticeConfEntityInterfaceMockRecorder
}

// MockNoticeConfEntityInterfaceMockRecorder is the mock recorder for MockNoticeConfEntityInterface.
type MockNoticeConfEntityInterfaceMockRecorder struct {
	mock *MockNoticeConfEntityInterface
}

// NewMockNoticeConfEntityInterface creates a new mock instance.
func NewMockNoticeConfEntityInterface(ctrl *gomock.Controller) *MockNoticeConfEntityInterface {
	mock := &MockNoticeConfEntityInterface{ctrl: ctrl}
	mock.recorder = &MockNoticeConfEntityInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNoticeConfEntityInterface) EXPECT() *MockNoticeConfEntityInterfaceMockRecorder {
	return m.recorder
}

// GetNoticeConfEmail mocks base method.
func (m *MockNoticeConfEntityInterface) GetNoticeConfEmail() (*vo.NoticeConfGetEmailRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNoticeConfEmail")
	ret0, _ := ret[0].(*vo.NoticeConfGetEmailRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNoticeConfEmail indicates an expected call of GetNoticeConfEmail.
func (mr *MockNoticeConfEntityInterfaceMockRecorder) GetNoticeConfEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNoticeConfEmail", reflect.TypeOf((*MockNoticeConfEntityInterface)(nil).GetNoticeConfEmail))
}

// ModNoticeConfEmail mocks base method.
func (m *MockNoticeConfEntityInterface) ModNoticeConfEmail(data map[string]interface{}, updatedBy string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModNoticeConfEmail", data, updatedBy)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModNoticeConfEmail indicates an expected call of ModNoticeConfEmail.
func (mr *MockNoticeConfEntityInterfaceMockRecorder) ModNoticeConfEmail(data, updatedBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModNoticeConfEmail", reflect.TypeOf((*MockNoticeConfEntityInterface)(nil).ModNoticeConfEmail), data, updatedBy)
}
