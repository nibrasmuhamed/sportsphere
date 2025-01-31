// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/operator_repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/repository/operator_repository.go -destination=./mocks/mock_operator_repository.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	db "github.com/nibrasmuhamed/sportsphere/db"
	models "github.com/nibrasmuhamed/sportsphere/pkg/model"
	gomock "go.uber.org/mock/gomock"
)

// MockOperatorRepository is a mock of OperatorRepository interface.
type MockOperatorRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorRepositoryMockRecorder
	isgomock struct{}
}

// MockOperatorRepositoryMockRecorder is the mock recorder for MockOperatorRepository.
type MockOperatorRepositoryMockRecorder struct {
	mock *MockOperatorRepository
}

// NewMockOperatorRepository creates a new mock instance.
func NewMockOperatorRepository(ctrl *gomock.Controller) *MockOperatorRepository {
	mock := &MockOperatorRepository{ctrl: ctrl}
	mock.recorder = &MockOperatorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorRepository) EXPECT() *MockOperatorRepositoryMockRecorder {
	return m.recorder
}

// CreateOperator mocks base method.
func (m *MockOperatorRepository) CreateOperator(arg0 context.Context, arg1 models.Operator, arg2 db.UnitOfWork) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOperator", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOperator indicates an expected call of CreateOperator.
func (mr *MockOperatorRepositoryMockRecorder) CreateOperator(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOperator", reflect.TypeOf((*MockOperatorRepository)(nil).CreateOperator), arg0, arg1, arg2)
}

// GetOperatorByName mocks base method.
func (m *MockOperatorRepository) GetOperatorByName(arg0 string) (models.Operator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorByName", arg0)
	ret0, _ := ret[0].(models.Operator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorByName indicates an expected call of GetOperatorByName.
func (mr *MockOperatorRepositoryMockRecorder) GetOperatorByName(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorByName", reflect.TypeOf((*MockOperatorRepository)(nil).GetOperatorByName), arg0)
}
