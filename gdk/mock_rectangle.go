package gdk

import "github.com/stretchr/testify/mock"

type MockRectangle struct {
	mock.Mock
}

func (m *MockRectangle) GetY() int {
	return m.Called().Int(0)
}
