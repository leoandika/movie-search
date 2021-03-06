// Code generated by MockGen. DO NOT EDIT.
// Source: routes.go

// Package rest is a generated GoMock package.
package rest

import (
	gomock "github.com/golang/mock/gomock"
	mux "github.com/gorilla/mux"
	http "net/http"
	reflect "reflect"
)

// MockIRoutes is a mock of IRoutes interface.
type MockIRoutes struct {
	ctrl     *gomock.Controller
	recorder *MockIRoutesMockRecorder
}

// MockIRoutesMockRecorder is the mock recorder for MockIRoutes.
type MockIRoutesMockRecorder struct {
	mock *MockIRoutes
}

// NewMockIRoutes creates a new mock instance.
func NewMockIRoutes(ctrl *gomock.Controller) *MockIRoutes {
	mock := &MockIRoutes{ctrl: ctrl}
	mock.recorder = &MockIRoutesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRoutes) EXPECT() *MockIRoutesMockRecorder {
	return m.recorder
}

// InitRoutes mocks base method.
func (m *MockIRoutes) InitRoutes(router *mux.Router) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitRoutes", router)
}

// InitRoutes indicates an expected call of InitRoutes.
func (mr *MockIRoutesMockRecorder) InitRoutes(router interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitRoutes", reflect.TypeOf((*MockIRoutes)(nil).InitRoutes), router)
}

// SearchMovie mocks base method.
func (m *MockIRoutes) SearchMovie(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SearchMovie", w, r)
}

// SearchMovie indicates an expected call of SearchMovie.
func (mr *MockIRoutesMockRecorder) SearchMovie(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovie", reflect.TypeOf((*MockIRoutes)(nil).SearchMovie), w, r)
}
