// Code generated by mockery. DO NOT EDIT.

//go:build test
// +build test

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	rpc "go.woodpecker-ci.org/woodpecker/v2/pipeline/rpc"
)

// Peer is an autogenerated mock type for the Peer type
type Peer struct {
	mock.Mock
}

// Done provides a mock function with given fields: c, workflowID, state
func (_m *Peer) Done(c context.Context, workflowID string, state rpc.WorkflowState) error {
	ret := _m.Called(c, workflowID, state)

	if len(ret) == 0 {
		panic("no return value specified for Done")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, rpc.WorkflowState) error); ok {
		r0 = rf(c, workflowID, state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Extend provides a mock function with given fields: c, workflowID
func (_m *Peer) Extend(c context.Context, workflowID string) error {
	ret := _m.Called(c, workflowID)

	if len(ret) == 0 {
		panic("no return value specified for Extend")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, workflowID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Init provides a mock function with given fields: c, workflowID, state
func (_m *Peer) Init(c context.Context, workflowID string, state rpc.WorkflowState) error {
	ret := _m.Called(c, workflowID, state)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, rpc.WorkflowState) error); ok {
		r0 = rf(c, workflowID, state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Log provides a mock function with given fields: c, logEntry
func (_m *Peer) Log(c context.Context, logEntry *rpc.LogEntry) error {
	ret := _m.Called(c, logEntry)

	if len(ret) == 0 {
		panic("no return value specified for Log")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *rpc.LogEntry) error); ok {
		r0 = rf(c, logEntry)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Next provides a mock function with given fields: c, f
func (_m *Peer) Next(c context.Context, f rpc.Filter) (*rpc.Workflow, error) {
	ret := _m.Called(c, f)

	if len(ret) == 0 {
		panic("no return value specified for Next")
	}

	var r0 *rpc.Workflow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, rpc.Filter) (*rpc.Workflow, error)); ok {
		return rf(c, f)
	}
	if rf, ok := ret.Get(0).(func(context.Context, rpc.Filter) *rpc.Workflow); ok {
		r0 = rf(c, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.Workflow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, rpc.Filter) error); ok {
		r1 = rf(c, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterAgent provides a mock function with given fields: ctx, platform, backend, version, capacity
func (_m *Peer) RegisterAgent(ctx context.Context, platform string, backend string, version string, capacity int) (int64, error) {
	ret := _m.Called(ctx, platform, backend, version, capacity)

	if len(ret) == 0 {
		panic("no return value specified for RegisterAgent")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, int) (int64, error)); ok {
		return rf(ctx, platform, backend, version, capacity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, int) int64); ok {
		r0 = rf(ctx, platform, backend, version, capacity)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, int) error); ok {
		r1 = rf(ctx, platform, backend, version, capacity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReportHealth provides a mock function with given fields: c
func (_m *Peer) ReportHealth(c context.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for ReportHealth")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnregisterAgent provides a mock function with given fields: ctx
func (_m *Peer) UnregisterAgent(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for UnregisterAgent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: c, workflowID, state
func (_m *Peer) Update(c context.Context, workflowID string, state rpc.StepState) error {
	ret := _m.Called(c, workflowID, state)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, rpc.StepState) error); ok {
		r0 = rf(c, workflowID, state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Version provides a mock function with given fields: c
func (_m *Peer) Version(c context.Context) (*rpc.Version, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Version")
	}

	var r0 *rpc.Version
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*rpc.Version, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *rpc.Version); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.Version)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Wait provides a mock function with given fields: c, workflowID
func (_m *Peer) Wait(c context.Context, workflowID string) error {
	ret := _m.Called(c, workflowID)

	if len(ret) == 0 {
		panic("no return value specified for Wait")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, workflowID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPeer creates a new instance of Peer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPeer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Peer {
	mock := &Peer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}