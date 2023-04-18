// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/hongdangcseiu/go-back-end/domain"
	mock "github.com/stretchr/testify/mock"
)

// RefreshTokenUsecase is an autogenerated mock type for the RefreshTokenUsecase type
type RefreshTokenUsecase struct {
	mock.Mock
}

// CreateAccessToken provides a mock function with given fields: user, secret, expiry
func (_m *RefreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	ret := _m.Called(user, secret, expiry)

	var r0 string
	if rf, ok := ret.Get(0).(func(*domain.User, string, int) string); ok {
		r0 = rf(user, secret, expiry)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.User, string, int) error); ok {
		r1 = rf(user, secret, expiry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRefreshToken provides a mock function with given fields: user, secret, expiry
func (_m *RefreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	ret := _m.Called(user, secret, expiry)

	var r0 string
	if rf, ok := ret.Get(0).(func(*domain.User, string, int) string); ok {
		r0 = rf(user, secret, expiry)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.User, string, int) error); ok {
		r1 = rf(user, secret, expiry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExtractIDFromToken provides a mock function with given fields: requestToken, secret
func (_m *RefreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	ret := _m.Called(requestToken, secret)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(requestToken, secret)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(requestToken, secret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: c, id
func (_m *RefreshTokenUsecase) GetUserByID(c context.Context, id string) (domain.User, error) {
	ret := _m.Called(c, id)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRefreshTokenUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewRefreshTokenUsecase creates a new instance of RefreshTokenUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRefreshTokenUsecase(t mockConstructorTestingTNewRefreshTokenUsecase) *RefreshTokenUsecase {
	mock := &RefreshTokenUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
