// Code generated by mockery v2.42.2. DO NOT EDIT.

package repositories

import (
	context "context"
	entities "nutri-plans-api/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockMealTypeRepository is an autogenerated mock type for the MealTypeRepository type
type MockMealTypeRepository struct {
	mock.Mock
}

// GetMealTypes provides a mock function with given fields: ctx
func (_m *MockMealTypeRepository) GetMealTypes(ctx context.Context) (*[]entities.MealType, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetMealTypes")
	}

	var r0 *[]entities.MealType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*[]entities.MealType, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *[]entities.MealType); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.MealType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockMealTypeRepository creates a new instance of MockMealTypeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMealTypeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMealTypeRepository {
	mock := &MockMealTypeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}