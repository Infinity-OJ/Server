// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/infinity-oj/server/internal/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// ProductsService is an autogenerated mock type for the ProductsService type
type ProductsService struct {
	mock.Mock
}

// Get provides a mock function with given fields: c, ID
func (_m *ProductsService) Get(c context.Context, ID uint64) (*models.User, error) {
	ret := _m.Called(c, ID)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.User); ok {
		r0 = rf(c, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
