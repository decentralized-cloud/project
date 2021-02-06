// Code generated by MockGen. DO NOT EDIT.
// Source: ../contract.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	repository "github.com/decentralized-cloud/project/services/repository"
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

// CreateProject mocks base method
func (m *MockRepositoryContract) CreateProject(ctx context.Context, request *repository.CreateProjectRequest) (*repository.CreateProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", ctx, request)
	ret0, _ := ret[0].(*repository.CreateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject
func (mr *MockRepositoryContractMockRecorder) CreateProject(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockRepositoryContract)(nil).CreateProject), ctx, request)
}

// ReadProject mocks base method
func (m *MockRepositoryContract) ReadProject(ctx context.Context, request *repository.ReadProjectRequest) (*repository.ReadProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadProject", ctx, request)
	ret0, _ := ret[0].(*repository.ReadProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadProject indicates an expected call of ReadProject
func (mr *MockRepositoryContractMockRecorder) ReadProject(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadProject", reflect.TypeOf((*MockRepositoryContract)(nil).ReadProject), ctx, request)
}

// UpdateProject mocks base method
func (m *MockRepositoryContract) UpdateProject(ctx context.Context, request *repository.UpdateProjectRequest) (*repository.UpdateProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", ctx, request)
	ret0, _ := ret[0].(*repository.UpdateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject
func (mr *MockRepositoryContractMockRecorder) UpdateProject(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockRepositoryContract)(nil).UpdateProject), ctx, request)
}

// DeleteProject mocks base method
func (m *MockRepositoryContract) DeleteProject(ctx context.Context, request *repository.DeleteProjectRequest) (*repository.DeleteProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", ctx, request)
	ret0, _ := ret[0].(*repository.DeleteProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject
func (mr *MockRepositoryContractMockRecorder) DeleteProject(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockRepositoryContract)(nil).DeleteProject), ctx, request)
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
