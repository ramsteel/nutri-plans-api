// Code generated by mockery v2.42.2. DO NOT EDIT.

package repositories

import (
	context "context"
	entities "nutri-plans-api/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *MockUserRepository) CreateUser(ctx context.Context, user *entities.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}