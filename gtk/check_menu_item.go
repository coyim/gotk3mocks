package gtk

type MockCheckMenuItem struct {
	MockMenuItem
}

func (m *MockCheckMenuItem) GetActive() bool {
	return m.Called().Bool(0)
}

func (m *MockCheckMenuItem) SetActive(v1 bool) {
	m.Called(v1)
}
