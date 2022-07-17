// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	v1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

// AccelDevice is an autogenerated mock type for the AccelDevice type
type AccelDevice struct {
	mock.Mock
}

// GetAPIDevice provides a mock function with given fields:
func (_m *AccelDevice) GetAPIDevice() *v1beta1.Device {
	ret := _m.Called()

	var r0 *v1beta1.Device
	if rf, ok := ret.Get(0).(func() *v1beta1.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Device)
		}
	}

	return r0
}

// GetDeviceCode provides a mock function with given fields:
func (_m *AccelDevice) GetDeviceCode() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetDeviceSpecs provides a mock function with given fields:
func (_m *AccelDevice) GetDeviceSpecs() []*v1beta1.DeviceSpec {
	ret := _m.Called()

	var r0 []*v1beta1.DeviceSpec
	if rf, ok := ret.Get(0).(func() []*v1beta1.DeviceSpec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.DeviceSpec)
		}
	}

	return r0
}

// GetDriver provides a mock function with given fields:
func (_m *AccelDevice) GetDriver() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetEnvVal provides a mock function with given fields:
func (_m *AccelDevice) GetEnvVal() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetMounts provides a mock function with given fields:
func (_m *AccelDevice) GetMounts() []*v1beta1.Mount {
	ret := _m.Called()

	var r0 []*v1beta1.Mount
	if rf, ok := ret.Get(0).(func() []*v1beta1.Mount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.Mount)
		}
	}

	return r0
}

// GetPciAddr provides a mock function with given fields:
func (_m *AccelDevice) GetPciAddr() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetPfPciAddr provides a mock function with given fields:
func (_m *AccelDevice) GetPfPciAddr() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetVFID provides a mock function with given fields:
func (_m *AccelDevice) GetVFID() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetVendor provides a mock function with given fields:
func (_m *AccelDevice) GetVendor() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsSriovPF provides a mock function with given fields:
func (_m *AccelDevice) IsSriovPF() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type NewAccelDeviceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccelDevice creates a new instance of AccelDevice. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccelDevice(t NewAccelDeviceT) *AccelDevice {
	mock := &AccelDevice{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
