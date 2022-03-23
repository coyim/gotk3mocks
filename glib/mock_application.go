package glib

type MockApplication struct {
	MockObject
}

func (m *MockApplication) Quit() {
	m.Called()
}

func (m *MockApplication) Run(v1 []string) int {
	return m.Called(v1).Int(0)
}
