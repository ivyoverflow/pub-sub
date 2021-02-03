// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	model "github.com/ivyoverflow/pub-sub/book/internal/model"
	reflect "reflect"
)

// MockBookerService is a mock of Booker interface
type MockBookerService struct {
	ctrl     *gomock.Controller
	recorder *MockBookerServiceMockRecorder
}

// MockBookerServiceMockRecorder is the mock recorder for MockBookerService
type MockBookerServiceMockRecorder struct {
	mock *MockBookerService
}

// NewMockBookerService creates a new mock instance
func NewMockBookerService(ctrl *gomock.Controller) *MockBookerService {
	mock := &MockBookerService{ctrl: ctrl}
	mock.recorder = &MockBookerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBookerService) EXPECT() *MockBookerServiceMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockBookerService) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, book)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockBookerServiceMockRecorder) Insert(ctx, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockBookerService)(nil).Insert), ctx, book)
}

// Get mocks base method
func (m *MockBookerService) Get(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, bookID)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockBookerServiceMockRecorder) Get(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBookerService)(nil).Get), ctx, bookID)
}

// Update mocks base method
func (m *MockBookerService) Update(ctx context.Context, bookID uuid.UUID, book *model.Book) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, bookID, book)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockBookerServiceMockRecorder) Update(ctx, bookID, book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookerService)(nil).Update), ctx, bookID, book)
}

// Delete mocks base method
func (m *MockBookerService) Delete(ctx context.Context, bookID uuid.UUID) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, bookID)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockBookerServiceMockRecorder) Delete(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBookerService)(nil).Delete), ctx, bookID)
}

// MockGeneratorService is a mock of Generator interface
type MockGeneratorService struct {
	ctrl     *gomock.Controller
	recorder *MockGeneratorServiceMockRecorder
}

// MockGeneratorServiceMockRecorder is the mock recorder for MockGeneratorService
type MockGeneratorServiceMockRecorder struct {
	mock *MockGeneratorService
}

// NewMockGeneratorService creates a new mock instance
func NewMockGeneratorService(ctrl *gomock.Controller) *MockGeneratorService {
	mock := &MockGeneratorService{ctrl: ctrl}
	mock.recorder = &MockGeneratorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGeneratorService) EXPECT() *MockGeneratorServiceMockRecorder {
	return m.recorder
}

// GenerateUUID mocks base method
func (m *MockGeneratorService) GenerateUUID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUUID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// GenerateUUID indicates an expected call of GenerateUUID
func (mr *MockGeneratorServiceMockRecorder) GenerateUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUUID", reflect.TypeOf((*MockGeneratorService)(nil).GenerateUUID))
}
