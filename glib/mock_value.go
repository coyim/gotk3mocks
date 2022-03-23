package glib

import "github.com/stretchr/testify/mock"

type MockValue struct {
	mock.Mock
}

func (m *MockValue) GetString() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockValue) GoValue() (interface{}, error) {
	args := m.Called()
	return args.Get(0), args.Error(1)
}
