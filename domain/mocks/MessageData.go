// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "wa/domain"

	mock "github.com/stretchr/testify/mock"
)

// MessageData is an autogenerated mock type for the MessageData type
type MessageData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: IDMessage
func (_m *MessageData) Delete(IDMessage int) bool {
	ret := _m.Called(IDMessage)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(IDMessage)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Insert provides a mock function with given fields: insertMessage
func (_m *MessageData) Insert(insertMessage domain.Message) domain.Message {
	ret := _m.Called(insertMessage)

	var r0 domain.Message
	if rf, ok := ret.Get(0).(func(domain.Message) domain.Message); ok {
		r0 = rf(insertMessage)
	} else {
		r0 = ret.Get(0).(domain.Message)
	}

	return r0
}

type mockConstructorTestingTNewMessageData interface {
	mock.TestingT
	Cleanup(func())
}

// NewMessageData creates a new instance of MessageData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMessageData(t mockConstructorTestingTNewMessageData) *MessageData {
	mock := &MessageData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}