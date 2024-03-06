// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// AuthRepositoryInterface is an autogenerated mock type for the AuthRepositoryInterface type
type AuthRepositoryInterface struct {
	mock.Mock
}

// DeleteOTP provides a mock function with given fields: OTP
func (_m *AuthRepositoryInterface) DeleteOTP(OTP *entity.OTPModels) error {
	ret := _m.Called(OTP)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOTP")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.OTPModels) error); ok {
		r0 = rf(OTP)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindValidOTP provides a mock function with given fields: userID, OTP
func (_m *AuthRepositoryInterface) FindValidOTP(userID int, OTP string) (*entity.OTPModels, error) {
	ret := _m.Called(userID, OTP)

	if len(ret) == 0 {
		panic("no return value specified for FindValidOTP")
	}

	var r0 *entity.OTPModels
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string) (*entity.OTPModels, error)); ok {
		return rf(userID, OTP)
	}
	if rf, ok := ret.Get(0).(func(int, string) *entity.OTPModels); ok {
		r0 = rf(userID, OTP)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.OTPModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(userID, OTP)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveOTP provides a mock function with given fields: OTP
func (_m *AuthRepositoryInterface) SaveOTP(OTP *entity.OTPModels) (*entity.OTPModels, error) {
	ret := _m.Called(OTP)

	if len(ret) == 0 {
		panic("no return value specified for SaveOTP")
	}

	var r0 *entity.OTPModels
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.OTPModels) (*entity.OTPModels, error)); ok {
		return rf(OTP)
	}
	if rf, ok := ret.Get(0).(func(*entity.OTPModels) *entity.OTPModels); ok {
		r0 = rf(OTP)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.OTPModels)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.OTPModels) error); ok {
		r1 = rf(OTP)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: user
func (_m *AuthRepositoryInterface) SignUp(user *entity.UserModels) (*entity.UserModels, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for SignUp")
	}

	var r0 *entity.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.UserModels) (*entity.UserModels, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*entity.UserModels) *entity.UserModels); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.UserModels) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthRepositoryInterface creates a new instance of AuthRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthRepositoryInterface {
	mock := &AuthRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
