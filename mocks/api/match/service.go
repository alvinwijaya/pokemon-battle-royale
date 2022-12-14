// Code generated by MockGen. DO NOT EDIT.
// Source: api/match/service.go

// Package mock_match is a generated GoMock package.
package mock_match

import (
	reflect "reflect"

	filter "github.com/alvinwijaya/pokemon-battle-royale/api/match/filter"
	custom_error "github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	match "github.com/alvinwijaya/pokemon-battle-royale/model/match"
	pagination_model "github.com/alvinwijaya/pokemon-battle-royale/model/pagination_model"
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

// GetAllInPagination mocks base method.
func (m *MockService) GetAllInPagination(paging pagination_model.Paging, f *filter.GetAllInPaginationFilter) (*pagination_model.PaginationData, *custom_error.ServiceError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllInPagination", paging, f)
	ret0, _ := ret[0].(*pagination_model.PaginationData)
	ret1, _ := ret[1].(*custom_error.ServiceError)
	return ret0, ret1
}

// GetAllInPagination indicates an expected call of GetAllInPagination.
func (mr *MockServiceMockRecorder) GetAllInPagination(paging, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllInPagination", reflect.TypeOf((*MockService)(nil).GetAllInPagination), paging, f)
}

// Store mocks base method.
func (m *MockService) Store(p *match.StoreMatchDTO) (string, *custom_error.ServiceError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", p)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*custom_error.ServiceError)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockServiceMockRecorder) Store(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockService)(nil).Store), p)
}
