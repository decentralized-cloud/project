// Code generated by MockGen. DO NOT EDIT.
// Source: ../contract.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	repository "github.com/decentralized-cloud/tenant/services/repository"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepositoryContract is a mock of RepositoryContract interface
type MockRepositoryContract struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryContractMockRecorder
}

// MockRepositoryContractMockRecorder is the mock recorder for MockRepositoryContract
type MockRepositoryContractMockRecorder struct {
	mock *MockRepositoryContract
}

// NewMockRepositoryContract creates a new mock instance
func NewMockRepositoryContract(ctrl *gomock.Controller) *MockRepositoryContract {
	mock := &MockRepositoryContract{ctrl: ctrl}
	mock.recorder = &MockRepositoryContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepositoryContract) EXPECT() *MockRepositoryContractMockRecorder {
	return m.recorder
}

// CreateTenant mocks base method
func (m *MockRepositoryContract) CreateTenant(ctx context.Context, request *repository.CreateTenantRequest) (*repository.CreateTenantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTenant", ctx, request)
	ret0, _ := ret[0].(*repository.CreateTenantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTenant indicates an expected call of CreateTenant
func (mr *MockRepositoryContractMockRecorder) CreateTenant(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTenant", reflect.TypeOf((*MockRepositoryContract)(nil).CreateTenant), ctx, request)
}

// ReadTenant mocks base method
func (m *MockRepositoryContract) ReadTenant(ctx context.Context, request *repository.ReadTenantRequest) (*repository.ReadTenantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadTenant", ctx, request)
	ret0, _ := ret[0].(*repository.ReadTenantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadTenant indicates an expected call of ReadTenant
func (mr *MockRepositoryContractMockRecorder) ReadTenant(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadTenant", reflect.TypeOf((*MockRepositoryContract)(nil).ReadTenant), ctx, request)
}

// UpdateTenant mocks base method
func (m *MockRepositoryContract) UpdateTenant(ctx context.Context, request *repository.UpdateTenantRequest) (*repository.UpdateTenantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTenant", ctx, request)
	ret0, _ := ret[0].(*repository.UpdateTenantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTenant indicates an expected call of UpdateTenant
func (mr *MockRepositoryContractMockRecorder) UpdateTenant(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTenant", reflect.TypeOf((*MockRepositoryContract)(nil).UpdateTenant), ctx, request)
}

// DeleteTenant mocks base method
func (m *MockRepositoryContract) DeleteTenant(ctx context.Context, request *repository.DeleteTenantRequest) (*repository.DeleteTenantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTenant", ctx, request)
	ret0, _ := ret[0].(*repository.DeleteTenantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTenant indicates an expected call of DeleteTenant
func (mr *MockRepositoryContractMockRecorder) DeleteTenant(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTenant", reflect.TypeOf((*MockRepositoryContract)(nil).DeleteTenant), ctx, request)
}

// Search mocks base method
func (m *MockRepositoryContract) Search(ctx context.Context, request *repository.SearchRequest) (*repository.SearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, request)
	ret0, _ := ret[0].(*repository.SearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryContractMockRecorder) Search(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepositoryContract)(nil).Search), ctx, request)
}
