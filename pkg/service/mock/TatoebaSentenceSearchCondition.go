// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// TatoebaSentenceSearchCondition is an autogenerated mock type for the TatoebaSentenceSearchCondition type
type TatoebaSentenceSearchCondition struct {
	mock.Mock
}

// GetKeyword provides a mock function with given fields:
func (_m *TatoebaSentenceSearchCondition) GetKeyword() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetPageNo provides a mock function with given fields:
func (_m *TatoebaSentenceSearchCondition) GetPageNo() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetPageSize provides a mock function with given fields:
func (_m *TatoebaSentenceSearchCondition) GetPageSize() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// IsRandom provides a mock function with given fields:
func (_m *TatoebaSentenceSearchCondition) IsRandom() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewTatoebaSentenceSearchCondition creates a new instance of TatoebaSentenceSearchCondition. It also registers a cleanup function to assert the mocks expectations.
func NewTatoebaSentenceSearchCondition(t testing.TB) *TatoebaSentenceSearchCondition {
	mock := &TatoebaSentenceSearchCondition{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
