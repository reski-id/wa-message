// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "wa/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserUseCase is an autogenerated mock type for the UserUseCase type
type UserUseCase struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: newUser
func (_m *UserUseCase) AddUser(newUser domain.User) (domain.User, error) {
	ret := _m.Called(newUser)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: id
func (_m *UserUseCase) DeleteUser(id int) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfile provides a mock function with given fields: id
func (_m *UserUseCase) GetProfile(id int) (domain.User, error) {
	ret := _m.Called(id)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(int) domain.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: userLogin
func (_m *UserUseCase) LoginUser(userLogin domain.User) (int, domain.User, error) {
	ret := _m.Called(userLogin)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.User) int); ok {
		r0 = rf(userLogin)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 domain.User
	if rf, ok := ret.Get(1).(func(domain.User) domain.User); ok {
		r1 = rf(userLogin)
	} else {
		r1 = ret.Get(1).(domain.User)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(domain.User) error); ok {
		r2 = rf(userLogin)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateUser provides a mock function with given fields: id, updateProfile
func (_m *UserUseCase) UpdateUser(id int, updateProfile domain.User) (domain.User, error) {
	ret := _m.Called(id, updateProfile)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(int, domain.User) domain.User); ok {
		r0 = rf(id, updateProfile)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, domain.User) error); ok {
		r1 = rf(id, updateProfile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUseCase creates a new instance of UserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUseCase(t mockConstructorTestingTNewUserUseCase) *UserUseCase {
	mock := &UserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
