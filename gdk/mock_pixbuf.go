package gdk_mock

import "github.com/stretchr/testify/mock"

type MockPixbuf struct {
	mock.Mock
}

func (m *MockPixbuf) SavePNG(v1 string, v2 int) error {
	return m.Called(v1, v2).Error(0)
}
