// Code generated by MockGen. DO NOT EDIT.
// Source: ./store/store.go

// Package mockstore is a generated GoMock package.
package mockstore

import (
	types "ToDoApp/types"
	reflect "reflect"
	gomock "github.com/golang/mock/gomock"
	
)

// MockdbInterface is a mock of dbInterface interface.
type MockdbInterface struct {
	ctrl     *gomock.Controller
	recorder *MockdbInterfaceMockRecorder
}

// MockdbInterfaceMockRecorder is the mock recorder for MockdbInterface.
type MockdbInterfaceMockRecorder struct {
	mock *MockdbInterface
}

// NewMockdbInterface creates a new mock instance.
func NewMockdbInterface(ctrl *gomock.Controller) *MockdbInterface {
	mock := &MockdbInterface{ctrl: ctrl}
	mock.recorder = &MockdbInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdbInterface) EXPECT() *MockdbInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockdbInterface) Delete(s *types.Todo, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", s, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockdbInterfaceMockRecorder) Delete(s, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockdbInterface)(nil).Delete), s, id)
}

// GetAllTodo mocks base method.
func (m *MockdbInterface) GetAllTodo() []types.Todo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTodo")
	ret0, _ := ret[0].([]types.Todo)
	return ret0
}

// GetAllTodo indicates an expected call of GetAllTodo.
func (mr *MockdbInterfaceMockRecorder) GetAllTodo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTodo", reflect.TypeOf((*MockdbInterface)(nil).GetAllTodo))
}

// GetOne mocks base method.
func (m *MockdbInterface) GetOne(s *types.Todo, id int) types.Todo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", s, id)
	ret0, _ := ret[0].(types.Todo)
	return ret0
}

// GetOne indicates an expected call of GetOne.
func (mr *MockdbInterfaceMockRecorder) GetOne(s, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockdbInterface)(nil).GetOne), s, id)
}

// InsertNewTodo mocks base method.
func (m *MockdbInterface) InsertNewTodo(s *types.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewTodo", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertNewTodo indicates an expected call of InsertNewTodo.
func (mr *MockdbInterfaceMockRecorder) InsertNewTodo(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewTodo", reflect.TypeOf((*MockdbInterface)(nil).InsertNewTodo), s)
}
