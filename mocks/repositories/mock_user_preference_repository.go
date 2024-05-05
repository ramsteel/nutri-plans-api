// Code generated by mockery v2.42.2. DO NOT EDIT.

package repositories

import (
	context "context"
	entities "nutri-plans-api/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockUserPreferenceRepository is an autogenerated mock type for the UserPreferenceRepository type
type MockUserPreferenceRepository struct {
	mock.Mock
}

// UpdateUserPreference provides a mock function with given fields: ctx, userPreference
func (_m *MockUserPreferenceRepository) UpdateUserPreference(ctx context.Context, userPreference *entities.UserPreference) error {
	ret := _m.Called(ctx, userPreference)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserPreference")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.UserPreference) error); ok {
		r0 = rf(ctx, userPreference)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockUserPreferenceRepository creates a new instance of MockUserPreferenceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserPreferenceRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserPreferenceRepository {
	mock := &MockUserPreferenceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
