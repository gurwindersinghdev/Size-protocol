// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	feeds "github.com/smartcontractkit/chainlink/v2/core/services/feeds"
	mock "github.com/stretchr/testify/mock"

	proto "github.com/smartcontractkit/chainlink/v2/core/services/feeds/proto"
)

// ConnectionsManager is an autogenerated mock type for the ConnectionsManager type
type ConnectionsManager struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ConnectionsManager) Close() {
	_m.Called()
}

// Connect provides a mock function with given fields: opts
func (_m *ConnectionsManager) Connect(opts feeds.ConnectOpts) {
	_m.Called(opts)
}

// Disconnect provides a mock function with given fields: id
func (_m *ConnectionsManager) Disconnect(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetClient provides a mock function with given fields: id
func (_m *ConnectionsManager) GetClient(id int64) (proto.FeedsManagerClient, error) {
	ret := _m.Called(id)

	var r0 proto.FeedsManagerClient
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (proto.FeedsManagerClient, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) proto.FeedsManagerClient); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(proto.FeedsManagerClient)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsConnected provides a mock function with given fields: id
func (_m *ConnectionsManager) IsConnected(id int64) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewConnectionsManager creates a new instance of ConnectionsManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConnectionsManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConnectionsManager {
	mock := &ConnectionsManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
