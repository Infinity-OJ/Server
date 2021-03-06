// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FilesService is an autogenerated mock type for the FilesService type
type FilesService struct {
	mock.Mock
}

// FetchFile provides a mock function with given fields: fileSpace, fileName
func (_m *FilesService) FetchFile(fileSpace string, fileName string) ([]byte, error) {
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
