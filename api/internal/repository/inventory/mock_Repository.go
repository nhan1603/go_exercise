// Code generated by mockery v2.10.0. DO NOT EDIT.

package inventory

import (
	context "context"
	model "gobase/api/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) CreateProduct(_a0 context.Context, _a1 model.Product) (model.Product, error) {
	ret := _m.Called(_a0, _a1)

	var r0 model.Product
	if rf, ok := ret.Get(0).(func(context.Context, model.Product) model.Product); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Product) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveProductsCountFromDB provides a mock function with given fields: ctx
func (_m *MockRepository) GetActiveProductsCountFromDB(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProducts provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) ListProducts(_a0 context.Context, _a1 ProductsFilter) ([]model.Product, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []model.Product
	if rf, ok := ret.Get(0).(func(context.Context, ProductsFilter) []model.Product); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ProductsFilter) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProductStatus provides a mock function with given fields: _a0, _a1
func (_m *MockRepository) UpdateProductStatus(_a0 context.Context, _a1 model.Product) (model.Product, error) {
	ret := _m.Called(_a0, _a1)

	var r0 model.Product
	if rf, ok := ret.Get(0).(func(context.Context, model.Product) model.Product); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Product) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
