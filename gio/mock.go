package gio

import (
	"github.com/coyim/gotk3adapter/gioi"
	"github.com/stretchr/testify/mock"
)

type getter interface {
	Get(int) interface{}
}

func ret[T any](args getter, index int) T {
	ret, _ := args.Get(index).(T)
	return ret
}

type Mock struct {
	mock.Mock
}

func (m *Mock) LoadResource(v1 string) (gioi.Resource, error) {
	returns := m.Called(v1)
	return ret[gioi.Resource](returns, 0), returns.Error(1)
}

func (m *Mock) NewResourceFromData(v1 []byte) (gioi.Resource, error) {
	returns := m.Called(v1)
	return ret[gioi.Resource](returns, 0), returns.Error(1)
}

func (m *Mock) RegisterResource(v1 gioi.Resource) {
	m.Called(v1)
}

func (m *Mock) UnregisterResource(v1 gioi.Resource) {
	m.Called(v1)
}
