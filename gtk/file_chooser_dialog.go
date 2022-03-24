package gtk

type MockFileChooserDialog struct {
	MockDialog
}

func (m *MockFileChooserDialog) GetFilename() string {
	return m.Called().String(0)
}

func (m *MockFileChooserDialog) SetCurrentName(v1 string) {
	m.Called(v1)
}

func (m *MockFileChooserDialog) SetDoOverwriteConfirmation(v1 bool) {
	m.Called(v1)
}
