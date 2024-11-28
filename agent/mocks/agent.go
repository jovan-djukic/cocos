// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	agent "github.com/ultravioletrs/cocos/agent"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// Algo provides a mock function with given fields: ctx, algorithm
func (_m *Service) Algo(ctx context.Context, algorithm agent.Algorithm) error {
	ret := _m.Called(ctx, algorithm)

	if len(ret) == 0 {
		panic("no return value specified for Algo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, agent.Algorithm) error); ok {
		r0 = rf(ctx, algorithm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_Algo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Algo'
type Service_Algo_Call struct {
	*mock.Call
}

// Algo is a helper method to define mock.On call
//   - ctx context.Context
//   - algorithm agent.Algorithm
func (_e *Service_Expecter) Algo(ctx interface{}, algorithm interface{}) *Service_Algo_Call {
	return &Service_Algo_Call{Call: _e.mock.On("Algo", ctx, algorithm)}
}

func (_c *Service_Algo_Call) Run(run func(ctx context.Context, algorithm agent.Algorithm)) *Service_Algo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(agent.Algorithm))
	})
	return _c
}

func (_c *Service_Algo_Call) Return(_a0 error) *Service_Algo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_Algo_Call) RunAndReturn(run func(context.Context, agent.Algorithm) error) *Service_Algo_Call {
	_c.Call.Return(run)
	return _c
}

// Attestation provides a mock function with given fields: ctx, reportData
func (_m *Service) Attestation(ctx context.Context, reportData [64]byte) ([]byte, error) {
	ret := _m.Called(ctx, reportData)

	if len(ret) == 0 {
		panic("no return value specified for Attestation")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [64]byte) ([]byte, error)); ok {
		return rf(ctx, reportData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [64]byte) []byte); ok {
		r0 = rf(ctx, reportData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [64]byte) error); ok {
		r1 = rf(ctx, reportData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_Attestation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Attestation'
type Service_Attestation_Call struct {
	*mock.Call
}

// Attestation is a helper method to define mock.On call
//   - ctx context.Context
//   - reportData [64]byte
func (_e *Service_Expecter) Attestation(ctx interface{}, reportData interface{}) *Service_Attestation_Call {
	return &Service_Attestation_Call{Call: _e.mock.On("Attestation", ctx, reportData)}
}

func (_c *Service_Attestation_Call) Run(run func(ctx context.Context, reportData [64]byte)) *Service_Attestation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([64]byte))
	})
	return _c
}

func (_c *Service_Attestation_Call) Return(_a0 []byte, _a1 error) *Service_Attestation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_Attestation_Call) RunAndReturn(run func(context.Context, [64]byte) ([]byte, error)) *Service_Attestation_Call {
	_c.Call.Return(run)
	return _c
}

// Data provides a mock function with given fields: ctx, dataset
func (_m *Service) Data(ctx context.Context, dataset agent.Dataset) error {
	ret := _m.Called(ctx, dataset)

	if len(ret) == 0 {
		panic("no return value specified for Data")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, agent.Dataset) error); ok {
		r0 = rf(ctx, dataset)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type Service_Data_Call struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
//   - ctx context.Context
//   - dataset agent.Dataset
func (_e *Service_Expecter) Data(ctx interface{}, dataset interface{}) *Service_Data_Call {
	return &Service_Data_Call{Call: _e.mock.On("Data", ctx, dataset)}
}

func (_c *Service_Data_Call) Run(run func(ctx context.Context, dataset agent.Dataset)) *Service_Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(agent.Dataset))
	})
	return _c
}

func (_c *Service_Data_Call) Return(_a0 error) *Service_Data_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_Data_Call) RunAndReturn(run func(context.Context, agent.Dataset) error) *Service_Data_Call {
	_c.Call.Return(run)
	return _c
}

// Result provides a mock function with given fields: ctx
func (_m *Service) Result(ctx context.Context) ([]byte, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Result")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]byte, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []byte); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_Result_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Result'
type Service_Result_Call struct {
	*mock.Call
}

// Result is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Service_Expecter) Result(ctx interface{}) *Service_Result_Call {
	return &Service_Result_Call{Call: _e.mock.On("Result", ctx)}
}

func (_c *Service_Result_Call) Run(run func(ctx context.Context)) *Service_Result_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Service_Result_Call) Return(_a0 []byte, _a1 error) *Service_Result_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_Result_Call) RunAndReturn(run func(context.Context) ([]byte, error)) *Service_Result_Call {
	_c.Call.Return(run)
	return _c
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
