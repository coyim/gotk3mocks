package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockDialog struct {
	MockWindow
}

func (m *MockDialog) Run() int {
	return m.Called().Int(0)
}

func (m *MockDialog) SetDefaultResponse(v1 gtki.ResponseType) {
	m.Called(v1)
}
