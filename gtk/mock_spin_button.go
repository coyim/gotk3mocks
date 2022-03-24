package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockSpinButton struct {
	MockEntry
}

func (m *MockSpinButton) GetValueAsInt() int {
	return m.Called().Int(0)
}

func (m *MockSpinButton) SetValue(v1 float64) {
	m.Called(v1)
}

func (m *MockSpinButton) GetValue() float64 {
	return ret[float64](m.Called(), 0)
}

func (m *MockSpinButton) GetAdjustment() gtki.Adjustment {
	return ret[gtki.Adjustment](m.Called(), 0)
}

func (m *MockSpinButton) SetRange(v1 float64, v2 float64) {
	m.Called(v1, v2)
}

func (m *MockSpinButton) SetIncrements(v1 float64, v2 float64) {
	m.Called(v1, v2)
}
