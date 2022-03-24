package gtk

type MockRevealer struct {
	MockBin
}

func (m *MockRevealer) SetRevealChild(v1 bool) {
	m.Called(v1)
}

func (m *MockRevealer) GetRevealChild() bool {
	return m.Called().Bool(0)
}
