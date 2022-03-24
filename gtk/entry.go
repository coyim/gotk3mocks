package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockEntry struct {
	MockWidget
}

func (m *MockEntry) GetText() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockEntry) SetHasFrame(v1 bool) {
	m.Called(v1)
}

func (m *MockEntry) GetVisibility() bool {
	return m.Called().Bool(0)
}

func (m *MockEntry) SetVisibility(v1 bool) {
	m.Called(v1)
}

func (m *MockEntry) SetText(v1 string) {
	m.Called(v1)
}

func (m *MockEntry) SetEditable(v1 bool) {
	m.Called(v1)
}

func (m *MockEntry) SetWidthChars(v1 int) {
	m.Called(v1)
}

func (m *MockEntry) GetAlignment() float32 {
	return ret[float32](m.Called(), 0)
}

func (m *MockEntry) SetAlignment(v1 float32) {
	m.Called(v1)
}

func (m *MockEntry) SetPosition(v1 int) {
	m.Called(v1)
}

func (m *MockEntry) GetPosition() int {
	return m.Called().Int(0)
}

func (m *MockEntry) SetCompletion(v1 gtki.EntryCompletion) {
	m.Called(v1)
}

func (m *MockEntry) SetPlaceholderText(v1 string) {
	m.Called(v1)
}
