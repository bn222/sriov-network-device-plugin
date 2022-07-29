// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	v1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
)

// NadUtils is an autogenerated mock type for the NadUtils type
type NadUtils struct {
	mock.Mock
}

// CleanDeviceInfoFile provides a mock function with given fields: resourceName, deviceID
func (_m *NadUtils) CleanDeviceInfoFile(resourceName string, deviceID string) error {
	ret := _m.Called(resourceName, deviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(resourceName, deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDeviceInfoFile provides a mock function with given fields: resourceName, deviceID, devInfo
func (_m *NadUtils) SaveDeviceInfoFile(resourceName string, deviceID string, devInfo *v1.DeviceInfo) error {
	ret := _m.Called(resourceName, deviceID, devInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *v1.DeviceInfo) error); ok {
		r0 = rf(resourceName, deviceID, devInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewNadUtilsT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNadUtils creates a new instance of NadUtils. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNadUtils(t NewNadUtilsT) *NadUtils {
	mock := &NadUtils{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
