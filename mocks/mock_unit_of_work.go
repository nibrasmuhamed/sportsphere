// Code generated by MockGen. DO NOT EDIT.
// Source: ./db/unit_of_work.go

// Package mock_db is a generated GoMock package.
package mocks

import (
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
        db "github.com/mohdjishin/sportsphere/db"
)

// MockUnitOfWork is a mock of UnitOfWork interface.
type MockUnitOfWork struct {
        ctrl     *gomock.Controller
        recorder *MockUnitOfWorkMockRecorder
}

// MockUnitOfWorkMockRecorder is the mock recorder for MockUnitOfWork.
type MockUnitOfWorkMockRecorder struct {
        mock *MockUnitOfWork
}

// NewMockUnitOfWork creates a new mock instance.
func NewMockUnitOfWork(ctrl *gomock.Controller) *MockUnitOfWork {
        mock := &MockUnitOfWork{ctrl: ctrl}
        mock.recorder = &MockUnitOfWorkMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnitOfWork) EXPECT() *MockUnitOfWorkMockRecorder {
        return m.recorder
}

// Abort mocks base method.
func (m *MockUnitOfWork) Abort() error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Abort")
        ret0, _ := ret[0].(error)
        return ret0
}

// Abort indicates an expected call of Abort.
func (mr *MockUnitOfWorkMockRecorder) Abort() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*MockUnitOfWork)(nil).Abort))
}

// Commit mocks base method.
func (m *MockUnitOfWork) Commit() error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Commit")
        ret0, _ := ret[0].(error)
        return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockUnitOfWorkMockRecorder) Commit() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockUnitOfWork)(nil).Commit))
}

// QueueCreate mocks base method.
func (m *MockUnitOfWork) QueueCreate(container db.DataContainer, object any) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "QueueCreate", container, object)
}

// QueueCreate indicates an expected call of QueueCreate.
func (mr *MockUnitOfWorkMockRecorder) QueueCreate(container, object interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueCreate", reflect.TypeOf((*MockUnitOfWork)(nil).QueueCreate), container, object)
}

// QueueCreateMany mocks base method.
func (m *MockUnitOfWork) QueueCreateMany(container db.DataContainer, object []any) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "QueueCreateMany", container, object)
}

// QueueCreateMany indicates an expected call of QueueCreateMany.
func (mr *MockUnitOfWorkMockRecorder) QueueCreateMany(container, object interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueCreateMany", reflect.TypeOf((*MockUnitOfWork)(nil).QueueCreateMany), container, object)
}

// QueueDelete mocks base method.
func (m *MockUnitOfWork) QueueDelete(container db.DataContainer, keyName string, keyValue any, operatorId string) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "QueueDelete", container, keyName, keyValue, operatorId)
}

// QueueDelete indicates an expected call of QueueDelete.
func (mr *MockUnitOfWorkMockRecorder) QueueDelete(container, keyName, keyValue, operatorId interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDelete", reflect.TypeOf((*MockUnitOfWork)(nil).QueueDelete), container, keyName, keyValue, operatorId)
}

// QueueUpdate mocks base method.
func (m *MockUnitOfWork) QueueUpdate(container db.DataContainer, keyName string, keyValue, object any, operatorId string) {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "QueueUpdate", container, keyName, keyValue, object, operatorId)
}

// QueueUpdate indicates an expected call of QueueUpdate.
func (mr *MockUnitOfWorkMockRecorder) QueueUpdate(container, keyName, keyValue, object, operatorId interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueUpdate", reflect.TypeOf((*MockUnitOfWork)(nil).QueueUpdate), container, keyName, keyValue, object, operatorId)
}