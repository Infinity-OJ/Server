// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	repositories "github.com/infinity-oj/server/internal/app/submissions/repositories"
	models "github.com/infinity-oj/server/internal/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// SubmissionRepository is an autogenerated mock type for the SubmissionRepository type
type SubmissionRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: submitterID, problemId, userSpace
func (_m *SubmissionRepository) Create(submitterID uint64, problemId string, userSpace string) (*models.Submission, error) {
	ret := _m.Called(submitterID, problemId, userSpace)

	var r0 *models.Submission
	if rf, ok := ret.Get(0).(func(uint64, string, string) *models.Submission); ok {
		r0 = rf(submitterID, problemId, userSpace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Submission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, string, string) error); ok {
		r1 = rf(submitterID, problemId, userSpace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProcess provides a mock function with given fields: s
func (_m *SubmissionRepository) CreateProcess(s *models.Submission) *repositories.Process {
	ret := _m.Called(s)

	var r0 *repositories.Process
	if rf, ok := ret.Get(0).(func(*models.Submission) *repositories.Process); ok {
		r0 = rf(s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repositories.Process)
		}
	}

	return r0
}

// FetchSubmissionBySubmissionId provides a mock function with given fields: submissionId
func (_m *SubmissionRepository) FetchSubmissionBySubmissionId(submissionId string) (*models.Submission, error) {
	ret := _m.Called(submissionId)

	var r0 *models.Submission
	if rf, ok := ret.Get(0).(func(string) *models.Submission); ok {
		r0 = rf(submissionId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Submission)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(submissionId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchProcess provides a mock function with given fields: submissionId
func (_m *SubmissionRepository) FetchProcess(submissionId string) *repositories.Process {
	ret := _m.Called(submissionId)

	var r0 *repositories.Process
	if rf, ok := ret.Get(0).(func(string) *repositories.Process); ok {
		r0 = rf(submissionId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repositories.Process)
		}
	}

	return r0
}

// Update provides a mock function with given fields: s
func (_m *SubmissionRepository) Update(s *models.Submission) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Submission) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
