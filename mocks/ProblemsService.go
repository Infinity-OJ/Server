// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/infinity-oj/server/internal/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// ProblemsService is an autogenerated mock type for the ProblemsService type
type ProblemsService struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ProblemID
func (_m *ProblemsService) Fetch(ProblemID string) (*models.Problem, error) {
	ret := _m.Called(ProblemID)

	var r0 *models.Problem
	if rf, ok := ret.Get(0).(func(string) *models.Problem); ok {
		r0 = rf(ProblemID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Problem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ProblemID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
