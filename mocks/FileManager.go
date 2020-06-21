// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FileManager is an autogenerated mock type for the FileManager type
type FileManager struct {
	mock.Mock
}

// CreateDirectory provides a mock function with given fields: fileName
func (_m *FileManager) CreateDirectory(fileName string) error {
	ret := _m.Called(fileName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(fileName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateFile provides a mock function with given fields: fileName, bytes
func (_m *FileManager) CreateFile(fileName string, bytes []byte) error {
	ret := _m.Called(fileName, bytes)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(fileName, bytes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchFile provides a mock function with given fields: fileName
func (_m *FileManager) FetchFile(fileName string) ([]byte, error) {
	ret := _m.Called(fileName)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(fileName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBase provides a mock function with given fields:
func (_m *FileManager) GetBase() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsDirectoryExists provides a mock function with given fields: fileName
func (_m *FileManager) IsDirectoryExists(fileName string) (bool, error) {
	ret := _m.Called(fileName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(fileName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsFileExists provides a mock function with given fields: fileName
func (_m *FileManager) IsFileExists(fileName string) (bool, error) {
	ret := _m.Called(fileName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(fileName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetBase provides a mock function with given fields: base
func (_m *FileManager) SetBase(base string) {
	_m.Called(base)
}
