package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockTextView struct {
	MockContainer
}

func (m *MockTextView) SetEditable(v1 bool) {
	m.Called(v1)
}

func (m *MockTextView) SetCursorVisible(v1 bool) {
	m.Called(v1)
}

func (m *MockTextView) SetBuffer(v1 gtki.TextBuffer) {
	m.Called(v1)
}

func (m *MockTextView) GetBuffer() (gtki.TextBuffer, error) {
	args := m.Called()
	return ret[gtki.TextBuffer](args, 0), args.Error(1)
}

func (m *MockTextView) ForwardDisplayLine(v1 gtki.TextIter) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextView) BackwardDisplayLine(v1 gtki.TextIter) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextView) ForwardDisplayLineEnd(v1 gtki.TextIter) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextView) BackwardDisplayLineStart(v1 gtki.TextIter) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextView) StartsDisplayLine(v1 gtki.TextIter) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextView) SetJustification(v1 gtki.Justification) {
	m.Called(v1)
}

func (m *MockTextView) GetJustification() gtki.Justification {
	return ret[gtki.Justification](m.Called(), 0)
}

func (m *MockTextView) MoveVisually(v1 gtki.TextIter, v2 int) bool {
	return m.Called(v1).Bool(0)
}
