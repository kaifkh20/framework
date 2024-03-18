// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	schema "github.com/goravel/framework/contracts/database/schema"
	mock "github.com/stretchr/testify/mock"
)

// Processor is an autogenerated mock type for the Processor type
type Processor struct {
	mock.Mock
}

type Processor_Expecter struct {
	mock *mock.Mock
}

func (_m *Processor) EXPECT() *Processor_Expecter {
	return &Processor_Expecter{mock: &_m.Mock}
}

// ProcessColumns provides a mock function with given fields: columns
func (_m *Processor) ProcessColumns(columns []schema.Column) []schema.Column {
	ret := _m.Called(columns)

	if len(ret) == 0 {
		panic("no return value specified for ProcessColumns")
	}

	var r0 []schema.Column
	if rf, ok := ret.Get(0).(func([]schema.Column) []schema.Column); ok {
		r0 = rf(columns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]schema.Column)
		}
	}

	return r0
}

// Processor_ProcessColumns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessColumns'
type Processor_ProcessColumns_Call struct {
	*mock.Call
}

// ProcessColumns is a helper method to define mock.On call
//   - columns []schema.Column
func (_e *Processor_Expecter) ProcessColumns(columns interface{}) *Processor_ProcessColumns_Call {
	return &Processor_ProcessColumns_Call{Call: _e.mock.On("ProcessColumns", columns)}
}

func (_c *Processor_ProcessColumns_Call) Run(run func(columns []schema.Column)) *Processor_ProcessColumns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]schema.Column))
	})
	return _c
}

func (_c *Processor_ProcessColumns_Call) Return(_a0 []schema.Column) *Processor_ProcessColumns_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Processor_ProcessColumns_Call) RunAndReturn(run func([]schema.Column) []schema.Column) *Processor_ProcessColumns_Call {
	_c.Call.Return(run)
	return _c
}

// NewProcessor creates a new instance of Processor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProcessor(t interface {
	mock.TestingT
	Cleanup(func())
}) *Processor {
	mock := &Processor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}