package gtk

type MockEventBox struct {
	MockBin
}

func (m *MockEventBox) SetAboveChild(v1 bool) {
	m.Called(v1)
}

func (m *MockEventBox) GetAboveChild() bool {
	return m.Called().Bool(0)
}

func (m *MockEventBox) SetVisibleWindow(v1 bool) {
	m.Called(v1)
}

func (m *MockEventBox) GetVisibleWindow() bool {
	return m.Called().Bool(0)
}
