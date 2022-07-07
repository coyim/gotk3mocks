package gtk

import "github.com/coyim/gotk3mocks/glib"

type MockIconTheme struct {
	glib.MockObject
}

func (m *MockIconTheme) AddResourcePath(v1 string) {
	m.Called(v1)
}

func (m *MockIconTheme) AppendSearchPath(v1 string) {
	m.Called(v1)
}

func (m *MockIconTheme) GetExampleIconName() string {
	return m.Called().String(0)
}

func (m *MockIconTheme) HasIcon(v1 string) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockIconTheme) PrependSearchPath(v1 string) {
	m.Called(v1)
}
