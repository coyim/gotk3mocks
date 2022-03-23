package glib

import "github.com/stretchr/testify/mock"
import "github.com/coyim/gotk3adapter/glibi"

type MockObject struct {
	mock.Mock
}

func (m *MockObject) Connect(v1 string, v2 interface{}) glibi.SignalHandle {
	args := m.Called(v1, v2)
	return ret[glibi.SignalHandle](args, 0)
}

func (m *MockObject) ConnectAfter(v1 string, v2 interface{}) glibi.SignalHandle {
	args := m.Called(v1, v2)
	return ret[glibi.SignalHandle](args, 0)
}

func (m *MockObject) Emit(v1 string, v2 ...interface{}) (interface{}, error) {
	args := m.Called(v1, v2)
	return args.Get(0), args.Error(1)
}

func (m *MockObject) GetProperty(v1 string) (interface{}, error) {
	args := m.Called(v1)
	return args.Get(0), args.Error(1)
}

func (m *MockObject) SetProperty(v1 string, v2 interface{}) error {
	args := m.Called(v1, v2)
	return args.Error(0)
}

func (m *MockObject) Ref() {
	m.Called()
}

func (m *MockObject) Unref() {
	m.Called()
}
