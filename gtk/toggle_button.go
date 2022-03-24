package gtk

type MockToggleButton struct {
	MockButton
}

func (m *MockToggleButton) GetActive() bool {
	return m.Called().Bool(0)
}

func (m *MockToggleButton) SetActive(v1 bool) {
	m.Called(v1)
}
