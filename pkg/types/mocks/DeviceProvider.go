// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	ghw "github.com/jaypipes/ghw"
	mock "github.com/stretchr/testify/mock"

	types "github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/types"
)

// DeviceProvider is an autogenerated mock type for the DeviceProvider type
type DeviceProvider struct {
	mock.Mock
}

// AddTargetDevices provides a mock function with given fields: _a0, _a1
func (_m *DeviceProvider) AddTargetDevices(_a0 []*ghw.PCIDevice, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*ghw.PCIDevice, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDevices provides a mock function with given fields: _a0
func (_m *DeviceProvider) GetDevices(_a0 *types.ResourceConfig) []types.PciDevice {
	ret := _m.Called(_a0)

	var r0 []types.PciDevice
	if rf, ok := ret.Get(0).(func(*types.ResourceConfig) []types.PciDevice); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.PciDevice)
		}
	}

	return r0
}

// GetDiscoveredDevices provides a mock function with given fields:
func (_m *DeviceProvider) GetDiscoveredDevices() []*ghw.PCIDevice {
	ret := _m.Called()

	var r0 []*ghw.PCIDevice
	if rf, ok := ret.Get(0).(func() []*ghw.PCIDevice); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ghw.PCIDevice)
		}
	}

	return r0
}

// GetFilteredDevices provides a mock function with given fields: _a0, _a1
func (_m *DeviceProvider) GetFilteredDevices(_a0 []types.PciDevice, _a1 *types.ResourceConfig) ([]types.PciDevice, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []types.PciDevice
	if rf, ok := ret.Get(0).(func([]types.PciDevice, *types.ResourceConfig) []types.PciDevice); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.PciDevice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]types.PciDevice, *types.ResourceConfig) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidConfig provides a mock function with given fields: _a0
func (_m *DeviceProvider) ValidConfig(_a0 *types.ResourceConfig) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*types.ResourceConfig) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type NewDeviceProviderT interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeviceProvider creates a new instance of DeviceProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeviceProvider(t NewDeviceProviderT) *DeviceProvider {
	mock := &DeviceProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
