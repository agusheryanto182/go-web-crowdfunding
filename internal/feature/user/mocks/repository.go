// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// UserRepositoryInterface is an autogenerated mock type for the UserRepositoryInterface type
type UserRepositoryInterface struct {
	mock.Mock
}

// FindUserByEmail provides a mock function with given fields: email
func (_m *UserRepositoryInterface) FindUserByEmail(email string) (*entity.UserModels, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for FindUserByEmail")
	}

	var r0 *entity.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.UserModels, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.UserModels); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: userID
func (_m *UserRepositoryInterface) GetByID(userID int) (*entity.UserModels, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *entity.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*entity.UserModels, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(int) *entity.UserModels); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: _a0
func (_m *UserRepositoryInterface) UpdateUser(_a0 *entity.UserModels) (*entity.UserModels, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 *entity.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.UserModels) (*entity.UserModels, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*entity.UserModels) *entity.UserModels); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.UserModels) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepositoryInterface creates a new instance of UserRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepositoryInterface {
	mock := &UserRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
