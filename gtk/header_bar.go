package gtk

type MockHeaderBar struct {
	MockContainer
}

func (m *MockHeaderBar) SetSubtitle(v1 string) {
	m.Called(v1)
}

func (m *MockHeaderBar) SetShowCloseButton(v1 bool) {
	m.Called(v1)
}

func (m *MockHeaderBar) GetShowCloseButton() bool {
	return m.Called().Bool(0)
}
