// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/italomlaino/gobank/domain"
	mock "github.com/stretchr/testify/mock"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: documentNumber
func (_m *AccountRepository) Create(documentNumber int64) (*domain.Account, error) {
	ret := _m.Called(documentNumber)

	var r0 *domain.Account
	if rf, ok := ret.Get(0).(func(int64) *domain.Account); ok {
		r0 = rf(documentNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(documentNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *AccountRepository) Delete(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchByDocumentNumber provides a mock function with given fields: documentNumber
func (_m *AccountRepository) FetchByDocumentNumber(documentNumber int64) (*domain.Account, error) {
	ret := _m.Called(documentNumber)

	var r0 *domain.Account
	if rf, ok := ret.Get(0).(func(int64) *domain.Account); ok {
		r0 = rf(documentNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(documentNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByID provides a mock function with given fields: id
func (_m *AccountRepository) FetchByID(id int64) (*domain.Account, error) {
	ret := _m.Called(id)

	var r0 *domain.Account
	if rf, ok := ret.Get(0).(func(int64) *domain.Account); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
