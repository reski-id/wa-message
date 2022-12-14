// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "wa/domain"

	mock "github.com/stretchr/testify/mock"
)

// MessageUseCase is an autogenerated mock type for the MessageUseCase type
type MessageUseCase struct {
	mock.Mock
}

// AddMessage provides a mock function with given fields: IDUser, useMessage
func (_m *MessageUseCase) AddMessage(IDUser int, useMessage domain.Message) (domain.Message, error) {
	ret := _m.Called(IDUser, useMessage)

	var r0 domain.Message
	if rf, ok := ret.Get(0).(func(int, domain.Message) domain.Message); ok {
		r0 = rf(IDUser, useMessage)
	} else {
		r0 = ret.Get(0).(domain.Message)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, domain.Message) error); ok {
		r1 = rf(IDUser, useMessage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DelMessage provides a mock function with given fields: IDMessage
func (_m *MessageUseCase) DelMessage(IDMessage int) (bool, error) {
	ret := _m.Called(IDMessage)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(IDMessage)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(IDMessage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMessageUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewMessageUseCase creates a new instance of MessageUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMessageUseCase(t mockConstructorTestingTNewMessageUseCase) *MessageUseCase {
	mock := &MessageUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
