// Code generated by mockery v2.42.2. DO NOT EDIT.

package openai

import mock "github.com/stretchr/testify/mock"

// MockOpenAIClient is an autogenerated mock type for the OpenAIClient type
type MockOpenAIClient struct {
	mock.Mock
}

// GetRecommendation provides a mock function with given fields: prompt, histories
func (_m *MockOpenAIClient) GetRecommendation(prompt string, histories []string) (string, error) {
	ret := _m.Called(prompt, histories)

	if len(ret) == 0 {
		panic("no return value specified for GetRecommendation")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string) (string, error)); ok {
		return rf(prompt, histories)
	}
	if rf, ok := ret.Get(0).(func(string, []string) string); ok {
		r0 = rf(prompt, histories)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(prompt, histories)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockOpenAIClient creates a new instance of MockOpenAIClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOpenAIClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOpenAIClient {
	mock := &MockOpenAIClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}