// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/kujilabo/cocotola-tatoeba-api/pkg/domain"
	mock "github.com/stretchr/testify/mock"

	testing "testing"

	time "time"
)

// TatoebaSentenceAddParameter is an autogenerated mock type for the TatoebaSentenceAddParameter type
type TatoebaSentenceAddParameter struct {
	mock.Mock
}

// GetAuthor provides a mock function with given fields:
func (_m *TatoebaSentenceAddParameter) GetAuthor() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetLang provides a mock function with given fields:
func (_m *TatoebaSentenceAddParameter) GetLang() domain.Lang3 {
	ret := _m.Called()

	var r0 domain.Lang3
	if rf, ok := ret.Get(0).(func() domain.Lang3); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Lang3)
		}
	}

	return r0
}

// GetSentenceNumber provides a mock function with given fields:
func (_m *TatoebaSentenceAddParameter) GetSentenceNumber() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetText provides a mock function with given fields:
func (_m *TatoebaSentenceAddParameter) GetText() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetUpdatedAt provides a mock function with given fields:
func (_m *TatoebaSentenceAddParameter) GetUpdatedAt() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// NewTatoebaSentenceAddParameter creates a new instance of TatoebaSentenceAddParameter. It also registers a cleanup function to assert the mocks expectations.
func NewTatoebaSentenceAddParameter(t testing.TB) *TatoebaSentenceAddParameter {
	mock := &TatoebaSentenceAddParameter{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
