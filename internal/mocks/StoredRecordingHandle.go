package mocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// StoredRecordingHandle is an autogenerated mock type for the StoredRecordingHandle type
type StoredRecordingHandle struct {
	mock.Mock
}

// Copy provides a mock function with given fields: dest
func (_m *StoredRecordingHandle) Copy(dest string) (ari.StoredRecordingHandle, error) {
	ret := _m.Called(dest)

	var r0 ari.StoredRecordingHandle
	if rf, ok := ret.Get(0).(func(string) ari.StoredRecordingHandle); ok {
		r0 = rf(dest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.StoredRecordingHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Data provides a mock function with given fields:
func (_m *StoredRecordingHandle) Data() (*ari.StoredRecordingData, error) {
	ret := _m.Called()

	var r0 *ari.StoredRecordingData
	if rf, ok := ret.Get(0).(func() *ari.StoredRecordingData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.StoredRecordingData)
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
func (_m *StoredRecordingHandle) Delete() error {
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
func (_m *StoredRecordingHandle) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
