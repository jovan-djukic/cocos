// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	agent "github.com/ultravioletrs/cocos/agent"
)

// AgentServer is an autogenerated mock type for the AgentServer type
type AgentServer struct {
	mock.Mock
}

type AgentServer_Expecter struct {
	mock *mock.Mock
}

func (_m *AgentServer) EXPECT() *AgentServer_Expecter {
	return &AgentServer_Expecter{mock: &_m.Mock}
}

// Start provides a mock function with given fields: cfg, cmp
func (_m *AgentServer) Start(cfg agent.AgentConfig, cmp agent.Computation) error {
	ret := _m.Called(cfg, cmp)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(agent.AgentConfig, agent.Computation) error); ok {
		r0 = rf(cfg, cmp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AgentServer_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type AgentServer_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - cfg agent.AgentConfig
//   - cmp agent.Computation
func (_e *AgentServer_Expecter) Start(cfg interface{}, cmp interface{}) *AgentServer_Start_Call {
	return &AgentServer_Start_Call{Call: _e.mock.On("Start", cfg, cmp)}
}

func (_c *AgentServer_Start_Call) Run(run func(cfg agent.AgentConfig, cmp agent.Computation)) *AgentServer_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(agent.AgentConfig), args[1].(agent.Computation))
	})
	return _c
}

func (_c *AgentServer_Start_Call) Return(_a0 error) *AgentServer_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentServer_Start_Call) RunAndReturn(run func(agent.AgentConfig, agent.Computation) error) *AgentServer_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with no fields
func (_m *AgentServer) Stop() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AgentServer_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type AgentServer_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *AgentServer_Expecter) Stop() *AgentServer_Stop_Call {
	return &AgentServer_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *AgentServer_Stop_Call) Run(run func()) *AgentServer_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AgentServer_Stop_Call) Return(_a0 error) *AgentServer_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AgentServer_Stop_Call) RunAndReturn(run func() error) *AgentServer_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewAgentServer creates a new instance of AgentServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAgentServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *AgentServer {
	mock := &AgentServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
