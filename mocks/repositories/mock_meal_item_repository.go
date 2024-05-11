// Code generated by mockery v2.42.2. DO NOT EDIT.

package repositories

import (
	context "context"
	entities "nutri-plans-api/entities"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
)

// MockMealItemRepository is an autogenerated mock type for the MealItemRepository type
type MockMealItemRepository struct {
	mock.Mock
}

// DeleteMealItem provides a mock function with given fields: ctx, id
func (_m *MockMealItemRepository) DeleteMealItem(ctx context.Context, id uint64) (*entities.MealItem, *gorm.DB, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMealItem")
	}

	var r0 *entities.MealItem
	var r1 *gorm.DB
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*entities.MealItem, *gorm.DB, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *entities.MealItem); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.MealItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) *gorm.DB); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gorm.DB)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64) error); ok {
		r2 = rf(ctx, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetCalculatedNutrients provides a mock function with given fields: ctx, id, start, end
func (_m *MockMealItemRepository) GetCalculatedNutrients(ctx context.Context, id uuid.UUID, start time.Time, end time.Time) (*entities.CalculatedNutrients, error) {
	ret := _m.Called(ctx, id, start, end)

	if len(ret) == 0 {
		panic("no return value specified for GetCalculatedNutrients")
	}

	var r0 *entities.CalculatedNutrients
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, time.Time, time.Time) (*entities.CalculatedNutrients, error)); ok {
		return rf(ctx, id, start, end)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, time.Time, time.Time) *entities.CalculatedNutrients); ok {
		r0 = rf(ctx, id, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.CalculatedNutrients)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, time.Time, time.Time) error); ok {
		r1 = rf(ctx, id, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMealItemByID provides a mock function with given fields: ctx, id
func (_m *MockMealItemRepository) GetMealItemByID(ctx context.Context, id uint64) (*entities.MealItem, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetMealItemByID")
	}

	var r0 *entities.MealItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*entities.MealItem, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *entities.MealItem); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.MealItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockMealItemRepository creates a new instance of MockMealItemRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMealItemRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMealItemRepository {
	mock := &MockMealItemRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}