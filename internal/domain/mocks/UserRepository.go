// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/ultimathul3/notes-backend/internal/domain"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *UserRepository) Create(ctx context.Context, user domain.User) (int64, error) {
	ret := _m.Called(ctx, user)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) (int64, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) int64); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetID provides a mock function with given fields: ctx, login, passwordHash
func (_m *UserRepository) GetID(ctx context.Context, login string, passwordHash string) (int64, error) {
	ret := _m.Called(ctx, login, passwordHash)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (int64, error)); ok {
		return rf(ctx, login, passwordHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) int64); ok {
		r0 = rf(ctx, login, passwordHash)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, login, passwordHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
