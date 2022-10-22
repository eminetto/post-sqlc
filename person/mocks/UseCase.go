// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	person "github.com/eminetto/post-sqlc/person"
	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, firstName, lastName
func (_m *UseCase) Create(ctx context.Context, firstName string, lastName string) (person.ID, error) {
	ret := _m.Called(ctx, firstName, lastName)

	var r0 person.ID
	if rf, ok := ret.Get(0).(func(context.Context, string, string) person.ID); ok {
		r0 = rf(ctx, firstName, lastName)
	} else {
		r0 = ret.Get(0).(person.ID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, firstName, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *UseCase) Delete(ctx context.Context, id person.ID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, person.ID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *UseCase) Get(ctx context.Context, id person.ID) (*person.Person, error) {
	ret := _m.Called(ctx, id)

	var r0 *person.Person
	if rf, ok := ret.Get(0).(func(context.Context, person.ID) *person.Person); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*person.Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, person.ID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *UseCase) List(ctx context.Context) ([]*person.Person, error) {
	ret := _m.Called(ctx)

	var r0 []*person.Person
	if rf, ok := ret.Get(0).(func(context.Context) []*person.Person); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*person.Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, query
func (_m *UseCase) Search(ctx context.Context, query string) ([]*person.Person, error) {
	ret := _m.Called(ctx, query)

	var r0 []*person.Person
	if rf, ok := ret.Get(0).(func(context.Context, string) []*person.Person); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*person.Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, e
func (_m *UseCase) Update(ctx context.Context, e *person.Person) error {
	ret := _m.Called(ctx, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *person.Person) error); ok {
		r0 = rf(ctx, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}