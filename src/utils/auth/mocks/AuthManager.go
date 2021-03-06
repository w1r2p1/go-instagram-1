// Code generated by mockery v1.0.0
package mocks

import auth "github.com/hieven/go-instagram/src/utils/auth"
import mock "github.com/stretchr/testify/mock"

// AuthManager is an autogenerated mock type for the AuthManager type
type AuthManager struct {
	mock.Mock
}

// GenerateRankToken provides a mock function with given fields: userID
func (_m *AuthManager) GenerateRankToken(userID int64) string {
	ret := _m.Called(userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GenerateSignature provides a mock function with given fields: payload
func (_m *AuthManager) GenerateSignature(payload *auth.SignaturePayload) (string, string, error) {
	ret := _m.Called(payload)

	var r0 string
	if rf, ok := ret.Get(0).(func(*auth.SignaturePayload) string); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(*auth.SignaturePayload) string); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*auth.SignaturePayload) error); ok {
		r2 = rf(payload)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GenerateUUID provides a mock function with given fields:
func (_m *AuthManager) GenerateUUID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
