package gtk

type MockComboBoxText struct {
	MockComboBox
}

func (m *MockComboBoxText) AppendText(v1 string) {
	m.Called(v1)
}

func (m *MockComboBoxText) GetActiveText() string {
	return m.Called().String(0)
}

func (m *MockComboBoxText) RemoveAll() {
	m.Called()
}
