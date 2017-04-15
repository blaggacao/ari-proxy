package mocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// Data provides a mock function with given fields: configClass, objectType, id
func (_m *Config) Data(configClass string, objectType string, id string) (*ari.ConfigData, error) {
	ret := _m.Called(configClass, objectType, id)

	var r0 *ari.ConfigData
	if rf, ok := ret.Get(0).(func(string, string, string) *ari.ConfigData); ok {
		r0 = rf(configClass, objectType, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.ConfigData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(configClass, objectType, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: configClass, objectType, id
func (_m *Config) Delete(configClass string, objectType string, id string) error {
	ret := _m.Called(configClass, objectType, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(configClass, objectType, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: configClass, objectType, id
func (_m *Config) Get(configClass string, objectType string, id string) ari.ConfigHandle {
	ret := _m.Called(configClass, objectType, id)

	var r0 ari.ConfigHandle
	if rf, ok := ret.Get(0).(func(string, string, string) ari.ConfigHandle); ok {
		r0 = rf(configClass, objectType, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.ConfigHandle)
		}
	}

	return r0
}

// Update provides a mock function with given fields: configClass, objectType, id, tuples
func (_m *Config) Update(configClass string, objectType string, id string, tuples []ari.ConfigTuple) error {
	ret := _m.Called(configClass, objectType, id, tuples)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, []ari.ConfigTuple) error); ok {
		r0 = rf(configClass, objectType, id, tuples)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
