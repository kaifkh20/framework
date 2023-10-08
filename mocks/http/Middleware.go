// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	http "github.com/goravel/framework/contracts/http"
	mock "github.com/stretchr/testify/mock"
)

// Middleware is an autogenerated mock type for the Middleware type
type Middleware struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *Middleware) Execute(_a0 http.Context) {
	_m.Called(_a0)
}

// NewMiddleware creates a new instance of Middleware. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMiddleware(t interface {
	mock.TestingT
	Cleanup(func())
}) *Middleware {
	mock := &Middleware{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}