// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

type Cache_Expecter struct {
	mock *mock.Mock
}

func (_m *Cache) EXPECT() *Cache_Expecter {
	return &Cache_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: key
func (_m *Cache) Get(key string) (string, bool) {
	ret := _m.Called(key)

	var r0 string
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (string, bool)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Cache_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Cache_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
func (_e *Cache_Expecter) Get(key interface{}) *Cache_Get_Call {
	return &Cache_Get_Call{Call: _e.mock.On("Get", key)}
}

func (_c *Cache_Get_Call) Run(run func(key string)) *Cache_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Cache_Get_Call) Return(_a0 string, _a1 bool) *Cache_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Cache_Get_Call) RunAndReturn(run func(string) (string, bool)) *Cache_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields:
func (_m *Cache) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Cache_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type Cache_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
func (_e *Cache_Expecter) Init() *Cache_Init_Call {
	return &Cache_Init_Call{Call: _e.mock.On("Init")}
}

func (_c *Cache_Init_Call) Run(run func()) *Cache_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Cache_Init_Call) Return(_a0 error) *Cache_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Cache_Init_Call) RunAndReturn(run func() error) *Cache_Init_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: key, val
func (_m *Cache) Set(key string, val string) error {
	ret := _m.Called(key, val)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(key, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Cache_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type Cache_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - key string
//   - val string
func (_e *Cache_Expecter) Set(key interface{}, val interface{}) *Cache_Set_Call {
	return &Cache_Set_Call{Call: _e.mock.On("Set", key, val)}
}

func (_c *Cache_Set_Call) Run(run func(key string, val string)) *Cache_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Cache_Set_Call) Return(_a0 error) *Cache_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Cache_Set_Call) RunAndReturn(run func(string, string) error) *Cache_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewCache creates a new instance of Cache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *Cache {
	mock := &Cache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
