// Code generated by mockery v2.10.0. DO NOT EDIT.

package products

import (
	context "context"
	model "gobase/api/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// MockController is an autogenerated mock type for the Controller type
type MockController struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, inp
func (_m *MockController) Create(ctx context.Context, inp CreateInput) (model.Product, error) {
	ret := _m.Called(ctx, inp)

	var r0 model.Product
	if rf, ok := ret.Get(0).(func(context.Context, CreateInput) model.Product); ok {
		r0 = rf(ctx, inp)
	} else {
		r0 = ret.Get(0).(model.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, CreateInput) error); ok {
		r1 = rf(ctx, inp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, extID
func (_m *MockController) Delete(ctx context.Context, extID string) error {
	ret := _m.Called(ctx, extID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, extID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, extID
func (_m *MockController) Get(ctx context.Context, extID string) (model.Product, error) {
	ret := _m.Called(ctx, extID)

	var r0 model.Product
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Product); ok {
		r0 = rf(ctx, extID)
	} else {
		r0 = ret.Get(0).(model.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, extID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *MockController) List(ctx context.Context) ([]model.Product, error) {
	ret := _m.Called(ctx)

	var r0 []model.Product
	if rf, ok := ret.Get(0).(func(context.Context) []model.Product); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
