package gtk

type MockAboutDialog struct {
	MockDialog
}

func (m *MockAboutDialog) SetAuthors(v1 []string) {
	m.Called(v1)
}

func (m *MockAboutDialog) SetProgramName(v1 string) {
	m.Called(v1)
}

func (m *MockAboutDialog) SetVersion(v1 string) {
	m.Called(v1)
}

func (m *MockAboutDialog) SetLicense(v1 string) {
	m.Called(v1)
}

func (m *MockAboutDialog) SetWrapLicense(v1 bool) {
	m.Called(v1)
}
