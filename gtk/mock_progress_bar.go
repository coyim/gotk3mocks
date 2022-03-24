package gtk

// MockProgressBar is a mock of the representation of GTK's GtkProgressBar.
type MockProgressBar struct {
	MockWidget
}

func (m *MockProgressBar) SetFraction(v1 float64) {
	m.Called(v1)
}

func (m *MockProgressBar) GetFraction() float64 {
	return ret[float64](m.Called(), 0)
}

func (m *MockProgressBar) SetShowText(v1 bool) {
	m.Called(v1)
}

func (m *MockProgressBar) GetShowText() bool {
	return m.Called().Bool(0)
}

func (m *MockProgressBar) SetText(v1 string) {
	m.Called(v1)
}
