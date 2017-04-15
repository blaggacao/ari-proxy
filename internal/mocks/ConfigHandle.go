package mocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// ConfigHandle is an autogenerated mock type for the ConfigHandle type
type ConfigHandle struct {
	mock.Mock
}

// Data provides a mock function with given fields:
func (_m *ConfigHandle) Data() (*ari.ConfigData, error) {
	ret := _m.Called()

	var r0 *ari.ConfigData
	if rf, ok := ret.Get(0).(func() *ari.ConfigData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.ConfigData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields:
func (_m *ConfigHandle) Delete() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *ConfigHandle) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Update provides a mock function with given fields: tuples
func (_m *ConfigHandle) Update(tuples []ari.ConfigTuple) error {
	ret := _m.Called(tuples)

	var r0 error
	if rf, ok := ret.Get(0).(func([]ari.ConfigTuple) error); ok {
		r0 = rf(tuples)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
