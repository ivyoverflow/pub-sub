// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"

	model "github.com/ivyoverflow/pub-sub/api/internal/model"
)

// MockBookerRepository is a mock of Booker interface
type MockBookerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookerRepositoryMockRecorder
}

// MockBookerRepositoryMockRecorder is the mock recorder for MockBookerRepository
type MockBookerRepositoryMockRecorder struct {
	mock *MockBookerRepository
}

// NewMockBookerRepository creates a new mock instance
func NewMockBookerRepository(ctrl *gomock.Controller) *MockBookerRepository {
	mock := &MockBookerRepository{ctrl: ctrl}
	mock.recorder = &MockBookerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookerRepository) EXPECT() *MockBookerRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockBookerRepository) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, book)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockBookerRepositoryMockRecorder) Insert(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockBookerRepository)(nil).Insert), ctx, book)
}

// Get mocks base method
func (m *MockBookerRepository) Get(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, bookID)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockBookerRepositoryMockRecorder) Get(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBookerRepository)(nil).Get), ctx, bookID)
}

// Update mocks base method
func (m *MockBookerRepository) Update(ctx context.Context, bookID uuid.UUID, book *model.Book) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, bookID, book)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockBookerRepositoryMockRecorder) Update(ctx, bookID, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookerRepository)(nil).Update), ctx, bookID, book)
}

// Delete mocks base method
func (m *MockBookerRepository) Delete(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, bookID)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockBookerRepositoryMockRecorder) Delete(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBookerRepository)(nil).Delete), ctx, bookID)
}