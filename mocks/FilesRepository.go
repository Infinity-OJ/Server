// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FilesRepository is an autogenerated mock type for the FilesRepository type
type FilesRepository struct {
	mock.Mock
}

// CreateDirectory provides a mock function with given fields: fileSpace, directory
func (_m *FilesRepository) CreateDirectory(fileSpace string, directory string) error {
	ret := _m.Called(fileSpace, directory)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(fileSpace, directory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateFile provides a mock function with given fields: fileSpace, fileName, data
func (_m *FilesRepository) CreateFile(fileSpace string, fileName string, data []byte) error {
	ret := _m.Called(fileSpace, fileName, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []byte) error); ok {
		r0 = rf(fileSpace, fileName, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateFileSpace provides a mock function with given fields: fileSpace
func (_m *FilesRepository) CreateFileSpace(fileSpace string) error {
	ret := _m.Called(fileSpace)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(fileSpace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchFile provides a mock function with given fields: fileSpace, fileName
func (_m *FilesRepository) FetchFile(fileSpace string, fileName string) ([]byte, error) {
	ret := _m.Called(fileSpace, fileName)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(fileSpace, fileName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(fileSpace, fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsFileExists provides a mock function with given fields: fileSpace, fileName
func (_m *FilesRepository) IsFileExists(fileSpace string, fileName string) bool {
	ret := _m.Called(fileSpace, fileName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(fileSpace, fileName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
