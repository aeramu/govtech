// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/alam/govtech/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// GetProduct provides a mock function with given fields: ctx, id
func (_m *ProductRepository) GetProduct(ctx context.Context, id int64) (model.Product, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (model.Product, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.Product); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Product)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductBySKU provides a mock function with given fields: ctx, sku
func (_m *ProductRepository) GetProductBySKU(ctx context.Context, sku string) (model.Product, error) {
	ret := _m.Called(ctx, sku)

	var r0 model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.Product, error)); ok {
		return rf(ctx, sku)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Product); ok {
		r0 = rf(ctx, sku)
	} else {
		r0 = ret.Get(0).(model.Product)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sku)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductList provides a mock function with given fields: ctx, filter
func (_m *ProductRepository) GetProductList(ctx context.Context, filter model.GetProductListFilter) ([]model.Product, error) {
	ret := _m.Called(ctx, filter)

	var r0 []model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetProductListFilter) ([]model.Product, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetProductListFilter) []model.Product); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetProductListFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertProduct provides a mock function with given fields: ctx, product
func (_m *ProductRepository) InsertProduct(ctx context.Context, product model.Product) error {
	ret := _m.Called(ctx, product)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Product) error); ok {
		r0 = rf(ctx, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProduct provides a mock function with given fields: ctx, id, product
func (_m *ProductRepository) UpdateProduct(ctx context.Context, id int64, product model.Product) error {
	ret := _m.Called(ctx, id, product)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, model.Product) error); ok {
		r0 = rf(ctx, id, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProductRating provides a mock function with given fields: ctx, id, rating
func (_m *ProductRepository) UpdateProductRating(ctx context.Context, id int64, rating float64) error {
	ret := _m.Called(ctx, id, rating)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, float64) error); ok {
		r0 = rf(ctx, id, rating)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
