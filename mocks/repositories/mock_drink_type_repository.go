// Code generated by mockery v2.42.2. DO NOT EDIT.

package repositories

import (
	context "context"
	entities "nutri-plans-api/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockDrinkTypeRepository is an autogenerated mock type for the DrinkTypeRepository type
type MockDrinkTypeRepository struct {
	mock.Mock
}

// GetDrinkTypes provides a mock function with given fields: ctx
func (_m *MockDrinkTypeRepository) GetDrinkTypes(ctx context.Context) (*[]entities.DrinkType, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetDrinkTypes")
	}

	var r0 *[]entities.DrinkType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*[]entities.DrinkType, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *[]entities.DrinkType); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.DrinkType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockDrinkTypeRepository creates a new instance of MockDrinkTypeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDrinkTypeRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDrinkTypeRepository {
	mock := &MockDrinkTypeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}