// Code generated by mockery v2.14.0. DO NOT EDIT.

package config

import (
	context "context"

	metadata "github.com/goharbor/harbor/src/lib/config/metadata"
	mock "github.com/stretchr/testify/mock"

	models "github.com/goharbor/harbor/src/common/models"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, key
func (_m *Manager) Get(ctx context.Context, key string) *metadata.ConfigureValue {
	ret := _m.Called(ctx, key)

	var r0 *metadata.ConfigureValue
	if rf, ok := ret.Get(0).(func(context.Context, string) *metadata.ConfigureValue); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metadata.ConfigureValue)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *Manager) GetAll(ctx context.Context) map[string]interface{} {
	ret := _m.Called(ctx)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]interface{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// GetDatabaseCfg provides a mock function with given fields:
func (_m *Manager) GetDatabaseCfg() *models.Database {
	ret := _m.Called()

	var r0 *models.Database
	if rf, ok := ret.Get(0).(func() *models.Database); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Database)
		}
	}

	return r0
}

// GetUserCfgs provides a mock function with given fields: ctx
func (_m *Manager) GetUserCfgs(ctx context.Context) map[string]interface{} {
	ret := _m.Called(ctx)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]interface{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// Load provides a mock function with given fields: ctx
func (_m *Manager) Load(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: ctx
func (_m *Manager) Save(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Set provides a mock function with given fields: ctx, key, value
func (_m *Manager) Set(ctx context.Context, key string, value interface{}) {
	_m.Called(ctx, key, value)
}

// UpdateConfig provides a mock function with given fields: ctx, cfgs
func (_m *Manager) UpdateConfig(ctx context.Context, cfgs map[string]interface{}) error {
	ret := _m.Called(ctx, cfgs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) error); ok {
		r0 = rf(ctx, cfgs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateCfg provides a mock function with given fields: ctx, cfgs
func (_m *Manager) ValidateCfg(ctx context.Context, cfgs map[string]interface{}) error {
	ret := _m.Called(ctx, cfgs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) error); ok {
		r0 = rf(ctx, cfgs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewManager(t mockConstructorTestingTNewManager) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
