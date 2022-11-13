// Code generated by MockGen. DO NOT EDIT.
// Source: api/pokemon/service.go

// Package mock_pokemon is a generated GoMock package.
package mock_pokemon

import (
	reflect "reflect"

	custom_error "github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetCount mocks base method.
func (m *MockService) GetCount() (uint64, *custom_error.ServiceError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(*custom_error.ServiceError)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockServiceMockRecorder) GetCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockService)(nil).GetCount))
}
