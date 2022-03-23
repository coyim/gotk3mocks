package gtk

import "github.com/coyim/gotk3mocks/glib"

type MockAdjustment struct {
	glib.MockObject
}

func (m *MockAdjustment) GetLower() float64 {
	return ret[float64](m.Called(), 0)
}

func (m *MockAdjustment) GetPageSize() float64 {
	return ret[float64](m.Called(), 0)
}

func (m *MockAdjustment) GetUpper() float64 {
	return ret[float64](m.Called(), 0)
}

func (m *MockAdjustment) GetValue() float64 {
	return ret[float64](m.Called(), 0)
}

func (m *MockAdjustment) SetValue(v1 float64) {
	m.Called(v1)
}
