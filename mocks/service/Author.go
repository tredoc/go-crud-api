// Code generated by mockery v2.42.0. DO NOT EDIT.

package mockservice

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/tredoc/go-crud-api/pkg/types"
)

// Author is an autogenerated mock type for the Author type
type Author struct {
	mock.Mock
}

// CreateAuthor provides a mock function with given fields: _a0, _a1
func (_m *Author) CreateAuthor(_a0 context.Context, _a1 *types.Author) (*types.Author, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateAuthor")
	}

	var r0 *types.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Author) (*types.Author, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.Author) *types.Author); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Author)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.Author) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAuthor provides a mock function with given fields: _a0, _a1
func (_m *Author) DeleteAuthor(_a0 context.Context, _a1 int64) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAuthor")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllAuthors provides a mock function with given fields: _a0
func (_m *Author) GetAllAuthors(_a0 context.Context) ([]*types.Author, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetAllAuthors")
	}

	var r0 []*types.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*types.Author, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*types.Author); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Author)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthorByID provides a mock function with given fields: _a0, _a1
func (_m *Author) GetAuthorByID(_a0 context.Context, _a1 int64) (*types.Author, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorByID")
	}

	var r0 *types.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*types.Author, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *types.Author); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Author)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthorByName provides a mock function with given fields: _a0, _a1, _a2
func (_m *Author) GetAuthorByName(_a0 context.Context, _a1 string, _a2 string) (*types.Author, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorByName")
	}

	var r0 *types.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*types.Author, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *types.Author); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Author)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthorsByIDs provides a mock function with given fields: _a0, _a1
func (_m *Author) GetAuthorsByIDs(_a0 context.Context, _a1 []int64) ([]*types.Author, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorsByIDs")
	}

	var r0 []*types.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []int64) ([]*types.Author, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []int64) []*types.Author); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Author)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAuthor provides a mock function with given fields: _a0, _a1, _a2
func (_m *Author) UpdateAuthor(_a0 context.Context, _a1 int64, _a2 *types.UpdateAuthor) (*types.Author, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAuthor")
	}

	var r0 *types.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *types.UpdateAuthor) (*types.Author, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, *types.UpdateAuthor) *types.Author); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Author)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, *types.UpdateAuthor) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthor creates a new instance of Author. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthor(t interface {
	mock.TestingT
	Cleanup(func())
}) *Author {
	mock := &Author{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
