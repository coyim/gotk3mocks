package gdk_mock

import "github.com/stretchr/testify/mock"

type MockRgba struct {
	mock.Mock
}

func (m *MockRgba) String() string {
	return m.Called().String(0)
}

func (m *MockRgba) GetRed() float64 {
	return ret[float64](m.Called(), 0)
}

func (m *MockRgba) GetGreen() float64 {
	return ret[float64](m.Called(), 0)
}

func (m *MockRgba) GetBlue() float64 {
	return ret[float64](m.Called(), 0)
}

func (m *MockRgba) GetAlpha() float64 {
	return ret[float64](m.Called(), 0)
}

func (m *MockRgba) SetRed(c float64) {
	m.Called(c)
}

func (m *MockRgba) SetGreen(c float64) {
	m.Called(c)
}

func (m *MockRgba) SetBlue(c float64) {
	m.Called(c)
}

func (m *MockRgba) SetAlpha(c float64) {
	m.Called(c)
}

func (m *MockRgba) Colors() (r, g, b, a float64) {
	args := m.Called()
	return ret[float64](args, 0),
		ret[float64](args, 1),
		ret[float64](args, 2),
		ret[float64](args, 3)
}

func (m *MockRgba) SetColors(r, g, b, a float64) {
	m.Called(r, g, b, a)
}

func (m *MockRgba) Parse(spec string) bool {
	return m.Called(spec).Bool(0)
}
