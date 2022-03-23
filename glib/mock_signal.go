package glib

import "github.com/stretchr/testify/mock"

type MockSignal struct {
	mock.Mock
}

func (m *MockSignal) String() string {
	return m.Called().String(0)
}
