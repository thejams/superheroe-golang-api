// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	config "superheroe-api/superheroe-golang-api/src/config"

	entity "superheroe-api/superheroe-golang-api/src/entity"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Add provides a mock function with given fields: c, ctx
func (_m *Repository) Add(c *entity.Character, ctx context.Context) (*entity.Character, error) {
	/*
		args := mock.Called()
		result := args.Get(0)
		return result.([]entity.Post), args.Error(1)

		args := mock.Called()
		result := args.Get(0)
		return result.(*entity.Post), args.Error(1)
	*/
	ret := _m.Called(c, ctx)

	var r0 *entity.Character
	var r1 error
	
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entity.Character)
	}
	
	r1 = ret.Error(1)
	

	return r0, r1
}

// Close provides a mock function with given fields: _a0
func (_m *Repository) Close(_a0 context.Context) error {
	var r0 error
	ret := _m.Called(_a0)

	r0 = ret.Error(0)
	return r0
}

// Conn provides a mock function with given fields: _a0, _a1
func (_m *Repository) Conn(_a0 context.Context, _a1 *config.APPConfig) error {
	var r0 error
	ret := _m.Called(_a0, _a1)

	r0 = ret.Error(0)
	return r0
}

// Delete provides a mock function with given fields: id, ctx
func (_m *Repository) Delete(id string, ctx context.Context) (string, error) {
	ret := _m.Called(id, ctx)

	var r0 string
	r0 = ret.Get(0).(string)

	var r1 error
	r1 = ret.Error(1)

	return r0, r1
}

// Edit provides a mock function with given fields: id, c, ctx
func (_m *Repository) Edit(id string, c *entity.Character, ctx context.Context) (*entity.Character, error) {
	ret := _m.Called(id, c, ctx)

	var r0 *entity.Character
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entity.Character)
	}

	var r1 error
	r1 = ret.Error(1)
	
	return r0, r1
}

// Get provides a mock function with given fields: id, ctx
func (_m *Repository) Get(id string, ctx context.Context) (*entity.Character, error) {
	ret := _m.Called(id, ctx)

	var r0 *entity.Character
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entity.Character)
	}

	var r1 error
	r1 = ret.Error(1)

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]entity.Character, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Character
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]entity.Character)
	}

	var r1 error
	r1 = ret.Error(1)

	return r0, r1
}
