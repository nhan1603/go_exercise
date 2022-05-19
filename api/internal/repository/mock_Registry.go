// Code generated by mockery v2.10.0. DO NOT EDIT.

package repository

import (
	context "context"

	backoff "github.com/cenkalti/backoff/v4"

	inventory "gobase/api/internal/repository/inventory"

	mock "github.com/stretchr/testify/mock"

	system "gobase/api/internal/repository/system"
)

// MockRegistry is an autogenerated mock type for the Registry type
type MockRegistry struct {
	mock.Mock
}

// DoInTx provides a mock function with given fields: ctx, txFunc, overrideBackoffPolicy
func (_m *MockRegistry) DoInTx(ctx context.Context, txFunc func(context.Context, Registry) error, overrideBackoffPolicy backoff.BackOff) error {
	ret := _m.Called(ctx, txFunc, overrideBackoffPolicy)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context, Registry) error, backoff.BackOff) error); ok {
		r0 = rf(ctx, txFunc, overrideBackoffPolicy)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Inventory provides a mock function with given fields:
func (_m *MockRegistry) Inventory() inventory.Repository {
	ret := _m.Called()

	var r0 inventory.Repository
	if rf, ok := ret.Get(0).(func() inventory.Repository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(inventory.Repository)
		}
	}

	return r0
}

// System provides a mock function with given fields:
func (_m *MockRegistry) System() system.Repository {
	ret := _m.Called()

	var r0 system.Repository
	if rf, ok := ret.Get(0).(func() system.Repository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(system.Repository)
		}
	}

	return r0
}
